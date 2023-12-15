package supplierdebtreportdetailmodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/importnote/importnotemodel"
)

type SupplierDebtReportDetail struct {
	ReportId   string                         `json:"-" gorm:"column:reportId;"`
	SupplierId string                         `json:"-" gorm:"column:supplierId;"`
	Supplier   importnotemodel.SimpleSupplier `json:"supplier"`
	Initial    int                            `json:"initial" gorm:"column:initial;" example:"100000"`
	Debt       int                            `json:"debt" gorm:"column:debt;" example:"-40000"`
	Pay        int                            `json:"pay" gorm:"column:pay;" example:"20000"`
	Final      int                            `json:"final" gorm:"column:final;" example:"80000"`
}

func (*SupplierDebtReportDetail) TableName() string {
	return common.TableSupplierDebtReportDetail
}
