package userstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/user/usermodel"
	"context"
)

func (s *sqlStore) GetAllUser(
	ctx context.Context,
	moreKeys ...string) ([]usermodel.SimpleUser, error) {
	var result []usermodel.SimpleUser
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
