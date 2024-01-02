package importnotestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/importnote/importnotemodel"
	"context"
)

func (s *sqlStore) GetAllIdImportNote(
	ctx context.Context) ([]importnotemodel.ImportNoteId, error) {
	var result []importnotemodel.ImportNoteId
	db := s.db

	db = db.Table(common.TableImportNote)

	if err := db.
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
