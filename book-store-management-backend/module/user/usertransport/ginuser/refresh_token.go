package ginuser

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/tokenprovider/jwt"
	"book-store-management-backend/module/user/userbiz"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Response 200 {object} common.ResSuccess "status of response"
// @Response 400 {object} common.AppError "error"
// @Router /refreshToken [post]
func RefreshToken(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken := c.MustGet(common.RefreshTokenStr).(string)

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())

		business := userbiz.NewRefreshTokenBiz(
			appCtx, common.MaxAgeAccessToken, tokenProvider)
		account, err := business.RefreshToken(c.Request.Context(), refreshToken)
		if err != nil {
			panic(err)
		}

		c.SetCookie(
			common.AccessTokenStrInCookie, account.AccessToken.Token, common.MaxAgeAccessToken,
			"/", "", true, true)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
