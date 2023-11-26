package supplierdebtmodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/common/enum"
	"errors"
	"time"
)

type SupplierDebt struct {
	Id         string         `json:"id" gorm:"column:id;"`
	SupplierId string         `json:"supplierId" gorm:"column:supplierId;"`
	Amount     float32        `json:"amount" gorm:"column:amount;"`
	AmountLeft float32        `json:"amountLeft" gorm:"column:amountLeft;"`
	DebtType   *enum.DebtType `json:"type" gorm:"column:type;"`
	CreateBy   string         `json:"createBy" gorm:"column:createBy;"`
	CreateAt   *time.Time     `json:"createAt" gorm:"column:createAt;"`
}

func (*SupplierDebt) TableName() string {
	return common.TableSupplierDebt
}

var (
	ErrSupplierDebtIdSupplierInvalid = common.NewCustomError(
		errors.New("id of supplier is invalid"),
		"id of supplier is invalid",
		"ErrSupplierDebtIdSupplierInvalid",
	)
	ErrSupplierDebtAmountIsNotNegativeNumber = common.NewCustomError(
		errors.New("amount is not negative number"),
		"amount is not negative number",
		"ErrSupplierDebtAmountIsNotNegativeNumber",
	)
	ErrSupplierDebtTypeEmpty = common.NewCustomError(
		errors.New("debt type is empty"),
		"debt type is empty",
		"ErrSupplierDebtTypeEmpty",
	)
	ErrSupplierDebtViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view supplier"),
	)
)
