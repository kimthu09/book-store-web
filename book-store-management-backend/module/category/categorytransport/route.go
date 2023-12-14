package categorytransport

import (
	"book-store-management-backend/component/appctx"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	categories := router.Group("/categories")
	{
		categories.GET("", ListCategory(appCtx))
		categories.POST("", CreateCategory(appCtx))
	}
}
