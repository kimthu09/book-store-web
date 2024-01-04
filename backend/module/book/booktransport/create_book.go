package booktransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/book/bookbiz"
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/book/bookrepo"
	"book-store-management-backend/module/book/bookstore"
	"book-store-management-backend/module/booktitle/booktitlestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateBook
// @BasePath /v1
// @Security BearerAuth
// @Summary Create Book
// @Tags books
// @Accept json
// @Produce json
// @Param book body bookmodel.ReqCreateBook true "Create Book"
// @Response 200 {object} bookmodel.ResCreateBook "book id"
// @Router /books [post]
func CreateBook(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqData bookmodel.ReqCreateBook
		if err := c.ShouldBind(&reqData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection().Begin()

		gen := generator.NewShortIdGenerator()
		bookStore := bookstore.NewSQLStore(db)
		bookTitleStore := booktitlestore.NewSQLStore(db)

		repo := bookrepo.NewCreateBookRepo(bookStore, bookTitleStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := bookbiz.NewCreateBookBiz(gen, repo, requester)

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
