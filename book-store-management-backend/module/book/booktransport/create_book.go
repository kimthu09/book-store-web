package booktransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/author/authorrepo"
	"book-store-management-backend/module/author/authorstore"
	"book-store-management-backend/module/book/bookbiz"
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/book/bookrepo"
	"book-store-management-backend/module/book/bookstore"
	"book-store-management-backend/module/category/categoryrepo"
	"book-store-management-backend/module/category/categorystore"
	"book-store-management-backend/module/publisher/publisherrepo"
	"book-store-management-backend/module/publisher/publisherstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateBook
// @BasePath /v1
// @Security BearerAuth
// @Summary Create book name, desc, authors, categories, publisher, etc.
// @Tags books
// @Accept json
// @Produce json
// @Param book body bookmodel.ReqCreateBook true "Create book"
// @Response 200 {object} bookmodel.ResCreateBook "book id"
// @Router /books [post]
func CreateBook(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqData bookmodel.ReqCreateBook
		if err := c.ShouldBind(&reqData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := bookstore.NewSQLStore(db)
		authorStore := authorstore.NewSQLStore(db)
		publisherStore := publisherstore.NewSQLStore(db)
		categoryStore := categorystore.NewSQLStore(db)

		repo := bookrepo.NewCreateBookRepo(store)
		authorRepo := authorrepo.NewExistAuthorRepo(authorStore)
		publisherRepo := publisherrepo.NewExistPublisherRepo(publisherStore)
		categoryRepo := categoryrepo.NewExistCategoryRepo(categoryStore)

		gen := generator.NewShortIdGenerator()

		biz := bookbiz.NewCreateBookBiz(gen, repo, authorRepo, publisherRepo, categoryRepo, requester)

		var resData bookmodel.ResCreateBook

		if err := biz.CreateBook(c.Request.Context(), &reqData, &resData); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(resData))
	}
}
