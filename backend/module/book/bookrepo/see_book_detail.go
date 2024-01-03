package bookrepo

import (
	"book-store-management-backend/module/book/bookmodel"
	"context"
	"strings"
)

type FindBookStore interface {
	FindDetailBook(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*bookmodel.ResDetailUnitBook, error)
}

type seeBookDetailRepo struct {
	bookStore     FindBookStore
	categoryStore GetCategoryStore
	authorStore   GetAuthorStore
}

func NewSeeBookDetailRepo(
	bookStore FindBookStore,
	categoryStore GetCategoryStore,
	authorStore GetAuthorStore) *seeBookDetailRepo {
	return &seeBookDetailRepo{
		bookStore:     bookStore,
		categoryStore: categoryStore,
		authorStore:   authorStore,
	}
}

func (repo *seeBookDetailRepo) SeeBookDetail(
	ctx context.Context,
	bookId string) (*bookmodel.ResDetailUnitBook, error) {
	book, errBook := repo.bookStore.FindDetailBook(
		ctx, map[string]interface{}{"id": bookId},
		"BookTitle", "Publisher")
	if errBook != nil {
		return nil, errBook
	}

	authors := book.BookTitle.Authors
	authorIds := strings.Split(book.BookTitle.AuthorIDs, "|")
	for _, authorId := range authorIds {
		author, errGetAuthor := repo.authorStore.FindAuthor(
			ctx, map[string]interface{}{"id": authorId})
		if errGetAuthor != nil {
			return nil, errGetAuthor
		}
		authors = append(authors, *author)
	}
	book.BookTitle.Authors = authors

	categories := book.BookTitle.Categories
	categoryIds := strings.Split(book.BookTitle.CategoryIDs, "|")
	for _, categoryId := range categoryIds {
		category, errGetCategory := repo.categoryStore.FindCategory(
			ctx, map[string]interface{}{"id": categoryId})
		if errGetCategory != nil {
			return nil, errGetCategory
		}
		categories = append(categories, *category)
	}
	book.BookTitle.Categories = categories

	return book, nil
}
