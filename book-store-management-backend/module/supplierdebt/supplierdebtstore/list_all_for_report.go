package supplierdebtstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/supplierdebt/supplierdebtmodel"
	"context"
	"time"
)

func (s *sqlStore) ListAllSupplierDebtForReport(
	ctx context.Context,
	supplierId string,
	timeFrom time.Time,
	timeTo time.Time) ([]supplierdebtmodel.SupplierDebt, error) {
	var result []supplierdebtmodel.SupplierDebt
	db := s.db

	db = db.Table(common.TableSupplierDebt)

	if err := db.
		Where("supplierId = ?", supplierId).
		Where("createdAt >= ?", timeFrom).
		Where("createdAt <= ?", timeTo).
		Order("createdAt").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
