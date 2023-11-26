package common

// TABLE NAME
const (
	TableAuthor                   string = "Author"
	TableUser                     string = "MUser"
	TableBook                     string = "Book"
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

	BookCreateFeatureCode = "BOOK_CREATE"
	BookViewFeatureCode   = "BOOK_VIEW"

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
