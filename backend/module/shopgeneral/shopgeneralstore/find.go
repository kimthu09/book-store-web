package shopgeneralstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/shopgeneral/shopgeneralmodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) FindShopGeneral(
	ctx context.Context) (*shopgeneralmodel.ShopGeneral, error) {
	var data shopgeneralmodel.ShopGeneral
	db := s.db

	if err := db.First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
