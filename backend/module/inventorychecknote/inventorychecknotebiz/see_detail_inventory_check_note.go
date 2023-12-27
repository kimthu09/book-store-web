package inventorychecknotebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/inventorychecknote/inventorychecknotemodel"
	"context"
)

type SeeDetailInventoryCheckNoteRepo interface {
	SeeDetailInventoryCheckNote(
		ctx context.Context,
		inventoryCheckNoteId string,
	) (*inventorychecknotemodel.ResDetailInventoryCheckNote, error)
}

type seeDetailInventoryCheckNoteBiz struct {
	repo      SeeDetailInventoryCheckNoteRepo
	requester middleware.Requester
}

func NewSeeDetailImportNoteBiz(
	repo SeeDetailInventoryCheckNoteRepo,
	requester middleware.Requester) *seeDetailInventoryCheckNoteBiz {
	return &seeDetailInventoryCheckNoteBiz{repo: repo, requester: requester}
}

func (biz *seeDetailInventoryCheckNoteBiz) SeeDetailInventoryCheckNote(
	ctx context.Context,
	inventoryCheckNoteId string) (*inventorychecknotemodel.ResDetailInventoryCheckNote, error) {
	if !biz.requester.IsHasFeature(common.InventoryCheckNoteViewFeatureCode) {
		return nil, inventorychecknotemodel.ErrInventoryCheckNoteViewNoPermission
	}

	inventoryCheckNote, err := biz.repo.SeeDetailInventoryCheckNote(
		ctx,
		inventoryCheckNoteId)

	if err != nil {
		return nil, err
	}

	return inventoryCheckNote, nil
}
