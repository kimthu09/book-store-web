package bookmodel

type Filter struct {
	SearchKey        string   `json:"searchKey,omitempty" form:"search" example:""`
	MinSellPrice     *float64 `json:"minSellPrice,omitempty" form:"minSellPrice" example:"10"`
	MaxSellPrice     *float64 `json:"maxSellPrice,omitempty" form:"maxSellPrice" example:"1000"`
	DateFromCreateAt *int64   `json:"createdAtFrom,omitempty" form:"createdAtFrom" example:"1709500431"`
	DateToCreateAt   *int64   `json:"createdAtTo,omitempty" form:"createdAtTo" example:"1709500431"`
}
