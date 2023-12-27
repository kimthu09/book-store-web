package usermodel

import "book-store-management-backend/common"

type ReqUpdateStatusUsers struct {
	UserIds  []string `json:"userIds" gorm:"-"`
	IsActive *bool    `json:"isActive" gorm:"column:isActive;" example:"true"`
}

func (*ReqUpdateStatusUsers) TableName() string {
	return common.TableUser
}

func (data *ReqUpdateStatusUsers) Validate() error {
	if data.IsActive == nil {
		return ErrUserStatusEmpty
	}
	for _, v := range data.UserIds {
		if !common.ValidateNotNilId(&v) {
			return ErrUserIdInvalid
		}
	}
	return nil
}
