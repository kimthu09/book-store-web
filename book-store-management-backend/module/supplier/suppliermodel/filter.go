package suppliermodel

type Filter struct {
	SearchKey string   `json:"searchKey,omitempty" form:"search"`
	MinDebt   *float32 `json:"minDebt,omitempty" form:"minDebt"`
	MaxDebt   *float32 `json:"maxDebt,omitempty" form:"maxDebt"`
}
