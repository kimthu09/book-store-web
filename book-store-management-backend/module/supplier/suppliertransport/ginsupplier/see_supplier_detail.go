package ginsupplier

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/supplier/supplierbiz"
	"book-store-management-backend/module/supplier/supplierrepo"
	"book-store-management-backend/module/supplier/supplierstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary See detail of supplier
// @Tags suppliers
// @Accept json
// @Produce json
// @Param id path string true "supplier id"
// @Response 200 {object} suppliermodel.Supplier "supplier"
// @Response 400 {object} common.AppError "error"
// @Router /suppliers/{id} [get]
func SeeSupplierDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		supplierStore := supplierstore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := supplierrepo.NewSeeSupplierDetailRepo(supplierStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := supplierbiz.NewSeeSupplierDetailBiz(repo, requester)

		result, err := biz.SeeSupplierDetail(
			c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
