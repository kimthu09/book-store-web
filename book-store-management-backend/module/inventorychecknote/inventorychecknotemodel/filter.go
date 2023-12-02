package inventorychecknotemodel

type Filter struct {
	SearchKey        string  `json:"searchKey,omitempty" form:"search"`
	DateFromCreateAt *int64  `json:"createAtFrom,omitempty" form:"createAtFrom"`
	DateToCreateAt   *int64  `json:"createAtTo,omitempty" form:"createAtTo"`
	CreateBy         *string `json:"createBy,omitempty" form:"createBy"`
}
