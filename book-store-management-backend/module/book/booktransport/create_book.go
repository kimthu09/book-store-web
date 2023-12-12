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
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateBookTitle
// @BasePath /v1
// @Security BearerAuth
// @Summary Create Book
// @Tags books
// @Accept json
// @Produce json
// @Param booktitle body bookmodel.ReqCreateBook true "Create Book"
// @Response 200 {object} bookmodel.ResCreateBook "book id"
// @Router /books [post]
func CreateBook(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqData bookmodel.ReqCreateBook
		if err := c.ShouldBind(&reqData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		gen := generator.NewShortIdGenerator()
		store := bookstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := bookrepo.NewCreateBookRepo(store)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := bookbiz.NewCreateBookBiz(gen, repo, requester)

		var resData bookmodel.ResCreateBook
		if err := biz.CreateBook(c.Request.Context(), &reqData, &resData); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(resData))
	}
}
