package ginuser

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/user/userbiz"
	"book-store-management-backend/module/user/userrepo"
	"book-store-management-backend/module/user/userstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary See profile
// @Tags auth
// @Accept json
// @Produce json
// @Response 200 {object} usermodel.User "user"
// @Response 400 {object} common.AppError "error"
// @Router /profile [get]
func SeeProfile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		userStore := userstore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := userrepo.NewSeeUserDetailRepo(userStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := userbiz.NewSeeProfileBiz(repo, requester)

		result, err := biz.SeeProfile(c.Request.Context())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
