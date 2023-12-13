package userstore

import (
	"book-store-management-backend/common"
	"context"
)

func (s *sqlStore) UpdatePasswordUser(
	ctx context.Context,
	id string,
	password string) error {
	db := s.db

	if err := db.
		Table(common.TableUser).
		Where("id = ?", id).
		Update("password", password).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
