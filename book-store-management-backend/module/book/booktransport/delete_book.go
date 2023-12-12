package booktransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/book/bookbiz"
	"book-store-management-backend/module/book/bookrepo"
	"book-store-management-backend/module/book/bookstore"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Delete book by id
// @Tags books
// @Accept json
// @Produce json
// @Response 200 {object} common.ResSuccess
// @Router /books/:id [delete]
func DeleteBook(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookId := c.Param("bookId")
		if bookId == "" {
			panic(common.ErrInvalidRequest(nil))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		store := bookstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := bookrepo.NewDeleteBookRepo(store)

		biz := bookbiz.NewDeleteBookBiz(requester, repo)

		fmt.Println(biz)

		if err := biz.DeleteBook(c.Request.Context(), bookId); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.ResSuccess{IsSuccess: true})
	}
}
