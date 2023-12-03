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
		Preload("CreateByUser").
		Preload("CloseByUser").
		Order("createAt desc").
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
		if filter.DateFromCreateAt != nil {
			timeFrom := time.Unix(*filter.DateFromCreateAt, 0)
			db = db.Where("createAt >= ?", timeFrom)
		}
		if filter.DateToCreateAt != nil {
			timeTo := time.Unix(*filter.DateToCreateAt, 0)
			db = db.Where("createAt <= ?", timeTo)
		}
		if filter.DateFromCloseAt != nil {
			timeFrom := time.Unix(*filter.DateFromCloseAt, 0)
			db = db.Where("closeAt >= ?", timeFrom)
		}
		if filter.DateToCloseAt != nil {
			timeTo := time.Unix(*filter.DateToCloseAt, 0)
			db = db.Where("closeAt <= ?", timeTo)
		}
		if filter.Supplier != nil {
			db = db.
				Joins("JOIN Supplier ON ResDetailImportNote.supplierId = Supplier.id").
				Where("Supplier.name LIKE ?", "%"+*filter.Supplier+"%")
		}
		if filter.CreateBy != nil {
			db = db.
				Joins("JOIN MUser AS CreateByUser ON ResDetailImportNote.createBy = CreateByUser.id").
				Where("CloseByUser.name LIKE ?", "%"+*filter.CreateBy+"%")
		}
		if filter.CloseBy != nil {
			db = db.
				Joins("JOIN MUser AS CloseByUser ON ResDetailImportNote.closeBy = CloseByUser.id").
				Where("CloseByUser.name LIKE ?", "%"+*filter.CloseBy+"%")
		}
	}
}
