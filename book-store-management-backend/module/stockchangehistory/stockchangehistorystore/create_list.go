package stockchangehistorystore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/stockchangehistory/stockchangehistorymodel"

	"context"
)

func (s *sqlStore) CreateLisStockChangeHistory(
	ctx context.Context,
	data []stockchangehistorymodel.StockChangeHistory) error {
	db := s.db
	if err := db.CreateInBatches(data, len(data)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
