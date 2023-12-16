package invoicestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/invoice/invoicemodel"
	"context"
)

func (s *sqlStore) CreateInvoice(
	ctx context.Context,
	data *invoicemodel.ReqCreateInvoice) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
