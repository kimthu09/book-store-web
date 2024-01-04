package booktitlestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (store *sqlStore) FindBookTitle(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*booktitlemodel.SimpleBookTitle, error) {
	var data booktitlemodel.SimpleBookTitle
	db := store.db

	db = db.Table(common.TableBookTitle)

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
