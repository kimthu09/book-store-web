package bookbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/book/bookmodel"
	"context"
)

type DeleteBookRepo interface {
	DeleteBook(ctx context.Context, id string) error
}

type DeleteRepo interface {
	DeleteBook(ctx context.Context, id string) error
}

type deleteBookBiz struct {
	requester middleware.Requester
	repo      DeleteRepo
}

func NewDeleteBookBiz(requester middleware.Requester, repo DeleteRepo) *deleteBookBiz {
	return &deleteBookBiz{
		requester: requester,
		repo:      repo,
	}
}

func (biz *deleteBookBiz) DeleteBook(ctx context.Context, id string) error {
	if !biz.requester.IsHasFeature(common.BookDeleteFeatureCode) {
		return bookmodel.ErrBookDeleteNoPermission
	}

	return biz.repo.DeleteBook(ctx, id)
}
