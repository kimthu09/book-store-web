package featurestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/feature/featuremodel"
	"context"
)

func (s *sqlStore) ListFeature(
	ctx context.Context) ([]featuremodel.Feature, error) {
	var result []featuremodel.Feature
	db := s.db

	db = db.Table(common.TableFeature)

	if err := db.
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
