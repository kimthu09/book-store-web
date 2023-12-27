package ginsupplier

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/module/supplier/supplierbiz"
	"book-store-management-backend/module/supplier/supplierrepo"
	"book-store-management-backend/module/supplier/supplierstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Get all supplier
// @Tags suppliers
// @Accept json
// @Produce json
// @Response 200 {object} suppliermodel.ResGetAllSupplier "list supplier"
// @Response 400 {object} common.AppError "error"
// @Router /suppliers/all [get]
func GetAllSupplier(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := supplierstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := supplierrepo.NewGetAllSupplierRepo(store)

		biz := supplierbiz.GetAllSupplierRepo(repo)

		result, err := biz.GetAllSupplier(c.Request.Context())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
