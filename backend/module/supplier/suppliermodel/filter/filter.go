package filter

type Filter struct {
	SearchKey string   `json:"searchKey,omitempty" form:"search" example:"id, name, phone, email"`
	MinDebt   *float32 `json:"minDebt,omitempty" form:"minDebt" example:"10"`
	MaxDebt   *float32 `json:"maxDebt,omitempty" form:"maxDebt" example:"100000"`
}
