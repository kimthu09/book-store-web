package ginimportnote

import (
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	importNotes := router.Group("/importNotes", middleware.RequireAuth(appCtx))
	{
		importNotes.GET("", ListImportNote(appCtx))
		importNotes.GET("/:id", SeeDetailImportNote(appCtx))
		importNotes.POST("", CreateImportNote(appCtx))
		importNotes.PATCH("/:id", ChangeStatusImportNote(appCtx))
	}
}
