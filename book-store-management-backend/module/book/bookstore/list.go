package bookstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/book/bookmodel"
	"context"
	"gorm.io/gorm"
	"time"
)

func (store *sqlStore) ListBook(ctx context.Context, filter *bookmodel.Filter, propertiesContainSearchKey []string, paging *common.Paging) ([]BookDBModel, error) {
	var result []BookDBModel
	db := store.db

	db = db.Table(common.TableBook)

	handleFilter(db, filter, propertiesContainSearchKey)

	dbTemp, errPaging := common.HandlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	if err := db.
		Limit(int(paging.Limit)).
		Order("createdAt").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func handleFilter(
	db *gorm.DB,
	filter *bookmodel.Filter,
	propertiesContainSearchKey []string) {
	if filter != nil {
		if filter.SearchKey != "" {
			db = common.GetWhereClause(db, filter.SearchKey, propertiesContainSearchKey)
		}
		if filter.DateFromCreateAt != nil {
			timeFrom := time.Unix(*filter.DateFromCreateAt, 0)
			db = db.Where("createdAt >= ?", timeFrom)
		}
		if filter.DateToCreateAt != nil {
			timeTo := time.Unix(*filter.DateToCreateAt, 0)
			db = db.Where("createdAt <= ?", timeTo)
		}
		db = db.Where("isActive = ?", 1)
	}
}
