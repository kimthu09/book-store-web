package booktitlemodel

import (
	"book-store-management-backend/common"
	"time"
)

type Book struct {
	ID          *string  `json:"id" gorm:"column:id;primaryKey"`
	Name        string   `json:"name" gorm:"column:name"`
	Description string   `json:"desc" gorm:"column:desc"`
	AuthorIDs   []string `json:"authorIds" gorm:"column:authorIds"`
	CategoryIDs []string `json:"categoryIds" gorm:"column:categoryIds"`

	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"createdAt; column:createdAt;"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"updatedAt; column:updatedAt;"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"deletedAt; column:deletedAt;"`
	IsActive  *int       `json:"isActive,omitempty" gorm:"isActive; column:isActive; default:1"`
}

func (*Book) TableName() string {
	return common.TableBookTitle
}

func (data *Book) Validate() *common.AppError {
	if common.ValidateEmptyString(data.Name) {
		return ErrBookTitleNameEmpty
	}

	if len(data.AuthorIDs) == 0 {
		return ErrBookTitleAuthorIdsEmpty
	}

	if len(data.CategoryIDs) == 0 {
		return ErrBookTitleCategoryIdsEmpty
	}

	return nil
}
