package common

import "time"

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

	ImportNoteViewFeatureCode           = "IMP_VIEW"
	ImportNoteCreateFeatureCode         = "IMP_CREATE"
	ImportNoteChangeStatusFeatureCode   = "IMP_UP_STATE"
	InventoryCheckNoteViewFeatureCode   = "INV_VIEW"
	InventoryCheckNoteCreateFeatureCode = "INV_CREATE"
	SupplierViewFeatureCode             = "SUP_VIEW"
	SupplierCreateFeatureCode           = "SUP_CREATE"
	SupplierPayFeatureCode              = "SUP_PAY"
	SupplierUpdateInfoFeatureCode       = "SUP_UP_INFO"
)

const MaxLengthIdCanGenerate = 12

const RoleAdminId = "admin"

const DefaultPass = "app123"

const CurrentUserStr = "current_user"

var (
	VietNamLocation = time.FixedZone("UTC+7", 7*60*60)
)
