package ginrole

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/feature/featurestore"
	"book-store-management-backend/module/role/rolebiz"
	"book-store-management-backend/module/role/rolerepo"
	"book-store-management-backend/module/role/rolestore"
	"book-store-management-backend/module/rolefeature/rolefeaturestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary See detail information of role
// @Tags roles
// @Accept json
// @Produce json
// @Param id path string true "role id"
// @Response 200 {object} rolemodel.ResSeeDetailRole "detailed information of role"
// @Response 400 {object} common.AppError "error"
// @Router /roles/{id} [get]
func SeeDetailRole(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		db := appCtx.GetMainDBConnection()
		roleStore := rolestore.NewSQLStore(db)
		roleFeatureStore := rolefeaturestore.NewSQLStore(db)
		featureStore := featurestore.NewSQLStore(db)

		repo := rolerepo.NewSeeRoleDetailRepo(roleStore, roleFeatureStore, featureStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := rolebiz.NewSeeDetailRoleBiz(repo, requester)

		result, err := biz.SeeDetailRole(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
