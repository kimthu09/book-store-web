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
// @Summary Create list author
// @Tags authors
// @Accept json
// @Produce json
// @Param author body authormodel.ReqCreateListAuthor true "list name of author"
// @Response 200 {object} common.ResSuccess "status of response"
// @Response 400 {object} common.AppError "error"
// @Router /authors/many [post]
func CreateListAuthor(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data authormodel.ReqCreateListAuthor
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		db := appCtx.GetMainDBConnection().Begin()
		store := authorstore.NewSQLStore(db)
		repo := authorrepo.NewCreateAuthorRepo(store)

		gen := generator.NewShortIdGenerator()

		business := authorbiz.NewCreateListAuthorBiz(gen, repo, requester)

		var tmpData []authormodel.Author
		for _, v := range data.Names {
			tmpData = append(tmpData, authormodel.Author{Name: v})
		}

		if err := business.CreateListAuthor(c.Request.Context(), tmpData); err != nil {
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
