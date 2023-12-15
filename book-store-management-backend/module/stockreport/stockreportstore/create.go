package stockreportstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/stockreport/stockreportmodel"
	"context"
)

func (s *sqlStore) CreateStockReport(
	ctx context.Context,
	data *stockreportmodel.ReqFindStockReport) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
