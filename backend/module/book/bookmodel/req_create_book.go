package bookmodel

import (
	"book-store-management-backend/common"
)

type ReqCreateBook struct {
	ID          *string `json:"id" gorm:"column:id"`
	Image       *string `json:"image" gorm:"column:imgUrl"`
	BookTitleID *string `json:"bookTitleId" gorm:"column:booktitleid,fk"`
	PublisherID *string `json:"publisherId" gorm:"column:publisherid,fk"`
	Edition     *int    `json:"edition" gorm:"column:edition"`
	ListedPrice *int    `json:"listedPrice" gorm:"column:listedPrice"`
	SellPrice   *int    `json:"sellPrice" gorm:"column:sellPrice"`
}

func (*ReqCreateBook) TableName() string {
	return common.TableBook
}

func (b *ReqCreateBook) Validate() error {
	if b.ID != nil && common.ValidateEmptyString(*b.ID) {
		b.ID = nil
	}
	if b.BookTitleID == nil || common.ValidateEmptyString(*b.BookTitleID) {
		return ErrBookTitleIdInvalid
	}
	if b.PublisherID == nil || common.ValidateEmptyString(*b.PublisherID) {
		return ErrPublisherIdInvalid
	}
	if b.Edition == nil || !common.ValidatePositiveNumber(*b.Edition) {
		return ErrBookEditionInvalid
	}
	if b.ListedPrice == nil || !common.ValidatePositiveNumber(*b.ListedPrice) {
		return ErrBookListedPriceInvalid
	}
	if b.SellPrice == nil || !common.ValidatePositiveNumber(*b.SellPrice) {
		return ErrBookSellPriceInvalid
	}
	if b.Image == nil || !common.ValidateImage(b.Image, common.DefaultImageBook) {
		return ErrBookImageInvalid
	}
	return nil
}
