package booktransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/book/bookrepo"
	"book-store-management-backend/module/book/bookstore"
	"github.com/gin-gonic/gin"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Get all books
// @Tags books
// @Accept json
// @Produce json
// @Param page query common.Paging false "page"
// @Param filter query bookmodel.Filter false "filter"
// @Response 200 {object} bookmodel.ResListBook
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

		store := bookstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := bookrepo.NewListBookRepo(store)

		response, err := repo.ListBook(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSuccessResponse(response, paging, filter))
	}
}
