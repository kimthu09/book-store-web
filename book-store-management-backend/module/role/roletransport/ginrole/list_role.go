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
// @Summary List role
// @Tags roles
// @Accept json
// @Produce json
// @Response 200 {object} rolemodel.ResListRole "list role"
// @Response 400 {object} common.AppError "error"
// @Router /roles [get]
func ListRole(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleStore := rolestore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := rolerepo.NewListRoleRepo(roleStore)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := rolebiz.NewListRoleBiz(repo, requester)

		result, err := biz.ListRole(c.Request.Context())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
