package importnotedetailstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/importnotedetail/importnotedetailmodel"
	"context"
)

func (s *sqlStore) CreateListImportNoteDetail(
	ctx context.Context,
	data []importnotedetailmodel.ImportNoteDetailCreate) error {
	db := s.db

	if err := db.CreateInBatches(data, len(data)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
