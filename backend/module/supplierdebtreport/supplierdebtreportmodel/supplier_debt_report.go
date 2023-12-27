package supplierdebtreportmodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/supplierdebtreportdetail/supplierdebtreportdetailmodel"
	"errors"
	"time"
)

type SupplierDebtReport struct {
	Id       string                                                   `json:"id" gorm:"column:id;"`
	TimeFrom time.Time                                                `json:"timeFrom" gorm:"-" example:"2023-12-03T15:02:19.62113565Z"`
	TimeTo   time.Time                                                `json:"timeTo" gorm:"-" example:"2023-12-03T15:02:19.62113565Z"`
	Initial  int                                                      `json:"initial" gorm:"initial" example:"-100000"`
	Debt     int                                                      `json:"debt" gorm:"debt" example:"-40000"`
	Pay      int                                                      `json:"pay" gorm:"pay" example:"20000"`
	Final    int                                                      `json:"final" gorm:"final" example:"-120000"`
	Details  []supplierdebtreportdetailmodel.SupplierDebtReportDetail `json:"details" gorm:"foreignkey:ReportId;association_foreignkey:id"`
}

func (*SupplierDebtReport) TableName() string {
	return common.TableSupplierDebtReport
}

var (
	ErrSupplierDebtReportDateInvalid = common.NewCustomError(
		errors.New("date report is invalid"),
		"Ngày bạn chọn không hợp lệ",
		"ErrSaleReportDateInvalid",
	)
	ErrSupplierDebtReportDateIsInFuture = common.NewCustomError(
		errors.New("date report is in future"),
		"Chưa tới thời điểm báo cáo",
		"ErrSupplierDebtReportDateIsInFuture",
	)
	ErrSupplierDebtReportViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem báo cáo nợ"),
	)
)
