package booktitletransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/author/authorrepo"
	"book-store-management-backend/module/author/authorstore"
	"book-store-management-backend/module/booktitle/booktitlebiz"
	"book-store-management-backend/module/booktitle/booktitlerepo"
	"book-store-management-backend/module/booktitle/booktitlestore"
	"book-store-management-backend/module/category/categoryrepo"
	"book-store-management-backend/module/category/categorystore"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Get all booktitle detail
// @Tags booktitles
// @Accept json
// @Produce json
// @Param page query common.Paging false "page"
// @Response 200 {object} booktitlemodel.ResBookTitleDetail
// @Router /booktitles [get]
func GetBookTitleDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := strings.Trim(c.Param("id"), " ")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "id is invalid",
			})
			return
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection()
		store := booktitlestore.NewSQLStore(db)
		authorStore := authorstore.NewSQLStore(db)
		categoryStore := categorystore.NewSQLStore(db)

		repo := booktitlerepo.NewDetailBookTitleRepo(store)
		authorRepo := authorrepo.NewAuthorPublicRepo(authorStore)
		categoryRepo := categoryrepo.NewCategoryPublicRepo(categoryStore)

		biz := booktitlebiz.NewGetBookTitleDetailBiz(repo, authorRepo, categoryRepo, requester)
		data, err := biz.GetBookTitleDetail(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
