package publishertransport

import (
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	publishers := router.Group("/publishers", middleware.RequireAuth(appCtx))
	{
		publishers.GET("", ListPublisher(appCtx))
		publishers.GET("/all", GetAllPublisher(appCtx))
		publishers.POST("", CreatePublisher(appCtx))
		publishers.POST("/many", CreateListPublisher(appCtx))
		publishers.PATCH("/:id", UpdatePublisher(appCtx))
	}
}
