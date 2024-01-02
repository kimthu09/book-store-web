package publishermodel

import "book-store-management-backend/common"

type ReqUpdatePublisher struct {
	Name *string `json:"name" gorm:"column:name;" example:"tên đã đổi"`
}

func (*ReqUpdatePublisher) TableName() string {
	return common.TablePublisher
}

func (data *ReqUpdatePublisher) Validate() *common.AppError {
	if data.Name != nil && common.ValidateEmptyString(*data.Name) {
		return ErrPublisherNameEmpty
	}
	return nil
}
