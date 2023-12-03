package supplierrepo

import (
	"book-store-management-backend/module/supplier/suppliermodel"
	"book-store-management-backend/module/supplierdebt/supplierdebtmodel"
	"context"
)

type PaySupplierStore interface {
	FindSupplier(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*suppliermodel.Supplier, error)
	UpdateSupplierDebt(
		ctx context.Context,
		id string,
		data *suppliermodel.ReqUpdateDebtSupplier,
	) error
}

type CreateSupplierDebtStore interface {
	CreateSupplierDebt(
		ctx context.Context,
		data *supplierdebtmodel.SupplierDebtCreate) error
}

type paySupplierRepo struct {
	supplierStore     PaySupplierStore
	supplierDebtStore CreateSupplierDebtStore
}

func NewUpdatePayRepo(
	supplierStore PaySupplierStore,
	supplierDebtStore CreateSupplierDebtStore) *paySupplierRepo {
	return &paySupplierRepo{
		supplierStore:     supplierStore,
		supplierDebtStore: supplierDebtStore,
	}
}

func (repo *paySupplierRepo) GetDebtSupplier(
	ctx context.Context,
	supplierId string) (*float32, error) {
	supplier, err := repo.supplierStore.FindSupplier(
		ctx, map[string]interface{}{"id": supplierId},
	)
	if err != nil {
		return nil, err
	}

	return &supplier.Debt, nil
}

func (repo *paySupplierRepo) CreateSupplierDebt(
	ctx context.Context,
	data *supplierdebtmodel.SupplierDebtCreate) error {
	if err := repo.supplierDebtStore.CreateSupplierDebt(ctx, data); err != nil {
		return err
	}
	return nil
}

func (repo *paySupplierRepo) UpdateDebtSupplier(
	ctx context.Context,
	supplierId string,
	data *suppliermodel.ReqUpdateDebtSupplier) error {
	if err := repo.supplierStore.UpdateSupplierDebt(ctx, supplierId, data); err != nil {
		return err
	}
	return nil
}
