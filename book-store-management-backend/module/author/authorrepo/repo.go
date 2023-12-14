package authorrepo

import (
	"book-store-management-backend/module/author/authormodel"
	"context"
)

type AuthorStore interface {
	GetByListId(ctx context.Context, idList []string) ([]authormodel.Author, error)
}

type AuthorPublicRepo interface {
	GetByListId(ctx context.Context, ids []string) ([]authormodel.Author, error)
}

type authorPublicRepo struct {
	store AuthorStore
}

func NewAuthorPublicRepo(store AuthorStore) *authorPublicRepo {
	return &authorPublicRepo{store: store}
}

func (repo *authorPublicRepo) GetByListId(ctx context.Context, ids []string) ([]authormodel.Author, error) {
	result, err := repo.store.GetByListId(ctx, ids)
	if err != nil {
		return nil, err
	}
	return result, nil
}
