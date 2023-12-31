package publishertransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/publisher/publisherbiz"
	"book-store-management-backend/module/publisher/publishermodel"
	"book-store-management-backend/module/publisher/publisherrepo"
	"book-store-management-backend/module/publisher/publisherstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Update publisher
// @Tags publishers
// @Accept json
// @Produce json
// @Param id path string true "publisher id"
// @Param publisher body publishermodel.ReqUpdatePublisher true "publisher info to update"
// @Response 200 {object} common.ResSuccess "status of response"
// @Response 400 {object} common.AppError "error"
// @Router /publishers/{id} [patch]
func UpdatePublisher(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data publishermodel.ReqUpdatePublisher

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := publisherstore.NewSQLStore(db)
		repo := publisherrepo.NewUpdatePublisherRepo(store)

		business := publisherbiz.NewUpdatePublisherBiz(repo, requester)

		if err := business.UpdatePublisher(c.Request.Context(), id, &data); err != nil {
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
