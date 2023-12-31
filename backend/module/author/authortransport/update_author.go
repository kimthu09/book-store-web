package authortransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/author/authorbiz"
	"book-store-management-backend/module/author/authormodel"
	"book-store-management-backend/module/author/authorrepo"
	"book-store-management-backend/module/author/authorstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Update author
// @Tags authors
// @Accept json
// @Produce json
// @Param id path string true "author id"
// @Param author body authormodel.ReqUpdateAuthor true "author info to update"
// @Response 200 {object} common.ResSuccess "status of response"
// @Response 400 {object} common.AppError "error"
// @Router /authors/{id} [patch]
func UpdateAuthor(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data authormodel.ReqUpdateAuthor

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := authorstore.NewSQLStore(db)
		repo := authorrepo.NewUpdateAuthorRepo(store)

		business := authorbiz.NewUpdateAuthorBiz(repo, requester)

		if err := business.UpdateAuthor(c.Request.Context(), id, &data); err != nil {
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
