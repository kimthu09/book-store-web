package shopgeneralbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/shopgeneral/shopgeneralmodel"
	"context"
)

type UpdateGeneralShopStore interface {
	UpdateGeneralShop(
		ctx context.Context,
		data *shopgeneralmodel.ReqUpdateShopGeneral) error
}

type updateGeneralShopBiz struct {
	store     UpdateGeneralShopStore
	requester middleware.Requester
}

func NewUpdateGeneralShopBiz(
	store UpdateGeneralShopStore,
	requester middleware.Requester) *updateGeneralShopBiz {
	return &updateGeneralShopBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *updateGeneralShopBiz) UpdateGeneralShop(
	ctx context.Context,
	data *shopgeneralmodel.ReqUpdateShopGeneral) error {
	if biz.requester.GetRoleId() != common.RoleAdminId {
		return shopgeneralmodel.ErrGeneralShopUpdateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.UpdateGeneralShop(ctx, data); err != nil {
		return err
	}

	return nil
}
