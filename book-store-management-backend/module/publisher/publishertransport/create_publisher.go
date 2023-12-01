package publishertransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/generator"
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
// @Summary Create publisher with name
// @Tags publishers
// @Accept json
// @Produce json
// @Param publisher body publishermodel.ReqCreatePublisher true "Create publisher"
// @Response 200 {object} publishermodel.ResCreatePublisher "publisher id"
// @Router /publishers [post]
func CreatePublisher(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data publishermodel.ReqCreatePublisher
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		db := appCtx.GetMainDBConnection().Begin()
		store := publisherstore.NewSQLStore(db)
		repo := publisherrepo.NewCreatePublisherRepo(store)

		gen := generator.NewShortIdGenerator()

		business := publisherbiz.NewCreatePublisherBiz(gen, repo, requester)

		tmpData := publishermodel.Publisher{
			Name: data.Name,
		}
		if err := business.CreatePublisher(c.Request.Context(), &tmpData); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, publishermodel.ResCreatePublisher{Id: tmpData.Id})
	}
}
