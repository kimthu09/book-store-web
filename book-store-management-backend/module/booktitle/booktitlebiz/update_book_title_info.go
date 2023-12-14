package booktitlebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/author/authorrepo"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"book-store-management-backend/module/booktitle/booktitlerepo"
	booktitlestore "book-store-management-backend/module/booktitle/booktitlestore"
	"book-store-management-backend/module/category/categoryrepo"
	"context"
)

type updateBookTitleBiz struct {
	requester    middleware.Requester
	repo         booktitlerepo.UpdateBookTitleRepo
	authorRepo   authorrepo.AuthorPublicRepo
	categoryRepo categoryrepo.CategoryPublicRepo
}

func NewUpdateBookBiz(
	repo booktitlerepo.UpdateBookTitleRepo,
	authorRepo authorrepo.AuthorPublicRepo,
	categoryRepo categoryrepo.CategoryPublicRepo,
	requester middleware.Requester,
) *updateBookTitleBiz {
	return &updateBookTitleBiz{
		repo:         repo,
		authorRepo:   authorRepo,
		categoryRepo: categoryRepo,
		requester:    requester,
	}
}

func (biz *updateBookTitleBiz) UpdateBookTitle(ctx context.Context, id string, reqData *booktitlemodel.ReqUpdateBookInfo) error {
	if !biz.requester.IsHasFeature(common.BookTitleUpdateFeatureCode) {
		return booktitlemodel.ErrBookTitleUpdateNoPermission
	}

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
