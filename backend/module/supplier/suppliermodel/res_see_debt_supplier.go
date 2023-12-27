package suppliermodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/supplier/suppliermodel/filter"
	"book-store-management-backend/module/supplierdebt/supplierdebtmodel"
)

type ResSeeDebtSupplier struct {
	// Data contains the detailed information about supplier's debts.
	Data []supplierdebtmodel.SupplierDebt `json:"data"`
	// Paging provides information about pagination.
	Paging common.Paging `json:"paging,omitempty"`
	// Filter contains the filter parameters used to retrieve debts.
	Filter filter.SupplierDebtFilter `json:"filter,omitempty"`
}
