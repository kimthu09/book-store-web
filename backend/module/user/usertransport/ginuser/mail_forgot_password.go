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
// @Summary Mail forgot password
// @Tags auth
// @Accept json
// @Produce json
// @Param email body usermodel.ReqMailForgotPassword true "email"
// @Response 200 {object} common.ResSuccess "status of response"
// @Response 400 {object} common.AppError "error"
// @Router /forgotPassword [post]
func MailForgotPassword(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.ReqMailForgotPassword
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())

		business := userbiz.NewMailForgetPasswordBiz(tokenProvider)
		err := business.MailForgetPassword(
			c.Request.Context(),
			data.Email,
			appCtx.GetEmailTo(),
			appCtx.GetSMTPPass(),
			appCtx.GetSMTPUser(),
			appCtx.GetSMTPHost(),
			appCtx.GetSMTPPort())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
