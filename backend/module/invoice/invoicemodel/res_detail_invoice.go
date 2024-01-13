package invoicemodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/shopgeneral/shopgeneralmodel"
)

type ResDetailInvoice struct {
	Invoice     Invoice                      `json:"invoice" gorm:"foreignkey:invoiceId;association_foreignkey:id"`
	ShopGeneral shopgeneralmodel.ShopGeneral `json:"shop" gorm:"foreignkey:invoiceId;association_foreignkey:id"`
}

func (*ResDetailInvoice) TableName() string {
	return common.TableInvoice
}
