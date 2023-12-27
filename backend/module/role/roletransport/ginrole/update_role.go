package ginrole

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/role/rolebiz"
	"book-store-management-backend/module/role/rolemodel"
	"book-store-management-backend/module/role/rolerepo"
	"book-store-management-backend/module/role/rolestore"
	"book-store-management-backend/module/rolefeature/rolefeaturestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Update info role
// @Tags roles
// @Accept json
// @Produce json
// @Param id path string true "role id"
// @Param role body rolemodel.ReqUpdateRole true "role info to update"
// @Response 200 {object} common.ResSuccess "status of response"
// @Response 400 {object} common.AppError "error"
// @Router /roles/{id} [patch]
func UpdateRole(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data rolemodel.ReqUpdateRole

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		roleStore := rolestore.NewSQLStore(db)
		roleFeatureStore := rolefeaturestore.NewSQLStore(db)

		repo := rolerepo.NewUpdateRoleRepo(
			roleStore,
			roleFeatureStore,
		)

		business := rolebiz.NewUpdateRoleBiz(repo, requester)

		if err := business.UpdateRole(c.Request.Context(), id, &data); err != nil {
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
