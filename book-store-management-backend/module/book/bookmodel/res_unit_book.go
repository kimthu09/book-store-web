package bookmodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"book-store-management-backend/module/publisher/publishermodel"
)

type ResUnitBook struct {
	ID          string                         `json:"id" gorm:"column:id" example:"bookId"`
	Name        string                         `json:"name" gorm:"column:name" example:"Cho tui 1 vé về tuổi thơ"`
	BookTitleID string                         `json:"-" gorm:"column:booktitleid"`
	BookTitle   booktitlemodel.SimpleBookTitle `json:"bookTitle"`
	PublisherID string                         `json:"-" gorm:"column:publisherid"`
	Publisher   publishermodel.Publisher       `json:"publisher"`
	Edition     int                            `json:"edition" gorm:"column:edition" example:"1"`
	Quantity    int                            `json:"quantity" gorm:"column:quantity" example:"100"`
	ListedPrice int                            `json:"listedPrice" gorm:"column:listedPrice" example:"100000"`
	SellPrice   int                            `json:"sellPrice" gorm:"column:sellPrice" example:"120000"`
	ImportPrice int                            `json:"importPrice" gorm:"column:importPrice" example:"100000"`
	ImgUrl      string                         `json:"img,omitempty" gorm:"column:imgUrl" example:"https://picsum.photos/200"`
}

func (*ResUnitBook) TableName() string {
	return common.TableBook
}
