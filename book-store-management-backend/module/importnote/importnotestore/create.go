package importnotestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/importnote/importnotemodel"
	"context"
)

func (s *sqlStore) CreateImportNote(
	ctx context.Context,
	data *importnotemodel.ImportNoteCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
