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
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Create book name, desc, authors, categories, publisher, .etc
// @Tags books
// @Accept json
// @Produce json
// @Param book body bookmodel.ReqCreateBook true "Create book"
// @Response 200 {object} bookmodel.ResCreateBook "book id"
// @Router /books [post]
func CreateBook(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data bookmodel.ReqCreateBook
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := bookstore.NewSQLStore(db)
		repo := bookrepo.NewCreateBookRepo(store)

		gen := generator.NewShortIdGenerator()

		fmt.Println(requester, repo)
		business := bookbiz.NewCreateBookBiz(gen, repo, requester)

		fmt.Println(business)
		//fmt.Print(data)
		//if err := business.CreateBook(c.Request.Context(), &data); err != nil {
		//	db.Rollback()
		//	panic(err)
		//}
		//
		//if err := db.Commit().Error; err != nil {
		//	db.Rollback()
		//	panic(err)
		//}
		//
		//c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.ID))
	}
}
