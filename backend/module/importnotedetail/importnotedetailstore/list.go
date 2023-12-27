package importnotedetailstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/importnotedetail/importnotedetailmodel"
	"context"
)

func (s *sqlStore) ListImportNoteDetail(
	ctx context.Context,
	importNoteId string) ([]importnotedetailmodel.ImportNoteDetail, error) {
	var result []importnotedetailmodel.ImportNoteDetail
	db := s.db

	db = db.Table(common.TableImportNoteDetail)

	db = db.Where("importNoteId = ?", importNoteId)

	if err := db.
		Preload("Book").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
