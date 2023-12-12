package booktitlebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	booktitlestore "book-store-management-backend/module/booktitle/booktitlestore"
	"context"
)

type updateBookRepo interface {
	UpdateBook(ctx context.Context, id string, data *booktitlestore.BookTitleDBModel) error
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

func (biz *updateBookBiz) UpdateBook(ctx context.Context, id string, reqData *booktitlemodel.ReqUpdateBookInfo) error {
	if !biz.requester.IsHasFeature(common.BookTitleUpdateFeatureCode) {
		return booktitlemodel.ErrBookTitleUpdateNoPermission
	}

	err := biz.repo.UpdateBook(ctx, id, &booktitlestore.BookTitleDBModel{
		Name:        reqData.Name,
		Description: reqData.Description,
		AuthorIDs:   nil,
		CategoryIDs: nil,
	})

	if err != nil {
		return err
	}

	return nil
}
