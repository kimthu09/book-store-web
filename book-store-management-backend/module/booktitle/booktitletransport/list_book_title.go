package booktitletransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/author/authorrepo"
	"book-store-management-backend/module/author/authorstore"
	"book-store-management-backend/module/booktitle/booktitlebiz"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"book-store-management-backend/module/booktitle/booktitlerepo"
	"book-store-management-backend/module/booktitle/booktitlestore"
	"book-store-management-backend/module/category/categoryrepo"
	"book-store-management-backend/module/category/categorystore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListBookTitle
// @BasePath /v1
// @Security BearerAuth
// @Summary Get all booktitles
// @Tags booktitles
// @Accept json
// @Produce json
// @Param page query common.Paging false "page"
// @Param filter query booktitlemodel.Filter false "filter"
// @Response 200 {object} booktitlemodel.ResListBookTitle
// @Router /booktitles [get]
func ListBookTitle(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter booktitlemodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		db := appCtx.GetMainDBConnection()
		store := booktitlestore.NewSQLStore(db)
		authorStore := authorstore.NewSQLStore(db)
		categoryStore := categorystore.NewSQLStore(db)

		repo := booktitlerepo.NewListBookTitleRepo(store)
		authorRepo := authorrepo.NewAuthorPublicRepo(authorStore)
		categoryRepo := categoryrepo.NewCategoryPublicRepo(categoryStore)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := booktitlebiz.NewListBookTitleBiz(repo, authorRepo, categoryRepo, requester)
		data, err := biz.ListBookTitle(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, booktitlemodel.NewResListBookTitle(data, paging, filter))
	}
}
