package bookmodel

type Filter struct {
	SearchKey        string   `json:"searchKey,omitempty" form:"search" example:""`
	MinPrice         *float32 `json:"minPrice,omitempty" form:"minPrice" example:"10"`
	MaxPrice         *float32 `json:"maxPrice,omitempty" form:"maxPrice" example:"1000"`
	DateFromCreateAt *int64   `json:"createdAtFrom,omitempty" form:"createdAtFrom" example:"1709500431"`
	DateToCreateAt   *int64   `json:"createdAtTo,omitempty" form:"createdAtTo" example:"1709500431"`
}
