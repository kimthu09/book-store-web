package invoicemodel

import (
	"book-store-management-backend/common"
)

type ResListInvoice struct {
	// Data contains list of invoice.
	Data []Invoice `json:"data"`
	// Paging provides information about pagination.
	Paging common.Paging `json:"paging,omitempty"`
	// Filter contains the filter parameters used to retrieve invoice.
	Filter Filter `json:"filter,omitempty"`
}
