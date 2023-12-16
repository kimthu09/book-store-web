package invoicemodel

import (
	"book-store-management-backend/module/invoicedetail/invoicedetailmodel"
	"book-store-management-backend/module/user/usermodel"
)

type ResCreateInvoice struct {
	ResCreateInvoiceData ResCreateInvoiceData `json:"data"`
}

type ResCreateInvoiceData struct {
	Id        string                                      `json:"id" example:"123"`
	Details   []invoicedetailmodel.ReqCreateInvoiceDetail `json:"details"`
	Total     int                                         `json:"total" example:"120000"`
	CreatedBy usermodel.SimpleUser                        `json:"createdBy"`
}
