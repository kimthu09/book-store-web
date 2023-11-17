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
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateInventoryCheckNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data inventorychecknotemodel.InventoryCheckNoteCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		data.CreateBy = requester.GetUserId()

		db := appCtx.GetMainDBConnection().Begin()

		inventoryCheckNoteStore := inventorychecknotestore.NewSQLStore(db)
		inventoryCheckNoteDetailStore := inventorychecknotedetailstore.NewSQLStore(db)
		bookStore := bookstore.NewSQLStore(db)

		repo := inventorychecknoterepo.NewCreateInventoryCheckNoteRepo(
			inventoryCheckNoteStore,
			inventoryCheckNoteDetailStore,
			bookStore,
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
