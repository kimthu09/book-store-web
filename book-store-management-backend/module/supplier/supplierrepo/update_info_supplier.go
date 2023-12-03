package supplierrepo

import (
	"book-store-management-backend/module/supplier/suppliermodel"
	"context"
)

type UpdateInfoSupplierStore interface {
	FindSupplier(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*suppliermodel.Supplier, error)
	UpdateSupplierInfo(
		ctx context.Context,
		id string,
		data *suppliermodel.ReqUpdateInfoSupplier,
	) error
}

type updateInfoSupplierRepo struct {
	store UpdateInfoSupplierStore
}

func NewUpdateInfoSupplierRepo(store UpdateInfoSupplierStore) *updateInfoSupplierRepo {
	return &updateInfoSupplierRepo{store: store}
}

func (repo *updateInfoSupplierRepo) CheckExist(
	ctx context.Context,
	supplierId string) error {
	if _, err := repo.store.FindSupplier(
		ctx,
		map[string]interface{}{
			"id": supplierId,
		},
	); err != nil {
		return err
	}
	return nil
}

func (repo *updateInfoSupplierRepo) UpdateSupplierInfo(
	ctx context.Context,
	supplierId string,
	data *suppliermodel.ReqUpdateInfoSupplier) error {
	if err := repo.store.UpdateSupplierInfo(ctx, supplierId, data); err != nil {
		return err
	}
	return nil
}
