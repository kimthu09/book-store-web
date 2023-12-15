package supplierstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/importnote/importnotemodel"
	"context"
)

func (s *sqlStore) ListAllSupplier(
	ctx context.Context) ([]importnotemodel.SimpleSupplier, error) {
	var result []importnotemodel.SimpleSupplier
	db := s.db

	db = db.Table(common.TableSupplier)

	if err := db.
		Order("name").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
