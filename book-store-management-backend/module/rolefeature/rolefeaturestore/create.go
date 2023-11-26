package rolefeaturestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/rolefeature/rolefeaturemodel"
	"context"
)

func (s *sqlStore) CreateRoleFeature(
	ctx context.Context,
	data *rolefeaturemodel.RoleFeature) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
