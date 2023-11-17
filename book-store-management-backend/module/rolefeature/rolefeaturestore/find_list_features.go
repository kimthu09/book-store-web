package rolefeaturestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/rolefeature/rolefeaturemodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) FindListFeatures(
	ctx context.Context,
	roleId string) ([]rolefeaturemodel.RoleFeature, error) {
	var data []rolefeaturemodel.RoleFeature
	db := s.db

	if err := db.
		Table(common.TableRoleFeature).
		Where("roleId = ?", roleId).
		Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return data, nil
}
