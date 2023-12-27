package supplierstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/supplier/suppliermodel"
	"context"
)

func (s *sqlStore) CreateSupplier(ctx context.Context, data *suppliermodel.ReqCreateSupplier) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("PRIMARY", "phone"); key {
			case "PRIMARY":
				return suppliermodel.ErrSupplierIdDuplicate
			case "phone":
				return suppliermodel.ErrSupplierPhoneDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
