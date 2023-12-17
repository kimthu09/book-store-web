package ginuser

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/module/user/userbiz"
	"book-store-management-backend/module/user/userrepo"
	"book-store-management-backend/module/user/userstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Get all user
// @Tags users
// @Accept json
// @Produce json
// @Response 200 {object} usermodel.ResGetAllUser "list user"
// @Response 400 {object} common.AppError "error"
// @Router /users/all [get]
func GetAllUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := userstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := userrepo.NewGetAllUserRepo(store)

		biz := userbiz.NewGetAllUserBiz(repo)

		result, err := biz.GetAllUser(c.Request.Context())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
