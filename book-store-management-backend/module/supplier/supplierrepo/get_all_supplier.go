package supplierrepo

import (
	"book-store-management-backend/module/importnote/importnotemodel"
	"context"
)

type GetAllSupplierStore interface {
	GetAllSupplier(
		ctx context.Context,
		moreKeys ...string) ([]importnotemodel.SimpleSupplier, error)
}

type getAllSupplierRepo struct {
	store GetAllSupplierStore
}

func NewGetAllSupplierRepo(store GetAllSupplierStore) *getAllSupplierRepo {
	return &getAllSupplierRepo{store: store}
}

func (repo *getAllSupplierRepo) GetAllSupplier(
	ctx context.Context) ([]importnotemodel.SimpleSupplier, error) {
	result, err := repo.store.GetAllSupplier(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}
