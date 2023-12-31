package authormodel

import "book-store-management-backend/common"

type ReqUpdateAuthor struct {
	Name *string `json:"name" gorm:"column:name;" example:"tên đã đổi"`
}

func (*ReqUpdateAuthor) TableName() string {
	return common.TableAuthor
}

func (data *ReqUpdateAuthor) Validate() *common.AppError {
	if data.Name != nil && common.ValidateEmptyString(*data.Name) {
		return ErrAuthorNameEmpty
	}
	return nil
}
