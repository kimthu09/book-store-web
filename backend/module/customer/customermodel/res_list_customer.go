package customermodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/supplier/suppliermodel/filter"
)

type ResListCustomer struct {
	// Data contains list of customer.
	Data []Customer `json:"data"`
	// Paging provides information about pagination.
	Paging common.Paging `json:"paging,omitempty"`
	// Filter contains the filter parameters used to retrieve customer.
	Filter filter.SupplierImportFilter `json:"filter,omitempty"`
}
