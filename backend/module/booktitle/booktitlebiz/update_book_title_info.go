package booktitlebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/author/authorrepo"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"book-store-management-backend/module/category/categoryrepo"
	"context"
)

type UpdateBookTitleRepo interface {
	UpdateBookTitle(ctx context.Context, id string, data *booktitlemodel.BookTitle) error
}

type updateBookTitleBiz struct {
	repo         UpdateBookTitleRepo
	authorRepo   authorrepo.AuthorPublicRepo
	categoryRepo categoryrepo.CategoryPublicRepo
	requester    middleware.Requester
}

func NewUpdateBookBiz(
	repo UpdateBookTitleRepo,
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

	data := &booktitlemodel.BookTitle{
		ID:          &id,
		Name:        reqData.Name,
		Description: reqData.Description,
		AuthorIDs:   reqData.AuthorIDs,
		CategoryIDs: reqData.CategoryIDs,
	}

	if data.AuthorIDs != nil && len(*data.AuthorIDs) > 0 {
		tmpAuthorIDs := common.RemoveDuplicateStringValues(*data.AuthorIDs)
		data.AuthorIDs = &tmpAuthorIDs
		if err := validateAuthors(ctx, biz.authorRepo, *data.AuthorIDs); err != nil {
			return err
		}
	} else {
		data.AuthorIDs = nil
	}

	if data.CategoryIDs != nil && len(*data.CategoryIDs) > 0 {
		tmpCategoryIDs := common.RemoveDuplicateStringValues(*data.CategoryIDs)
		data.CategoryIDs = &tmpCategoryIDs
		if err := validateCategories(ctx, biz.categoryRepo, *data.CategoryIDs); err != nil {
			return err
		}
	} else {
		data.CategoryIDs = nil
	}

	if err := biz.repo.UpdateBookTitle(ctx, id, data); err != nil {
		return err
	}

	return nil
}
