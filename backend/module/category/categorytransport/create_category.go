package categorytransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/category/categorybiz"
	"book-store-management-backend/module/category/categorymodel"
	"book-store-management-backend/module/category/categoryrepo"
	"book-store-management-backend/module/category/categorystore"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Create category with name
// @Tags categories
// @Accept json
// @Produce json
// @Param category body categorymodel.ReqCreateCategory true "Create category"
// @Response 200 {object} categorymodel.ResCreateCategory "category id"
// @Router /categories [post]
func CreateCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data categorymodel.ReqCreateCategory
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		db := appCtx.GetMainDBConnection().Begin()
		store := categorystore.NewSQLStore(db)
		repo := categoryrepo.NewCreateCategoryRepo(store)

		gen := generator.NewShortIdGenerator()

		business := categorybiz.NewCreateCategoryBiz(gen, repo, requester)

		tmpData := categorymodel.Category{
			Name: data.Name,
		}
		if err := business.CreateCategory(c.Request.Context(), &tmpData); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, categorymodel.ResCreateCategory{Id: tmpData.Id})
	}
}
