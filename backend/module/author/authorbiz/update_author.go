package authorbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/author/authormodel"
	"context"
)

type UpdateAuthorRepo interface {
	UpdateAuthorInfo(
		ctx context.Context,
		supplierId string,
		data *authormodel.ReqUpdateAuthor) error
}

type updateAuthorBiz struct {
	repo      UpdateAuthorRepo
	requester middleware.Requester
}

func NewUpdateAuthorBiz(
	repo UpdateAuthorRepo,
	requester middleware.Requester) *updateAuthorBiz {
	return &updateAuthorBiz{repo: repo, requester: requester}
}

func (biz *updateAuthorBiz) UpdateAuthor(
	ctx context.Context,
	id string,
	data *authormodel.ReqUpdateAuthor) error {
	if !biz.requester.IsHasFeature(common.AuthorUpdateFeatureCode) {
		return authormodel.ErrAuthorUpdateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.repo.UpdateAuthorInfo(ctx, id, data); err != nil {
		return err
	}

	return nil
}
