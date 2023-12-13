package rolemodel

import "book-store-management-backend/common"

type ReqCreateRole struct {
	Id       string   `json:"-" gorm:"column:id;"`
	Name     string   `json:"name" gorm:"column:name;" example:"user"`
	Features []string `json:"features" gorm:"-"`
}

func (*ReqCreateRole) TableName() string {
	return common.TableRole
}

func (data *ReqCreateRole) Validate() *common.AppError {
	if common.ValidateEmptyString(data.Name) {
		return ErrRoleNameEmpty
	}
	if data.Features != nil && len(data.Features) != 0 {
		for _, v := range data.Features {
			if !common.ValidateFeatureCode(&v) {
				return ErrRoleFeatureInvalid
			}
		}
	}
	return nil
}
