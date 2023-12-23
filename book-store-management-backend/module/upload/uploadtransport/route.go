package uploadtransport

import (
	"book-store-management-backend/component/appctx"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	uploadFile := router.Group("/upload")
	uploadFile.POST("", UploadFile(appCtx))
}
