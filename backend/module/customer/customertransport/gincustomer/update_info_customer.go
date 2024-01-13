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
// @Summary Update info customer
// @Tags customers
// @Accept json
// @Produce json
// @Param id path string true "customer id"
// @Param customer body customermodel.ReqUpdateInfoCustomer true "customer info to update"
// @Response 200 {object} common.ResSuccess "status of response"
// @Response 400 {object} common.AppError "error"
// @Router /customers/{id} [patch]
func UpdateInfoCustomer(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data customermodel.ReqUpdateInfoCustomer

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := customerstore.NewSQLStore(db)
		repo := customerrepo.NewUpdateInfoCustomerRepo(store)
		biz := customerbiz.NewUpdateInfoCustomerBiz(repo, requester)

		if err := biz.UpdateInfoCustomer(c.Request.Context(), id, &data); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
