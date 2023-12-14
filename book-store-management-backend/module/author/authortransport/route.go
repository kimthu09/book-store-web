package authortransport

import (
	"book-store-management-backend/component/appctx"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	authors := router.Group("/authors")
	{
		authors.GET("", ListAuthor(appCtx))
		authors.POST("", CreateAuthor(appCtx))
	}
}
