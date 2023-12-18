package bookstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/book/bookmodel"
	"context"
)

func (s *sqlStore) GetAllBook(
	ctx context.Context,
	justGetAllActiveBook bool,
	moreKeys ...string) ([]bookmodel.ResUnitBook, error) {
	var result []bookmodel.ResUnitBook
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Where("isActive = ?", justGetAllActiveBook).
		Order("name").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
