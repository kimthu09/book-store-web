package booktitlestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"context"
)

func (store *sqlStore) DeleteBook(ctx context.Context, id string) error {
	db := store.db.Table(common.TableBookTitle).Where("id = ?", id)

	if db.Error != nil {
		return common.ErrDB(db.Error)
	}

	db = db.Updates(map[string]interface{}{
		"isActive":  "0",
		"deletedAt": db.NowFunc(),
	})

	if db.Error != nil {
		return common.ErrDB(db.Error)
	}
	if db.RowsAffected == 0 {
		return booktitlemodel.ErrBookTitleNotFound
	}
	return db.Error
}
