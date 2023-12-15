package salereportmodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/salereportdetail/salereportetailmodel"
	"errors"
	"time"
)

type SaleReport struct {
	TimeFrom time.Time                               `json:"timeFrom" gorm:"column:timeFrom" example:"2023-12-03T15:02:19.62113565Z"`
	TimeTo   time.Time                               `json:"timeTo" gorm:"column:timeTo" example:"2023-12-03T15:02:19.62113565Z"`
	Total    int                                     `json:"total" example:"100000"`
	Details  []salereportetailmodel.SaleReportDetail `json:"details"`
}

var (
	ErrSaleReportDateInvalid = common.NewCustomError(
		errors.New("date report is invalid"),
		"Các ngày bạn chọn không hợp lệ",
		"ErrSaleReportDateInvalid",
	)
	ErrSaleReportViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem báo cáo doanh thu"),
	)
)
