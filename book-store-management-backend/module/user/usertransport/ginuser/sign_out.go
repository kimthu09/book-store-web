package ginuser

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Summary SignOut
// @Tags auth
// @Accept json
// @Produce json
// @Response 200 {object} common.ResSuccess "status of response"
// @Response 400 {object} common.AppError "error"
// @Router /signOut [post]
func SignOut(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie(
			common.AccessTokenStrInCookie, "", -1,
			"/", appCtx.GetDomain(), true, true)
		c.SetCookie(
			common.RefreshTokenStrInCookie, "", -1,
			"/", appCtx.GetDomain(), true, true)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
