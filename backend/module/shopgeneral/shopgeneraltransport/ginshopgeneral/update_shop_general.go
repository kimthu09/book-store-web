package ginshopgeneral

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/shopgeneral/shopgeneralbiz"
	"book-store-management-backend/module/shopgeneral/shopgeneralmodel"
	"book-store-management-backend/module/shopgeneral/shopgeneralstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Update shop general
// @Tags shop
// @Accept json
// @Produce json
// @Param shop body shopgeneralmodel.ReqUpdateShopGeneral true "shop info to update"
// @Response 200 {object} common.ResSuccess "status of response"
// @Response 400 {object} common.AppError "error"
// @Router /shop [patch]
func UpdateShopGeneral(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data shopgeneralmodel.ReqUpdateShopGeneral

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := shopgeneralstore.NewSQLStore(db)

		business := shopgeneralbiz.NewUpdateGeneralShopBiz(
			store,
			requester,
		)

		if err := business.UpdateGeneralShop(c.Request.Context(), &data); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
