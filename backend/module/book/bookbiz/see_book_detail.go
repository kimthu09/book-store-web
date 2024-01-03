package bookbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/book/bookmodel"
	"context"
)

type SeeBookDetailRepo interface {
	SeeBookDetail(
		ctx context.Context,
		bookId string) (*bookmodel.ResDetailUnitBook, error)
}

type seeBookDetailBiz struct {
	repo      SeeBookDetailRepo
	requester middleware.Requester
}

func NewSeeBookDetailBiz(
	repo SeeBookDetailRepo,
	requester middleware.Requester) *seeBookDetailBiz {
	return &seeBookDetailBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *seeBookDetailBiz) SeeBookDetail(
	ctx context.Context,
	bookId string) (*bookmodel.ResDetailUnitBook, error) {
	if !biz.requester.IsHasFeature(common.BookViewFeatureCode) {
		return nil, bookmodel.ErrBookViewNoPermission
	}

	book, err := biz.repo.SeeBookDetail(
		ctx, bookId)
	if err != nil {
		return nil, err
	}

	return book, nil
}
