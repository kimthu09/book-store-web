package booktitletransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/module/author/authorstore"
	"book-store-management-backend/module/booktitle/booktitlebiz"
	"book-store-management-backend/module/booktitle/booktitlerepo"
	"book-store-management-backend/module/booktitle/booktitlestore"
	"book-store-management-backend/module/category/categorystore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllBookTitle
// @BasePath /v1
// @Security BearerAuth
// @Summary Get all book title
// @Tags booktitles
// @Accept json
// @Produce json
// @Response 200 {object} booktitlemodel.ResGetAllBookTitle "list book title"
// @Response 400 {object} common.AppError "error"
// @Router /booktitles/all [get]
func GetAllBookTitle(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection().Begin()

		bookTitleStore := booktitlestore.NewSQLStore(db)
		authorStore := authorstore.NewSQLStore(db)
		categoryStore := categorystore.NewSQLStore(db)

		repo := booktitlerepo.NewGetAllBookTitleRepo(bookTitleStore, categoryStore, authorStore)

		biz := booktitlebiz.NewGetAllBookTitleBiz(repo)

		books, err := biz.GetAllBookTitle(c.Request.Context())
		if err != nil {
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(books))
	}
}
