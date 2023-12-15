package booktransport

import (
	"book-store-management-backend/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListBook(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "list book",
		})
	}
}
