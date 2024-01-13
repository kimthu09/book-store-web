package customerbiz

import (
	"book-store-management-backend/module/invoice/invoicemodel"
	"context"
)

type GetAllCustomerRepo interface {
	GetAllCustomer(
		ctx context.Context,
	) ([]invoicemodel.SimpleCustomer, error)
}

type getAllCustomerBiz struct {
	repo GetAllCustomerRepo
}

func NewGetAllCustomerBiz(
	repo GetAllCustomerRepo) *getAllCustomerBiz {
	return &getAllCustomerBiz{repo: repo}
}

func (biz *getAllCustomerBiz) GetAllCustomer(
	ctx context.Context) ([]invoicemodel.SimpleCustomer, error) {
	result, err := biz.repo.GetAllCustomer(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
