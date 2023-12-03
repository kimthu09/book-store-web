package supplierstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/supplier/suppliermodel"
	"context"
)

func (s *sqlStore) UpdateSupplierInfo(
	ctx context.Context,
	id string,
	data *suppliermodel.ReqUpdateInfoSupplier) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("phone"); key {
			case "phone":
				return suppliermodel.ErrSupplierPhoneDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
