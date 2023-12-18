package booktitlebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/author/authorrepo"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"book-store-management-backend/module/booktitle/booktitlerepo"
	"book-store-management-backend/module/category/categoryrepo"
	"context"
)

type listBookTitleBiz struct {
	repo         booktitlerepo.ListBookTitleRepo
	authorRepo   authorrepo.AuthorPublicRepo
	categoryRepo categoryrepo.CategoryPublicRepo
	requester    middleware.Requester
}

func NewListBookTitleBiz(repo booktitlerepo.ListBookTitleRepo, authorRepo authorrepo.AuthorPublicRepo, categoryRepo categoryrepo.CategoryPublicRepo, requester middleware.Requester) *listBookTitleBiz {
	return &listBookTitleBiz{repo: repo, authorRepo: authorRepo, categoryRepo: categoryRepo, requester: requester}
}

func (biz *listBookTitleBiz) ListBookTitle(ctx context.Context, filter *booktitlemodel.Filter, paging *common.Paging) ([]booktitlemodel.BookTitleDetail, error) {
	if !biz.requester.IsHasFeature(common.BookTitleViewFeatureCode) {
		return nil, booktitlemodel.ErrBookTitleViewNoPermission
	}

	booktitles, err := biz.repo.ListBookTitle(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	result := make([]booktitlemodel.BookTitleDetail, len(booktitles))

	for i, booktitle := range booktitles {
		result[i].CreatedAt = booktitle.CreatedAt
		result[i].UpdatedAt = booktitle.UpdatedAt
		result[i].DeletedAt = booktitle.DeletedAt
		result[i].IsActive = booktitle.IsActive

		result[i].ID = booktitle.ID
		result[i].Name = *booktitle.Name
		result[i].Description = *booktitle.Description

		authors, err := biz.authorRepo.GetByListId(ctx, *booktitle.AuthorIDs)
		if err != nil {
			return nil, err
		}
		result[i].Authors = authors
		for j := range result[i].Authors {
			result[i].Authors[j].CreatedAt = nil
			result[i].Authors[j].UpdatedAt = nil
			result[i].Authors[j].DeletedAt = nil
			result[i].Authors[j].IsActive = nil
		}

		categories, err := biz.categoryRepo.GetByListId(ctx, *booktitle.CategoryIDs)
		if err != nil {
			return nil, err
		}
		result[i].Categories = categories
		for j := range result[i].Categories {
			result[i].Categories[j].CreatedAt = nil
			result[i].Categories[j].UpdatedAt = nil
			result[i].Categories[j].DeletedAt = nil
			result[i].Categories[j].IsActive = nil
		}
	}

	return result, nil
}
