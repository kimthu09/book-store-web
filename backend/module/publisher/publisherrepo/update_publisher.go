package publisherrepo

import (
	"book-store-management-backend/module/publisher/publishermodel"
	"context"
)

type UpdatePublisherStore interface {
	UpdatePublisher(
		ctx context.Context,
		id string,
		data *publishermodel.ReqUpdatePublisher) error
}

type updatePublisherRepo struct {
	store UpdatePublisherStore
}

func NewUpdatePublisherRepo(store UpdatePublisherStore) *updatePublisherRepo {
	return &updatePublisherRepo{store: store}
}

func (repo *updatePublisherRepo) UpdatePublisherInfo(
	ctx context.Context,
	supplierId string,
	data *publishermodel.ReqUpdatePublisher) error {
	if err := repo.store.UpdatePublisher(ctx, supplierId, data); err != nil {
		return err
	}
	return nil
}
