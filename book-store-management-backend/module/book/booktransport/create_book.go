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
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data bookmodel.ReqCreateBook
		//c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		fmt.Println("begin log")
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		fmt.Println("end log")

		db := appCtx.GetMainDBConnection().Begin()

		store := bookstore.NewSQLStore(db)
		repo := bookrepo.NewCreateBookRepo(store)

		gen := generator.NewShortIdGenerator()

		business := bookbiz.NewCreateBookBiz(gen, repo, requester)

		if err := business.CreateBook(c.Request.Context(), &data); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
