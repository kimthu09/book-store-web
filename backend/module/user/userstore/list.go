package userstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/user/usermodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) ListUser(
	ctx context.Context,
	userSearch string,
	filter *usermodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging,
	moreKeys ...string) ([]usermodel.ResUser, error) {
	var result []usermodel.ResUser
	db := s.db

	db = db.Table(common.TableUser)

	handleFilter(db, filter, propertiesContainSearchKey)

	db = db.Where("id <> ?", userSearch)

	dbTemp, errPaging := common.HandlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

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

func handleFilter(
	db *gorm.DB,
	filter *usermodel.Filter,
	propertiesContainSearchKey []string) {
	if filter != nil {
		if filter.SearchKey != "" {
			db = common.GetWhereClause(db, filter.SearchKey, propertiesContainSearchKey)
		}
		if filter.IsActive != nil {
			if *filter.IsActive {
				db = db.Where("isActive = ?", true)
			} else {
				db = db.Where("isActive = ?", false)
			}
		}
	}
}
