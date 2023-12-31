package importnoterepo

import (
	"book-store-management-backend/common/enum"
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/importnote/importnotemodel"
	"book-store-management-backend/module/importnotedetail/importnotedetailmodel"
	"book-store-management-backend/module/stockchangehistory/stockchangehistorymodel"
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
		data *importnotemodel.ReqUpdateImportNote,
	) error
}

type GetImportNoteDetailStore interface {
	FindListImportNoteDetail(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) ([]importnotedetailmodel.ImportNoteDetail, error)
}

type UpdateQuantityBookStore interface {
	UpdateQuantityBook(
		ctx context.Context,
		id string,
		data *bookmodel.BookUpdateQuantity,
	) error
	FindBook(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*bookmodel.Book, error)
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
		data *suppliermodel.ReqUpdateDebtSupplier,
	) error
}

type CreateSupplierDebtStore interface {
	CreateSupplierDebt(
		ctx context.Context,
		data *supplierdebtmodel.SupplierDebtCreate,
	) error
}
type StockChangeHistoryStore interface {
	CreateLisStockChangeHistory(
		ctx context.Context,
		data []stockchangehistorymodel.StockChangeHistory) error
}

type changeStatusImportNoteRepo struct {
	importNoteStore         ChangeStatusImportNoteStore
	importNoteDetailStore   GetImportNoteDetailStore
	bookStore               UpdateQuantityBookStore
	supplierStore           UpdateDebtOfSupplierStore
	supplierDebtStore       CreateSupplierDebtStore
	stockChangeHistoryStore StockChangeHistoryStore
}

func NewChangeStatusImportNoteRepo(
	importNoteStore ChangeStatusImportNoteStore,
	importNoteDetailStore GetImportNoteDetailStore,
	bookStore UpdateQuantityBookStore,
	supplierStore UpdateDebtOfSupplierStore,
	supplierDebtStore CreateSupplierDebtStore,
	stockChangeHistoryStore StockChangeHistoryStore) *changeStatusImportNoteRepo {
	return &changeStatusImportNoteRepo{
		importNoteStore:         importNoteStore,
		importNoteDetailStore:   importNoteDetailStore,
		bookStore:               bookStore,
		supplierStore:           supplierStore,
		supplierDebtStore:       supplierDebtStore,
		stockChangeHistoryStore: stockChangeHistoryStore,
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
	data *importnotemodel.ReqUpdateImportNote) error {
	if err := repo.importNoteStore.UpdateImportNote(
		ctx, importNoteId, data); err != nil {
		return err
	}
	return nil
}

func (repo *changeStatusImportNoteRepo) CreateSupplierDebt(
	ctx context.Context,
	supplierDebtId string,
	importNote *importnotemodel.ReqUpdateImportNote) error {
	supplier, err := repo.supplierStore.FindSupplier(
		ctx,
		map[string]interface{}{"id": importNote.SupplierId})
	if err != nil {
		return err
	}

	qtyBorrow := importNote.TotalPrice
	qtyLeft := supplier.Debt + qtyBorrow

	debtType := enum.Debt
	supplierDebtCreate := supplierdebtmodel.SupplierDebtCreate{
		Id:           supplierDebtId,
		SupplierId:   importNote.SupplierId,
		Quantity:     qtyBorrow,
		QuantityLeft: qtyLeft,
		DebtType:     &debtType,
		CreatedBy:    importNote.ClosedBy,
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
	importNote *importnotemodel.ReqUpdateImportNote) error {
	qtyUpdate := importNote.TotalPrice
	supplierUpdateDebt := suppliermodel.ReqUpdateDebtSupplier{
		Id:             &importNote.Id,
		QuantityUpdate: &qtyUpdate,
		CreatedBy:      importNote.ClosedBy,
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

func (repo *changeStatusImportNoteRepo) HandleBookQuantity(
	ctx context.Context,
	importNoteId string,
	bookTotalQuantityNeedUpdate map[string]int) error {
	var history []stockchangehistorymodel.StockChangeHistory
	for key, value := range bookTotalQuantityNeedUpdate {
		book, errFindBook := repo.bookStore.FindBook(
			ctx, map[string]interface{}{"id": key})
		if errFindBook != nil {
			return errFindBook
		}

		bookUpdate := bookmodel.BookUpdateQuantity{QuantityUpdate: value}
		if err := repo.bookStore.UpdateQuantityBook(
			ctx, key, &bookUpdate,
		); err != nil {
			return err
		}

		typeChange := stockchangehistorymodel.Import
		stockChangeHistory := stockchangehistorymodel.StockChangeHistory{
			Id:           importNoteId,
			BookId:       key,
			Quantity:     value,
			QuantityLeft: value + *book.Quantity,
			Type:         &typeChange,
		}
		history = append(history, stockChangeHistory)
	}

	if err := repo.stockChangeHistoryStore.CreateLisStockChangeHistory(
		ctx, history); err != nil {
		return err
	}

	return nil
}
