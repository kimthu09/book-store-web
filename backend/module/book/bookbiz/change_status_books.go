package bookbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/book/bookmodel"
	"context"
)

type ChangeStatusBookRepo interface {
	UpdateStatusBooks(
		ctx context.Context,
		data *bookmodel.ReqUpdateStatusBooks,
	) error
}

type changeStatusBookBiz struct {
	repo      ChangeStatusBookRepo
	requester middleware.Requester
}

func NewChangeStatusBooksBiz(
	repo ChangeStatusBookRepo,
	requester middleware.Requester) *changeStatusBookBiz {
	return &changeStatusBookBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *changeStatusBookBiz) ChangeStatusBooks(
	ctx context.Context,
	data *bookmodel.ReqUpdateStatusBooks) error {
	if !biz.requester.IsHasFeature(common.BookUpdateStatusFeatureCode) {
		return bookmodel.ErrBookUpdateStatusNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.repo.UpdateStatusBooks(ctx, data); err != nil {
		return err
	}

	return nil
}
