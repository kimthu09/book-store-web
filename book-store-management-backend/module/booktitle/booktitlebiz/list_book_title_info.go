package booktitlebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"context"
)

type ListBookTitleRepo interface {
	ListBookTitle(ctx context.Context, filter *booktitlemodel.Filter, paging *common.Paging) ([]booktitlemodel.BookTitle, error)
}

type listBookTitleBiz struct {
	repo      ListBookTitleRepo
	requester middleware.Requester
}

func NewListBookTitleBiz(repo ListBookTitleRepo, requester middleware.Requester) *listBookTitleBiz {
	return &listBookTitleBiz{repo: repo, requester: requester}
}

func (biz *listBookTitleBiz) ListBookTitle(ctx context.Context, filter *booktitlemodel.Filter, paging *common.Paging) ([]booktitlemodel.BookTitle, error) {
	if !biz.requester.IsHasFeature(common.BookTitleViewFeatureCode) {
		return nil, booktitlemodel.ErrBookTitleViewNoPermission
	}

	result, err := biz.repo.ListBookTitle(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
