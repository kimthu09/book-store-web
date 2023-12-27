package authorstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/author/authormodel"
	"context"
)

func (s *sqlStore) GetByListId(ctx context.Context, ids []string) ([]authormodel.Author, error) {
	var result []authormodel.Author
	db := s.db

	if err := db.
		Table(common.TableAuthor).
		Where("id IN (?)", ids).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
