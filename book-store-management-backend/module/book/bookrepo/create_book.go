package bookrepo

import (
	"book-store-management-backend/module/book/bookmodel"
	"context"
	"fmt"
)

type CreateBookStore interface {
	CreateBook(ctx context.Context, bookGeneral *bookmodel.Book) error
}

type createBookRepo struct {
	store CreateBookStore
}

func NewCreateBookRepo(store CreateBookStore) *createBookRepo {
	return &createBookRepo{store: store}
}

func (biz *createBookRepo) CreateBook(ctx context.Context, data *bookmodel.Book) error {
	fmt.Println("=====================================\nRepo Book\n=====================================\n")

	// if err := biz.store.CreateBook(ctx, data); err != nil {
	// 	return err
	// }

	return nil
}
