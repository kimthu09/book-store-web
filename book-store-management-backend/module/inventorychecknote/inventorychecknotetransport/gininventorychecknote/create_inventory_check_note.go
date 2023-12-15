package gininventorychecknote

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/book/bookstore"
	"book-store-management-backend/module/inventorychecknote/inventorychecknotebiz"
	"book-store-management-backend/module/inventorychecknote/inventorychecknotemodel"
	"book-store-management-backend/module/inventorychecknote/inventorychecknoterepo"
	"book-store-management-backend/module/inventorychecknote/inventorychecknotestore"
	"book-store-management-backend/module/inventorychecknotedetail/inventorychecknotedetailstore"
	"book-store-management-backend/module/stockchangehistory/stockchangehistorystore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Create inventory check note
// @Tags inventoryCheckNotes
// @Accept json
// @Produce json
// @Param inventoryCheckNote body inventorychecknotemodel.ReqCreateInventoryCheckNote true "inventory check note need to create"
// @Response 200 {object} inventorychecknotemodel.ResCreateInventoryCheckNote "inventory check note id"
// @Response 400 {object} common.AppError "error"
// @Router /inventoryCheckNotes [post]
func CreateInventoryCheckNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data inventorychecknotemodel.ReqCreateInventoryCheckNote

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		data.CreatedBy = requester.GetUserId()

		db := appCtx.GetMainDBConnection().Begin()

		inventoryCheckNoteStore := inventorychecknotestore.NewSQLStore(db)
		inventoryCheckNoteDetailStore := inventorychecknotedetailstore.NewSQLStore(db)
		bookStore := bookstore.NewSQLStore(db)
		stockChangeHistoryStore := stockchangehistorystore.NewSQLStore(db)

		repo := inventorychecknoterepo.NewCreateInventoryCheckNoteRepo(
			inventoryCheckNoteStore,
			inventoryCheckNoteDetailStore,
			bookStore,
			stockChangeHistoryStore,
		)

		gen := generator.NewShortIdGenerator()

		business := inventorychecknotebiz.NewCreateInventoryCheckNoteBiz(gen, repo, requester)

		if err := business.CreateInventoryCheckNote(c.Request.Context(), &data); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
