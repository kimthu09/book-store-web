package booktitlerepo

import (
	"book-store-management-backend/module/author/authormodel"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"book-store-management-backend/module/category/categorymodel"
	"context"
	"strings"
)

type GetAllBookTitleStore interface {
	GetAllBookTitle(
		ctx context.Context,
		moreKeys ...string) ([]booktitlemodel.SimpleBookTitle, error)
}

type GetAuthorStore interface {
	FindAuthor(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*authormodel.SimpleAuthor, error)
}

type GetCategoryStore interface {
	FindCategory(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*categorymodel.SimpleCategory, error)
}

type getAllBookTitleRepo struct {
	bookTitleStore GetAllBookTitleStore
	categoryStore  GetCategoryStore
	authorStore    GetAuthorStore
}

func NewGetAllBookTitleRepo(
	bookTitleStore GetAllBookTitleStore,
	categoryStore GetCategoryStore,
	authorStore GetAuthorStore) *getAllBookTitleRepo {
	return &getAllBookTitleRepo{
		bookTitleStore: bookTitleStore,
		categoryStore:  categoryStore,
		authorStore:    authorStore,
	}
}

func (repo *getAllBookTitleRepo) GetAllBookTitle(
	ctx context.Context) ([]booktitlemodel.SimpleBookTitle, error) {

	bookTitles, errGetAllBookTitle := repo.bookTitleStore.GetAllBookTitle(ctx)
	if errGetAllBookTitle != nil {
		return nil, errGetAllBookTitle
	}

	authorMap := make(map[string]*authormodel.SimpleAuthor)
	categoryMap := make(map[string]*categorymodel.SimpleCategory)
	for i, bookTitle := range bookTitles {
		authors := bookTitle.Authors
		authorIds := strings.Split(bookTitle.AuthorIDs, "|")
		for _, authorId := range authorIds {
			if authorMap[authorId] == nil {
				author, errGetAuthor := repo.authorStore.FindAuthor(
					ctx, map[string]interface{}{"id": authorId})
				if errGetAuthor != nil {
					return nil, errGetAuthor
				}

				authorMap[authorId] = author
			}
			authors = append(authors, *authorMap[authorId])
		}
		bookTitles[i].Authors = authors

		categories := bookTitle.Categories
		categoryIds := strings.Split(bookTitle.CategoryIDs, "|")
		for _, categoryId := range categoryIds {
			if authorMap[categoryId] == nil {
				category, errGetCategory := repo.categoryStore.FindCategory(
					ctx, map[string]interface{}{"id": categoryId})
				if errGetCategory != nil {
					return nil, errGetCategory
				}

				categoryMap[categoryId] = category
			}
			categories = append(categories, *categoryMap[categoryId])
		}
		bookTitles[i].Categories = categories
	}

	return bookTitles, nil
}
