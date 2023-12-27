package supplierdebtreportstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/supplierdebtreport/supplierdebtreportmodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) FindSupplierDebtReport(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*supplierdebtreportmodel.SupplierDebtReport, error) {
	var data supplierdebtreportmodel.SupplierDebtReport
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
