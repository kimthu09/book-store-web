package customerstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/customer/customermodel"
	"context"
)

func (s *sqlStore) UpdateCustomerInfo(
	ctx context.Context,
	id string,
	data *customermodel.ReqUpdateInfoCustomer) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("phone"); key {
			case "phone":
				return customermodel.ErrCustomerPhoneDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
