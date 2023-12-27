package gininventorychecknote

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/inventorychecknote/inventorychecknotebiz"
	"book-store-management-backend/module/inventorychecknote/inventorychecknoterepo"
	"book-store-management-backend/module/inventorychecknote/inventorychecknotestore"
	"book-store-management-backend/module/inventorychecknotedetail/inventorychecknotedetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary See detail inventory check note
// @Tags inventoryCheckNotes
// @Accept json
// @Produce json
// @Param id path string true "inventory check note id"
// @Response 200 {object} inventorychecknotemodel.ResSeeDetailInventoryCheckNote "inventory check note"
// @Response 400 {object} common.AppError "error"
// @Router /inventoryCheckNotes/{id} [get]
func SeeDetailInventoryCheckNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		inventoryCheckNoteId := c.Param("id")

		inventoryCheckNoteStore :=
			inventorychecknotestore.NewSQLStore(appCtx.GetMainDBConnection())
		inventoryCheckNoteDetailStore :=
			inventorychecknotedetailstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := inventorychecknoterepo.NewSeeDetailInventoryCheckNoteRepo(
			inventoryCheckNoteStore, inventoryCheckNoteDetailStore)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := inventorychecknotebiz.NewSeeDetailImportNoteBiz(repo, requester)

		result, err := biz.SeeDetailInventoryCheckNote(
			c.Request.Context(), inventoryCheckNoteId)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
