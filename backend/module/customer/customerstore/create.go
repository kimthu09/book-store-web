package customerstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/customer/customermodel"
	"context"
)

func (s *sqlStore) CreateCustomer(
	ctx context.Context,
	data *customermodel.ReqCreateCustomer) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("PRIMARY", "phone"); key {
			case "PRIMARY":
				return customermodel.ErrCustomerIdDuplicate
			case "phone":
				return customermodel.ErrCustomerPhoneDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
