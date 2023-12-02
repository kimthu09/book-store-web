package inventorychecknoterepo

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/inventorychecknote/inventorychecknotemodel"
	"context"
)

type ListInventoryCheckNoteStore interface {
	ListInventoryCheckNote(
		ctx context.Context,
		filter *inventorychecknotemodel.Filter,
		propertiesContainSearchKey []string,
		paging *common.Paging) ([]inventorychecknotemodel.InventoryCheckNote, error)
}

type listInventoryCheckNoteRepo struct {
	store ListInventoryCheckNoteStore
}

func NewListInventoryCheckNoteRepo(store ListInventoryCheckNoteStore) *listInventoryCheckNoteRepo {
	return &listInventoryCheckNoteRepo{store: store}
}

func (repo *listInventoryCheckNoteRepo) ListInventoryCheckNote(
	ctx context.Context,
	filter *inventorychecknotemodel.Filter,
	paging *common.Paging) ([]inventorychecknotemodel.InventoryCheckNote, error) {
	result, err := repo.store.ListInventoryCheckNote(
		ctx,
		filter,
		[]string{"InventoryCheckNote.id"},
		paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
