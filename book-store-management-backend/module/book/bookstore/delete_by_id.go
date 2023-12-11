package bookstore

import (
	"book-store-management-backend/common"
	"context"
)

func (store *sqlStore) DeleteBook(ctx context.Context, id string) error {
	db := store.db.Table(common.TableBook).Where("id = ?", id)

	if db.Error != nil {
		return common.ErrDB(db.Error)
	}

	db = db.Updates(map[string]interface{}{
		"isActive":  "0",
		"deletedAt": db.NowFunc(),
	})
	return db.Error
}
