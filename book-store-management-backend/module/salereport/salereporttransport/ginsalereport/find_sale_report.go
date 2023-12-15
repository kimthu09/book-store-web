package ginsalereport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/invoice/invoicestore"
	"book-store-management-backend/module/invoicedetail/invoicedetailstore"
	"book-store-management-backend/module/salereport/salereportbiz"
	"book-store-management-backend/module/salereport/salereportmodel"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Find sale report
// @Tags reports
// @Accept json
// @Produce json
// @Param condition body salereportmodel.ReqFindSaleReport true "time from and time to"
// @Response 200 {object} salereportmodel.ResFindSaleReport "sale report"
// @Response 400 {object} common.AppError "error"
// @Router /reports/sale [post]
func FindSaleReport(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data salereportmodel.ReqFindSaleReport

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		invoiceStore := invoicestore.NewSQLStore(db)
		invoiceDetailStore := invoicedetailstore.NewSQLStore(db)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		business := salereportbiz.NewFindSaleReportBiz(
			invoiceStore, invoiceDetailStore, requester,
		)

		report, err := business.FindSaleReport(
			c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(report))
	}
}
