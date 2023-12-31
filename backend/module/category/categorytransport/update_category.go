package categorytransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/category/categorybiz"
	"book-store-management-backend/module/category/categorymodel"
	"book-store-management-backend/module/category/categoryrepo"
	"book-store-management-backend/module/category/categorystore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Update category
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "category id"
// @Param category body categorymodel.ReqUpdateCategory true "category info to update"
// @Response 200 {object} common.ResSuccess "status of response"
// @Response 400 {object} common.AppError "error"
// @Router /categories/{id} [patch]
func UpdateCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data categorymodel.ReqUpdateCategory

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := categorystore.NewSQLStore(db)
		repo := categoryrepo.NewUpdateCategoryRepo(store)

		business := categorybiz.NewUpdateCategoryBiz(repo, requester)

		if err := business.UpdateCategory(c.Request.Context(), id, &data); err != nil {
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
