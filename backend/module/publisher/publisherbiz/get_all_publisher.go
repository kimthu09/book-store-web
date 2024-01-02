package publisherbiz

import (
	"book-store-management-backend/module/publisher/publishermodel"
	"context"
)

type GetAllPublisherRepo interface {
	GetAllPublisher(ctx context.Context) ([]publishermodel.Publisher, error)
}

type getAllPublisherBiz struct {
	repo GetAllPublisherRepo
}

func NewGetAllPublisherBiz(repo GetAllPublisherRepo) *getAllPublisherBiz {
	return &getAllPublisherBiz{repo: repo}
}

func (biz *getAllPublisherBiz) GetAllPublisher(
	ctx context.Context) ([]publishermodel.Publisher, error) {
	result, err := biz.repo.GetAllPublisher(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
