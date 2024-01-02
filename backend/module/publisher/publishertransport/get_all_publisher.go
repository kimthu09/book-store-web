package publishertransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/module/publisher/publisherbiz"
	"book-store-management-backend/module/publisher/publisherrepo"
	"book-store-management-backend/module/publisher/publisherstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Get all publishers
// @Tags publishers
// @Accept json
// @Produce json
// @Response 200 {object} publishermodel.ResGetAllPublisher
// @Router /publishers/all [get]
func GetAllPublisher(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := publisherstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := publisherrepo.NewGetAllPublisherRepo(store)

		biz := publisherbiz.NewGetAllPublisherBiz(repo)

		result, err := biz.GetAllPublisher(c.Request.Context())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
