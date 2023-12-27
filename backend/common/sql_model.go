package common

import (
	"time"
)

type SQLModel struct {
	CreatedAt *time.Time `json:"createdAt" gorm:"createdAt; column:createdAt;"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"updatedAt; column:updatedAt;"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"deletedAt; column:deletedAt;"`
	IsActive  *bool      `json:"isActive" gorm:"isActive; column:isActive; default:1"`
}
