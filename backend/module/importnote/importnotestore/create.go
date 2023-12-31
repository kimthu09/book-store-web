package importnotestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/importnote/importnotemodel"
	"context"
)

func (s *sqlStore) CreateImportNote(
	ctx context.Context,
	data *importnotemodel.ReqCreateImportNote) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("PRIMARY"); key {
			case "PRIMARY":
				return importnotemodel.ErrImportNoteIdDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
