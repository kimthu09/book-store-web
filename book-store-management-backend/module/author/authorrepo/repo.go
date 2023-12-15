package authorrepo

import (
	"book-store-management-backend/module/author/authormodel"
	"context"
)

type AuthorStore interface {
	CheckExistByID(ctx context.Context, id string) (bool, error)
	GetByListId(ctx context.Context, idList []string) ([]authormodel.Author, error)
}

type AuthorPublicRepo interface {
	GetByListId(ctx context.Context, ids []string) ([]authormodel.Author, error)
	IsExistAuthorId(ctx context.Context, authorId string) bool
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

func (repo *authorPublicRepo) IsExistAuthorId(ctx context.Context, authorId string) bool {
	isExist, err := repo.store.CheckExistByID(ctx, authorId)

	if err != nil {
		return false
	}

	return isExist
}
