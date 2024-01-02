package authortransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/module/author/authorbiz"
	"book-store-management-backend/module/author/authorrepo"
	"book-store-management-backend/module/author/authorstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Get all authors
// @Tags authors
// @Accept json
// @Produce json
// @Response 200 {object} authormodel.ResGetAllAuthor
// @Router /authors/all [get]
func GetAllAuthor(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := authorstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := authorrepo.NewGetAllAuthorRepo(store)

		biz := authorbiz.NewGetAllAuthorBiz(repo)

		result, err := biz.GetAllAuthor(c.Request.Context())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
