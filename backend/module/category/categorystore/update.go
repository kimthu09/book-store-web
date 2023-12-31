package categorystore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/category/categorymodel"
	"context"
)

func (s *sqlStore) UpdateCategory(
	ctx context.Context,
	id string,
	data *categorymodel.ReqUpdateCategory) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("name"); key {
			case "name":
				return categorymodel.ErrCategoryNameDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
