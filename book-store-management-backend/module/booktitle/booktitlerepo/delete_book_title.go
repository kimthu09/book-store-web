package booktitlerepo

import (
	"context"
)

type DeleteBookTileRepo interface {
	DeleteBookTitle(ctx context.Context, id string) error
}

type DeleteBookStore interface {
	DeleteBookTitle(ctx context.Context, id string) error
}

type deleteBookRepo struct {
	store DeleteBookStore
}

func NewDeleteBookTitleRepo(store DeleteBookStore) *deleteBookRepo {
	return &deleteBookRepo{store: store}
}

func (repo *deleteBookRepo) DeleteBookTitle(ctx context.Context, id string) error {
	err := repo.store.DeleteBookTitle(ctx, id)
	return err
}
