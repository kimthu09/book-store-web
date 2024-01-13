package usermodel

type ReqMailForgotPassword struct {
	Email string `json:"email" gorm:"column:email;"`
}
