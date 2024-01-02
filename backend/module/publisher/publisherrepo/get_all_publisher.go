package publisherrepo

import (
	"book-store-management-backend/module/publisher/publishermodel"
	"context"
)

type GetAllPublisherStore interface {
	GetAllPublisher(ctx context.Context) ([]publishermodel.Publisher, error)
}

type getAllPublisherStore struct {
	store GetAllPublisherStore
}

func NewGetAllPublisherRepo(store GetAllPublisherStore) *getAllPublisherStore {
	return &getAllPublisherStore{store: store}
}

func (repo *getAllPublisherStore) GetAllPublisher(
	ctx context.Context) ([]publishermodel.Publisher, error) {
	result, err := repo.store.GetAllPublisher(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
