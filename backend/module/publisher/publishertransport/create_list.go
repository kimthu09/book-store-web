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
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Create list publisher
// @Tags publishers
// @Accept json
// @Produce json
// @Param publisher body publishermodel.ReqCreateListPublisher true "list name of publisher"
// @Response 200 {object} common.ResSuccess "status of response"
// @Response 400 {object} common.AppError "error"
// @Router /publishers/many [post]
func CreateListPublisher(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data publishermodel.ReqCreateListPublisher
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		db := appCtx.GetMainDBConnection().Begin()
		store := publisherstore.NewSQLStore(db)
		repo := publisherrepo.NewCreatePublisherRepo(store)

		gen := generator.NewShortIdGenerator()

		business := publisherbiz.NewCreateListPublisherBiz(gen, repo, requester)

		var tmpData []publishermodel.Publisher
		for _, v := range data.Names {
			tmpData = append(tmpData, publishermodel.Publisher{Name: v})
		}

		if err := business.CreateListPublisher(c.Request.Context(), tmpData); err != nil {
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
