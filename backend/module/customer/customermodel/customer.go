package customermodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/invoice/invoicemodel"
	"errors"
)

type Customer struct {
	Id       string                 `json:"id" gorm:"column:id;"`
	Name     string                 `json:"name" gorm:"column:name;"`
	Email    string                 `json:"email" gorm:"column:email;"`
	Phone    string                 `json:"phone" gorm:"column:phone;"`
	Point    int                    `json:"point" gorm:"column:point;"`
	Invoices []invoicemodel.Invoice `json:"invoices,omitempty"`
}

func (*Customer) TableName() string {
	return common.TableCustomer
}

var (
	ErrCustomerIdInvalid = common.NewCustomError(
		errors.New("id of customer is invalid"),
		"Mã của khách hàng không hợp lệ",
		"ErrCustomerIdInvalid",
	)

	ErrCustomerNameEmpty = common.NewCustomError(
		errors.New("name of customer is empty"),
		"Tên của khách hàng đang trống",
		"ErrCustomerNameEmpty",
	)

	ErrCustomerPhoneInvalid = common.NewCustomError(
		errors.New("phone of customer is invalid"),
		"Số điện thoại của khách hàng không hợp lệ",
		"ErrCustomerPhoneInvalid",
	)

	ErrCustomerEmailInvalid = common.NewCustomError(
		errors.New("email of customer is invalid"),
		"Email của khách hàng không hợp lệ",
		"ErrCustomerEmailInvalid",
	)

	ErrCustomerIdDuplicate = common.ErrDuplicateKey(
		errors.New("Khách hàng đã tồn tại"),
	)
	ErrCustomerPhoneDuplicate = common.ErrDuplicateKey(
		errors.New("Số điện thoại của khách hàng đã tồn tại"),
	)
	ErrCustomerCreateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền thêm khách hàng mới"),
	)
	ErrCustomerUpdateInfoNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền chỉnh sửa thông tin khách hàng"),
	)
	ErrCustomerViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem thông tin khách hàng"),
	)
)
