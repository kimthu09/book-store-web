package gincustomer

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/customer/customerbiz"
	"book-store-management-backend/module/customer/customermodel"
	"book-store-management-backend/module/customer/customerrepo"
	"book-store-management-backend/module/customer/customerstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary List customer
// @Tags customers
// @Accept json
// @Produce json
// @Param page query common.Paging false "page"
// @Param filter query customermodel.Filter false "filter"
// @Response 200 {object} customermodel.ResListCustomer "list customer"
// @Response 400 {object} common.AppError "error"
// @Router /customers [get]
func ListCustomer(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter customermodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := customerstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := customerrepo.NewListCustomerRepo(store)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := customerbiz.NewListCustomerBiz(repo, requester)

		result, err := biz.ListCustomer(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
