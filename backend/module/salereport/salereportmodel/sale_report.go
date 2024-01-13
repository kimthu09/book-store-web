package salereportmodel

import (
	"book-store-management-backend/common"
	"errors"
	"time"
)

type SaleReport struct {
	TimeFrom time.Time `json:"timeFrom" gorm:"-" example:"2023-12-03T15:02:19.62113565Z"`
	TimeTo   time.Time `json:"timeTo" gorm:"-" example:"2023-12-03T15:02:19.62113565Z"`
	Total    int       `json:"total" gorm:"-" example:"100000"`
	Amount   int       `json:"amount" gorm:"-" example:"10"`
	Details  Details   `json:"details"`
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
