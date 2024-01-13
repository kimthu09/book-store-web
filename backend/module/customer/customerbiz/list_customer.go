package customerbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/customer/customermodel"
	"context"
)

type ListCustomerRepo interface {
	ListCustomer(
		ctx context.Context,
		filter *customermodel.Filter,
		paging *common.Paging,
	) ([]customermodel.Customer, error)
}

type listCustomerBiz struct {
	repo      ListCustomerRepo
	requester middleware.Requester
}

func NewListCustomerBiz(
	repo ListCustomerRepo,
	requester middleware.Requester) *listCustomerBiz {
	return &listCustomerBiz{repo: repo, requester: requester}
}

func (biz *listCustomerBiz) ListCustomer(
	ctx context.Context,
	filter *customermodel.Filter,
	paging *common.Paging) ([]customermodel.Customer, error) {
	if !biz.requester.IsHasFeature(common.CustomerViewFeatureCode) {
		return nil, customermodel.ErrCustomerViewNoPermission
	}

	result, err := biz.repo.ListCustomer(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
