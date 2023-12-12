package booktitletransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/module/booktitle/booktitlerepo"
	"book-store-management-backend/module/booktitle/booktitlestore"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Get all booktitle detail
// @Tags booktitles
// @Accept json
// @Produce json
// @Param page query common.Paging false "page"
// @Response 200 {object} booktitlemodel.ResBookTitleDetail
// @Router /booktitles [get]
func GetBookTitleDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := strings.Trim(c.Param("id"), " ")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "id is invalid",
			})
			return
		}

		store := booktitlestore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := booktitlerepo.NewDetailBookTitleRepo(store)

		data, err := repo.DetailBookTitle(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(data))
	}
}
