package ginsupplier

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/importnote/importnotestore"
	"book-store-management-backend/module/supplier/supplierbiz"
	"book-store-management-backend/module/supplier/suppliermodel/filter"
	"book-store-management-backend/module/supplier/supplierrepo"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary See import notes of supplier
// @Tags suppliers
// @Accept json
// @Produce json
// @Param id path string true "supplier id"
// @Param page query common.Paging false "page"
// @Param filter query filter.SupplierImportFilter false "filter"
// @Response 200 {object} suppliermodel.ResSeeImportNoteSupplier "supplier"
// @Response 400 {object} common.AppError "error"
// @Router /suppliers/{id}/importNotes [get]
func SeeSupplierImportNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var importSupplierFilter filter.SupplierImportFilter
		if err := c.ShouldBind(&importSupplierFilter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		importNoteStore := importnotestore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := supplierrepo.NewSeeSupplierImportNoteRepo(importNoteStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := supplierbiz.NewSeeSupplierImportNoteBiz(repo, requester)

		result, err := biz.SeeSupplierImportNote(
			c.Request.Context(), id, &importSupplierFilter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
