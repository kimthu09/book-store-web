package authorstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/author/authormodel"
	"context"
)

func (s *sqlStore) GetAllAuthor(ctx context.Context) ([]authormodel.SimpleAuthor, error) {
	var result []authormodel.SimpleAuthor
	db := s.db

	db = db.Table(common.TableAuthor)

	if err := db.
		Order("name").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
