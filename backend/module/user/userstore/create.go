package userstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/user/usermodel"
	"context"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *usermodel.ReqCreateUser) error {
	db := s.db

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
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
