package bookmodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"book-store-management-backend/module/publisher/publishermodel"
)

type ResDetailUnitBook struct {
	ID          string                         `json:"id" gorm:"column:id" example:"bookId"`
	Name        string                         `json:"name" gorm:"column:name" example:"Cho tui 1 vé về tuổi thơ"`
	BookTitleID string                         `json:"-" gorm:"column:booktitleid"`
	BookTitle   booktitlemodel.SimpleBookTitle `json:"bookTitle"`
	Image       string                         `json:"image" gorm:"column:imgUrl" example:"https://cdn.com/abc.jpg"`
	PublisherID string                         `json:"-" gorm:"column:publisherid"`
	Publisher   publishermodel.Publisher       `json:"publisher"`
	Edition     int                            `json:"edition" gorm:"column:edition" example:"1"`
	Quantity    int                            `json:"quantity" gorm:"column:quantity" example:"100"`
	ListedPrice int                            `json:"listedPrice" gorm:"column:listedPrice" example:"100000"`
	SellPrice   int                            `json:"sellPrice" gorm:"column:sellPrice" example:"120000"`
	ImportPrice int                            `json:"importPrice" gorm:"column:importPrice" example:"100000"`
	IsActive    bool                           `json:"isActive,omitempty" gorm:"isActive; column:isActive; default:1" example:"true"`
}

func (*ResDetailUnitBook) TableName() string {
	return common.TableBook
}
