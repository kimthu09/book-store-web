package gininvoice

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/invoice/invoicebiz"
	"book-store-management-backend/module/invoice/invoicerepo"
	"book-store-management-backend/module/invoice/invoicestore"
	"book-store-management-backend/module/invoicedetail/invoicedetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary See detail invoice
// @Tags invoices
// @Accept json
// @Produce json
// @Param id path string true "invoice id"
// @Response 200 {object} invoicemodel.ResSeeDetailInvoice "invoice"
// @Response 400 {object} common.AppError "error"
// @Router /invoices/{id} [get]
func SeeInvoiceDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		db := appCtx.GetMainDBConnection()
		invoiceStore := invoicestore.NewSQLStore(db)
		invoiceDetailStore := invoicedetailstore.NewSQLStore(db)

		repo := invoicerepo.NewSeeInvoiceDetailRepo(invoiceStore, invoiceDetailStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := invoicebiz.NewSeeInvoiceDetailBiz(
			repo, requester)

		result, err := biz.SeeInvoiceDetail(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
