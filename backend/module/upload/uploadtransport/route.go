package uploadtransport

import (
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	staticRouter := router.Group("/static")
	staticRouter.Static("", appCtx.GetStaticPath())

	uploadFile := router.Group("/upload", middleware.RequireAuth(appCtx))
	uploadFile.POST("", UploadFile(appCtx, staticRouter.BasePath()))
}
