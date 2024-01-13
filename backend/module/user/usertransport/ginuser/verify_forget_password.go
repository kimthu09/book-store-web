package ginuser

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/hasher"
	"book-store-management-backend/component/tokenprovider/jwt"
	"book-store-management-backend/module/user/userbiz"
	"book-store-management-backend/module/user/usermodel"
	"book-store-management-backend/module/user/userrepo"
	"book-store-management-backend/module/user/userstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Verify forget password
// @Tags auth
// @Accept json
// @Produce json
// @Param token path string true "token"
// @Param user body usermodel.ReqForgetPassword true "new password"
// @Response 200 {object} common.ResSuccess "status of response"
// @Response 400 {object} common.AppError "error"
// @Router /forgetPassword/{token} [post]
func VerifyForgetPassword(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Param("token")

		var data usermodel.ReqForgetPassword
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection().Begin()

		store := userstore.NewSQLStore(db)
		repo := userrepo.NewVerifyForgetPasswordRepo(store)

		md5 := hasher.NewMd5Hash()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())
		business := userbiz.NewVerifyForgetPasswordBiz(repo, md5, tokenProvider)

		if err := business.VerifyForgetPassword(c.Request.Context(), token, &data); err != nil {
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
