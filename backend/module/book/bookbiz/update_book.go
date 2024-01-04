package bookbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/book/bookmodel"
	"context"
)

type UpdateBookRepo interface {
	UpdateBook(ctx context.Context, data *bookmodel.Book) error
}

type updateBookBiz struct {
	repo      UpdateBookRepo
	requester middleware.Requester
}

func NewUpdateBookInfoBiz(repo UpdateBookRepo, requester middleware.Requester) *updateBookBiz {
	return &updateBookBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *updateBookBiz) UpdateBook(ctx context.Context, id string, reqData *bookmodel.ReqUpdateBook) error {
	if !biz.requester.IsHasFeature(common.BookUpdateFeatureCode) {
		return bookmodel.ErrBookUpdateNoPermission
	}

	if err := reqData.Validate(); err != nil {
		return err
	}

	data := &bookmodel.Book{
		ID:          &id,
		Image:       reqData.Image,
		BookTitleID: reqData.BookTitleID,
		PublisherID: reqData.PublisherID,
		Edition:     reqData.Edition,
		ListedPrice: reqData.ListedPrice,
		SellPrice:   reqData.SellPrice,
	}

	if err := biz.repo.UpdateBook(ctx, data); err != nil {
		return err
	}

	return nil
}
