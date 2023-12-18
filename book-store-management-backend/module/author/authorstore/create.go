package authorstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/author/authormodel"
	"context"
)

func (s *sqlStore) CreateAuthor(ctx context.Context, data *authormodel.Author) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("PRIMARY", "name"); key {
			case "PRIMARY":
				return authormodel.ErrAuthorIdDuplicate
			case "name":
				return authormodel.ErrAuthorNameDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
