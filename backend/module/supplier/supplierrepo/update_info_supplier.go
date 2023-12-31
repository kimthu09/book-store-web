package supplierrepo

import (
	"book-store-management-backend/module/supplier/suppliermodel"
	"context"
)

type UpdateInfoSupplierStore interface {
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

func (repo *updateInfoSupplierRepo) UpdateSupplierInfo(
	ctx context.Context,
	supplierId string,
	data *suppliermodel.ReqUpdateInfoSupplier) error {
	if err := repo.store.UpdateSupplierInfo(ctx, supplierId, data); err != nil {
		return err
	}
	return nil
}
