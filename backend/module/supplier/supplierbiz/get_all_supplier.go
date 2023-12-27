package supplierbiz

import (
	"book-store-management-backend/module/importnote/importnotemodel"
	"context"
)

type GetAllSupplierRepo interface {
	GetAllSupplier(
		ctx context.Context) ([]importnotemodel.SimpleSupplier, error)
}

type getAllSupplierBiz struct {
	repo GetAllSupplierRepo
}

func NewGetAllSupplierBiz(
	repo GetAllSupplierRepo) *getAllSupplierBiz {
	return &getAllSupplierBiz{repo: repo}
}

func (biz *getAllSupplierBiz) GetAllUser(
	ctx context.Context) ([]importnotemodel.SimpleSupplier, error) {
	result, err := biz.repo.GetAllSupplier(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
