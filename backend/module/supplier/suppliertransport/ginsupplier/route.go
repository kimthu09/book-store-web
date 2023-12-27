package ginsupplier

import (
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	suppliers := router.Group("/suppliers", middleware.RequireAuth(appCtx))
	{
		suppliers.GET("", ListSupplier(appCtx))
		suppliers.GET("/all", GetAllSupplier(appCtx))
		suppliers.POST("", CreateSupplier(appCtx))
		suppliers.GET("/:id", SeeSupplierDetail(appCtx))
		suppliers.GET("/:id/importNotes", SeeSupplierImportNote(appCtx))
		suppliers.GET("/:id/debts", SeeSupplierDebt(appCtx))
		suppliers.PATCH("/:id", UpdateInfoSupplier(appCtx))
		suppliers.POST("/:id/pay", PaySupplier(appCtx))
	}
}
