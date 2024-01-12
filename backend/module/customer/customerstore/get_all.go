package customerstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/invoice/invoicemodel"
	"context"
)

func (s *sqlStore) GetAllCustomer(
	ctx context.Context) ([]invoicemodel.SimpleCustomer, error) {
	var result []invoicemodel.SimpleCustomer
	db := s.db

	db = db.Table(common.TableCustomer)

	if err := db.
		Order("name").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
