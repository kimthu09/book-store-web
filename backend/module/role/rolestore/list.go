package rolestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/role/rolemodel"
	"context"
)

func (s *sqlStore) ListRole(
	ctx context.Context) ([]rolemodel.SimpleRole, error) {
	var result []rolemodel.SimpleRole
	db := s.db

	db = db.Table(common.TableRole)

	if err := db.
		Find(&result).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
