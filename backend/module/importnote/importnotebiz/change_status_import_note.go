package importnotebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/importnote/importnotemodel"
	"book-store-management-backend/module/importnotedetail/importnotedetailmodel"
	"context"
)

type ChangeStatusImportNoteRepo interface {
	FindImportNote(
		ctx context.Context,
		importNoteId string) (*importnotemodel.ImportNote, error)
	UpdateImportNote(
		ctx context.Context,
		importNoteId string,
		data *importnotemodel.ReqUpdateImportNote) error
	CreateSupplierDebt(
		ctx context.Context,
		supplierDebtId string,
		importNote *importnotemodel.ReqUpdateImportNote) error
	UpdateDebtSupplier(
		ctx context.Context,
		importNote *importnotemodel.ReqUpdateImportNote) error
	FindListImportNoteDetail(
		ctx context.Context,
		importNoteId string,
	) ([]importnotedetailmodel.ImportNoteDetail, error)
	HandleBookQuantity(
		ctx context.Context,
		importNoteId string,
		bookTotalQuantityNeedUpdate map[string]int) error
}

type changeStatusImportNoteRepo struct {
	gen       generator.IdGenerator
	repo      ChangeStatusImportNoteRepo
	requester middleware.Requester
}

func NewChangeStatusImportNoteBiz(
	gen generator.IdGenerator,
	repo ChangeStatusImportNoteRepo,
	requester middleware.Requester) *changeStatusImportNoteRepo {
	return &changeStatusImportNoteRepo{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *changeStatusImportNoteRepo) ChangeStatusImportNote(
	ctx context.Context,
	importNoteId string,
	data *importnotemodel.ReqUpdateImportNote) error {
	if !biz.requester.IsHasFeature(common.ImportNoteChangeStatusFeatureCode) {
		return importnotemodel.ErrImportNoteChangeStatusNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	importNote, errGetImportNote := biz.repo.FindImportNote(ctx, importNoteId)
	if errGetImportNote != nil {
		return errGetImportNote
	}
	data.Id = importNoteId
	data.TotalPrice = importNote.TotalPrice
	data.SupplierId = importNote.SupplierId
	data.ClosedBy = biz.requester.GetUserId()

	if *importNote.Status != importnotemodel.InProgress {
		return importnotemodel.ErrImportNoteClosed
	}

	if *data.Status == importnotemodel.Done {
		supplierDebtId, errGenerateId := biz.gen.IdProcess(&data.Id)
		if errGenerateId != nil {
			return errGenerateId
		}

		if err := biz.repo.CreateSupplierDebt(ctx, *supplierDebtId, data); err != nil {
			return err
		}

		if err := biz.repo.UpdateDebtSupplier(ctx, data); err != nil {
			return err
		}

		importNoteDetails, errGetImportNoteDetails := biz.repo.FindListImportNoteDetail(
			ctx,
			importNoteId)
		if errGetImportNoteDetails != nil {
			return errGetImportNoteDetails
		}

		mapBookQuantity := getMapBookTotalQuantityNeedUpdated(importNoteDetails)
		if err := biz.repo.HandleBookQuantity(ctx, importNoteId, mapBookQuantity); err != nil {
			return err
		}
	}
	if err := biz.repo.UpdateImportNote(ctx, importNoteId, data); err != nil {
		return err
	}
	return nil
}

func getMapBookTotalQuantityNeedUpdated(
	importNoteDetails []importnotedetailmodel.ImportNoteDetail) map[string]int {
	result := make(map[string]int)
	for _, v := range importNoteDetails {
		result[v.BookId] += v.QuantityImport
	}
	return result
}
