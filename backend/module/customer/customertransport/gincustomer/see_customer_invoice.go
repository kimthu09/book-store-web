package gincustomer

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/customer/customerbiz"
	"book-store-management-backend/module/customer/customermodel"
	"book-store-management-backend/module/customer/customerrepo"
	"book-store-management-backend/module/invoice/invoicestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary See invoices of customer
// @Tags customers
// @Accept json
// @Produce json
// @Param id path string true "customer id"
// @Param page query common.Paging false "page"
// @Param filter query customermodel.FilterInvoice false "filter"
// @Response 200 {object} customermodel.ResSeeInvoiceCustomer "list invoice"
// @Response 400 {object} common.AppError "error"
// @Router /customers/{id}/invoices [get]
func SeeCustomerInvoice(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var filter customermodel.FilterInvoice
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		invoiceStore := invoicestore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := customerrepo.NewSeeCustomerInvoiceRepo(invoiceStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := customerbiz.NewSeeCustomerInvoiceBiz(repo, requester)

		result, err := biz.SeeCustomerInvoice(c.Request.Context(), id, &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
