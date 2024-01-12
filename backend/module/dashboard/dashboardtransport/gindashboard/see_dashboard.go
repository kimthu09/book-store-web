package gindashboard

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/dashboard/dashboardbiz"
	"book-store-management-backend/module/dashboard/dashboardmodel"
	"book-store-management-backend/module/dashboard/dashboardrepo"
	"book-store-management-backend/module/invoice/invoicestore"
	"book-store-management-backend/module/invoicedetail/invoicedetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary See dashboard
// @Tags dashboard
// @Accept json
// @Produce json
// @Param condition body dashboardmodel.ReqSeeDashboard true "time from and time to"
// @Response 200 {object} dashboardmodel.ResSeeDashboard "dashboard"
// @Response 400 {object} common.AppError "error"
// @Router /dashboard [post]
func SeeDashboard(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data dashboardmodel.ReqSeeDashboard

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		invoiceStore := invoicestore.NewSQLStore(db)
		invoiceDetailStore := invoicedetailstore.NewSQLStore(db)

		dashboardRepo := dashboardrepo.NewSeeDashboardBiz(invoiceStore, invoiceDetailStore)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		business := dashboardbiz.NewSeeDashboardBiz(
			dashboardRepo, requester,
		)

		dashboard, err := business.SeeDashboard(
			c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(dashboard))
	}
}
