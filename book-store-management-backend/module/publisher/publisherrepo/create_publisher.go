package publisherrepo

import (
	"book-store-management-backend/module/publisher/publishermodel"
	"context"
)

type CreatePublisher interface {
	CreatePublisher(ctx context.Context, data *publishermodel.Publisher) error
}

type createPublisherRepo struct {
	store CreatePublisher
}

func NewCreatePublisherRepo(store CreatePublisher) *createPublisherRepo {
	return &createPublisherRepo{store: store}
}

func (biz *createPublisherRepo) CreatePublisher(ctx context.Context, data *publishermodel.Publisher) error {
	if err := biz.store.CreatePublisher(ctx, data); err != nil {
		return err
	}

	return nil
}
