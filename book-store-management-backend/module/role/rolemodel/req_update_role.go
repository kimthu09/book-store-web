package rolemodel

import "book-store-management-backend/common"

type ReqUpdateRole struct {
	Name     *string   `json:"name" gorm:"column:name;" example:"nếu tên không đổi thì không cần trường này"`
	Features *[]string `json:"features" gorm:"-"`
}

func (*ReqUpdateRole) TableName() string {
	return common.TableRole
}

func (data *ReqUpdateRole) Validate() *common.AppError {
	if data.Name != nil && common.ValidateEmptyString(*data.Name) {
		return ErrRoleNameEmpty
	}
	if data.Features != nil {
		if len(*data.Features) == 0 {
			return ErrRoleFeaturesEmpty
		}
		for _, v := range *data.Features {
			if !common.ValidateFeatureCode(&v) {
				return ErrRoleFeatureInvalid
			}
		}
	}
	return nil
}
