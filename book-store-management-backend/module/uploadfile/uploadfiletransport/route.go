package uploadfiletransport

import (
	"book-store-management-backend/component/appctx"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	staticRoute := router.Group("/static")
	staticRoute.Static("/", "./static")

	uploadFile := staticRoute.Group("/upload")
	uploadFile.POST("", UploadFile(appCtx))
}
