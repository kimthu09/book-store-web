package authortransport

import (
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	authors := router.Group("/authors", middleware.RequireAuth(appCtx))
	{
		authors.GET("", ListAuthor(appCtx))
		authors.GET("/all", GetAllAuthor(appCtx))
		authors.POST("", CreateAuthor(appCtx))
		authors.POST("/many", CreateListAuthor(appCtx))
		authors.PATCH("/:id", UpdateAuthor(appCtx))
	}
}
