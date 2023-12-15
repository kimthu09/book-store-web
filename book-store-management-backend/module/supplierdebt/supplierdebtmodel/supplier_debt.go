package supplierdebtmodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/common/enum"
	"book-store-management-backend/module/user/usermodel"
	"time"
)

type SupplierDebt struct {
	Id            string               `json:"id" gorm:"column:id;" example:"debt001"`
	SupplierId    string               `json:"supplierId" gorm:"column:supplierId;" example:"123"`
	Quantity      int                  `json:"qty" gorm:"column:qty;" example:"-70000"`
	QuantityLeft  int                  `json:"qtyLeft" gorm:"column:qtyLeft;" example:"-100000"`
	DebtType      *enum.DebtType       `json:"type" gorm:"column:type;" example:"Debt"`
	CreatedBy     string               `json:"-" gorm:"column:createdBy;"`
	CreatedByUser usermodel.SimpleUser `json:"createdBy" gorm:"foreignKey:CreatedBy"`
	CreatedAt     *time.Time           `json:"createdAt" gorm:"column:createdAt;" example:"1709500431"`
}

func (*SupplierDebt) TableName() string {
	return common.TableSupplierDebt
}
