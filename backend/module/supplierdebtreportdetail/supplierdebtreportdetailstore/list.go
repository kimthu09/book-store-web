package supplierdebtreportdetailstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/supplierdebtreportdetail/supplierdebtreportdetailmodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) ListSupplierDebtReportDetail(
	ctx context.Context,
	supplierDebtReportDetailId string,
	paging *common.Paging,
) ([]supplierdebtreportdetailmodel.SupplierDebtReportDetail, error) {
	var data []supplierdebtreportdetailmodel.SupplierDebtReportDetail
	db := s.db

	dbTemp, errPaging := common.HandlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	if err := db.
		Where("supplierDebtReportDetailId", supplierDebtReportDetailId).
		Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return data, nil
}
