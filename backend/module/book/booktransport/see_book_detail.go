package booktransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/author/authorstore"
	"book-store-management-backend/module/book/bookbiz"
	"book-store-management-backend/module/book/bookrepo"
	"book-store-management-backend/module/book/bookstore"
	"book-store-management-backend/module/category/categorystore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary See detail of book
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "book id"
// @Response 200 {object} bookmodel.ResDetailUnitBook "book"
// @Response 400 {object} common.AppError "error"
// @Router /books/{id} [get]
func SeeBookDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		db := appCtx.GetMainDBConnection()

		bookStore := bookstore.NewSQLStore(db)
		authorStore := authorstore.NewSQLStore(db)
		categoryStore := categorystore.NewSQLStore(db)

		repo := bookrepo.NewSeeBookDetailRepo(bookStore, categoryStore, authorStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := bookbiz.NewSeeBookDetailBiz(repo, requester)

		result, err := biz.SeeBookDetail(
			c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
