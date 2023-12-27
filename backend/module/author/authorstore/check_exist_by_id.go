package authorstore

import (
	"book-store-management-backend/module/author/authormodel"
	"context"
)

func (s *sqlStore) CheckExistByID(ctx context.Context, id string) (bool, error) {
	var count int64
	err := s.db.Model(&authormodel.Author{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
