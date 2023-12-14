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

// @BasePath /v1
// @Security BearerAuth
// @Summary See detail import note
// @Tags importNotes
// @Accept json
// @Produce json
// @Param id path string true "import note id"
// @Response 200 {object} importnotemodel.ResSeeDetailImportNote "import note"
// @Response 400 {object} common.AppError "error"
// @Router /importNotes/{id} [get]
func SeeDetailImportNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		importNoteId := c.Param("id")

		importNoteDetailStore := importnotedetailstore.NewSQLStore(appCtx.GetMainDBConnection())
		importNoteStore := importnotestore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := importnoterepo.NewSeeDetailImportNoteRepo(
			importNoteStore, importNoteDetailStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := importnotebiz.NewSeeDetailImportNoteBiz(
			repo, requester)

		result, err := biz.SeeDetailImportNote(c.Request.Context(), importNoteId)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
