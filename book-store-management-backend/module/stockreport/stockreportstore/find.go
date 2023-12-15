package stockreportstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/stockreport/stockreportmodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) FindStockReport(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*stockreportmodel.StockReport, error) {
	var data stockreportmodel.StockReport
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
