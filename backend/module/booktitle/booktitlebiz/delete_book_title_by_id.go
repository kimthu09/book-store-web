package booktitlebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"book-store-management-backend/module/booktitle/booktitlerepo"
	"context"
)

type DeleteBookRepo interface {
	DeleteBookTitle(ctx context.Context, id string) error
}

type deleteBookTitleBiz struct {
	requester middleware.Requester
	repo      booktitlerepo.DeleteBookTileRepo
}

func NewDeleteBookTitleBiz(requester middleware.Requester, repo booktitlerepo.DeleteBookTileRepo) *deleteBookTitleBiz {
	return &deleteBookTitleBiz{
		requester: requester,
		repo:      repo,
	}
}

func (biz *deleteBookTitleBiz) DeleteBookTitle(ctx context.Context, id string) error {
	if !biz.requester.IsHasFeature(common.BookTitleDeleteFeatureCode) {
		return booktitlemodel.ErrBookTitleDeleteNoPermission
	}

	err := biz.repo.DeleteBookTitle(ctx, id)
	if err != nil {
		panic(err)
	}
	return nil
}
