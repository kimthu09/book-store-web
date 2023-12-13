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
		users.GET("", ListUser(appCtx))
		users.POST("", CreateUser(appCtx))
		users.PATCH("/:id/info", UpdateInfoUser(appCtx))
		users.PATCH("/status", ChangeStatusUsers(appCtx))
		users.PATCH("/:id/role", ChangeRoleUser(appCtx))
		users.PATCH("/:id/reset", ResetPassword(appCtx))
		users.PATCH("/:id/password", UpdatePassword(appCtx))
	}
}
