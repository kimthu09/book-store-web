package booktitlebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/author/authorrepo"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"book-store-management-backend/module/booktitle/booktitlerepo"
	"book-store-management-backend/module/category/categoryrepo"
	"context"
)

type createBookTitleBiz struct {
	gen          generator.IdGenerator
	repo         booktitlerepo.CreateBookTitleRepo
	authorRepo   authorrepo.AuthorPublicRepo
	categoryRepo categoryrepo.CategoryPublicRepo
	requester    middleware.Requester
}

func NewCreateBookTitleBiz(
	gen generator.IdGenerator,
	repo booktitlerepo.CreateBookTitleRepo,
	authorRepo authorrepo.AuthorPublicRepo,
	categoryRepo categoryrepo.CategoryPublicRepo,
	requester middleware.Requester) *createBookTitleBiz {
	return &createBookTitleBiz{
		gen:          gen,
		repo:         repo,
		authorRepo:   authorRepo,
		categoryRepo: categoryRepo,
		requester:    requester,
	}
}

func (biz *createBookTitleBiz) CreateBookTitle(ctx context.Context, reqData *booktitlemodel.ReqCreateBookTitle, resData *booktitlemodel.ResCreateBookTitle) error {
	if !biz.requester.IsHasFeature(common.BookTitleCreateFeatureCode) {
		return booktitlemodel.ErrBookTitleCreateNoPermission
	}

	data := &booktitlemodel.BookTitle{
		ID:          nil,
		Name:        &reqData.Name,
		Description: &reqData.Description,
		AuthorIDs:   &reqData.AuthorIDs,
		CategoryIDs: &reqData.CategoryIDs,
	}
	if reqData.Id != "" {
		data.ID = &reqData.Id
	}

	if err := data.Validate(); err != nil {
		return err
	}
	tmpAuthorIDs := common.RemoveDuplicateStringValues(*data.AuthorIDs)
	data.AuthorIDs = &tmpAuthorIDs
	if err := validateAuthors(ctx, biz.authorRepo, *data.AuthorIDs); err != nil {
		return err
	}
	tmpCategoryIDs := common.RemoveDuplicateStringValues(*data.CategoryIDs)
	data.CategoryIDs = &tmpCategoryIDs
	if err := validateCategories(ctx, biz.categoryRepo, *data.CategoryIDs); err != nil {
		return err
	}
	if err := handleBookTitleId(biz.gen, data); err != nil {
		return err
	}

	if err := biz.repo.CreateBookTitle(ctx, data); err != nil {
		return common.ErrDB(err)
	}
	resData.Id = *data.ID
	return nil
}

func handleBookTitleId(gen generator.IdGenerator, data *booktitlemodel.BookTitle) error {
	if data.ID != nil && *data.ID != "" {
		return nil
	}

	id, err := gen.IdProcess(data.ID)
	if err != nil {
		return err
	}
	data.ID = id
	return nil
}
