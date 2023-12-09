package categorystore

import (
	"book-store-management-backend/module/category/categorymodel"
	"context"
)

func (s *sqlStore) CheckExistByID(ctx context.Context, id string) (bool, error) {
	var count int64
	err := s.db.Model(&categorymodel.Category{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
