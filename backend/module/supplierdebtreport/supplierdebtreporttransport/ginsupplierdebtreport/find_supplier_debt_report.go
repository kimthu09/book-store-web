package ginsupplierdebtreport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/supplier/supplierstore"
	"book-store-management-backend/module/supplierdebt/supplierdebtstore"
	"book-store-management-backend/module/supplierdebtreport/supplierdebtreportbiz"
	"book-store-management-backend/module/supplierdebtreport/supplierdebtreportmodel"
	"book-store-management-backend/module/supplierdebtreport/supplierdebtreportstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Find supplier debt report
// @Tags reports
// @Accept json
// @Produce json
// @Param condition body supplierdebtreportmodel.ReqFindSupplierDebtReport true "time from and time to"
// @Response 200 {object} supplierdebtreportmodel.ResFindSupplierDebtReport "supplier debt report"
// @Response 400 {object} common.AppError "error"
// @Router /reports/debt [post]
func FindSupplierDebtReport(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data supplierdebtreportmodel.ReqFindSupplierDebtReport

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		supplierStore := supplierstore.NewSQLStore(db)
		supplierDebtStore := supplierdebtstore.NewSQLStore(db)
		supplierDebtReportStore := supplierdebtreportstore.NewSQLStore(db)

		gen := generator.NewShortIdGenerator()
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		business := supplierdebtreportbiz.NewFindSupplierDebtReportBiz(
			gen,
			supplierStore,
			supplierDebtStore,
			supplierDebtReportStore,
			requester,
		)

		report, err := business.FindSupplierDebtReport(c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(report))
	}
}
