package usermodel

type ResGetAllUser struct {
	// Data contains list of user.
	Data []SimpleUser `json:"data"`
}
