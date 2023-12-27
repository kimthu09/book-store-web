package importnotestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/importnote/importnotemodel"
	"book-store-management-backend/module/supplier/suppliermodel/filter"
	"context"
	"gorm.io/gorm"
	"time"
)

func (s *sqlStore) ListAllImportNoteBySupplier(
	supplierId string,
	filter *filter.SupplierImportFilter,
	ctx context.Context,
	paging *common.Paging,
	moreKeys ...string) ([]importnotemodel.ImportNote, error) {
	var result []importnotemodel.ImportNote
	db := s.db

	db = db.Table(common.TableImportNote)

	handleSupplierImportFilter(db, filter)

	dbTemp, errPaging := common.HandlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Where("supplierId = ?", supplierId).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func handleSupplierImportFilter(
	db *gorm.DB,
	filterSupplierDebt *filter.SupplierImportFilter) {
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
