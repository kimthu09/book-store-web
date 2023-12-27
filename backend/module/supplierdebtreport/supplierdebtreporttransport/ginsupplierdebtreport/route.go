package ginsupplierdebtreport

import (
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	stock := router.Group("/debt", middleware.RequireAuth(appCtx))
	{
		stock.POST("", FindSupplierDebtReport(appCtx))
	}
}
