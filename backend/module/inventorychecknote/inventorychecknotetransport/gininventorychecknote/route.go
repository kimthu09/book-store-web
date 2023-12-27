package gininventorychecknote

import (
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	inventoryCheckNotes := router.Group("/inventoryCheckNotes", middleware.RequireAuth(appCtx))
	{
		inventoryCheckNotes.GET("", ListInventoryCheckNote(appCtx))
		inventoryCheckNotes.GET("/:id", SeeDetailInventoryCheckNote(appCtx))
		inventoryCheckNotes.POST("", CreateInventoryCheckNote(appCtx))
	}
}
