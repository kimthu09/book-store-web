package bookstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/book/bookmodel"
	"context"
)

func (s *sqlStore) ListAllBook(
	ctx context.Context) ([]bookmodel.SimpleBook, error) {
	var result []bookmodel.SimpleBook
	db := s.db

	db = db.Table(common.TableBook)

	if err := db.
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
