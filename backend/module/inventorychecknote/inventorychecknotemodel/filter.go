package inventorychecknotemodel

type Filter struct {
	SearchKey         string  `json:"searchKey,omitempty" form:"search" example:"note id"`
	DateFromCreatedAt *int64  `json:"createdAtFrom,omitempty" form:"createdAtFrom" example:"1709500431"`
	DateToCreatedAt   *int64  `json:"createdAtTo,omitempty" form:"createdAtTo" example:"1709500431"`
	CreatedBy         *string `json:"createdBy,omitempty" form:"createdBy" example:"user id"`
}
