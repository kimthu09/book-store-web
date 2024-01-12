package customermodel

import (
	"book-store-management-backend/module/invoice/invoicemodel"
)

type ResGetAllCustomer struct {
	// Data contains list of customers.
	Data []invoicemodel.SimpleCustomer `json:"data"`
}
