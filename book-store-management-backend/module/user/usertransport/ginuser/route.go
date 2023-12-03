package ginuser

import (
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	router.POST("/login", Login(appCtx))
	users := router.Group("/users", middleware.RequireAuth(appCtx))
	{
		users.POST("", CreateUser(appCtx))
	}
}
