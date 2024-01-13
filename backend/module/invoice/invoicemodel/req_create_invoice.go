package invoicemodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/invoicedetail/invoicedetailmodel"
)

type ReqCreateInvoice struct {
	Id                  string                                      `json:"-" gorm:"column:id;"`
	CustomerId          string                                      `json:"customerId" gorm:"column:customerId;"`
	IsUsePoint          bool                                        `json:"isUsePoint" gorm:"-"`
	TotalPrice          int                                         `json:"-" gorm:"column:totalPrice;"`
	TotalImportPrice    int                                         `json:"-" gorm:"column:totalImportPrice;"`
	AmountReceived      int                                         `json:"-" gorm:"column:amountReceived"`
	AmountPriceUsePoint int                                         `json:"-" gorm:"column:amountPriceUsePoint"`
	PointUse            int                                         `json:"-" gorm:"column:pointUse;"`
	PointReceive        int                                         `json:"-" gorm:"column:pointReceive;"`
	CreatedBy           string                                      `json:"-" gorm:"column:createdBy;"`
	InvoiceDetails      []invoicedetailmodel.ReqCreateInvoiceDetail `json:"details" gorm:"-"`
}

func (*ReqCreateInvoice) TableName() string {
	return common.TableInvoice
}

func (data *ReqCreateInvoice) Validate() *common.AppError {
	if data.InvoiceDetails == nil || len(data.InvoiceDetails) == 0 {
		return ErrInvoiceDetailsEmpty
	}

	for _, invoiceDetail := range data.InvoiceDetails {
		if err := invoiceDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}
