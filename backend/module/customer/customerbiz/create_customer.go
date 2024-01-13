package customerbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/customer/customermodel"
	"context"
)

type CreateCustomerRepo interface {
	CreateCustomer(
		ctx context.Context,
		data *customermodel.ReqCreateCustomer,
	) error
}

type createCustomerBiz struct {
	gen       generator.IdGenerator
	repo      CreateCustomerRepo
	requester middleware.Requester
}

func NewCreateCustomerBiz(
	gen generator.IdGenerator,
	repo CreateCustomerRepo,
	requester middleware.Requester) *createCustomerBiz {
	return &createCustomerBiz{gen: gen, repo: repo, requester: requester}
}

func (biz *createCustomerBiz) CreateCustomer(
	ctx context.Context,
	data *customermodel.ReqCreateCustomer) error {
	if !biz.requester.IsHasFeature(common.CustomerCreateFeatureCode) {
		return customermodel.ErrCustomerCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := handleCustomerId(biz.gen, data); err != nil {
		return err
	}

	if err := biz.repo.CreateCustomer(ctx, data); err != nil {
		return err
	}

	return nil
}

func handleCustomerId(gen generator.IdGenerator, data *customermodel.ReqCreateCustomer) error {
	idAddress, err := gen.IdProcess(data.Id)
	if err != nil {
		return err
	}

	data.Id = idAddress
	return nil
}
