package booktransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/module/author/authorstore"
	"book-store-management-backend/module/book/bookbiz"
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/book/bookrepo"
	"book-store-management-backend/module/book/bookstore"
	"book-store-management-backend/module/category/categorystore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Get all book
// @Tags books
// @Accept json
// @Produce json
// @Response 200 {object} bookmodel.ResGetAllBook "list book"
// @Response 400 {object} common.AppError "error"
// @Router /books/all [get]
func GetAllBook(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqData bookmodel.ReqCreateBook
		if err := c.ShouldBind(&reqData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection().Begin()

		bookStore := bookstore.NewSQLStore(db)
		authorStore := authorstore.NewSQLStore(db)
		categoryStore := categorystore.NewSQLStore(db)

		repo := bookrepo.NewGetAllBookRepo(bookStore, categoryStore, authorStore)

		biz := bookbiz.NewGetAllBookBiz(repo)

		books, err := biz.GetAllBook(c.Request.Context())
		if err != nil {
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(books))
	}
}
