package authorrepo

import (
	"book-store-management-backend/module/author/authormodel"
	"context"
)

type GetAllAuthorStore interface {
	GetAllAuthor(ctx context.Context) ([]authormodel.SimpleAuthor, error)
}

type getAllAuthorStore struct {
	store GetAllAuthorStore
}

func NewGetAllAuthorRepo(store GetAllAuthorStore) *getAllAuthorStore {
	return &getAllAuthorStore{store: store}
}

func (repo *getAllAuthorStore) GetAllAuthor(
	ctx context.Context) ([]authormodel.SimpleAuthor, error) {
	result, err := repo.store.GetAllAuthor(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
