package categorystore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/category/categorymodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) FindCategory(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*categorymodel.SimpleCategory, error) {
	var data categorymodel.SimpleCategory
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
