package tokenprovider

import (
	"book-store-management-backend/common"
	"errors"
	"time"
)

type Token struct {
	Token   string    `json:"token" example:"token"`
	Created time.Time `json:"created" example:"2023-12-03T15:02:19.62113565Z"`
	Expiry  int       `json:"expiry" example:"2592000"`
}

type TokenPayload struct {
	UserId string `json:"userId"`
	Role   string `json:"role"`
}

type TokenPayloadEmail struct {
	Email string `json:"email"`
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ErrNotFound",
	)

	ErrEncodingToken = common.NewCustomError(errors.New("error encoding the token"),
		"error encoding the token",
		"ErrEncodingToken",
	)

	ErrInvalidToken = common.NewCustomError(errors.New("invalid token provided"),
		"invalid token provided",
		"ErrInvalidToken",
	)
)
