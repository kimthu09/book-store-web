package common

// TABLE NAME
const (
	TableAuthor                   string = "Author"
	TableCategory                 string = "Category"
	TableUser                     string = "MUser"
	TableBook                     string = "Book"
	TableBookTitle                string = "BookTitle"
	TablePublisher                string = "Publisher"
	TableSupplier                 string = "Supplier"
	TableSupplierDebt             string = "SupplierDebt"
	TableImportNote               string = "ImportNote"
	TableImportNoteDetail         string = "ImportNoteDetail"
	TableInventoryCheckNote       string = "InventoryCheckNote"
	TableInventoryCheckNoteDetail string = "InventoryCheckNoteDetail"
	TableRole                     string = "Role"
	TableRoleFeature              string = "RoleFeature"
	TableFeature                  string = "Feature"
	TableStockChangeHistory       string = "StockChangeHistory"
	TableInvoice                  string = "Invoice"
	TableInvoiceDetail            string = "InvoiceDetail"
	TableSupplierDebtReport       string = "SupplierDebtReport"
	TableSupplierDebtReportDetail string = "SupplierDebtReportDetail"
	TableStockReport              string = "StockReport"
	TableStockReportDetail        string = "StockReportDetail"
)

// FEATURE CODE
const (
	AuthorCreateFeatureCode = "AUTHOR_CREATE"
	AuthorViewFeatureCode   = "AUTHOR_VIEW"
	AuthorUpdateFeatureCode = "AUTHOR_UPDATE"
	AuthorDeleteFeatureCode = "AUTHOR_DELETE"

	CategoryCreateFeatureCode = "CATEGORY_CREATE"
	CategoryViewFeatureCode   = "CATEGORY_VIEW"
	CategoryUpdateFeatureCode = "CATEGORY_UPDATE"
	CategoryDeleteFeatureCode = "CATEGORY_DELETE"

	BookCreateFeatureCode = "BOOK_CREATE"
	BookUpdateFeatureCode = "BOOK_UPDATE"
	BookViewFeatureCode   = "BOOK_VIEW"
	BookDeleteFeatureCode = "BOOK_DELETE"

	BookTitleCreateFeatureCode = "BOOK_TITLE_CREATE"
	BookTitleUpdateFeatureCode = "BOOK_TITLE_UPDATE"
	BookTitleViewFeatureCode   = "BOOK_TITLE_VIEW"
	BookTitleDeleteFeatureCode = "BOOK_TITLE_DELETE"

	PublisherCreateFeatureCode = "PUBLISHER_CREATE"
	PublisherViewFeatureCode   = "PUBLISHER_VIEW"

	ImportNoteViewFeatureCode         = "IMPORT_NOTE_VIEW"
	ImportNoteCreateFeatureCode       = "IMPORT_NOTE_CREATE"
	ImportNoteChangeStatusFeatureCode = "IMPORT_NOTE_STATUS"

	InventoryCheckNoteViewFeatureCode   = "INVENTORY_NOTE_VIEW"
	InventoryCheckNoteCreateFeatureCode = "INVENTORY_NOTE_CREATE"

	SupplierViewFeatureCode       = "SUPPLIER_VIEW"
	SupplierCreateFeatureCode     = "SUPPLIER_CREATE"
	SupplierPayFeatureCode        = "SUPPLIER_PAY"
	SupplierUpdateInfoFeatureCode = "SUPPLIER_UPDATE_INFO"

	UserViewFeatureCode         = "USER_VIEW"
	UserUpdateInfoFeatureCode   = "USER_UPDATE_INFO"
	UserUpdateStatusFeatureCode = "USER_UPDATE_STATE"

	InvoiceViewFeatureCode   = "INVOICE_VIEW"
	InvoiceCreateFeatureCode = "INVOICE_CREATE"

	SupplierDebtReportViewFeatureCode = "REPORT_VIEW_SUPPLIER"
	StockReportViewFeatureCode        = "REPORT_VIEW_STOCK"
	SaleReportViewFeatureCode         = "REPORT_VIEW_SALE"
)

const MaxLengthIdCanGenerate = 12
const MaxLengthOfFeatureCode = 30

const RoleAdminId = "admin"

const DefaultPass = "app123"

const CurrentUserStr = "current_user"
