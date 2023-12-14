package supplierdebtmodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/common/enum"
	"time"
)

type SupplierDebt struct {
	Id           string         `json:"id" gorm:"column:id;" example:"debt001"`
	SupplierId   string         `json:"supplierId" gorm:"column:supplierId;" example:"123"`
	Quantity     float32        `json:"qty" gorm:"column:qty;" example:"-70000"`
	QuantityLeft float32        `json:"qtyLeft" gorm:"column:qtyLeft;" example:"-100000"`
	DebtType     *enum.DebtType `json:"type" gorm:"column:type;" example:"Debt"`
	CreatedBy    string         `json:"createdBy" gorm:"column:createdBy;" example:"user id"`
	CreatedAt    *time.Time     `json:"createdAt" gorm:"column:createdAt;" example:"1709500431"`
}

func (*SupplierDebt) TableName() string {
	return common.TableSupplierDebt
}