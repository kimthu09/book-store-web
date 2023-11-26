package rolemodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/rolefeature/rolefeaturemodel"
	"errors"
)

type Role struct {
	Id           string           `json:"id" gorm:"column:id;"`
	Name         string           `json:"name" gorm:"column:name;"`
	RoleFeatures ListRoleFeatures `json:"features"`
}

func (*Role) TableName() string {
	return common.TableRole
}

type ListRoleFeatures []rolefeaturemodel.RoleFeature

func (*ListRoleFeatures) TableName() string {
	return common.TableRoleFeature
}

var (
	ErrRoleNameEmpty = common.NewCustomError(
		errors.New("name of role is empty"),
		"name of role is empty",
		"ErrRoleNameEmpty",
	)
	ErrRoleFeaturesEmpty = common.NewCustomError(
		errors.New("features of role is empty"),
		"features of role is empty",
		"ErrRoleFeaturesEmpty",
	)
	ErrRoleFeatureInvalid = common.NewCustomError(
		errors.New("features of role is invalid"),
		"features of role is invalid",
		"ErrRoleFeatureInvalid",
	)
	ErrRoleCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create role"),
	)
	ErrRoleUpdateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to update role"),
	)
	ErrRoleViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view role"),
	)
)
