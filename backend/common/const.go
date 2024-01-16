package common

// TABLE NAME
const (
	TableImage                    string = "Image"
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
	TableShopGeneral              string = "ShopGeneral"
	TableCustomer                 string = "Customer"
)

// FEATURE CODE
const (
	AuthorCreateFeatureCode = "AUTHOR_CREATE"
	AuthorViewFeatureCode   = "AUTHOR_VIEW"
	AuthorUpdateFeatureCode = "AUTHOR_UPDATE"

	CategoryCreateFeatureCode = "CATEGORY_CREATE"
	CategoryViewFeatureCode   = "CATEGORY_VIEW"
	CategoryUpdateFeatureCode = "CATEGORY_UPDATE"

	BookCreateFeatureCode       = "BOOK_CREATE"
	BookUpdateFeatureCode       = "BOOK_UPDATE"
	BookUpdateStatusFeatureCode = "BOOK_UPDATE"
	BookViewFeatureCode         = "BOOK_VIEW"

	BookTitleCreateFeatureCode = "BOOK_TITLE_CREATE"
	BookTitleUpdateFeatureCode = "BOOK_TITLE_UPDATE"
	BookTitleViewFeatureCode   = "BOOK_TITLE_VIEW"
	BookTitleDeleteFeatureCode = "BOOK_TITLE_DELETE"

	PublisherCreateFeatureCode = "PUBLISHER_CREATE"
	PublisherViewFeatureCode   = "PUBLISHER_VIEW"
	PublisherUpdateFeatureCode = "PUBLISHER_UPDATE"

	ImportNoteViewFeatureCode         = "IMPORT_NOTE_VIEW"
	ImportNoteCreateFeatureCode       = "IMPORT_NOTE_CREATE"
	ImportNoteChangeStatusFeatureCode = "IMPORT_NOTE_STATUS"

	InventoryCheckNoteViewFeatureCode   = "INVENTORY_NOTE_VIEW"
	InventoryCheckNoteCreateFeatureCode = "INVENTORY_NOTE_CREATE"

	SupplierViewFeatureCode       = "SUPPLIER_VIEW"
	SupplierCreateFeatureCode     = "SUPPLIER_CREATE"
	SupplierPayFeatureCode        = "SUPPLIER_PAY"
	SupplierUpdateInfoFeatureCode = "SUPPLIER_UPDATE_INFO"

	CustomerViewFeatureCode       = "CUSTOMER_VIEW"
	CustomerCreateFeatureCode     = "CUSTOMER_CREATE"
	CustomerUpdateInfoFeatureCode = "CUSTOMER_UPDATE_INFO"

	UserViewFeatureCode       = "USER_VIEW"
	UserUpdateInfoFeatureCode = "USER_UPDATE_INFO"

	InvoiceCreateFeatureCode = "INVOICE_CREATE"
	InvoiceViewFeatureCode   = "INVOICE_VIEW"

	SupplierDebtReportViewFeatureCode = "REPORT_VIEW_SUPPLIER"
	StockReportViewFeatureCode        = "REPORT_VIEW_STOCK"
	SaleReportViewFeatureCode         = "REPORT_VIEW_SALE"
)

const MaxLengthIdCanGenerate = 12
const MaxLengthOfFeatureCode = 30

const RoleAdminId = "admin"

const DefaultPass = "app123"

const CurrentUserStr = "current_user"

const MinuteVerifyEmail = 15

const (
	DefaultImageBook   string = "https://firebasestorage.googleapis.com/v0/b/coffee-shop-web.appspot.com/o/Default%2Fbook.jpg?alt=media"
	DefaultImageAvatar string = "https://firebasestorage.googleapis.com/v0/b/coffee-shop-web.appspot.com/o/Default%2Favatar.jpg?alt=media"
)
