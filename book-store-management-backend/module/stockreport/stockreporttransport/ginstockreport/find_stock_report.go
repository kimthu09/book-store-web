package ginstockreports

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/book/bookstore"
	"book-store-management-backend/module/stockchangehistory/stockchangehistorystore"
	"book-store-management-backend/module/stockreport/stockreportbiz"
	"book-store-management-backend/module/stockreport/stockreportmodel"
	"book-store-management-backend/module/stockreport/stockreportstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Find stock report
// @Tags reports
// @Accept json
// @Produce json
// @Param condition body stockreportmodel.ReqFindStockReport true "time from and time to"
// @Response 200 {object} stockreportmodel.ResFindStockReport "stock report"
// @Response 400 {object} common.AppError "error"
// @Router /reports/stock [post]
func FindStockReport(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data stockreportmodel.ReqFindStockReport

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		bookStore := bookstore.NewSQLStore(db)
		stockChangeHistoryStore := stockchangehistorystore.NewSQLStore(db)
		stockReportStore := stockreportstore.NewSQLStore(db)

		gen := generator.NewShortIdGenerator()
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		business := stockreportbiz.NewFindStockReportBiz(
			gen,
			bookStore,
			stockChangeHistoryStore,
			stockReportStore,
			requester,
		)

		report, err := business.FindStockReport(
			c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(report))
	}
}
