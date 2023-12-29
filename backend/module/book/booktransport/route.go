package booktransport

import (
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	books := router.Group("/books", middleware.RequireAuth(appCtx))
	{
		books.GET("", ListBook(appCtx))
		books.POST("", CreateBook(appCtx))
		books.GET("/all", GetAllBook(appCtx))
		books.PATCH("/:id/info", UpdateBookInfo(appCtx))
		//books.DELETE("/:id", DeleteBookTitle(appCtx))
	}
}
