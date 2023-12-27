package usermodel

type Filter struct {
	SearchKey string `json:"searchKey,omitempty" form:"search" example:"id, name, phone, email, address"`
	IsActive  *bool  `json:"active,omitempty" form:"active" example:"true"`
}
