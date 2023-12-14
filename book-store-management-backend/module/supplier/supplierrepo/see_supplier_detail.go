package supplierrepo

import (
	"book-store-management-backend/module/supplier/suppliermodel"
	"context"
)

type FindSupplierStore interface {
	FindSupplier(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*suppliermodel.Supplier, error)
}

type seeSupplierDetailRepo struct {
	supplierStore FindSupplierStore
}

func NewSeeSupplierDetailRepo(
	supplierStore FindSupplierStore) *seeSupplierDetailRepo {
	return &seeSupplierDetailRepo{
		supplierStore: supplierStore,
	}
}

func (biz *seeSupplierDetailRepo) SeeSupplierDetail(
	ctx context.Context,
	supplierId string) (*suppliermodel.Supplier, error) {
	supplier, errSupplier := biz.supplierStore.FindSupplier(
		ctx, map[string]interface{}{"id": supplierId})
	if errSupplier != nil {
		return nil, errSupplier
	}

	return supplier, nil
}
