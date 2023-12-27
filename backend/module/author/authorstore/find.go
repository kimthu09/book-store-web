package authorstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/author/authormodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) FindAuthor(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*authormodel.SimpleAuthor, error) {
	var data authormodel.SimpleAuthor
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
