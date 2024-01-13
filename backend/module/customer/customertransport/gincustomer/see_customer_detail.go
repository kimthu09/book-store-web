package gincustomer

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/customer/customerbiz"
	"book-store-management-backend/module/customer/customerrepo"
	"book-store-management-backend/module/customer/customerstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary See detail of customer
// @Tags customers
// @Accept json
// @Produce json
// @Param id path string true "customer id"
// @Response 200 {object} customermodel.Customer "customer"
// @Response 400 {object} common.AppError "error"
// @Router /customers/{id} [get]
func SeeCustomerDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		customerStore := customerstore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := customerrepo.NewSeeCustomerDetailRepo(customerStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := customerbiz.NewSeeCustomerDetailBiz(repo, requester)

		result, err := biz.SeeCustomerDetail(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
