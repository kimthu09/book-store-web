package importnotestore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/importnote/importnotemodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) FindImportNote(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*importnotemodel.ImportNote, error) {
	var data importnotemodel.ImportNote
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
