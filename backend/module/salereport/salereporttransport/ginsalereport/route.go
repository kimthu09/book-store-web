package ginsalereport

import (
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	stock := router.Group("/sale", middleware.RequireAuth(appCtx))
	{
		stock.POST("", FindSaleReport(appCtx))
	}
}
