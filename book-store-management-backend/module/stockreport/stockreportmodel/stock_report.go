package stockreportmodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/stockreportdetail/stockreportdetailmodel"
	"errors"
	"time"
)

type StockReport struct {
	Id       string                                     `json:"id" gorm:"column:id;" example:"report id"`
	TimeFrom time.Time                                  `json:"timeFrom" gorm:"column:timeFrom" example:"2023-12-03T15:02:19.62113565Z"`
	TimeTo   time.Time                                  `json:"timeTo" gorm:"column:timeTo" example:"2023-12-03T15:02:19.62113565Z"`
	Details  []stockreportdetailmodel.StockReportDetail `json:"details" gorm:"foreignkey:ReportId;association_foreignkey:id"`
}

func (*StockReport) TableName() string {
	return common.TableStockReport
}

var (
	ErrStockReportDateInvalid = common.NewCustomError(
		errors.New("date report is invalid"),
		"Ngày bạn chọn không hợp lệ",
		"ErrSaleReportDateInvalid",
	)
	ErrStockReportFutureDateInvalid = common.NewCustomError(
		errors.New("date report is in future"),
		"Chưa tới thời điểm báo cáo",
		"ErrStockReportFutureDateInvalid",
	)
	ErrStockReportViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem báo cáo tồn kho"),
	)
)
