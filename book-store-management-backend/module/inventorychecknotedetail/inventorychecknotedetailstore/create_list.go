package inventorychecknotedetailstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
	"context"
)

func (s *sqlStore) CreateListInventoryCheckNoteDetail(
	ctx context.Context,
	data []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate) error {
	db := s.db

	if err := db.CreateInBatches(data, len(data)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
