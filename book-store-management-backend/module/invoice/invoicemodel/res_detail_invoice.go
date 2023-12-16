package invoicemodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/invoicedetail/invoicedetailmodel"
	"book-store-management-backend/module/user/usermodel"
	"time"
)

type ResDetailInvoice struct {
	Id            string                             `json:"id" gorm:"column:id;" example:"123"`
	TotalPrice    float32                            `json:"totalPrice" gorm:"column:totalPrice;" example:"120000"`
	CreatedBy     string                             `json:"-" gorm:"column:createdBy;" example:"admin"`
	CreatedByUser usermodel.SimpleUser               `json:"createdBy" gorm:"foreignKey:CreatedBy"`
	CreatedAt     *time.Time                         `json:"createdAt" gorm:"column:createdAt;" example:"2023-12-03T15:02:19.62113565Z"`
	Details       []invoicedetailmodel.InvoiceDetail `json:"details"`
}

func (*ResDetailInvoice) TableName() string {
	return common.TableInvoice
}

func GetResDetailInvoiceFromInvoice(invoice *Invoice) *ResDetailInvoice {
	var src ResDetailInvoice
	src.Id = invoice.Id
	src.TotalPrice = invoice.TotalPrice
	src.CreatedBy = invoice.CreatedBy
	src.CreatedByUser = invoice.CreatedByUser
	src.CreatedAt = invoice.CreatedAt
	return &src
}
