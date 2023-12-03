package ginsupplier

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/generator"
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
// @Summary Create supplier
// @Tags suppliers
// @Accept json
// @Produce json
// @Param supplier body suppliermodel.ReqCreateSupplier true "supplier need to create"
// @Response 200 {object} suppliermodel.ResSupplierCreate "supplier id"
// @Response 400 {object} common.AppError "error"
// @Router /suppliers [post]
func CreateSupplier(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data suppliermodel.ReqCreateSupplier

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := supplierstore.NewSQLStore(db)
		repo := supplierrepo.NewCreateSupplierRepo(store)

		gen := generator.NewShortIdGenerator()

		business := supplierbiz.NewCreateSupplierBiz(gen, repo, requester)

		if err := business.CreateSupplier(c.Request.Context(), &data); err != nil {
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
