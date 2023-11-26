package authorbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/author/authormodel"
	"context"
)

type ListAuthorRepo interface {
	ListAuthor(ctx context.Context, filter *authormodel.Filter, paging *common.Paging) ([]authormodel.Author, error)
}

type listAuthorBiz struct {
	repo      ListAuthorRepo
	requester middleware.Requester
}

func NewListAuthorRepo(repo ListAuthorRepo, requester middleware.Requester) *listAuthorBiz {
	return &listAuthorBiz{repo: repo, requester: requester}
}

func (biz *listAuthorBiz) ListAuthor(ctx context.Context, filter *authormodel.Filter, paging *common.Paging) ([]authormodel.Author, error) {
	if !biz.requester.IsHasFeature(common.AuthorViewFeatureCode) {
		return nil, authormodel.ErrAuthorViewNoPermission
	}

	result, err := biz.repo.ListAuthor(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
