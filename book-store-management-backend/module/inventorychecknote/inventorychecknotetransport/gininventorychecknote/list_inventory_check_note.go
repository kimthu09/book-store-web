package gininventorychecknote

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/inventorychecknote/inventorychecknotebiz"
	"book-store-management-backend/module/inventorychecknote/inventorychecknotemodel"
	"book-store-management-backend/module/inventorychecknote/inventorychecknoterepo"
	"book-store-management-backend/module/inventorychecknote/inventorychecknotestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary List inventory check note
// @Tags inventoryCheckNotes
// @Accept json
// @Produce json
// @Param page query common.Paging false "page"
// @Param filter query inventorychecknotemodel.Filter false "filter"
// @Response 200 {object} inventorychecknotemodel.ResListInventoryCheckNote "list inventory check note"
// @Response 400 {object} common.AppError "error"
// @Router /inventoryCheckNotes [get]
func ListInventoryCheckNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter inventorychecknotemodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := inventorychecknotestore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := inventorychecknoterepo.NewListInventoryCheckNoteRepo(store)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := inventorychecknotebiz.NewListInventoryCheckNoteBiz(repo, requester)

		result, err := biz.ListInventoryCheckNote(
			c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
