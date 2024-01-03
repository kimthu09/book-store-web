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
		books.GET("/all", GetAllBookTitle(appCtx))
		books.POST("", CreateBookTitle(appCtx))
		books.DELETE("/:id", DeleteBookTitle(appCtx))

		books.GET("/:id", GetBookTitleDetail(appCtx))
		books.PATCH("/:id/info", UpdateBookTitleInfo(appCtx))
	}
}
