package booktransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/book/bookrepo"
	"book-store-management-backend/module/book/bookstore"
	"github.com/gin-gonic/gin"
)

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

		store := bookstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := bookrepo.NewListBookRepo(store)

		response, err := repo.ListBook(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSuccessResponse(response, paging, filter))
		//requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		//biz := bookbiz.NewListBookRepo(repo, requester)

		//result, err := biz.ListAuthor(c.Request.Context(), &filter, &paging)

		//if err != nil {
		//	panic(err)
		//}
		//
		//c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
