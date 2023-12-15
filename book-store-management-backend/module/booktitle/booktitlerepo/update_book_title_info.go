package booktitlerepo

import (
	"book-store-management-backend/module/booktitle/booktitlestore"
	"context"
)

type UpdateBookTitleRepo interface {
	UpdateBookTitle(ctx context.Context, id string, data *booktitlestore.BookTitleDBModel) error
}

type updateBookStore interface {
	UpdateBookTitle(ctx context.Context, id string, data *booktitlestore.BookTitleDBModel) error
}

type updateBookRepo struct {
	store updateBookStore
}

func NewUpdateBookRepo(store updateBookStore) *updateBookRepo {
	return &updateBookRepo{store: store}
}

func (repo *updateBookRepo) UpdateBookTitle(ctx context.Context, id string, data *booktitlestore.BookTitleDBModel) error {
	err := repo.store.UpdateBookTitle(ctx, id, data)
	return err
}
