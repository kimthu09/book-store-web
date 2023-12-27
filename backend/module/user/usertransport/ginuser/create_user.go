package ginuser

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/component/hasher"
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
// @Summary Create user
// @Tags users
// @Accept json
// @Produce json
// @Param user body usermodel.ReqCreateUser true "user need to create"
// @Response 200 {object} usermodel.ResCreateUser "user id"
// @Response 400 {object} common.AppError "error"
// @Router /users [post]
func CreateUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.ReqCreateUser

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		userStore := userstore.NewSQLStore(db)
		repo := userrepo.NewCreateUserRepo(userStore)

		md5 := hasher.NewMd5Hash()
		gen := generator.NewShortIdGenerator()
		biz := userbiz.NewCreateUserBiz(gen, repo, md5, requester)

		if err := biz.CreateUser(c.Request.Context(), &data); err != nil {
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
