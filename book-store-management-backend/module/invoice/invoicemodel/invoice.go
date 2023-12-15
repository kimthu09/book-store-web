package invoicemodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/invoicedetail/invoicedetailmodel"
	"book-store-management-backend/module/user/usermodel"
	"errors"
	"time"
)

type Invoice struct {
	Id            string                             `json:"id" gorm:"column:id;" example:"123"`
	TotalPrice    float32                            `json:"totalPrice" gorm:"column:totalPrice;" example:"120000"`
	CreatedBy     string                             `json:"-" gorm:"column:createdBy;" example:"admin"`
	CreatedByUser usermodel.SimpleUser               `json:"createdBy" gorm:"foreignKey:CreatedBy"`
	CreatedAt     *time.Time                         `json:"createdAt" gorm:"column:createdAt;" example:"2023-12-03T15:02:19.62113565Z"`
	Details       []invoicedetailmodel.InvoiceDetail `json:"-"`
}

func (*Invoice) TableName() string {
	return common.TableInvoice
}

var (
	ErrInvoiceDetailsEmpty = common.NewCustomError(
		errors.New("list import note details are empty"),
		"Danh sách sản phẩm cần thanh toán đang trống",
		"ErrInvoiceDetailsEmpty",
	)
	ErrInvoiceBookIsNotEnough = common.NewCustomError(
		errors.New("exist book in the stock is not enough for the invoice"),
		"Tồn tại 1 sách có số lượng trong kho đủ để bán",
		"ErrInvoiceBookIsNotEnough",
	)
	ErrInvoiceCreateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền tạo hóa đơn"),
	)
	ErrInvoiceViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem hóa đơn"),
	)
)
