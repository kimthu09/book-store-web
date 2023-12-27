package supplierstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/supplier/suppliermodel"
	"book-store-management-backend/module/supplier/suppliermodel/filter"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) ListSupplier(
	ctx context.Context,
	filter *filter.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging) ([]suppliermodel.Supplier, error) {
	var result []suppliermodel.Supplier
	db := s.db

	db = db.Table(common.TableSupplier)

	handleFilter(db, filter, propertiesContainSearchKey)

	dbTemp, errPaging := common.HandlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	if err := db.
		Order("name").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func handleFilter(
	db *gorm.DB,
	filter *filter.Filter,
	propertiesContainSearchKey []string) {
	if filter != nil {
		if filter.SearchKey != "" {
			db = common.GetWhereClause(db, filter.SearchKey, propertiesContainSearchKey)
		}
		if filter.MinDebt != nil {
			db = db.Where("debt >= ?", filter.MinDebt)
		}
		if filter.MaxDebt != nil {
			db = db.Where("debt <= ?", filter.MaxDebt)
		}
	}
}
