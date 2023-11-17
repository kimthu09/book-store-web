package importnoterepo

import (
	"book-store-management-backend/module/importnote/importnotemodel"
	"context"
)

type SeeDetailImportNoteStore interface {
	FindImportNote(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*importnotemodel.ImportNote, error)
}

type seeDetailImportNoteRepo struct {
	store SeeDetailImportNoteStore
}

func NewSeeDetailImportNoteRepo(store SeeDetailImportNoteStore) *seeDetailImportNoteRepo {
	return &seeDetailImportNoteRepo{store: store}
}

func (repo *seeDetailImportNoteRepo) SeeDetailImportNote(
	ctx context.Context,
	importNoteId string) (*importnotemodel.ImportNote, error) {
	importNote, err := repo.store.FindImportNote(
		ctx,
		map[string]interface{}{"id": importNoteId},
		"Supplier",
		"Details.Book")

	if err != nil {
		return nil, err
	}

	return importNote, nil
}
