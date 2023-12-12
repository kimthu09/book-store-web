package booktitlemodel

type Filter struct {
	SearchKey        string `json:"searchKey,omitempty" form:"search" example:""`
	DateFromCreateAt *int64 `json:"createdAtFrom,omitempty" form:"createdAtFrom" example:"1709500431"`
	DateToCreateAt   *int64 `json:"createdAtTo,omitempty" form:"createdAtTo" example:"1709500431"`
	IsActive         *bool  `json:"isActive,omitempty" form:"isActive" example:"true"`
}
