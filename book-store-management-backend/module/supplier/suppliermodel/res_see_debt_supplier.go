package suppliermodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/supplier/suppliermodel/filter"
)

type ResSeeDebtSupplier struct {
	// Data contains the detailed information about supplier's debts.
	Data ResDebtSupplier `json:"data"`
	// Paging provides information about pagination.
	Paging common.Paging `json:"paging,omitempty"`
	// Filter contains the filter parameters used to retrieve debts.
	Filter filter.SupplierDebtFilter `json:"filter,omitempty"`
}
