package ginimportnote

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/importnote/importnotebiz"
	"book-store-management-backend/module/importnote/importnoterepo"
	"book-store-management-backend/module/importnote/importnotestore"
	"book-store-management-backend/module/importnotedetail/importnotedetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeDetailImportNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		importNoteId := c.Param("id")

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		importNoteDetailStore := importnotedetailstore.NewSQLStore(appCtx.GetMainDBConnection())
		importNoteStore := importnotestore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := importnoterepo.NewSeeDetailImportNoteRepo(
			importNoteStore, importNoteDetailStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := importnotebiz.NewSeeDetailImportNoteBiz(
			repo, requester)

		result, err := biz.SeeDetailImportNote(c.Request.Context(), importNoteId, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, &paging, nil))
	}
}
