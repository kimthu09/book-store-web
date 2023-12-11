package bookstore

import (
	"book-store-management-backend/module/book/bookmodel"
	"context"
)

func (store *sqlStore) UpdateBook(ctx context.Context, id string, data *BookDBModel) error {
	data.ID = nil
	db := store.db.Table(data.TableName()).Where("id = ? and isActive = ?", id, "1").Updates(data)
	if err := db.Error; err != nil {
		return err
	}
	if db.RowsAffected == 0 {
		return bookmodel.ErrBookNotFound
	}
	return nil
}
