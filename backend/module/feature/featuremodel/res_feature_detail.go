package featuremodel

type ResFeatureDetail struct {
	Id          string `json:"id" gorm:"column:id;" example:"feature id"`
	Description string `json:"description" gorm:"column:description;" example:"Xem nhân viên"`
	GroupName   string `json:"groupName" gorm:"column:groupName" example:"Nhân viên"`
	IsHas       bool   `json:"isHas" gorm:"-" example:"true"`
}
