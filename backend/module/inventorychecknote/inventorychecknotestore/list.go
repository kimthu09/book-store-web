package inventorychecknotestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/inventorychecknote/inventorychecknotemodel"
	"context"
	"gorm.io/gorm"
	"time"
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

	dbTemp, errPaging := common.HandlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	if err := db.
		Preload("CreatedByUser").
		Order("createdAt desc").
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
		if filter.DateFromCreatedAt != nil {
			timeFrom := time.Unix(*filter.DateFromCreatedAt, 0)
			db = db.Where("InventoryCheckNote.createdAt >= ?", timeFrom)
		}
		if filter.DateToCreatedAt != nil {
			timeTo := time.Unix(*filter.DateToCreatedAt, 0)
			db = db.Where("InventoryCheckNote.createdAt <= ?", timeTo)
		}
		if filter.CreatedBy != nil {
			db = db.
				Joins("JOIN MUser ON InventoryCheckNote.createdBy = MUser.id").
				Where("MUser.id = ?", *filter.CreatedBy)
		}
	}
}
