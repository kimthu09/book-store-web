package authorbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/author/authormodel"
	"context"
)

type createListAuthorBiz struct {
	gen       generator.IdGenerator
	repo      CreateAuthorRepo
	requester middleware.Requester
}

func NewCreateListAuthorBiz(gen generator.IdGenerator, repo CreateAuthorRepo, requester middleware.Requester) *createListAuthorBiz {
	return &createListAuthorBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createListAuthorBiz) CreateListAuthor(
	ctx context.Context, data []authormodel.Author) error {
	if !biz.requester.IsHasFeature(common.AuthorCreateFeatureCode) {
		return authormodel.ErrAuthorCreateNoPermission
	}

	for _, v := range data {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	for i, _ := range data {
		if err := handleAuthorId(biz.gen, &data[i]); err != nil {
			return err
		}
	}

	for _, v := range data {
		if err := biz.repo.CreateAuthor(ctx, &v); err != nil {
			return err
		}
	}

	return nil
}
