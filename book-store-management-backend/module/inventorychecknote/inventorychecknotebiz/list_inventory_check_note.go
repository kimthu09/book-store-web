package inventorychecknotebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/inventorychecknote/inventorychecknotemodel"
	"context"
)

type ListInventoryCheckNoteRepo interface {
	ListInventoryCheckNote(
		ctx context.Context,
		filter *inventorychecknotemodel.Filter,
		paging *common.Paging,
	) ([]inventorychecknotemodel.InventoryCheckNote, error)
}

type listInventoryCheckNoteBiz struct {
	repo      ListInventoryCheckNoteRepo
	requester middleware.Requester
}

func NewListInventoryCheckNoteBiz(
	repo ListInventoryCheckNoteRepo,
	requester middleware.Requester) *listInventoryCheckNoteBiz {
	return &listInventoryCheckNoteBiz{repo: repo, requester: requester}
}

func (biz *listInventoryCheckNoteBiz) ListInventoryCheckNote(
	ctx context.Context,
	filter *inventorychecknotemodel.Filter,
	paging *common.Paging) ([]inventorychecknotemodel.InventoryCheckNote, error) {
	if !biz.requester.IsHasFeature(common.InventoryCheckNoteViewFeatureCode) {
		return nil, inventorychecknotemodel.ErrInventoryCheckNoteViewNoPermission
	}

	result, err := biz.repo.ListInventoryCheckNote(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
