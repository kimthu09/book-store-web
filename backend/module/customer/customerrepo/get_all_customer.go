package customerrepo

import (
	"book-store-management-backend/module/invoice/invoicemodel"
	"context"
)

type GetAllCustomerStore interface {
	GetAllCustomer(
		ctx context.Context,
	) ([]invoicemodel.SimpleCustomer, error)
}

type getAllCustomerRepo struct {
	store GetAllCustomerStore
}

func NewGetAllCustomerRepo(store GetAllCustomerStore) *getAllCustomerRepo {
	return &getAllCustomerRepo{store: store}
}

func (repo *getAllCustomerRepo) GetAllCustomer(
	ctx context.Context) ([]invoicemodel.SimpleCustomer, error) {
	result, err := repo.store.GetAllCustomer(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}
