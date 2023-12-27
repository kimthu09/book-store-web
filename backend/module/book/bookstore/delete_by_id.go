package bookstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/book/bookmodel"
	"context"
)

func (s *sqlStore) DeleteBook(ctx context.Context, id string) error {
	db := s.db.Table(common.TableBook).Where("id = ?", id)

	numRows := db.Find(&BookDBModel{}).RowsAffected
	if numRows == 0 {
		return bookmodel.ErrBookIdInvalid
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
