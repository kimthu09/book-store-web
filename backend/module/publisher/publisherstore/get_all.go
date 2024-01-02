package publisherstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/publisher/publishermodel"
	"context"
)

func (s *sqlStore) GetAllPublisher(ctx context.Context) ([]publishermodel.Publisher, error) {
	var result []publishermodel.Publisher
	db := s.db

	db = db.Table(common.TablePublisher)

	if err := db.
		Order("name").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
