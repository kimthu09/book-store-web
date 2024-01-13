package customerrepo

import (
	"book-store-management-backend/module/customer/customermodel"
	"context"
)

type CreateCustomerStore interface {
	CreateCustomer(
		ctx context.Context,
		data *customermodel.ReqCreateCustomer) error
}

type createCustomerRepo struct {
	store CreateCustomerStore
}

func NewCreateCustomerRepo(store CreateCustomerStore) *createCustomerRepo {
	return &createCustomerRepo{store: store}
}

func (repo *createCustomerRepo) CreateCustomer(
	ctx context.Context,
	data *customermodel.ReqCreateCustomer) error {
	if err := repo.store.CreateCustomer(ctx, data); err != nil {
		return err
	}
	return nil
}
