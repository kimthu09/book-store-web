package publishertransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/publisher/publisherbiz"
	"book-store-management-backend/module/publisher/publishermodel"
	"book-store-management-backend/module/publisher/publisherrepo"
	"book-store-management-backend/module/publisher/publisherstore"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary List publishers
// @Tags publishers
// @Accept json
// @Produce json
// @Response 200 {object} publishermodel.ResListPublisher
// @Router /publishers [get]
func ListPublisher(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter publishermodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := publisherstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := publisherrepo.NewListPublisherRepo(store)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := publisherbiz.NewListPublisherRepo(repo, requester)

		result, err := biz.ListPublisher(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
