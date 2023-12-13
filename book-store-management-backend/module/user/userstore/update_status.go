package userstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/user/usermodel"
	"context"
)

func (s *sqlStore) UpdateStatusUsers(
	ctx context.Context,
	id string,
	data *usermodel.ReqUpdateStatusUser) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
