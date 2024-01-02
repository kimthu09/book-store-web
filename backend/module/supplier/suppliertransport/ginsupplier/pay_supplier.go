package ginsupplier

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/importnote/importnotestore"
	"book-store-management-backend/module/supplier/supplierbiz"
	"book-store-management-backend/module/supplier/suppliermodel"
	"book-store-management-backend/module/supplier/supplierrepo"
	"book-store-management-backend/module/supplier/supplierstore"
	"book-store-management-backend/module/supplierdebt/supplierdebtstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Pay supplier
// @Tags suppliers
// @Accept json
// @Produce json
// @Param id path string true "supplier id"
// @Param supplier body suppliermodel.ReqUpdateDebtSupplier true "pay information"
// @Response 200 {object} common.ResSuccess "status of response"
// @Response 400 {object} common.AppError "error"
// @Router /suppliers/{id}/pay [post]
func PaySupplier(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data suppliermodel.ReqUpdateDebtSupplier

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		supplierStore := supplierstore.NewSQLStore(db)
		supplierDebtStore := supplierdebtstore.NewSQLStore(db)
		importNoteStore := importnotestore.NewSQLStore(db)
		repo := supplierrepo.NewUpdatePayRepo(
			supplierStore, supplierDebtStore, importNoteStore,
		)

		gen := generator.NewShortIdGenerator()

		business := supplierbiz.NewUpdatePayBiz(gen, repo, requester)

		idSupplierDebt, err := business.PaySupplier(c.Request.Context(), id, &data)

		if err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(idSupplierDebt))
	}
}
