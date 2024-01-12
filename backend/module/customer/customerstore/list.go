package customerstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/customer/customermodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) ListCustomer(
	ctx context.Context,
	filter *customermodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging) ([]customermodel.Customer, error) {
	var result []customermodel.Customer
	db := s.db

	db = db.Table(common.TableCustomer)

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
	filter *customermodel.Filter,
	propertiesContainSearchKey []string) {
	if filter != nil {
		if filter.SearchKey != "" {
			db = common.GetWhereClause(db, filter.SearchKey, propertiesContainSearchKey)
		}
		if filter.MinPoint != nil {
			db = db.Where("point >= ?", filter.MinPoint)
		}
		if filter.MaxPoint != nil {
			db = db.Where("point <= ?", filter.MaxPoint)
		}
	}
}
