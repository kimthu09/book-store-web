package gininvoice

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/book/bookstore"
	"book-store-management-backend/module/customer/customerstore"
	"book-store-management-backend/module/invoice/invoicebiz"
	"book-store-management-backend/module/invoice/invoicemodel"
	"book-store-management-backend/module/invoice/invoicerepo"
	"book-store-management-backend/module/invoice/invoicestore"
	"book-store-management-backend/module/invoicedetail/invoicedetailstore"
	"book-store-management-backend/module/shopgeneral/shopgeneralstore"
	"book-store-management-backend/module/stockchangehistory/stockchangehistorystore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Create invoice
// @Tags invoices
// @Accept json
// @Produce json
// @Param invoice body invoicemodel.ReqCreateInvoice true "invoice need to create"
// @Response 200 {object} invoicemodel.ResCreateInvoice "invoice id"
// @Response 400 {object} common.AppError "error"
// @Router /invoices [post]
func CreateInvoice(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data invoicemodel.ReqCreateInvoice

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		data.CreatedBy = requester.GetUserId()

		db := appCtx.GetMainDBConnection().Begin()

		invoiceStore := invoicestore.NewSQLStore(db)
		invoiceDetailStore := invoicedetailstore.NewSQLStore(db)
		customerStore := customerstore.NewSQLStore(db)
		bookStore := bookstore.NewSQLStore(db)
		stockChangeHistoryStore := stockchangehistorystore.NewSQLStore(db)
		shopGeneralStore := shopgeneralstore.NewSQLStore(db)

		repo := invoicerepo.NewCreateInvoiceRepo(
			invoiceStore,
			invoiceDetailStore,
			customerStore,
			bookStore,
			stockChangeHistoryStore,
			shopGeneralStore,
		)

		gen := generator.NewShortIdGenerator()

		biz := invoicebiz.NewCreateInvoiceBiz(gen, repo, requester)

		err := biz.CreateInvoice(c.Request.Context(), &data)
		if err != nil {
			db.Rollback()
			panic(err)
		}

		shopStore := shopgeneralstore.NewSQLStore(appCtx.GetMainDBConnection())

		seeDetailRepo := invoicerepo.NewSeeInvoiceDetailRepo(invoiceDetailStore, invoiceStore)

		seeDetailBiz := invoicebiz.NewSeeDetailInvoiceBiz(
			seeDetailRepo, shopStore, requester)

		result, err := seeDetailBiz.SeeDetailInvoice(c.Request.Context(), data.Id)
		if err != nil {
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
