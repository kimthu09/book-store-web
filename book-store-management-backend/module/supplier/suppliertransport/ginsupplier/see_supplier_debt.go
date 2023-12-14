package ginsupplier

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/supplier/supplierbiz"
	"book-store-management-backend/module/supplier/suppliermodel/filter"
	"book-store-management-backend/module/supplier/supplierrepo"
	"book-store-management-backend/module/supplierdebt/supplierdebtstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary See debts of supplier
// @Tags suppliers
// @Accept json
// @Produce json
// @Param id path string true "supplier id"
// @Param page query common.Paging false "page"
// @Param filter query filter.SupplierDebtFilter false "filter"
// @Response 200 {object} suppliermodel.ResSeeDebtSupplier "supplier"
// @Response 400 {object} common.AppError "error"
// @Router /suppliers/{id}/debts [get]
func SeeSupplierDebt(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var debtSupplierFilter filter.SupplierDebtFilter
		if err := c.ShouldBind(&debtSupplierFilter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		supplierDebtStore := supplierdebtstore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := supplierrepo.NewSeeSupplierDebtRepo(supplierDebtStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := supplierbiz.NewSeeSupplierDebtBiz(repo, requester)

		result, err := biz.SeeSupplierDebt(
			c.Request.Context(), id, &debtSupplierFilter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
