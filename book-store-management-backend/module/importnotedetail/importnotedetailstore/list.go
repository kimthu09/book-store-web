package importnotedetailstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/importnotedetail/importnotedetailmodel"
	"context"
)

func (s *sqlStore) ListImportNoteDetail(
	ctx context.Context,
	importNoteId string,
	paging *common.Paging) ([]importnotedetailmodel.ImportNoteDetail, error) {
	var result []importnotedetailmodel.ImportNoteDetail
	db := s.db

	db = db.Table(common.TableImportNoteDetail)

	db = db.Where("importNoteId = ?", importNoteId)

	dbTemp, errPaging := common.HandlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp.Limit(int(paging.Limit))

	if err := db.
		Preload("Book").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
