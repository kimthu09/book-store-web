package ginsupplier

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/supplier/supplierbiz"
	"book-store-management-backend/module/supplier/suppliermodel"
	"book-store-management-backend/module/supplier/supplierrepo"
	"book-store-management-backend/module/supplier/supplierstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Update info supplier
// @Tags suppliers
// @Accept json
// @Produce json
// @Param id path string true "supplier id"
// @Param supplier body suppliermodel.ReqUpdateInfoSupplier true "supplier info to update"
// @Response 200 {object} common.ResSuccess "status of response"
// @Response 400 {object} common.AppError "error"
// @Router /suppliers/{id} [patch]
func UpdateInfoSupplier(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data suppliermodel.ReqUpdateInfoSupplier

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := supplierstore.NewSQLStore(db)
		repo := supplierrepo.NewUpdateInfoSupplierRepo(store)

		business := supplierbiz.NewUpdateInfoSupplierBiz(repo, requester)

		if err := business.UpdateInfoSupplier(c.Request.Context(), id, &data); err != nil {
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
