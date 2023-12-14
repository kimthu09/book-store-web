package authorbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/author/authormodel"
	"context"
)

type ListAuthorRepo interface {
	ListAuthor(ctx context.Context, filter *authormodel.Filter, paging *common.Paging) ([]authormodel.Author, error)
}

type listAuthorBiz struct {
	repo ListAuthorRepo
}

func NewListAuthorRepo(repo ListAuthorRepo) *listAuthorBiz {
	return &listAuthorBiz{repo: repo}
}

func (biz *listAuthorBiz) ListAuthor(ctx context.Context, filter *authormodel.Filter, paging *common.Paging) ([]authormodel.Author, error) {
	result, err := biz.repo.ListAuthor(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
