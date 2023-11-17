package supplierstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/supplier/suppliermodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) UpdateSupplierDebt(
	ctx context.Context,
	id string,
	data *suppliermodel.SupplierUpdateDebt) error {
	db := s.db

	if err := db.Table(common.TableSupplier).
		Where("id = ?", id).
		Update("debt", gorm.Expr("debt + ?", data.Amount)).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
