package categorymodel

import "book-store-management-backend/common"

type ReqUpdateCategory struct {
	Name *string `json:"name" gorm:"column:name;" example:"tên đã đổi"`
}

func (*ReqUpdateCategory) TableName() string {
	return common.TableCategory
}

func (data *ReqUpdateCategory) Validate() *common.AppError {
	if data.Name != nil && common.ValidateEmptyString(*data.Name) {
		return ErrCategoryNameEmpty
	}
	return nil
}
