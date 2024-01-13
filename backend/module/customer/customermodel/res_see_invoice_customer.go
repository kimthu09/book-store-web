package customermodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/invoice/invoicemodel"
)

type ResSeeInvoiceCustomer struct {
	// Data contains the detailed information about customer's invoices.
	Data []invoicemodel.Invoice `json:"data"`
	// Paging provides information about pagination.
	Paging common.Paging `json:"paging,omitempty"`
	// Filter contains the filter parameters used to retrieve invoices.
	Filter FilterInvoice `json:"filter,omitempty"`
}
