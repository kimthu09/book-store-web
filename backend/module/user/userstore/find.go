package userstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/user/usermodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) (*usermodel.User, error) {
	db := s.db.Table(common.TableUser)

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}

		return nil, common.ErrDB(err)
	}

	return &user, nil
}
