package publisherrepo

import (
	"context"
)

type ExistPublisherStore interface {
	CheckExistByID(ctx context.Context, id string) (bool, error)
}

type existPublisherRepo struct {
	store ExistPublisherStore
}

func NewExistPublisherRepo(store ExistPublisherStore) *existPublisherRepo {
	return &existPublisherRepo{store: store}
}

func (repo *existPublisherRepo) IsExistPublisherId(ctx context.Context, publisherId string) bool {
	isExist, err := repo.store.CheckExistByID(ctx, publisherId)

	if err != nil {
		return false
	}

	return isExist
}
