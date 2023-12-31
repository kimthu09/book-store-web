package authorstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/author/authormodel"
	"context"
)

func (s *sqlStore) UpdateAuthor(
	ctx context.Context,
	id string,
	data *authormodel.ReqUpdateAuthor) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("name"); key {
			case "name":
				return authormodel.ErrAuthorNameDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
