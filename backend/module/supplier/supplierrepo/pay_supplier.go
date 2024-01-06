package supplierrepo

import (
	"book-store-management-backend/module/importnote/importnotemodel"
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

type GetAllIdImportNoteStore interface {
	GetAllIdImportNote(
		ctx context.Context) ([]importnotemodel.ImportNoteId, error)
}

type paySupplierRepo struct {
	supplierStore     PaySupplierStore
	supplierDebtStore CreateSupplierDebtStore
	importNoteStore   GetAllIdImportNoteStore
}

func NewUpdatePayRepo(
	supplierStore PaySupplierStore,
	supplierDebtStore CreateSupplierDebtStore,
	importNoteStore GetAllIdImportNoteStore) *paySupplierRepo {
	return &paySupplierRepo{
		supplierStore:     supplierStore,
		supplierDebtStore: supplierDebtStore,
		importNoteStore:   importNoteStore,
	}
}

func (repo *paySupplierRepo) GetDebtSupplier(
	ctx context.Context,
	supplierId string) (*int, error) {
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
	newUpdate := -*data.QuantityUpdate
	newData := suppliermodel.ReqUpdateDebtSupplier{
		Id:             data.Id,
		QuantityUpdate: &newUpdate,
		CreatedBy:      data.CreatedBy,
	}
	if err := repo.supplierStore.UpdateSupplierDebt(ctx, supplierId, &newData); err != nil {
		return err
	}
	return nil
}

func (repo *paySupplierRepo) GetAllImportNoteId(
	ctx context.Context) ([]string, error) {
	result, err := repo.importNoteStore.GetAllIdImportNote(ctx)
	if err != nil {
		return nil, err
	}

	var listId []string
	for _, v := range result {
		listId = append(listId, v.Id)
	}

	return listId, nil
}
