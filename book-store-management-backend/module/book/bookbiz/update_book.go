package bookbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/book/bookstore"
	"context"
)

type updateBookRepo interface {
	UpdateBook(ctx context.Context, id string, data *bookstore.BookDBModel) error
}

type updateBookBiz struct {
	requester     middleware.Requester
	repo          updateBookRepo
	authorRepo    authorRepo
	publisherRepo publisherRepo
	categoryRepo  categoryRepo
}

func NewUpdateBookBiz(
	repo updateBookRepo,
	authorRepo authorRepo,
	publisherRepo publisherRepo,
	categoryRepo categoryRepo,
	requester middleware.Requester,
) *updateBookBiz {
	return &updateBookBiz{
		repo:          repo,
		authorRepo:    authorRepo,
		publisherRepo: publisherRepo,
		categoryRepo:  categoryRepo,
		requester:     requester,
	}
}

func (biz *updateBookBiz) UpdateBook(ctx context.Context, id string, reqData *bookmodel.ReqUpdateBookInfo) error {
	if !biz.requester.IsHasFeature(common.BookUpdateFeatureCode) {
		return bookmodel.ErrBookUpdateNoPermission
	}

	err := biz.repo.UpdateBook(ctx, id, &bookstore.BookDBModel{
		Name:        reqData.Name,
		Description: reqData.Description,
		Edition:     reqData.Edition,
		Quantity:    reqData.Quantity,
		ListedPrice: reqData.ListedPrice,
		SellPrice:   reqData.SellPrice,
		PublisherID: reqData.PublisherID,
		AuthorIDs:   nil,
		CategoryIDs: nil,
	})

	if err != nil {
		return err
	}

	return nil
}
