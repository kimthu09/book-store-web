package bookstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/book/bookmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) UpdateAmountBook(
	ctx context.Context,
	id string,
	data *bookmodel.BookUpdateAmount) error {
	db := s.db

	if err := db.Table(common.TableBook).
		Where("id = ?", id).
		Update("amount", gorm.Expr("amount + ?", data.AmountUpdate)).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
