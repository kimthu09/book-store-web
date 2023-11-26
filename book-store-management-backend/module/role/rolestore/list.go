package rolestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/role/rolemodel"
	"context"
)

func (s *sqlStore) ListRole(
	ctx context.Context) ([]rolemodel.Role, error) {
	var result []rolemodel.Role
	db := s.db

	db = db.Table(common.TableRole)

	if err := db.
		Preload("RoleFeatures").
		Find(&result).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
