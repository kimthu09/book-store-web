package ginuser

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/tokenprovider/jwt"
	"book-store-management-backend/module/user/userbiz"
	"book-store-management-backend/module/user/usermodel"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param refreshToken body usermodel.ReqRefreshToken true "refreshToken"
// @Response 200 {object} usermodel.AccountWithoutRefresh "user token"
// @Response 400 {object} common.AppError "error"
// @Router /refreshToken [post]
func RefreshToken(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.ReqRefreshToken
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())

		business := userbiz.NewRefreshTokenBiz(60*60*24*15, tokenProvider)
		account, err := business.RefreshToken(c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
