package customerrepo

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/customer/customermodel"
	"context"
)

type ListCustomerStore interface {
	ListCustomer(
		ctx context.Context,
		filter *customermodel.Filter,
		propertiesContainSearchKey []string,
		paging *common.Paging,
	) ([]customermodel.Customer, error)
}

type listCustomerRepo struct {
	store ListCustomerStore
}

func NewListCustomerRepo(store ListCustomerStore) *listCustomerRepo {
	return &listCustomerRepo{store: store}
}

func (repo *listCustomerRepo) ListCustomer(
	ctx context.Context,
	filter *customermodel.Filter,
	paging *common.Paging) ([]customermodel.Customer, error) {
	result, err := repo.store.ListCustomer(
		ctx,
		filter,
		[]string{"id", "name", "email", "phone"},
		paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
