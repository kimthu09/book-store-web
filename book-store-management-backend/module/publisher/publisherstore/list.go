package publisherstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/publisher/publishermodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) ListPublisher(ctx context.Context, filter *publishermodel.Filter, propertiesContainSearchKey []string, paging *common.Paging) ([]publishermodel.Publisher, error) {
	var result []publishermodel.Publisher
	db := s.db

	db = db.Table(common.TablePublisher)

	handleFilter(db, filter, propertiesContainSearchKey)

	dbTemp, errPaging := common.HandlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	if err := db.
		Limit(int(paging.Limit)).
		Order("name").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func handleFilter(
	db *gorm.DB,
	filter *publishermodel.Filter,
	propertiesContainSearchKey []string) {
	if filter != nil {
		if filter.SearchKey != "" {
			db = common.GetWhereClause(db, filter.SearchKey, propertiesContainSearchKey)
		}
	}
}
