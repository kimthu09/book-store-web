package categorystore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/category/categorymodel"
	"context"
)

func (s *sqlStore) GetByListId(ctx context.Context, ids []string) ([]categorymodel.Category, error) {
	var result []categorymodel.Category
	db := s.db

	if err := db.
		Table(common.TableCategory).
		Where("id IN (?)", ids).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
