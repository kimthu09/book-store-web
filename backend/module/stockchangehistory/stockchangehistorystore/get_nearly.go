package stockchangehistorystore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/stockchangehistory/stockchangehistorymodel"
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

func (s *sqlStore) GetNearlyStockChangeHistory(
	ctx context.Context,
	bookId string,
	timeFrom time.Time) (*stockchangehistorymodel.StockChangeHistory, error) {
	var result stockchangehistorymodel.StockChangeHistory
	db := s.db

	db = db.Table(common.TableStockChangeHistory)

	timeRequest := timeFrom.Add(-time.Second)

	if err := db.
		Where("bookId = ?", bookId).
		Where("createdAt <= ?", timeRequest).
		Order("createdAt desc").
		First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
