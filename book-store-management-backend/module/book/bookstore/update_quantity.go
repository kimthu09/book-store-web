package bookstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/book/bookmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) UpdateQuantityBook(
	ctx context.Context,
	id string,
	data *bookmodel.BookUpdateQuantity) error {
	db := s.db

	if err := db.Table(common.TableBook).
		Where("id = ?", id).
		Update("qty", gorm.Expr("qty + ?", data.QuantityUpdate)).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
