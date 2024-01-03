package bookmodel

import "book-store-management-backend/common"

type ReqUpdateStatusBooks struct {
	BookIds  []string `json:"bookIds" gorm:"-"`
	IsActive *bool    `json:"isActive" gorm:"column:isActive;" example:"true"`
}

func (*ReqUpdateStatusBooks) TableName() string {
	return common.TableUser
}

func (data *ReqUpdateStatusBooks) Validate() error {
	if data.IsActive == nil {
		return ErrBookStatusEmpty
	}
	for _, v := range data.BookIds {
		if !common.ValidateNotNilId(&v) {
			return ErrBookIdInvalid
		}
	}
	return nil
}
