package supplierdebtmodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/common/enum"
)

type SupplierDebtCreate struct {
	Id           string         `json:"-" gorm:"column:id;"`
	SupplierId   string         `json:"supplierId" gorm:"column:supplierId;"`
	Quantity     float32        `json:"qty" gorm:"column:qty;"`
	QuantityLeft float32        `json:"-" gorm:"column:qtyLeft;"`
	DebtType     *enum.DebtType `json:"type" gorm:"column:type;"`
	CreateBy     string         `json:"-" gorm:"column:createBy;"`
}

func (*SupplierDebtCreate) TableName() string {
	return common.TableSupplierDebt
}
