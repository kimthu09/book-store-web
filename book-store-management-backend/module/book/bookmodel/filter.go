package bookmodel

type Filter struct {
	SearchKey    string  `json:"searchKey,omitempty" form:"search" example:"id, name"`
	AuthorIds    *string `json:"authorIds,omitempty" form:"authors" example:"authorId1|authorId2"`
	CategoryIds  *string `json:"categoryIds,omitempty" form:"categories" example:"categoryId1|categoryId2"`
	PublisherId  *string `json:"publisherId,omitempty" form:"publisher" example:"publisher id"`
	MinSellPrice *int    `json:"minSellPrice,omitempty" form:"minSellPrice" example:"0"`
	MaxSellPrice *int    `json:"maxSellPrice,omitempty" form:"maxSellPrice" example:"0"`
}
