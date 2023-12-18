package bookstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/book/bookmodel"
	"context"
	"log"
)

func (s *sqlStore) GetAllBook(
	ctx context.Context,
	moreKeys ...string) ([]bookmodel.ResUnitBook, error) {
	var result []bookmodel.ResUnitBook
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Order("name").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	log.Println(result[0].BookTitleID)
	return result, nil
}
