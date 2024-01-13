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
	ImgUrl   string         `json:"img" gorm:"column:imgUrl" example:"https://picsum.photos/200"`
	IsActive bool           `json:"isActive" gorm:"column:isActive;" example:"true"`
}

func (u *User) GetUserId() string {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetName() string {
	return u.Name
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
		"Mã của người dùng không hợp lệ",
		"ErrUserIdInvalid",
	)
	ErrUserNameEmpty = common.NewCustomError(
		errors.New("name of user is empty"),
		"Tên của người dùng đang trống",
		"ErrUserNameEmpty",
	)
	ErrUserEmailInvalid = common.NewCustomError(
		errors.New("email is invalid"),
		"Email của người dùng không hợp lệ",
		"ErrUserEmailInvalid",
	)
	ErrUserRoleInvalid = common.NewCustomError(
		errors.New("role is invalid"),
		"Quyền của người dùng không hợp lệ",
		"ErrUserRoleInvalid",
	)
	ErrUserPhoneInvalid = common.NewCustomError(
		errors.New("phone of user is invalid"),
		"Số điện thoại của người dùng không hợp lệ",
		"ErrUserPhoneInvalid",
	)
	ErrUserImageInvalid = common.NewCustomError(
		errors.New("avatar of user is invalid"),
		"Ảnh của người dùng không hợp lệ",
		"ErrUserImageInvalid",
	)
	ErrUserEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"Email hoặc mật khẩu không hợp lệ",
		"ErrUserEmailOrPasswordInvalid",
	)
	ErrUserEmailDuplicated = common.NewCustomError(
		errors.New("email is duplicated"),
		"Email người dùng đã tồn tại",
		"ErrUserEmailDuplicated",
	)
	ErrUserSenderPassInvalid = common.NewCustomError(
		errors.New("pass of user sender is invalid"),
		"Mật khẩu người gửi không hợp lệ",
		"ErrUserSenderPassInvalid",
	)
	ErrUserUpdatedPassInvalid = common.NewCustomError(
		errors.New("pass of user is invalid"),
		"Mật khẩu mới không hợp lệ",
		"ErrUserUpdatedPassInvalid",
	)
	ErrUserInactive = common.NewCustomError(
		errors.New("user has been inactive"),
		"Người dùng đã ngừng hoạt động",
		"ErrUserInactive",
	)
	ErrUserStatusEmpty = common.NewCustomError(
		errors.New("user status is empty"),
		"Trạng thái của người dùng đang trống",
		"ErrUserStatusEmpty",
	)
	ErrUserEmailNotExist = common.NewCustomError(
		errors.New("email of user is not existed"),
		"Người dùng không tồn tại trong hệ thống",
		"ErrUserEmailNotExist",
	)
	ErrUserCreateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền tạo người dùng mới"),
	)
	ErrUserUpdateInfoNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền chỉnh sửa thông tin người dùng"),
	)
	ErrUserUpdateRoleNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền chỉnh sửa quyền người dùng"),
	)
	ErrUserUpdateStatusNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền chỉnh sửa trạng thái người dùng"),
	)
	ErrUserResetPasswordNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền làm mới mật khẩu người dùng"),
	)
	ErrUserViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem người dùng"),
	)
	ErrUserSenderPasswordWrong = common.NewCustomError(
		errors.New("password of user sender is wrong"),
		"Mật khẩu bạn nhập đã sai",
		"ErrUserSenderPasswordWrong",
	)
)
