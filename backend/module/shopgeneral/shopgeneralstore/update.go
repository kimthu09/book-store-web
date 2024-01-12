package shopgeneralstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/shopgeneral/shopgeneralmodel"
	"context"
)

func (s *sqlStore) UpdateGeneralShop(
	ctx context.Context,
	data *shopgeneralmodel.ReqUpdateShopGeneral) error {
	db := s.db

	if err := db.Where("id = ?", "shop").Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
