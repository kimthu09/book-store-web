package supplierrepo

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/supplier/suppliermodel/filter"
	"book-store-management-backend/module/supplierdebt/supplierdebtmodel"
	"context"
)

type ListSupplierDebtStore interface {
	ListSupplierDebt(
		ctx context.Context,
		supplierId string,
		filterSupplierDebt *filter.SupplierDebtFilter,
		paging *common.Paging,
		moreKeys ...string) ([]supplierdebtmodel.SupplierDebt, error)
}

type seeSupplierDebtRepo struct {
	debtStore     ListSupplierDebtStore
	supplierStore FindSupplierStore
}

func NewSeeSupplierDebtRepo(
	debtStore ListSupplierDebtStore) *seeSupplierDebtRepo {
	return &seeSupplierDebtRepo{
		debtStore: debtStore,
	}
}

func (biz *seeSupplierDebtRepo) SeeSupplierDebt(
	ctx context.Context,
	supplierId string,
	filterSupplierDebt *filter.SupplierDebtFilter,
	paging *common.Paging) ([]supplierdebtmodel.SupplierDebt, error) {
	debts, errDebts := biz.debtStore.ListSupplierDebt(
		ctx,
		supplierId,
		filterSupplierDebt,
		paging,
		"CreatedByUser",
	)
	if errDebts != nil {
		return nil, errDebts
	}

	return debts, nil
}
