package common

import (
	"gorm.io/gorm"
	"strings"
)

func GetWhereClause(
	db *gorm.DB,
	searchKey string,
	propertiesContainSearchKey []string) *gorm.DB {
	conditions := make([]string, len(propertiesContainSearchKey))
	args := make([]interface{}, len(propertiesContainSearchKey))

	for i, prop := range propertiesContainSearchKey {
		conditions[i] = prop + " LIKE ?"
		args[i] = "%" + searchKey + "%"
	}

	whereClause := strings.Join(conditions, " OR ")

	return db.Where(whereClause, args...)
}

func HandlePaging(db *gorm.DB, paging *Paging) (*gorm.DB, error) {
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, ErrDB(err)
	}

	offset := (paging.Page - 1) * paging.Limit
	db = db.Offset(int(offset)).Limit(int(paging.Limit))

	return db, nil
}
