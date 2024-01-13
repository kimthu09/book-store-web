package ginshopgeneral

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/module/shopgeneral/shopgeneralbiz"
	"book-store-management-backend/module/shopgeneral/shopgeneralstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary See shop general
// @Tags shop
// @Accept json
// @Produce json
// @Response 200 {object} shopgeneralmodel.ShopGeneral "shop"
// @Response 400 {object} common.AppError "error"
// @Router /shop [post]
func SeeShopGeneral(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		store := shopgeneralstore.NewSQLStore(db)

		business := shopgeneralbiz.NewSeeShopGeneralBiz(
			store,
		)

		general, err := business.SeeShopGeneral(c.Request.Context())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(general))
	}
}
