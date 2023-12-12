package booktitletransport

import (
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	books := router.Group("/booktitles", middleware.RequireAuth(appCtx))
	{
		books.GET("", ListBookTitle(appCtx))
		books.POST("", CreateBookTitle(appCtx))
		books.PATCH("/:id", UpdateBookTitleInfo(appCtx))
		books.DELETE("/:id", DeleteBookTitle(appCtx))
	}
}
