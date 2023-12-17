package supplierstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/importnote/importnotemodel"
	"context"
)

func (s *sqlStore) GetAllSupplier(
	ctx context.Context,
	moreKeys ...string) ([]importnotemodel.SimpleSupplier, error) {
	var result []importnotemodel.SimpleSupplier
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Order("name").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
