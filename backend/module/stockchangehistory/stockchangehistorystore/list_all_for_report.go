package stockchangehistorystore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/stockchangehistory/stockchangehistorymodel"
	"context"
	"time"
)

func (s *sqlStore) ListAllStockChangeForReport(
	ctx context.Context,
	bookId string,
	timeFrom time.Time,
	timeTo time.Time) ([]stockchangehistorymodel.StockChangeHistory, error) {
	var result []stockchangehistorymodel.StockChangeHistory
	db := s.db

	db = db.Table(common.TableStockChangeHistory)

	if err := db.
		Where("bookId = ?", bookId).
		Where("createdAt >= ?", timeFrom).
		Where("createdAt <= ?", timeTo).
		Order("createdAt").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
