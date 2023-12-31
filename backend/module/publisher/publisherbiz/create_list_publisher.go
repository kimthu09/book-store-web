package publisherbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/publisher/publishermodel"
	"context"
)

type createListPublisherBiz struct {
	gen       generator.IdGenerator
	repo      CreatePublisherRepo
	requester middleware.Requester
}

func NewCreateListPublisherBiz(
	gen generator.IdGenerator, repo CreatePublisherRepo, requester middleware.Requester) *createListPublisherBiz {
	return &createListPublisherBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createListPublisherBiz) CreateListPublisher(
	ctx context.Context, data []publishermodel.Publisher) error {
	if !biz.requester.IsHasFeature(common.PublisherCreateFeatureCode) {
		return publishermodel.ErrPublisherCreateNoPermission
	}

	for _, v := range data {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	for i, _ := range data {
		if err := handlePublisherId(biz.gen, &data[i]); err != nil {
			return err
		}
	}

	for _, v := range data {
		if err := biz.repo.CreatePublisher(ctx, &v); err != nil {
			return err
		}
	}

	return nil
}
