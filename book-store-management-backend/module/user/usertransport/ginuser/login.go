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
// @Summary Login
// @Tags auth
// @Accept json
// @Produce json
// @Param user body usermodel.ReqLoginUser true "login information"
// @Response 200 {object} common.ResSuccess "status of response"
// @Response 400 {object} common.AppError "error"
// @Router /login [post]
func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.ReqLoginUser

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection().Begin()

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())

		store := userstore.NewSQLStore(db)
		repo := userrepo.NewLoginRepo(store)

		md5 := hasher.NewMd5Hash()

		business := userbiz.NewLoginBiz(
			appCtx, repo,
			common.MaxAgeAccessToken, common.MaxAgeRefreshToken,
			tokenProvider, md5)
		account, err := business.Login(c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.SetCookie(
			common.AccessTokenStrInCookie, account.AccessToken.Token, common.MaxAgeAccessToken,
			"/", "", true, true)
		c.SetCookie(
			common.RefreshTokenStrInCookie, account.RefreshToken.Token, common.MaxAgeRefreshToken,
			"/", "", true, true)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
