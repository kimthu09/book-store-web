package supplierbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/supplier/suppliermodel"
	"context"
)

type CreateSupplierRepo interface {
	CreateSupplier(
		ctx context.Context,
		data *suppliermodel.ReqCreateSupplier,
	) error
}

type createSupplierBiz struct {
	gen       generator.IdGenerator
	repo      CreateSupplierRepo
	requester middleware.Requester
}

func NewCreateSupplierBiz(
	gen generator.IdGenerator,
	repo CreateSupplierRepo,
	requester middleware.Requester) *createSupplierBiz {
	return &createSupplierBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createSupplierBiz) CreateSupplier(
	ctx context.Context,
	data *suppliermodel.ReqCreateSupplier) error {
	if !biz.requester.IsHasFeature(common.SupplierCreateFeatureCode) {
		return suppliermodel.ErrSupplierCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := handleSupplierId(biz.gen, data); err != nil {
		return err
	}
	if err := biz.repo.CreateSupplier(ctx, data); err != nil {
		return err
	}

	return nil
}

func handleSupplierId(gen generator.IdGenerator, data *suppliermodel.ReqCreateSupplier) error {
	id, err := gen.IdProcess(data.Id)
	if err != nil {
		return err
	}

	data.Id = id

	return nil
}
