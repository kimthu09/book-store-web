package bookstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/book/bookmodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) FindBook(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*bookmodel.Book, error) {
	var data bookmodel.Book
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
