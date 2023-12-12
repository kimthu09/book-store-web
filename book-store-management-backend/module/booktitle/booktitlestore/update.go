package booktitlestore

import (
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"context"
)

func (store *sqlStore) UpdateBookTitle(ctx context.Context, id string, data *BookTitleDBModel) error {
	data.ID = nil
	db := store.db.Table(data.TableName()).Where("id = ? and isActive = ?", id, "1").Updates(data)
	if err := db.Error; err != nil {
		return err
	}
	if db.RowsAffected == 0 {
		return booktitlemodel.ErrBookTitleNotFound
	}
	return nil
}
