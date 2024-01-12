package ginshopgeneral

import (
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	shopGenerals := router.Group("/shop", middleware.RequireAuth(appCtx))
	{
		shopGenerals.GET("", SeeShopGeneral(appCtx))
		shopGenerals.PATCH("", UpdateShopGeneral(appCtx))
	}
}
