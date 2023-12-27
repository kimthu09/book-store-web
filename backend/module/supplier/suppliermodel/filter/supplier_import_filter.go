package filter

type SupplierImportFilter struct {
	DateFrom *int64 `json:"from,omitempty" form:"from" example:"1709500431"`
	DateTo   *int64 `json:"to,omitempty" form:"to" example:"1709500431"`
}
