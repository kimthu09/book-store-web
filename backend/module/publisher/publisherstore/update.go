package publisherstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/publisher/publishermodel"
	"context"
)

func (s *sqlStore) UpdatePublisher(
	ctx context.Context,
	id string,
	data *publishermodel.ReqUpdatePublisher) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("name"); key {
			case "name":
				return publishermodel.ErrPublisherNameDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
