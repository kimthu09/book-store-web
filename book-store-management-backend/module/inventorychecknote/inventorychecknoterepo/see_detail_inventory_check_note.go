package inventorychecknoterepo

import (
	"book-store-management-backend/module/inventorychecknote/inventorychecknotemodel"
	"context"
)

type SeeDetailInventoryCheckNoteStore interface {
	FindInventoryCheckNote(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*inventorychecknotemodel.InventoryCheckNote, error)
}

type seeDetailInventoryCheckNoteRepo struct {
	store SeeDetailInventoryCheckNoteStore
}

func NewSeeDetailInventoryCheckNoteRepo(store SeeDetailInventoryCheckNoteStore) *seeDetailInventoryCheckNoteRepo {
	return &seeDetailInventoryCheckNoteRepo{store: store}
}

func (repo *seeDetailInventoryCheckNoteRepo) SeeDetailInventoryCheckNote(
	ctx context.Context,
	inventoryCheckNoteId string) (*inventorychecknotemodel.InventoryCheckNote, error) {
	inventoryCheckNote, err := repo.store.FindInventoryCheckNote(
		ctx,
		map[string]interface{}{"id": inventoryCheckNoteId},
		"Details.Book")

	if err != nil {
		return nil, err
	}

	return inventoryCheckNote, nil
}
