package publisherbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/publisher/publishermodel"
	"context"
)

type CreatePublisherRepo interface {
	CreatePublisher(ctx context.Context, data *publishermodel.Publisher) error
}

type createPublisherBiz struct {
	gen       generator.IdGenerator
	repo      CreatePublisherRepo
	requester middleware.Requester
}

func NewCreatePublisherBiz(gen generator.IdGenerator, repo CreatePublisherRepo, requester middleware.Requester) *createPublisherBiz {
	return &createPublisherBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createPublisherBiz) CreatePublisher(ctx context.Context, data *publishermodel.Publisher) error {
	if !biz.requester.IsHasFeature(common.PublisherCreateFeatureCode) {
		return publishermodel.ErrPublisherCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := handlePublisherId(biz.gen, data); err != nil {
		return err
	}
	if err := biz.repo.CreatePublisher(ctx, data); err != nil {
		return err
	}

	return nil
}

func handlePublisherId(gen generator.IdGenerator, data *publishermodel.Publisher) error {
	id, err := gen.IdProcess(&data.Id)
	if err != nil {
		return err
	}
	data.Id = *id
	return nil
}
