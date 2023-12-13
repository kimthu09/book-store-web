package usermodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/role/rolemodel"
	"errors"
)

type User struct {
	Id       string         `json:"id" gorm:"column:id;" example:"user id"`
	Name     string         `json:"name" gorm:"column:name;" example:"Nguyễn Văn B"`
	Email    string         `json:"email" gorm:"column:email;" example:"b@gmail.com"`
	Phone    string         `json:"phone" gorm:"column:phone;" example:"0919199112"`
	Address  string         `json:"address" gorm:"column:address;" example:"HCM"`
	Password string         `json:"-" gorm:"column:password;"`
	Salt     string         `json:"-" gorm:"column:salt;"`
	RoleId   string         `json:"-" gorm:"column:roleId;"`
	Role     rolemodel.Role `json:"role" gorm:"foreignkey:roleId"`
	IsActive bool           `json:"isActive" gorm:"column:isActive;" example:"true"`
}

func (u *User) GetUserId() string {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRoleId() string {
	return u.RoleId
}

func (u *User) IsHasFeature(featureCode string) bool {
	for _, v := range u.Role.RoleFeatures {
		if v.FeatureId == featureCode {
			return true
		}
	}
	return false
}

func (*User) TableName() string {
	return common.TableUser
}

var (
	ErrUserIdInvalid = common.NewCustomError(
		errors.New("id of user is invalid"),
		"id of user is empty",
		"ErrUserIdInvalid",
	)
	ErrUserNameEmpty = common.NewCustomError(
		errors.New("name of user is empty"),
		"name of user is empty",
		"ErrUserNameEmpty",
	)
	ErrUserEmailInvalid = common.NewCustomError(
		errors.New("email is invalid"),
		"email is invalid",
		"ErrUserEmailInvalid",
	)
	ErrUserRoleInvalid = common.NewCustomError(
		errors.New("role is invalid"),
		"role is invalid",
		"ErrUserRoleInvalid",
	)
	ErrUserPhoneInvalid = common.NewCustomError(
		errors.New("phone of user is invalid"),
		"phone of user is invalid",
		"ErrUserPhoneInvalid",
	)
	ErrUserEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrUserEmailOrPasswordInvalid",
	)
	ErrUserEmailDuplicated = common.NewCustomError(
		errors.New("email is duplicated"),
		"email is duplicated",
		"ErrUserEmailDuplicated",
	)
	ErrUserSenderPassInvalid = common.NewCustomError(
		errors.New("pass of user sender is invalid"),
		"pass of user sender is invalid",
		"ErrUserSenderPassInvalid",
	)
	ErrUserUpdatedPassInvalid = common.NewCustomError(
		errors.New("pass of user is invalid"),
		"pass of user is invalid",
		"ErrUserUpdatedPassInvalid",
	)
	ErrUserInactive = common.NewCustomError(
		errors.New("user has been inactive"),
		"user has been inactive",
		"ErrUserInactive",
	)
	ErrUserStatusEmpty = common.NewCustomError(
		errors.New("user status is empty"),
		"user status is empty",
		"ErrUserStatusEmpty",
	)
	ErrUserCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create user"),
	)
	ErrUserUpdateInfoNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to update info user"),
	)
	ErrUserUpdateRoleNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to update role user"),
	)
	ErrUserUpdateStatusNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to update status user"),
	)
	ErrUserResetPasswordNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to reset password user"),
	)
	ErrUserViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view user"),
	)
	ErrUserSenderPasswordWrong = common.NewCustomError(
		errors.New("password of user sender is wrong"),
		"your password is not right",
		"ErrUserSenderPasswordWrong",
	)
)
