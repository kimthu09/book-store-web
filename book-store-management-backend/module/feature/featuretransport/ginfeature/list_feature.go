package ginfeature

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/feature/featurebiz"
	"book-store-management-backend/module/feature/featurestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary List feature
// @Tags features
// @Accept json
// @Produce json
// @Response 200 {object} featuremodel.ResListFeature "list feature"
// @Response 400 {object} common.AppError "error"
// @Router /features [get]
func ListFeature(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := featurestore.NewSQLStore(appCtx.GetMainDBConnection())

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := featurebiz.NewListFeatureBiz(store, requester)

		result, err := biz.ListFeature(c.Request.Context())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, nil, nil))
	}
}
