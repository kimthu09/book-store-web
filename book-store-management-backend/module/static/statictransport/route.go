package statictransport

import (
	"book-store-management-backend/component/appctx"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	sr := router.Group("/static")
	{
		sr.Static("/images", "./images")
	}
}
