package inventorychecknotestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/inventorychecknote/inventorychecknotemodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) ListInventoryCheckNote(
	ctx context.Context,
	filter *inventorychecknotemodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging) ([]inventorychecknotemodel.InventoryCheckNote, error) {
	var result []inventorychecknotemodel.InventoryCheckNote
	db := s.db

	db = db.Table(common.TableInventoryCheckNote)

	handleFilter(db, filter, propertiesContainSearchKey)

	dbTemp, errPaging := handlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	if err := db.
		Limit(int(paging.Limit)).
		Order("createAt desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func handleFilter(
	db *gorm.DB,
	filter *inventorychecknotemodel.Filter,
	propertiesContainSearchKey []string) {
	if filter != nil {
		if filter.SearchKey != "" {
			db = common.GetWhereClause(db, filter.SearchKey, propertiesContainSearchKey)
		}
	}
}

func handlePaging(db *gorm.DB, paging *common.Paging) (*gorm.DB, error) {
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	offset := (paging.Page - 1) * paging.Limit
	db = db.Offset(int(offset))

	return db, nil
}
