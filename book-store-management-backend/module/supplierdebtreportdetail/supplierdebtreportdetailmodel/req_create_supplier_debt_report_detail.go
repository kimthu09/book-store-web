package supplierdebtreportdetailmodel

import (
	"book-store-management-backend/common"
)

type ReqCreateSupplierDebtReportDetail struct {
	ReportId   string `json:"-" gorm:"column:reportId;"`
	SupplierId string `json:"-" gorm:"column:supplierId;" example:"123"`
	Initial    int    `json:"-" gorm:"column:initial;" example:"-100000"`
	Debt       int    `json:"debt" gorm:"column:debt;" example:"-40000"`
	Pay        int    `json:"pay" gorm:"column:pay;" example:"20000"`
	Final      int    `json:"-" gorm:"column:final;" example:"-120000"`
}

func (*ReqCreateSupplierDebtReportDetail) TableName() string {
	return common.TableSupplierDebtReportDetail
}
