package booktitlestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"context"
)

func (s *sqlStore) GetAllBookTitle(
	ctx context.Context,
	moreKeys ...string) ([]booktitlemodel.SimpleBookTitle, error) {
	var result []booktitlemodel.SimpleBookTitle
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Order("name").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
