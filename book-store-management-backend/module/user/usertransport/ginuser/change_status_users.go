package ginuser

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/user/userbiz"
	"book-store-management-backend/module/user/usermodel"
	"book-store-management-backend/module/user/userrepo"
	"book-store-management-backend/module/user/userstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Change status users
// @Tags users
// @Accept json
// @Produce json
// @Param user body usermodel.ReqUpdateStatusUsers true "list user id and status want to be updated"
// @Response 200 {object} common.ResSuccess "status of response"
// @Response 400 {object} common.AppError "error"
// @Router /users/status [patch]
func ChangeStatusUsers(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.ReqUpdateStatusUsers

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := userstore.NewSQLStore(db)
		repo := userrepo.NewChangeStatusUserRepo(store)

		business := userbiz.NewChangeStatusUserBiz(repo, requester)

		if err := business.ChangeStatusUser(c.Request.Context(), &data); err != nil {
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
