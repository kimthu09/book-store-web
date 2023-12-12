package booktitlestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"golang.org/x/net/context"
)

func (*BookTitleDBModel) TableName() string {
	return common.TableBookTitle
}

func (store *sqlStore) CreateBookTitle(ctx context.Context, data *BookTitleDBModel) error {
	db := store.db

	if err := db.Create(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("PRIMARY"); key {
			case "PRIMARY":
				return booktitlemodel.ErrBookTitleIdDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
