package publisherrepo

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/publisher/publishermodel"
	"context"
)

type ListPublisherStore interface {
	ListPublisher(ctx context.Context, filter *publishermodel.Filter, propertiesContainSearchKey []string, paging *common.Paging) ([]publishermodel.Publisher, error)
}

type listPublisherRepo struct {
	store ListPublisherStore
}

func NewListPublisherRepo(store ListPublisherStore) *listPublisherRepo {
	return &listPublisherRepo{store: store}
}

func (repo *listPublisherRepo) ListPublisher(ctx context.Context, filter *publishermodel.Filter, paging *common.Paging) ([]publishermodel.Publisher, error) {
	result, err := repo.store.ListPublisher(ctx, filter, []string{"name"}, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
