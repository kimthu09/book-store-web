package gincustomer

import (
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	customers := router.Group("/customers", middleware.RequireAuth(appCtx))
	{
		customers.GET("", ListCustomer(appCtx))
		customers.GET("/all", GetAllCustomer(appCtx))
		customers.POST("", CreateCustomer(appCtx))
		customers.GET("/:id", SeeCustomerDetail(appCtx))
		customers.GET("/:id/invoices", SeeCustomerInvoice(appCtx))
		customers.PATCH("/:id", UpdateInfoCustomer(appCtx))
	}
}
