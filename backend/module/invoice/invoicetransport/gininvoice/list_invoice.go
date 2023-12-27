package gininvoice

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/invoice/invoicebiz"
	"book-store-management-backend/module/invoice/invoicemodel"
	"book-store-management-backend/module/invoice/invoicerepo"
	"book-store-management-backend/module/invoice/invoicestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary List invoice
// @Tags invoices
// @Accept json
// @Produce json
// @Param page query common.Paging false "page"
// @Param filter query invoicemodel.Filter false "filter"
// @Response 200 {object} invoicemodel.ResListInvoice "list invoice"
// @Response 400 {object} common.AppError "error"
// @Router /invoices [get]
func ListInvoice(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter invoicemodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := invoicestore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := invoicerepo.NewListImportNoteRepo(store)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := invoicebiz.NewListImportNoteBiz(repo, requester)

		result, err := biz.ListInvoice(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
