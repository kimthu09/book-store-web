package bookrepo

import (
	"book-store-management-backend/module/author/authormodel"
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/category/categorymodel"
	"context"
	"strings"
)

type GetAllBookStore interface {
	GetAllBook(
		ctx context.Context,
		justGetAllActiveBook bool,
		moreKeys ...string) ([]bookmodel.ResUnitBook, error)
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

type getAllBookRepo struct {
	bookStore     GetAllBookStore
	categoryStore GetCategoryStore
	authorStore   GetAuthorStore
}

func NewGetAllBookRepo(
	bookStore GetAllBookStore,
	categoryStore GetCategoryStore,
	authorStore GetAuthorStore) *getAllBookRepo {
	return &getAllBookRepo{
		bookStore:     bookStore,
		categoryStore: categoryStore,
		authorStore:   authorStore,
	}
}

func (repo *getAllBookRepo) GetAllBook(
	ctx context.Context) ([]bookmodel.ResUnitBook, error) {

	books, errGetAllBook := repo.bookStore.GetAllBook(
		ctx, true, "BookTitle", "Publisher")
	if errGetAllBook != nil {
		return nil, errGetAllBook
	}

	authorMap := make(map[string]*authormodel.SimpleAuthor)
	categoryMap := make(map[string]*categorymodel.SimpleCategory)
	for i, book := range books {
		authors := book.BookTitle.Authors
		authorIds := strings.Split(book.BookTitle.AuthorIDs, "|")
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
		books[i].BookTitle.Authors = authors

		categories := book.BookTitle.Categories
		categoryIds := strings.Split(book.BookTitle.CategoryIDs, "|")
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
		books[i].BookTitle.Categories = categories
	}

	return books, nil
}
