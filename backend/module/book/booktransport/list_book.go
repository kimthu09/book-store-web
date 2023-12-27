package booktransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/author/authorstore"
	"book-store-management-backend/module/book/bookbiz"
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/book/bookrepo"
	"book-store-management-backend/module/book/bookstore"
	"book-store-management-backend/module/category/categorystore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListBook
// @BasePath /v1
// @Security BearerAuth
// @Summary List book
// @Tags books
// @Accept json
// @Produce json
// @Param page query common.Paging false "page"
// @Param filter query bookmodel.Filter false "filter"
// @Response 200 {object} bookmodel.ResListBook "list book"
// @Response 400 {object} common.AppError "error"
// @Router /books [get]
func ListBook(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter bookmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		db := appCtx.GetMainDBConnection()
		bookStore := bookstore.NewSQLStore(db)
		authorStore := authorstore.NewSQLStore(db)
		categoryStore := categorystore.NewSQLStore(db)

		repo := bookrepo.NewListBookRepo(bookStore, categoryStore, authorStore)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := bookbiz.NewListBookBiz(repo, requester)

		result, err := biz.ListBook(
			c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(
			result, paging, filter))
	}
}
