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

type bookTitleDetailBiz struct {
	repo         booktitlerepo.DetailBookTitleRepo
	authorRepo   authorrepo.AuthorPublicRepo
	categoryRepo categoryrepo.CategoryPublicRepo
	requester    middleware.Requester
}

func NewGetBookTitleDetailBiz(repo booktitlerepo.DetailBookTitleRepo, authorRepo authorrepo.AuthorPublicRepo, categoryRepo categoryrepo.CategoryPublicRepo, requester middleware.Requester) *bookTitleDetailBiz {
	return &bookTitleDetailBiz{repo: repo, authorRepo: authorRepo, categoryRepo: categoryRepo, requester: requester}
}

func (biz *bookTitleDetailBiz) GetBookTitleDetail(ctx context.Context, id string) (*booktitlemodel.BookTitleDetail, error) {
	if !biz.requester.IsHasFeature(common.BookTitleViewFeatureCode) {
		return nil, booktitlemodel.ErrBookTitleViewNoPermission
	}

	rawBookTitle, err := biz.repo.DetailBookTitle(ctx, id)
	if err != nil {
		switch err.Error() {
		case "record not found":
			return nil, booktitlemodel.ErrBookTitleNotFound
		default:
			return nil, common.ErrDB(err)
		}
	}

	result := booktitlemodel.BookTitleDetail{
		ID:          rawBookTitle.ID,
		Name:        *rawBookTitle.Name,
		Description: *rawBookTitle.Description,
		CreatedAt:   rawBookTitle.CreatedAt,
		Authors:     nil,
		Categories:  nil,
		UpdatedAt:   rawBookTitle.UpdatedAt,
		DeletedAt:   rawBookTitle.DeletedAt,
		IsActive:    rawBookTitle.IsActive,
	}

	authors, err := biz.authorRepo.GetByListId(ctx, *rawBookTitle.AuthorIDs)
	if err != nil {
		return nil, err
	}
	result.Authors = authors
	for j := range result.Authors {
		result.Authors[j].CreatedAt = nil
		result.Authors[j].UpdatedAt = nil
		result.Authors[j].DeletedAt = nil
		result.Authors[j].IsActive = nil
	}

	categories, err := biz.categoryRepo.GetByListId(ctx, *rawBookTitle.CategoryIDs)
	if err != nil {
		return nil, err
	}
	result.Categories = categories
	for j := range result.Categories {
		result.Categories[j].CreatedAt = nil
		result.Categories[j].UpdatedAt = nil
		result.Categories[j].DeletedAt = nil
		result.Categories[j].IsActive = nil
	}

	return &result, nil
}
