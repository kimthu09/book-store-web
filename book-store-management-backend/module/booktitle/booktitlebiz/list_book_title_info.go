package booktitlebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"book-store-management-backend/module/category/categorymodel"
	"context"
)

type ListBookTitleRepo interface {
	ListBookTitle(ctx context.Context, filter *booktitlemodel.Filter, paging *common.Paging) ([]booktitlemodel.BookTitle, error)
}

type listBookTitleBiz struct {
	repo       ListBookTitleRepo
	authorRepo authorPublicRepo
	requester  middleware.Requester
}

func NewListBookTitleBiz(repo ListBookTitleRepo, authorRepo authorPublicRepo, requester middleware.Requester) *listBookTitleBiz {
	return &listBookTitleBiz{repo: repo, authorRepo: authorRepo, requester: requester}
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
		result[i].Name = booktitle.Name
		result[i].Description = booktitle.Description

		authors, err := biz.authorRepo.GetByListId(ctx, booktitle.AuthorIDs)
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

		result[i].Categories = make([]categorymodel.Category, len(booktitle.CategoryIDs))
		// TODO: get categories
		for j, categoryID := range booktitle.CategoryIDs {
			result[i].Categories[j].Id = categoryID
		}

	}

	return result, nil
}
