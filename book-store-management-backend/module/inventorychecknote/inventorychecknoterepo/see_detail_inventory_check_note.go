package inventorychecknoterepo

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/inventorychecknote/inventorychecknotemodel"
	"book-store-management-backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
	"context"
)

type SeeDetailInventoryCheckNoteStore interface {
	ListInventoryCheckNoteDetail(
		ctx context.Context,
		inventoryCheckNoteId string,
		paging *common.Paging) ([]inventorychecknotedetailmodel.InventoryCheckNoteDetail, error)
}

type FindInventoryCheckNoteStore interface {
	FindInventoryCheckNote(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*inventorychecknotemodel.InventoryCheckNote, error)
}

type seeDetailInventoryCheckNoteRepo struct {
	inventoryCheckNoteStore       FindInventoryCheckNoteStore
	inventoryCheckNoteDetailStore SeeDetailInventoryCheckNoteStore
}

func NewSeeDetailInventoryCheckNoteRepo(
	inventoryCheckNoteStore FindInventoryCheckNoteStore,
	inventoryCheckNoteDetailStore SeeDetailInventoryCheckNoteStore) *seeDetailInventoryCheckNoteRepo {
	return &seeDetailInventoryCheckNoteRepo{
		inventoryCheckNoteStore:       inventoryCheckNoteStore,
		inventoryCheckNoteDetailStore: inventoryCheckNoteDetailStore,
	}
}

func (repo *seeDetailInventoryCheckNoteRepo) SeeDetailInventoryCheckNote(
	ctx context.Context,
	inventoryCheckNoteId string,
	paging *common.Paging) (*inventorychecknotemodel.ResDetailInventoryCheckNote, error) {
	inventoryCheckNote, errInventoryCheckNote :=
		repo.inventoryCheckNoteStore.FindInventoryCheckNote(
			ctx,
			map[string]interface{}{"id": inventoryCheckNoteId},
			"CreateByUser")
	if errInventoryCheckNote != nil {
		return nil, errInventoryCheckNote
	}

	resDetailInventoryCheckNote := inventorychecknotemodel.GetResDetailInventoryCheckNoteFromInventoryCheckNote(inventoryCheckNote)

	details, errInventoryCheckNoteDetail := repo.inventoryCheckNoteDetailStore.ListInventoryCheckNoteDetail(
		ctx,
		inventoryCheckNoteId,
		paging,
	)
	if errInventoryCheckNoteDetail != nil {
		return nil, errInventoryCheckNoteDetail
	}

	resDetailInventoryCheckNote.Details = details

	return resDetailInventoryCheckNote, nil
}
