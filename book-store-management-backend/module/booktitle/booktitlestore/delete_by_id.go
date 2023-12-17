package booktitlestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"context"
)

func (store *sqlStore) DeleteBookTitle(ctx context.Context, id string) error {
	db := store.db.Table(common.TableBookTitle).Where("id = ?", id)

	numRows := db.Find(&BookTitleDBModel{}).RowsAffected
	if numRows == 0 {
		return booktitlemodel.ErrBookTitleIdInvalid
	}

	db = db.Updates(map[string]interface{}{
		"isActive":  "0",
		"deletedAt": db.NowFunc(),
	})

	if db.Error != nil {
		return common.ErrDB(db.Error)
	}
	return nil
}
