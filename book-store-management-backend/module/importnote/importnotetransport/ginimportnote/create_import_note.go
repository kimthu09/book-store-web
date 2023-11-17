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
	"book-store-management-backend/module/supplier/supplierstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateImportNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data importnotemodel.ImportNoteCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		data.CreateBy = requester.GetUserId()

		db := appCtx.GetMainDBConnection().Begin()

		importNoteStore := importnotestore.NewSQLStore(db)
		importNoteDetailStore := importnotedetailstore.NewSQLStore(db)
		bookStore := bookstore.NewSQLStore(db)
		supplierStore := supplierstore.NewSQLStore(db)

		repo := importnoterepo.NewCreateImportNoteRepo(
			importNoteStore,
			importNoteDetailStore,
			bookStore,
			supplierStore,
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
