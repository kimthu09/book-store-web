package userstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/user/usermodel"
	"context"
)

func (s *sqlStore) UpdateInfoUser(
	ctx context.Context,
	id string,
	data *usermodel.ReqUpdateInfoUser) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("email"); key {
			case "email":
				return usermodel.ErrUserEmailDuplicated
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
