package bookbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/book/bookmodel"
	"context"
)

type ListBookRepo interface {
	ListBook(
		ctx context.Context,
		filter *bookmodel.Filter,
		paging *common.Paging) ([]bookmodel.ResDetailUnitBook, error)
}

type listBookBiz struct {
	repo      ListBookRepo
	requester middleware.Requester
}

func NewListBookBiz(
	repo ListBookRepo,
	requester middleware.Requester) *listBookBiz {
	return &listBookBiz{repo: repo, requester: requester}
}

func (biz *listBookBiz) ListBook(
	ctx context.Context,
	filter *bookmodel.Filter,
	paging *common.Paging) ([]bookmodel.ResDetailUnitBook, error) {
	if !biz.requester.IsHasFeature(common.BookViewFeatureCode) {
		return nil, bookmodel.ErrBookViewNoPermission
	}

	result, err := biz.repo.ListBook(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
