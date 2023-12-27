package ginrole

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/generator"
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
// @Summary Create role
// @Tags roles
// @Accept json
// @Produce json
// @Param role body rolemodel.ReqCreateRole true "role need to create"
// @Response 200 {object} rolemodel.ResCreateRole "role id"
// @Response 400 {object} common.AppError "error"
// @Router /roles [post]
func CreateRole(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data rolemodel.ReqCreateRole

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		roleStore := rolestore.NewSQLStore(db)
		roleFeatureStore := rolefeaturestore.NewSQLStore(db)

		repo := rolerepo.NewCreateRoleRepo(
			roleStore,
			roleFeatureStore,
		)

		gen := generator.NewShortIdGenerator()

		business := rolebiz.NewCreateRoleStore(gen, repo, requester)

		if err := business.CreateRole(c.Request.Context(), &data); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
