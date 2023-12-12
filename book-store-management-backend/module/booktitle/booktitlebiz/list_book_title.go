package booktitlebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"context"
)

type ListBookRepo interface {
	ListBook(ctx context.Context, filter *booktitlemodel.Filter, paging *common.Paging) ([]booktitlemodel.Book, error)
}

type listBookBiz struct {
	repo      ListBookRepo
	requester middleware.Requester
}

func NewListBookBiz(repo ListBookRepo, requester middleware.Requester) *listBookBiz {
	return &listBookBiz{repo: repo, requester: requester}
}

func (biz *listBookBiz) ListBook(ctx context.Context, filter *booktitlemodel.Filter, paging *common.Paging) ([]booktitlemodel.Book, error) {
	if !biz.requester.IsHasFeature(common.BookTitleViewFeatureCode) {
		return nil, booktitlemodel.ErrBookTitleViewNoPermission
	}

	result, err := biz.repo.ListBook(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
