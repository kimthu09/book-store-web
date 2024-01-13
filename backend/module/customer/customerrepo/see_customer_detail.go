package customerrepo

import (
	"book-store-management-backend/module/customer/customermodel"
	"context"
)

type FindCustomerStore interface {
	FindCustomer(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*customermodel.Customer, error)
}

type seeCustomerDetailRepo struct {
	customerStore FindCustomerStore
}

func NewSeeCustomerDetailRepo(
	customerStore FindCustomerStore) *seeCustomerDetailRepo {
	return &seeCustomerDetailRepo{
		customerStore: customerStore,
	}
}

func (biz *seeCustomerDetailRepo) SeeCustomerDetail(
	ctx context.Context,
	customerId string) (*customermodel.Customer, error) {
	customer, errCustomer := biz.customerStore.FindCustomer(
		ctx, map[string]interface{}{"id": customerId})
	if errCustomer != nil {
		return nil, errCustomer
	}

	return customer, nil
}
