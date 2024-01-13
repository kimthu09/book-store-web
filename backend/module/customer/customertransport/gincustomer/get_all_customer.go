package gincustomer

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/module/customer/customerbiz"
	"book-store-management-backend/module/customer/customerrepo"
	"book-store-management-backend/module/customer/customerstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Get all customer
// @Tags customers
// @Accept json
// @Produce json
// @Response 200 {object} customermodel.ResGetAllCustomer "list customer"
// @Response 400 {object} common.AppError "error"
// @Router /customers/all [get]
func GetAllCustomer(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := customerstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := customerrepo.NewGetAllCustomerRepo(store)

		biz := customerbiz.NewGetAllCustomerBiz(repo)

		result, err := biz.GetAllCustomer(c.Request.Context())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
