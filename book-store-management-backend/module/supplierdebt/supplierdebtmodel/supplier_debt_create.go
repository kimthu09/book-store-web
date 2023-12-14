package supplierdebtmodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/common/enum"
)

type SupplierDebtCreate struct {
	Id           string         `json:"-" gorm:"column:id;"`
	SupplierId   string         `json:"supplierId" gorm:"column:supplierId;"`
	Quantity     int            `json:"qty" gorm:"column:qty;"`
	QuantityLeft int            `json:"-" gorm:"column:qtyLeft;"`
	DebtType     *enum.DebtType `json:"type" gorm:"column:type;"`
	CreatedBy    string         `json:"-" gorm:"column:createdBy;"`
}

func (*SupplierDebtCreate) TableName() string {
	return common.TableSupplierDebt
}
