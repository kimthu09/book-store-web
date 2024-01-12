package invoicestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/customer/customermodel"
	"book-store-management-backend/module/invoice/invoicemodel"
	"context"
	"gorm.io/gorm"
	"time"
)

func (s *sqlStore) ListAllInvoiceByCustomer(
	ctx context.Context,
	customerId string,
	filter *customermodel.FilterInvoice,
	paging *common.Paging,
	moreKeys ...string) ([]invoicemodel.Invoice, error) {
	var result []invoicemodel.Invoice
	db := s.db

	db = db.Table(common.TableInvoice)

	db = db.Where("customerId = ?", customerId)

	handleFilterInvoice(db, filter)

	dbTemp, errPaging := common.HandlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Order("createdAt desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func handleFilterInvoice(
	db *gorm.DB,
	filter *customermodel.FilterInvoice) {
	if filter != nil {
		if filter.DateFrom != nil {
			timeFrom := time.Unix(*filter.DateFrom, 0)
			db = db.Where("createdAt >= ?", timeFrom)
		}
		if filter.DateTo != nil {
			timeTo := time.Unix(*filter.DateTo, 0)
			db = db.Where("createdAt <= ?", timeTo)
		}
	}
}
