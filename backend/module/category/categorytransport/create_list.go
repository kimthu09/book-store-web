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
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Create list category
// @Tags categories
// @Accept json
// @Produce json
// @Param category body categorymodel.ReqCreateListCategory true "list name of category"
// @Response 200 {object} common.ResSuccess "status of response"
// @Response 400 {object} common.AppError "error"
// @Router /categories/many [post]
func CreateListCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data categorymodel.ReqCreateListCategory
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		db := appCtx.GetMainDBConnection().Begin()
		store := categorystore.NewSQLStore(db)
		repo := categoryrepo.NewCreateCategoryRepo(store)

		gen := generator.NewShortIdGenerator()

		business := categorybiz.NewCreateListCategoryBiz(gen, repo, requester)

		var tmpData []categorymodel.Category
		for _, v := range data.Names {
			tmpData = append(tmpData, categorymodel.Category{Name: v})
		}

		if err := business.CreateListCategory(c.Request.Context(), tmpData); err != nil {
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
