package ginimportnote

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/book/bookstore"
	"book-store-management-backend/module/importnote/importnotebiz"
	"book-store-management-backend/module/importnote/importnotemodel"
	"book-store-management-backend/module/importnote/importnoterepo"
	"book-store-management-backend/module/importnote/importnotestore"
	"book-store-management-backend/module/importnotedetail/importnotedetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Create import note
// @Tags importNotes
// @Accept json
// @Produce json
// @Param importNote body importnotemodel.ReqCreateImportNote true "import note need to create"
// @Response 200 {object} importnotemodel.ResCreateImportNote "import note id"
// @Response 400 {object} common.AppError "error"
// @Router /importNotes [post]
func CreateImportNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data importnotemodel.ReqCreateImportNote

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		data.CreatedBy = requester.GetUserId()

		db := appCtx.GetMainDBConnection().Begin()

		importNoteStore := importnotestore.NewSQLStore(db)
		importNoteDetailStore := importnotedetailstore.NewSQLStore(db)
		bookStore := bookstore.NewSQLStore(db)

		repo := importnoterepo.NewCreateImportNoteRepo(
			importNoteStore,
			importNoteDetailStore,
			bookStore,
		)

		gen := generator.NewShortIdGenerator()

		business := importnotebiz.NewCreateImportNoteBiz(gen, repo, requester)

		if err := business.CreateImportNote(c.Request.Context(), &data); err != nil {
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
