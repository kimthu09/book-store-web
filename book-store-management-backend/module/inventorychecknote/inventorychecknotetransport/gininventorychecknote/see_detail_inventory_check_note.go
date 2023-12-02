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

func SeeDetailInventoryCheckNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		inventoryCheckNoteId := c.Param("id")

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		inventoryCheckNoteStore :=
			inventorychecknotestore.NewSQLStore(appCtx.GetMainDBConnection())
		inventoryCheckNoteDetailStore :=
			inventorychecknotedetailstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := inventorychecknoterepo.NewSeeDetailInventoryCheckNoteRepo(
			inventoryCheckNoteStore, inventoryCheckNoteDetailStore)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := inventorychecknotebiz.NewSeeDetailImportNoteBiz(repo, requester)

		result, err := biz.SeeDetailInventoryCheckNote(
			c.Request.Context(), inventoryCheckNoteId, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, &paging, nil))
	}
}
