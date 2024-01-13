package dashboardbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/dashboard/dashboardmodel"
	"context"
)

type SeeDashboardRepo interface {
	SeeDashboard(
		ctx context.Context,
		data *dashboardmodel.ReqSeeDashboard) (*dashboardmodel.ResSeeDashboard, error)
}

type seeDashboardBiz struct {
	repo      SeeDashboardRepo
	requester middleware.Requester
}

func NewSeeDashboardBiz(
	repo SeeDashboardRepo,
	requester middleware.Requester) *seeDashboardBiz {
	return &seeDashboardBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *seeDashboardBiz) SeeDashboard(
	ctx context.Context,
	data *dashboardmodel.ReqSeeDashboard) (*dashboardmodel.ResSeeDashboard, error) {
	if !biz.requester.IsHasFeature(common.SaleReportViewFeatureCode) {
		return nil, dashboardmodel.ErrDashboardViewNoPermission
	}

	if err := data.Validate(); err != nil {
		return nil, err
	}

	dashboard, err := biz.repo.SeeDashboard(ctx, data)
	if err != nil {
		return nil, err
	}

	return dashboard, nil
}
