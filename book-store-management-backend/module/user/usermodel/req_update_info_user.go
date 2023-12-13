package usermodel

import "book-store-management-backend/common"

type ReqUpdateInfoUser struct {
	Name    *string `json:"name" gorm:"column:name;" example:"bỏ trường này nếu không muốn update tên"`
	Phone   *string `json:"phone" gorm:"column:phone;" example:"bỏ trường này nếu không muốn update sđt"`
	Address *string `json:"address" gorm:"column:address;" example:"bỏ trường này nếu không muốn update địa chỉ"`
}

func (*ReqUpdateInfoUser) TableName() string {
	return common.TableUser
}

func (data *ReqUpdateInfoUser) Validate() error {
	if data.Name != nil && common.ValidateEmptyString(*data.Name) {
		return ErrUserNameEmpty
	}
	if data.Phone != nil && len(*data.Phone) != 0 && !common.ValidatePhone(*data.Phone) {
		return ErrUserPhoneInvalid
	}
	return nil
}
