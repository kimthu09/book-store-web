package importnotemodel

type Filter struct {
	SearchKey        string   `json:"searchKey,omitempty" form:"search" example:""`
	MinPrice         *float32 `json:"minPrice,omitempty" form:"minPrice" example:"10"`
	MaxPrice         *float32 `json:"maxPrice,omitempty" form:"maxPrice" example:"1000"`
	DateFromCreateAt *int64   `json:"createAtFrom,omitempty" form:"createAtFrom" example:"1709500431"`
	DateToCreateAt   *int64   `json:"createAtTo,omitempty" form:"createAtTo" example:"1709500431"`
	DateFromCloseAt  *int64   `json:"closeAtFrom,omitempty" form:"closeAtFrom" example:"1709500431"`
	DateToCloseAt    *int64   `json:"closeAtTo,omitempty" form:"closeAtTo" example:"1709500431"`
	Supplier         *string  `json:"supplier,omitempty" form:"supplier" example:"supplier name"`
	CreateBy         *string  `json:"createBy,omitempty" form:"createBy" example:"user name"`
	CloseBy          *string  `json:"closeBy,omitempty" form:"closeBy" example:"user name"`
	Status           string   `json:"status,omitempty" form:"status" example:"Done"`
}
