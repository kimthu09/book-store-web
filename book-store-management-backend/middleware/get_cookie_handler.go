package middleware

import (
	"book-store-management-backend/common"
	"github.com/gin-gonic/gin"
)

func GetCookieHandler(cookiePath string, valuePath string, errReturn error) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		value, err := c.Cookie(cookiePath)
		if err != nil {
			panic(errReturn)
		}
		c.Set(valuePath, value)
		c.Next()
	}
}

func GetAccessTokenCookieHandler() func(ctx *gin.Context) {
	return GetCookieHandler(
		common.AccessTokenStrInCookie, common.AccessTokenStr, common.ErrTokenExpired)
}

func GetRefreshTokenCookieHandler() func(ctx *gin.Context) {
	return GetCookieHandler(
		common.RefreshTokenStrInCookie, common.RefreshTokenStr, common.ErrTokenExpired)
}
