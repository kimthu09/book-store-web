package invoicemodel

type Filter struct {
	SearchKey string   `json:"searchKey,omitempty" form:"search" example:"invoice id"`
	MinPrice  *float32 `json:"minPrice,omitempty" form:"minPrice"`
	MaxPrice  *float32 `json:"maxPrice,omitempty" form:"maxPrice"`
	CreatedBy *string  `json:"createdBy,omitempty" form:"createdBy" example:"user id"`
}
