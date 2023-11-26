package importnotestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/importnote/importnotemodel"
	"context"
	"gorm.io/gorm"
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

	dbTemp, errPaging := handlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	if err := db.
		Limit(int(paging.Limit)).
		Order("createAt desc").
		Preload("Supplier").
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
	}
}

func handlePaging(db *gorm.DB, paging *common.Paging) (*gorm.DB, error) {
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	offset := (paging.Page - 1) * paging.Limit
	db = db.Offset(int(offset))

	return db, nil
}
