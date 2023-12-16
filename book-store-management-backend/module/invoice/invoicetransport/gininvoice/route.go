package gininvoice

import (
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	invoices := router.Group("/invoices", middleware.RequireAuth(appCtx))
	{
		invoices.GET("", ListInvoice(appCtx))
		invoices.GET("/:id", SeeInvoiceDetail(appCtx))
		invoices.POST("", CreateInvoice(appCtx))
	}
}
