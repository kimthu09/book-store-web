package suppliermodel

import (
	"book-store-management-backend/common"
	"errors"
)

type Supplier struct {
	Id    string `json:"id" gorm:"column:id;" example:"123"`
	Name  string `json:"name" gorm:"column:name;" example:"Nguyễn Văn A"`
	Email string `json:"email" gorm:"column:email;" example:"a@gmail.com"`
	Phone string `json:"phone" gorm:"column:phone;" example:"0123456789"`
	Debt  int    `json:"debt" gorm:"column:debt;" example:"-100000"`
}

func (*Supplier) TableName() string {
	return common.TableSupplier
}

var (
	ErrSupplierIdInvalid = common.NewCustomError(
		errors.New("id of supplier is invalid"),
		"id of supplier is invalid",
		"ErrSupplierIdInvalid",
	)
	ErrSupplierNameEmpty = common.NewCustomError(
		errors.New("name of supplier is empty"),
		"name of supplier is empty",
		"ErrSupplierNameEmpty",
	)
	ErrSupplierPhoneInvalid = common.NewCustomError(
		errors.New("phone of supplier is invalid"),
		"phone of supplier is invalid",
		"ErrSupplierPhoneInvalid",
	)
	ErrSupplierEmailInvalid = common.NewCustomError(
		errors.New("email of supplier is invalid"),
		"email of supplier is invalid",
		"ErrSupplierEmailInvalid",
	)
	ErrSupplierInitDebtInvalid = common.NewCustomError(
		errors.New("init debt of supplier is invalid"),
		"init debt of supplier is invalid",
		"ErrSupplierInitDebtInvalid",
	)
	ErrSupplierDebtPayNotExist = common.NewCustomError(
		errors.New("debt pay is not exist"),
		"debt pay is not exist",
		"ErrSupplierDebtPayNotExist",
	)
	ErrSupplierDebtPayIsInvalid = common.NewCustomError(
		errors.New("debt pay is invalid"),
		"debt pay is invalid",
		"ErrSupplierDebtPayIsInvalid",
	)
	ErrSupplierIdDuplicate = common.ErrDuplicateKey(
		errors.New("id of supplier is duplicate"),
	)
	ErrSupplierPhoneDuplicate = common.ErrDuplicateKey(
		errors.New("phone of supplier is duplicate"),
	)
	ErrSupplierCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create supplier"),
	)
	ErrSupplierPayNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to pay supplier"),
	)
	ErrSupplierUpdateInfoNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to update info supplier"),
	)
	ErrSupplierViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view supplier"),
	)
)
