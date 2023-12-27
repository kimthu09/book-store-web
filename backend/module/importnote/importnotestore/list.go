package importnotestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/importnote/importnotemodel"
	"context"
	"gorm.io/gorm"
	"time"
)

func (s *sqlStore) ListImportNote(
	ctx context.Context,
	filter *importnotemodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging) ([]importnotemodel.ImportNote, error) {
	var result []importnotemodel.ImportNote
	db := s.db

	db = db.Table(common.TableImportNote)

	handleFilter(db, filter, propertiesContainSearchKey)

	dbTemp, errPaging := common.HandlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	if err := db.
		Preload("Supplier").
		Preload("CreatedByUser").
		Preload("ClosedByUser").
		Order("createdAt desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func handleFilter(
	db *gorm.DB,
	filter *importnotemodel.Filter,
	propertiesContainSearchKey []string) {
	if filter != nil {
		if filter.SearchKey != "" {
			db = common.GetWhereClause(db, filter.SearchKey, propertiesContainSearchKey)
		}
		if filter.Status != "" {
			db = db.Where("status = ?", filter.Status)
		}
		if filter.MinPrice != nil {
			db = db.Where("totalPrice >= ?", filter.MinPrice)
		}
		if filter.MaxPrice != nil {
			db = db.Where("totalPrice <= ?", filter.MaxPrice)
		}
		if filter.DateFromCreatedAt != nil {
			timeFrom := time.Unix(*filter.DateFromCreatedAt, 0)
			db = db.Where("createdAt >= ?", timeFrom)
		}
		if filter.DateToCreatedAt != nil {
			timeTo := time.Unix(*filter.DateToCreatedAt, 0)
			db = db.Where("createdAt <= ?", timeTo)
		}
		if filter.DateFromClosedAt != nil {
			timeFrom := time.Unix(*filter.DateFromClosedAt, 0)
			db = db.Where("closedAt >= ?", timeFrom)
		}
		if filter.DateToClosedAt != nil {
			timeTo := time.Unix(*filter.DateToClosedAt, 0)
			db = db.Where("closedAt <= ?", timeTo)
		}
		if filter.Supplier != nil {
			db = db.
				Joins("JOIN Supplier ON ImportNote.supplierId = Supplier.id").
				Where("Supplier.name LIKE ?", "%"+*filter.Supplier+"%")
		}
		if filter.CreatedBy != nil {
			db = db.
				Joins("JOIN MUser AS CreatedByUser ON ImportNote.createdBy = CreatedByUser.id").
				Where("CreatedByUser.name LIKE ?", "%"+*filter.CreatedBy+"%")
		}
		if filter.ClosedBy != nil {
			db = db.
				Joins("JOIN MUser AS ClosedByUser ON ImportNote.closedBy = ClosedByUser.id").
				Where("ClosedByUser.name LIKE ?", "%"+*filter.ClosedBy+"%")
		}
	}
}