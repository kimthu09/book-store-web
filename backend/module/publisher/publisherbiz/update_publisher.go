package publisherbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/publisher/publishermodel"
	"context"
)

type UpdatePublisherRepo interface {
	UpdatePublisherInfo(
		ctx context.Context,
		supplierId string,
		data *publishermodel.ReqUpdatePublisher) error
}

type updatePublisherBiz struct {
	repo      UpdatePublisherRepo
	requester middleware.Requester
}

func NewUpdatePublisherBiz(
	repo UpdatePublisherRepo,
	requester middleware.Requester) *updatePublisherBiz {
	return &updatePublisherBiz{repo: repo, requester: requester}
}

func (biz *updatePublisherBiz) UpdatePublisher(
	ctx context.Context,
	id string,
	data *publishermodel.ReqUpdatePublisher) error {
	if !biz.requester.IsHasFeature(common.PublisherUpdateFeatureCode) {
		return publishermodel.ErrPublisherUpdateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.repo.UpdatePublisherInfo(ctx, id, data); err != nil {
		return err
	}

	return nil
}
