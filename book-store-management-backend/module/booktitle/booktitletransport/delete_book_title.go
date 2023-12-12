package booktitletransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/booktitle/booktitlebiz"
	booktitlerepo "book-store-management-backend/module/booktitle/booktitlerepo"
	"book-store-management-backend/module/booktitle/booktitlestore"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Delete booktitle by id
// @Tags booktitles
// @Accept json
// @Produce json
// @Response 200 {object} common.ResSuccess
// @Router /booktitles/:id [delete]
func DeleteBookTitle(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookId := c.Param("id")
		if bookId == "" {
			panic(common.ErrInvalidRequest(nil))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		store := booktitlestore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := booktitlerepo.NewDeleteBookRepo(store)

		biz := booktitlebiz.NewDeleteBookBiz(requester, repo)

		fmt.Println(biz)

		if err := biz.DeleteBook(c.Request.Context(), bookId); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.ResSuccess{IsSuccess: true})
	}
}
