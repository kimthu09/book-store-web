package supplierrepo

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/importnote/importnotemodel"
	"book-store-management-backend/module/supplier/suppliermodel/filter"
	"context"
)

type ListSupplierImportNoteStore interface {
	ListAllImportNoteBySupplier(
		supplierId string,
		filter *filter.SupplierImportFilter,
		ctx context.Context,
		paging *common.Paging,
		moreKeys ...string) ([]importnotemodel.ImportNote, error)
}

type seeSupplierImportNoteRepo struct {
	importNoteStore ListSupplierImportNoteStore
}

func NewSeeSupplierImportNoteRepo(
	importNoteStore ListSupplierImportNoteStore) *seeSupplierImportNoteRepo {
	return &seeSupplierImportNoteRepo{
		importNoteStore: importNoteStore,
	}
}

func (biz *seeSupplierImportNoteRepo) SeeSupplierImportNote(
	ctx context.Context,
	supplierId string,
	filter *filter.SupplierImportFilter,
	paging *common.Paging) ([]importnotemodel.ImportNote, error) {
	importNotes, errImportNotes := biz.importNoteStore.ListAllImportNoteBySupplier(
		supplierId,
		filter,
		ctx,
		paging,
		"CreatedByUser", "ClosedByUser",
	)
	if errImportNotes != nil {
		return nil, errImportNotes
	}

	return importNotes, nil
}
