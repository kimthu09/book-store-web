package invoicestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/invoice/invoicemodel"
	"context"
)

func (s *sqlStore) GetNearest(
	ctx context.Context,
	amountNeed int,
	moreKeys ...string) ([]invoicemodel.Invoice, error) {
	var result []invoicemodel.Invoice
	db := s.db

	db = db.Table(common.TableInvoice)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Limit(amountNeed).
		Order("createdAt desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
