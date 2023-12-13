package ginrole

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/role/rolebiz"
	"book-store-management-backend/module/role/rolerepo"
	"book-store-management-backend/module/role/rolestore"
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

		roleStore := rolestore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := rolerepo.NewSeeRoleDetailRepo(roleStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := rolebiz.NewSeeDetailRoleBiz(repo, requester)

		result, err := biz.SeeDetailRole(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
