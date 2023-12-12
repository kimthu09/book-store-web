package booktransport

import (
	"book-store-management-backend/component/appctx"
	"github.com/gin-gonic/gin"
)

func ListBook(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "list book",
		})
	}
}
