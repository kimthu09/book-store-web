package dashboardmodel

import (
	"book-store-management-backend/common"
	"errors"
	"time"
)

type ReqSeeDashboard struct {
	TimeFrom int64 `json:"timeFrom" gorm:"-"`
	TimeTo   int64 `json:"timeTo" gorm:"-"`
}

func (data *ReqSeeDashboard) Validate() error {
	timeFrom := time.Unix(data.TimeFrom, 0)
	timeTo := time.Unix(data.TimeTo, 0)

	if timeFrom.After(timeTo) {
		return ErrDashboardDateInvalid
	}
	return nil
}

var (
	ErrDashboardDateInvalid = common.NewCustomError(
		errors.New("date see dashboard is invalid"),
		"Các ngày bạn chọn không hợp lệ",
		"ErrDashboardDateInvalid",
	)
	ErrDashboardViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem trang chủ"),
	)
)
