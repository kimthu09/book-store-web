package customerbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/customer/customermodel"
	"context"
)

type UpdateInfoCustomerRepo interface {
	UpdateCustomerInfo(
		ctx context.Context,
		customerId string,
		data *customermodel.ReqUpdateInfoCustomer,
	) error
}

type updateInfoCustomerBiz struct {
	repo      UpdateInfoCustomerRepo
	requester middleware.Requester
}

func NewUpdateInfoCustomerBiz(
	repo UpdateInfoCustomerRepo,
	requester middleware.Requester) *updateInfoCustomerBiz {
	return &updateInfoCustomerBiz{repo: repo, requester: requester}
}

func (biz *updateInfoCustomerBiz) UpdateInfoCustomer(
	ctx context.Context,
	id string,
	data *customermodel.ReqUpdateInfoCustomer) error {
	if !biz.requester.IsHasFeature(common.CustomerUpdateInfoFeatureCode) {
		return customermodel.ErrCustomerUpdateInfoNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.repo.UpdateCustomerInfo(ctx, id, data); err != nil {
		return err
	}

	return nil
}
