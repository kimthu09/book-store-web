package authorrepo

import (
	"context"
)

type ExistAuthorStore interface {
	CheckExistByID(ctx context.Context, id string) (bool, error)
}

type existAuthorRepo struct {
	store ExistAuthorStore
}

func NewExistAuthorRepo(store ExistAuthorStore) *existAuthorRepo {
	return &existAuthorRepo{store: store}
}

func (repo *existAuthorRepo) IsExistAuthorId(ctx context.Context, authorId string) bool {
	isExist, err := repo.store.CheckExistByID(ctx, authorId)

	if err != nil {
		return false
	}

	return isExist
}
