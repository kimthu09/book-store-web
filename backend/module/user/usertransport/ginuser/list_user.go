package ginuser

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/user/userbiz"
	"book-store-management-backend/module/user/usermodel"
	"book-store-management-backend/module/user/userrepo"
	"book-store-management-backend/module/user/userstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary List user
// @Tags users
// @Accept json
// @Produce json
// @Param page query common.Paging false "page"
// @Param filter query usermodel.Filter false "filter"
// @Response 200 {object} usermodel.ResListUser "list user"
// @Response 400 {object} common.AppError "error"
// @Router /users [get]
func ListUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter usermodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := userstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := userrepo.NewListUserRepo(store)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := userbiz.NewListUserBiz(repo, requester)

		result, err := biz.ListUser(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
