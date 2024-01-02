package categorytransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/module/category/categorybiz"
	"book-store-management-backend/module/category/categoryrepo"
	"book-store-management-backend/module/category/categorystore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Get all categories
// @Tags categories
// @Accept json
// @Produce json
// @Response 200 {object} categorymodel.ResGetAllCategory
// @Router /categories/all [get]
func GetAllCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := categorystore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := categoryrepo.NewGetAllCategoryRepo(store)

		biz := categorybiz.NewGetAllCategoryBiz(repo)

		result, err := biz.GetAllCategory(c.Request.Context())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
