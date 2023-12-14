package importnotemodel

type Filter struct {
	SearchKey         string   `json:"searchKey,omitempty" form:"search" example:""`
	MinPrice          *float32 `json:"minPrice,omitempty" form:"minPrice" example:"10"`
	MaxPrice          *float32 `json:"maxPrice,omitempty" form:"maxPrice" example:"1000"`
	DateFromCreatedAt *int64   `json:"createdAtFrom,omitempty" form:"createdAtFrom" example:"1709500431"`
	DateToCreatedAt   *int64   `json:"createdAtTo,omitempty" form:"createdAtTo" example:"1709500431"`
	DateFromClosedAt  *int64   `json:"closedAtFrom,omitempty" form:"closedAtFrom" example:"1709500431"`
	DateToClosedAt    *int64   `json:"closedAtTo,omitempty" form:"closedAtTo" example:"1709500431"`
	Supplier          *string  `json:"supplier,omitempty" form:"supplier" example:"supplier name"`
	CreatedBy         *string  `json:"createdBy,omitempty" form:"createdBy" example:"user name"`
	ClosedBy          *string  `json:"closedBy,omitempty" form:"closedBy" example:"user name"`
	Status            string   `json:"status,omitempty" form:"status" example:"Done"`
}
