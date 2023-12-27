package bookrepo

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/author/authormodel"
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/category/categorymodel"
	"context"
	"strings"
)

type ListBookStore interface {
	ListBook(
		ctx context.Context,
		filter *bookmodel.Filter,
		propertiesContainSearchKey []string,
		paging *common.Paging,
		moreKeys ...string) ([]bookmodel.ResDetailUnitBook, error)
}

type listBookRepo struct {
	bookStore     ListBookStore
	categoryStore GetCategoryStore
	authorStore   GetAuthorStore
}

func NewListBookRepo(
	bookStore ListBookStore,
	categoryStore GetCategoryStore,
	authorStore GetAuthorStore) *listBookRepo {
	return &listBookRepo{
		bookStore:     bookStore,
		categoryStore: categoryStore,
		authorStore:   authorStore,
	}
}

func (repo *listBookRepo) ListBook(
	ctx context.Context,
	filter *bookmodel.Filter,
	paging *common.Paging) ([]bookmodel.ResDetailUnitBook, error) {

	books, errGetAllBook := repo.bookStore.ListBook(
		ctx,
		filter,
		[]string{"Book.id", "Book.name"},
		paging,
		"BookTitle", "Publisher")
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
