package importnotemodel

type Filter struct {
	SearchKey        string   `json:"searchKey,omitempty" form:"search"`
	MinPrice         *float32 `json:"minPrice,omitempty" form:"minPrice"`
	MaxPrice         *float32 `json:"maxPrice,omitempty" form:"maxPrice"`
	DateFromCreateAt *int64   `json:"createAtFrom,omitempty" form:"createAtFrom"`
	DateToCreateAt   *int64   `json:"createAtTo,omitempty" form:"createAtTo"`
	DateFromCloseAt  *int64   `json:"closeAtFrom,omitempty" form:"closeAtFrom"`
	DateToCloseAt    *int64   `json:"closeAtTo,omitempty" form:"closeAtTo"`
	Supplier         *string  `json:"supplier,omitempty" form:"supplier"`
	CreateBy         *string  `json:"createBy,omitempty" form:"createBy"`
	CloseBy          *string  `json:"closeBy,omitempty" form:"closeBy"`
	Status           string   `json:"status,omitempty" form:"status"`
}
