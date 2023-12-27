package supplierdebtstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/supplierdebt/supplierdebtmodel"
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

func (s *sqlStore) GetNearlySupplierDebt(
	ctx context.Context,
	supplierId string,
	timeFrom time.Time) (*supplierdebtmodel.SupplierDebt, error) {
	var result *supplierdebtmodel.SupplierDebt
	db := s.db

	db = db.Table(common.TableSupplierDebt)

	timeRequest := timeFrom.Add(-time.Second)

	if err := db.
		Where("supplierId = ?", supplierId).
		Where("createdAt <= ?", timeRequest).
		Order("createdAt desc").
		First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return result, nil
}
