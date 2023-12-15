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
	"book-store-management-backend/module/stockchangehistory/stockchangehistorystore"
	"book-store-management-backend/module/supplier/supplierstore"
	"book-store-management-backend/module/supplierdebt/supplierdebtstore"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Change status import note
// @Tags importNotes
// @Accept json
// @Produce json
// @Param id path string true "import note id"
// @Param importNote body importnotemodel.ReqUpdateImportNote true "status need to update of import note"
// @Response 200 {object} common.ResSuccess "status of response"
// @Response 400 {object} common.AppError "error"
// @Router /importNotes/{id} [patch]
func ChangeStatusImportNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		idImportNote := c.Param("id")
		if idImportNote == "" {
			panic(common.ErrInvalidRequest(errors.New("param id not exist")))
		}

		var data importnotemodel.ReqUpdateImportNote

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		importNoteStore := importnotestore.NewSQLStore(db)
		importNoteDetailStore := importnotedetailstore.NewSQLStore(db)
		bookStore := bookstore.NewSQLStore(db)
		supplierStore := supplierstore.NewSQLStore(db)
		supplierDebtStore := supplierdebtstore.NewSQLStore(db)
		stockChangeHistoryStore := stockchangehistorystore.NewSQLStore(db)

		repo := importnoterepo.NewChangeStatusImportNoteRepo(
			importNoteStore,
			importNoteDetailStore,
			bookStore,
			supplierStore,
			supplierDebtStore,
			stockChangeHistoryStore,
		)

		gen := generator.NewShortIdGenerator()

		business := importnotebiz.NewChangeStatusImportNoteBiz(gen, repo, requester)

		if err := business.ChangeStatusImportNote(
			c.Request.Context(),
			idImportNote,
			&data,
		); err != nil {
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
