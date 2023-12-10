package bookstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/book/bookmodel"
	"context"
)

func (store *sqlStore) UpdatePriceBook(
	ctx context.Context,
	id string,
	data *bookmodel.BookUpdatePrice) error {
	db := store.db

	if err := db.Table(common.TableBook).
		Where("id = ?", id).
		Updates(data).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
