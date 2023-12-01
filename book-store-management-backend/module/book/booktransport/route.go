package booktransport

import (
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	books := router.Group("/books", middleware.RequireAuth(appCtx))
	{
		books.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "get all books",
			})
		})
		books.POST("", CreateBook(appCtx))
	}
}
