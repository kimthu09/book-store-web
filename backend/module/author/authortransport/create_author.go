package authortransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/author/authorbiz"
	"book-store-management-backend/module/author/authormodel"
	"book-store-management-backend/module/author/authorrepo"
	"book-store-management-backend/module/author/authorstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Create author with name
// @Tags authors
// @Accept json
// @Produce json
// @Param author body authormodel.ReqCreateAuthor true "Create author"
// @Response 200 {object} authormodel.ResCreateAuthor "author id"
// @Router /authors [post]
func CreateAuthor(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data authormodel.ReqCreateAuthor
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		db := appCtx.GetMainDBConnection().Begin()
		store := authorstore.NewSQLStore(db)
		repo := authorrepo.NewCreateAuthorRepo(store)

		gen := generator.NewShortIdGenerator()

		business := authorbiz.NewCreateAuthorBiz(gen, repo, requester)

		tmpData := authormodel.Author{
			Name: data.Name,
		}
		if err := business.CreateAuthor(c.Request.Context(), &tmpData); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, authormodel.ResCreateAuthor{Id: tmpData.Id})
	}
}
