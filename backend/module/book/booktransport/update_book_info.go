package booktransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/author/authorrepo"
	"book-store-management-backend/module/author/authorstore"
	"book-store-management-backend/module/booktitle/booktitlebiz"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"book-store-management-backend/module/booktitle/booktitlerepo"
	"book-store-management-backend/module/booktitle/booktitlestore"
	"book-store-management-backend/module/category/categoryrepo"
	"book-store-management-backend/module/category/categorystore"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func UpdateBookInfo(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Param("id"))
		id := strings.Trim(c.Param("id"), " ")
		var reqData booktitlemodel.ReqUpdateBookInfo

		if err := c.ShouldBind(&reqData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if id == "" {
			panic(common.ErrInvalidRequest(fmt.Errorf("id is empty")))
			return
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := booktitlestore.NewSQLStore(db)
		authorStore := authorstore.NewSQLStore(db)
		categoryStore := categorystore.NewSQLStore(db)

		repo := booktitlerepo.NewUpdateBookRepo(store)
		authorRepo := authorrepo.NewAuthorPublicRepo(authorStore)
		categoryRepo := categoryrepo.NewCategoryPublicRepo(categoryStore)

		biz := booktitlebiz.NewUpdateBookBiz(repo, authorRepo, categoryRepo, requester)

		err := biz.UpdateBookTitle(c.Request.Context(), id, &reqData)
		if err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.ResSuccess{IsSuccess: true})
	}
}
