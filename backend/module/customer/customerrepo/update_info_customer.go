package customerrepo

import (
	"book-store-management-backend/module/customer/customermodel"
	"context"
)

type UpdateInfoCustomerStore interface {
	UpdateCustomerInfo(
		ctx context.Context,
		id string,
		data *customermodel.ReqUpdateInfoCustomer,
	) error
}

type updateInfoCustomerRepo struct {
	store UpdateInfoCustomerStore
}

func NewUpdateInfoCustomerRepo(store UpdateInfoCustomerStore) *updateInfoCustomerRepo {
	return &updateInfoCustomerRepo{store: store}
}

func (repo *updateInfoCustomerRepo) UpdateCustomerInfo(
	ctx context.Context,
	customerId string,
	data *customermodel.ReqUpdateInfoCustomer) error {
	if err := repo.store.UpdateCustomerInfo(ctx, customerId, data); err != nil {
		return err
	}
	return nil
}
