package booktitlebiz

import (
	"book-store-management-backend/module/booktitle/booktitlemodel"
	booktitlestore "book-store-management-backend/module/booktitle/booktitlestore"
	"context"
)

type updateBookRepo interface {
	UpdateBookTitle(ctx context.Context, id string, data *booktitlestore.BookTitleDBModel) error
}

type updateBookBiz struct {
	repo         updateBookRepo
	authorRepo   authorRepo
	categoryRepo categoryRepo
}

func NewUpdateBookBiz(
	repo updateBookRepo,
	authorRepo authorRepo,
	categoryRepo categoryRepo) *updateBookBiz {
	return &updateBookBiz{
		repo:         repo,
		authorRepo:   authorRepo,
		categoryRepo: categoryRepo,
	}
}

func (biz *updateBookBiz) UpdateBookTitle(ctx context.Context, id string, reqData *booktitlemodel.ReqUpdateBookInfo) error {
	err := biz.repo.UpdateBookTitle(ctx, id, &booktitlestore.BookTitleDBModel{
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
