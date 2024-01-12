package gincustomer

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/generator"
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
// @Summary Create customer
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body customermodel.ReqCreateCustomer true "customer need to create"
// @Response 200 {object} customermodel.ResCreateCustomer "customer id"
// @Response 400 {object} common.AppError "error"
// @Router /customers [post]
func CreateCustomer(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data customermodel.ReqCreateCustomer

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := customerstore.NewSQLStore(db)
		repo := customerrepo.NewCreateCustomerRepo(store)

		gen := generator.NewShortIdGenerator()

		business := customerbiz.NewCreateCustomerBiz(gen, repo, requester)

		if err := business.CreateCustomer(c.Request.Context(), &data); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
