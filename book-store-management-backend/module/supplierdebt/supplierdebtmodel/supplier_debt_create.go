package supplierdebtmodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/common/enum"
)

type SupplierDebtCreate struct {
	Id         string         `json:"-" gorm:"column:id;"`
	SupplierId string         `json:"supplierId" gorm:"column:supplierId;"`
	Amount     float32        `json:"amount" gorm:"column:amount;"`
	AmountLeft float32        `json:"-" gorm:"column:amountLeft;"`
	DebtType   *enum.DebtType `json:"type" gorm:"column:type;"`
	CreateBy   string         `json:"-" gorm:"column:createBy;"`
}

func (*SupplierDebtCreate) TableName() string {
	return common.TableSupplierDebt
}

func (data *SupplierDebtCreate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.SupplierId) {
		return ErrSupplierDebtIdSupplierInvalid
	}
	if common.ValidateNotNegativeNumber(data.Amount) {
		return ErrSupplierDebtAmountIsNotNegativeNumber
	}
	if data.DebtType == nil {
		return ErrSupplierDebtTypeEmpty
	}
	return nil
}
