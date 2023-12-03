package inventorychecknotemodel

type Filter struct {
	SearchKey        string  `json:"searchKey,omitempty" form:"search" example:""`
	DateFromCreateAt *int64  `json:"createAtFrom,omitempty" form:"createAtFrom" example:"1709500431"`
	DateToCreateAt   *int64  `json:"createAtTo,omitempty" form:"createAtTo" example:"1709500431"`
	CreateBy         *string `json:"createBy,omitempty" form:"createBy" example:"user name"`
}
