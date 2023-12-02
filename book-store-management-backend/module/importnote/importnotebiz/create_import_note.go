package importnotebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/importnote/importnotemodel"
	"context"
)

type CreateImportNoteRepo interface {
	CheckBook(
		ctx context.Context,
		bookId string,
	) error
	CheckSupplier(
		ctx context.Context,
		supplierId string,
	) error
	HandleCreateImportNote(
		ctx context.Context,
		data *importnotemodel.ImportNoteCreate,
	) error
	UpdatePriceBook(
		ctx context.Context,
		bookId string,
		price float32,
	) error
}

type createImportNoteBiz struct {
	gen       generator.IdGenerator
	repo      CreateImportNoteRepo
	requester middleware.Requester
}

func NewCreateImportNoteBiz(
	gen generator.IdGenerator,
	repo CreateImportNoteRepo,
	requester middleware.Requester) *createImportNoteBiz {
	return &createImportNoteBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createImportNoteBiz) CreateImportNote(
	ctx context.Context,
	data *importnotemodel.ImportNoteCreate) error {
	if !biz.requester.IsHasFeature(common.ImportNoteCreateFeatureCode) {
		return importnotemodel.ErrImportNoteCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	data.Round()

	for _, v := range data.ImportNoteDetails {
		if err := biz.repo.CheckBook(ctx, v.BookId); err != nil {
			return err
		}
	}

	if err := handleImportNoteCreateId(biz.gen, data); err != nil {
		return err
	}

	if err := biz.repo.CheckSupplier(ctx, data.SupplierId); err != nil {
		return err
	}

	handleTotalPrice(data)

	if err := biz.repo.HandleCreateImportNote(ctx, data); err != nil {
		return err
	}

	for _, v := range data.ImportNoteDetails {
		if v.IsReplacePrice {
			if err := biz.repo.UpdatePriceBook(
				ctx, v.BookId, v.Price,
			); err != nil {
				return err
			}
		}
	}

	return nil
}

func handleImportNoteCreateId(
	gen generator.IdGenerator,
	data *importnotemodel.ImportNoteCreate) error {
	idImportNote, err := gen.IdProcess(data.Id)
	if err != nil {
		return err
	}
	data.Id = idImportNote

	for i := range data.ImportNoteDetails {
		data.ImportNoteDetails[i].ImportNoteId = *idImportNote
	}
	return nil
}

func handleTotalPrice(data *importnotemodel.ImportNoteCreate) {
	var totalPrice float32 = 0
	for _, importNoteDetail := range data.ImportNoteDetails {
		totalPrice += importNoteDetail.Price * importNoteDetail.QuantityImport
	}
	data.TotalPrice = totalPrice
}
