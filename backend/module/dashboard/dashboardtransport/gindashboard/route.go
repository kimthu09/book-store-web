package gindashboard

import (
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	dashboard := router.Group("/dashboard", middleware.RequireAuth(appCtx))
	{
		dashboard.POST("", SeeDashboard(appCtx))
	}
}
