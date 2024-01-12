package shopgeneralmodel

import (
	"book-store-management-backend/common"
	"errors"
)

type ShopGeneral struct {
	Name                   string  `json:"name" gorm:"column:name;"`
	Email                  string  `json:"email" gorm:"column:email;"`
	Phone                  string  `json:"phone" gorm:"column:phone;"`
	Address                string  `json:"address" gorm:"column:address;"`
	WifiPass               string  `json:"wifiPass" gorm:"column:wifiPass;"`
	AccumulatePointPercent float32 `json:"accumulatePointPercent" gorm:"column:accumulatePointPercent"`
	UsePointPercent        float32 `json:"usePointPercent" gorm:"column:usePointPercent"`
}

func (*ShopGeneral) TableName() string {
	return common.TableShopGeneral
}

var (
	ErrEmailInvalid = common.NewCustomError(
		errors.New("email of shop is invalid"),
		"Email cửa hàng không hợp lệ",
		"ErrEmailInvalid",
	)
	ErrPhoneInvalid = common.NewCustomError(
		errors.New("phone of shop is invalid"),
		"Số điện thoại cửa hàng không hợp lệ",
		"ErrPhoneInvalid",
	)
	ErrAccumulatePointPercentInvalid = common.NewCustomError(
		errors.New("accumulate point percent of shop is invalid"),
		"Tỷ lệ tích điểm không hợp lệ",
		"ErrAccumulatePointPercentInvalid",
	)
	ErrUsePointPercentInvalid = common.NewCustomError(
		errors.New("use point percent of shop is invalid"),
		"Tỷ lệ dùng điểm không hợp lệ",
		"ErrUsePointPercentInvalid",
	)
	ErrGeneralShopViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn có quyền xem thông tin cửa hàng"),
	)
	ErrGeneralShopUpdateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền chỉnh sửa thông tin cửa hàng"),
	)
)
