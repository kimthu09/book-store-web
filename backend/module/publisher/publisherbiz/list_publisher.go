package publisherbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/publisher/publishermodel"
	"context"
)

type ListPublisherRepo interface {
	ListPublisher(ctx context.Context, filter *publishermodel.Filter, paging *common.Paging) ([]publishermodel.Publisher, error)
}

type listPublisherBiz struct {
	repo      ListPublisherRepo
	requester middleware.Requester
}

func NewListPublisherRepo(repo ListPublisherRepo, requester middleware.Requester) *listPublisherBiz {
	return &listPublisherBiz{repo: repo, requester: requester}
}

func (biz *listPublisherBiz) ListPublisher(ctx context.Context, filter *publishermodel.Filter, paging *common.Paging) ([]publishermodel.Publisher, error) {
	if !biz.requester.IsHasFeature(common.PublisherViewFeatureCode) {
		return nil, publishermodel.ErrPublisherViewNoPermission
	}

	result, err := biz.repo.ListPublisher(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
