package ginimportnote

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/importnote/importnotebiz"
	"book-store-management-backend/module/importnote/importnoterepo"
	"book-store-management-backend/module/importnote/importnotestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeDetailImportNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		importNoteId := c.Param("id")

		store := importnotestore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := importnoterepo.NewSeeDetailImportNoteRepo(store)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := importnotebiz.NewSeeDetailImportNoteBiz(repo, requester)

		result, err := biz.SeeDetailImportNote(c.Request.Context(), importNoteId)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
