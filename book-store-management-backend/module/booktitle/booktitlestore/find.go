package booktitlestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (store *sqlStore) FindBook(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*booktitlemodel.Book, error) {
	var data booktitlemodel.Book
	db := store.db

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
