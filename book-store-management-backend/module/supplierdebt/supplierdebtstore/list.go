package supplierdebtstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/supplier/suppliermodel/filter"
	"book-store-management-backend/module/supplierdebt/supplierdebtmodel"
	"context"
	"gorm.io/gorm"
	"time"
)

func (s *sqlStore) ListSupplierDebt(
	ctx context.Context,
	supplierId string,
	filterSupplierDebt *filter.SupplierDebtFilter,
	paging *common.Paging,
	moreKeys ...string) ([]supplierdebtmodel.SupplierDebt, error) {
	var result []supplierdebtmodel.SupplierDebt
	db := s.db

	db = db.Table(common.TableSupplierDebt)

	db = db.Where("supplierId = ?", supplierId)

	handleFilter(db, filterSupplierDebt)

	dbTemp, errPaging := common.HandlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Order("createdAt desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func handleFilter(
	db *gorm.DB,
	filterSupplierDebt *filter.SupplierDebtFilter) {
	if filterSupplierDebt != nil {
		if filterSupplierDebt.DateFrom != nil {
			timeFrom := time.Unix(*filterSupplierDebt.DateFrom, 0)
			db = db.Where("createdAt >= ?", timeFrom)
		}
		if filterSupplierDebt.DateTo != nil {
			timeTo := time.Unix(*filterSupplierDebt.DateTo, 0)
			db = db.Where("createdAt <= ?", timeTo)
		}
	}
}
