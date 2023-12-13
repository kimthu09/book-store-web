package ginrole

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/feature/featurestore"
	"book-store-management-backend/module/role/rolebiz"
	"book-store-management-backend/module/role/rolerepo"
	"book-store-management-backend/module/rolefeature/rolefeaturestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary List feature by role
// @Tags roles
// @Accept json
// @Produce json
// @Param id path string true "role id"
// @Response 200 {object} rolemodel.ResListFeatureByRole "list feature by role"
// @Response 400 {object} common.AppError "error"
// @Router /roles/{id}/features [get]
func ListFeatureByRole(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		featureStore := featurestore.NewSQLStore(appCtx.GetMainDBConnection())
		roleFeatureStore := rolefeaturestore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := rolerepo.NewListFeatureByRoleRepo(roleFeatureStore, featureStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := rolebiz.NewListFeatureByRoleBiz(repo, requester)

		result, err := biz.ListFeatureByRole(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
