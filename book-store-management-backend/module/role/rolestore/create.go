package rolestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/role/rolemodel"
	"context"
)

func (s *sqlStore) CreateRole(
	ctx context.Context,
	data *rolemodel.ReqCreateRole) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("name"); key {
			case "name":
				return rolemodel.ErrRoleNameDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
