package customermodel

type Filter struct {
	SearchKey string   `json:"searchKey,omitempty" form:"search"`
	MinPoint  *float32 `json:"minPoint,omitempty" form:"minPoint"`
	MaxPoint  *float32 `json:"maxPoint,omitempty" form:"maxPoint"`
}
