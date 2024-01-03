package booktransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/book/bookbiz"
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/book/bookrepo"
	"book-store-management-backend/module/book/bookstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ChangeStatusBooks
// @BasePath /v1
// @Security BearerAuth
// @Summary Change status books
// @Tags books
// @Accept json
// @Produce json
// @Param book body bookmodel.ReqUpdateStatusBooks true "list book id and status want to be updated"
// @Response 200 {object} common.ResSuccess "status of response"
// @Response 400 {object} common.AppError "error"
// @Router /books/status [patch]
func ChangeStatusBooks(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data bookmodel.ReqUpdateStatusBooks

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := bookstore.NewSQLStore(db)
		repo := bookrepo.NewChangeStatusBookRepo(store)

		biz := bookbiz.NewChangeStatusBooksBiz(repo, requester)
		if err := biz.ChangeStatusBooks(c.Request.Context(), &data); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
