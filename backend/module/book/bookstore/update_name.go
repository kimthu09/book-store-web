package bookstore

import (
	"book-store-management-backend/common"
	"context"
)

func (s *sqlStore) UpdateName(
	ctx context.Context, bookTitleId string, name *string) error {
	db := s.db.Table(common.TableBook).
		Where("booktitleid = ?", bookTitleId).
		Update("name", name)
	if err := db.Error; err != nil {
		return err
	}
	return nil
}
