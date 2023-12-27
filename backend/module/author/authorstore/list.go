package authorstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/author/authormodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) ListAuthor(ctx context.Context, filter *authormodel.Filter, propertiesContainSearchKey []string, paging *common.Paging) ([]authormodel.Author, error) {
	var result []authormodel.Author
	db := s.db

	db = db.Table(common.TableAuthor)

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
	filter *authormodel.Filter,
	propertiesContainSearchKey []string) {
	if filter != nil {
		if filter.SearchKey != "" {
			db = common.GetWhereClause(db, filter.SearchKey, propertiesContainSearchKey)
		}
	}
}
