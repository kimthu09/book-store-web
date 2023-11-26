package rolemodel

import "book-store-management-backend/common"

type RoleCreate struct {
	Id       string   `json:"-" gorm:"column:id;"`
	Name     string   `json:"name" gorm:"column:name;"`
	Features []string `json:"features" gorm:"-"`
}

func (*RoleCreate) TableName() string {
	return common.TableRole
}

func (data *RoleCreate) Validate() *common.AppError {
	if common.ValidateEmptyString(data.Name) {
		return ErrRoleNameEmpty
	}
	if data.Features == nil || len(data.Features) == 0 {
		return ErrRoleFeaturesEmpty
	}
	for _, v := range data.Features {
		if !common.ValidateNotNilId(&v) {
			return ErrRoleFeatureInvalid
		}
	}
	return nil
}
