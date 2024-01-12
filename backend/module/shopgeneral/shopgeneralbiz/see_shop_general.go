package shopgeneralbiz

import (
	"book-store-management-backend/module/shopgeneral/shopgeneralmodel"
	"context"
)

type SeeShopGeneralStore interface {
	FindShopGeneral(
		ctx context.Context) (*shopgeneralmodel.ShopGeneral, error)
}

type seeShopGeneralBiz struct {
	store SeeShopGeneralStore
}

func NewSeeShopGeneralBiz(
	store SeeShopGeneralStore) *seeShopGeneralBiz {
	return &seeShopGeneralBiz{
		store: store,
	}
}

func (biz *seeShopGeneralBiz) SeeShopGeneral(
	ctx context.Context) (*shopgeneralmodel.ShopGeneral, error) {
	general, errGetGeneral := biz.store.FindShopGeneral(ctx)
	if errGetGeneral != nil {
		return nil, errGetGeneral
	}

	return general, nil
}
