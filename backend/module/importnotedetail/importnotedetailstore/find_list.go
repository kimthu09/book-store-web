package importnotedetailstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/importnotedetail/importnotedetailmodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) FindListImportNoteDetail(ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) ([]importnotedetailmodel.ImportNoteDetail, error) {
	var data []importnotedetailmodel.ImportNoteDetail
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Table(common.TableImportNoteDetail).
		Where(conditions).
		Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return data, nil
}
