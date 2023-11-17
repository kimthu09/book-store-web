package importnoterepo

import (
	"book-store-management-backend/common/enum"
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/importnote/importnotemodel"
	"book-store-management-backend/module/importnotedetail/importnotedetailmodel"
	"book-store-management-backend/module/supplier/suppliermodel"
	"book-store-management-backend/module/supplierdebt/supplierdebtmodel"
	"context"
)

type ChangeStatusImportNoteStore interface {
	FindImportNote(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*importnotemodel.ImportNote, error)
	UpdateImportNote(
		ctx context.Context,
		id string,
		data *importnotemodel.ImportNoteUpdate,
	) error
}

type GetImportNoteDetailStore interface {
	FindListImportNoteDetail(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) ([]importnotedetailmodel.ImportNoteDetail, error)
}

type UpdateAmountBookStore interface {
	UpdateAmountBook(
		ctx context.Context,
		id string,
		data *bookmodel.BookUpdateAmount,
	) error
}

type UpdateDebtOfSupplierStore interface {
	FindSupplier(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*suppliermodel.Supplier, error)
	UpdateSupplierDebt(
		ctx context.Context,
		id string,
		data *suppliermodel.SupplierUpdateDebt,
	) error
}

type CreateSupplierDebtStore interface {
	CreateSupplierDebt(
		ctx context.Context,
		data *supplierdebtmodel.SupplierDebtCreate,
	) error
}

type changeStatusImportNoteRepo struct {
	importNoteStore       ChangeStatusImportNoteStore
	importNoteDetailStore GetImportNoteDetailStore
	bookStore             UpdateAmountBookStore
	supplierStore         UpdateDebtOfSupplierStore
	supplierDebtStore     CreateSupplierDebtStore
}

func NewChangeStatusImportNoteRepo(
	importNoteStore ChangeStatusImportNoteStore,
	importNoteDetailStore GetImportNoteDetailStore,
	bookStore UpdateAmountBookStore,
	supplierStore UpdateDebtOfSupplierStore,
	supplierDebtStore CreateSupplierDebtStore) *changeStatusImportNoteRepo {
	return &changeStatusImportNoteRepo{
		importNoteStore:       importNoteStore,
		importNoteDetailStore: importNoteDetailStore,
		bookStore:             bookStore,
		supplierStore:         supplierStore,
		supplierDebtStore:     supplierDebtStore,
	}
}

func (repo *changeStatusImportNoteRepo) FindImportNote(
	ctx context.Context,
	importNoteId string) (*importnotemodel.ImportNote, error) {
	importNote, err := repo.importNoteStore.FindImportNote(
		ctx, map[string]interface{}{"id": importNoteId},
	)
	if err != nil {
		return nil, err
	}
	return importNote, nil
}

func (repo *changeStatusImportNoteRepo) UpdateImportNote(
	ctx context.Context,
	importNoteId string,
	data *importnotemodel.ImportNoteUpdate) error {
	if err := repo.importNoteStore.UpdateImportNote(
		ctx, importNoteId, data); err != nil {
		return err
	}
	return nil
}

func (repo *changeStatusImportNoteRepo) CreateSupplierDebt(
	ctx context.Context,
	supplierDebtId string,
	importNote *importnotemodel.ImportNoteUpdate) error {
	supplier, err := repo.supplierStore.FindSupplier(
		ctx,
		map[string]interface{}{"id": importNote.SupplierId})
	if err != nil {
		return err
	}

	amountBorrow := -importNote.TotalPrice
	amountLeft := supplier.Debt + amountBorrow

	debtType := enum.Debt
	supplierDebtCreate := supplierdebtmodel.SupplierDebtCreate{
		Id:         supplierDebtId,
		SupplierId: importNote.SupplierId,
		Amount:     amountBorrow,
		AmountLeft: amountLeft,
		DebtType:   &debtType,
		CreateBy:   importNote.CloseBy,
	}

	if err := repo.supplierDebtStore.CreateSupplierDebt(
		ctx, &supplierDebtCreate,
	); err != nil {
		return err
	}
	return nil
}

func (repo *changeStatusImportNoteRepo) UpdateDebtSupplier(
	ctx context.Context,
	importNote *importnotemodel.ImportNoteUpdate) error {
	amountUpdate := -importNote.TotalPrice
	supplierUpdateDebt := suppliermodel.SupplierUpdateDebt{
		Amount: &amountUpdate,
	}
	if err := repo.supplierStore.UpdateSupplierDebt(
		ctx, importNote.SupplierId, &supplierUpdateDebt,
	); err != nil {
		return err
	}
	return nil
}

func (repo *changeStatusImportNoteRepo) FindListImportNoteDetail(
	ctx context.Context,
	importNoteId string) ([]importnotedetailmodel.ImportNoteDetail, error) {
	importNoteDetails, errGetImportNoteDetails :=
		repo.importNoteDetailStore.FindListImportNoteDetail(
			ctx,
			map[string]interface{}{"importNoteId": importNoteId},
		)
	if errGetImportNoteDetails != nil {
		return nil, errGetImportNoteDetails
	}
	return importNoteDetails, nil
}

func (repo *changeStatusImportNoteRepo) HandleBookAmount(
	ctx context.Context,
	bookTotalAmountNeedUpdate map[string]float32) error {
	for key, value := range bookTotalAmountNeedUpdate {
		bookUpdate := bookmodel.BookUpdateAmount{AmountUpdate: value}
		if err := repo.bookStore.UpdateAmountBook(
			ctx, key, &bookUpdate,
		); err != nil {
			return err
		}
	}
	return nil
}
