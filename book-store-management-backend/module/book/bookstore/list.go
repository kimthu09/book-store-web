package bookstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/book/bookmodel"
	"context"
	"gorm.io/gorm"
	"strings"
)

func (s *sqlStore) ListBook(
	ctx context.Context,
	filter *bookmodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging,
	moreKeys ...string) ([]bookmodel.ResDetailUnitBook, error) {
	var result []bookmodel.ResDetailUnitBook
	db := s.db

	db = db.Table(common.TableBook)

	handleFilter(db, filter, propertiesContainSearchKey)

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
	filter *bookmodel.Filter,
	propertiesContainSearchKey []string) {
	if filter != nil {
		if filter.SearchKey != "" {
			db = common.GetWhereClause(db, filter.SearchKey, propertiesContainSearchKey)
		}
		if filter.MinSellPrice != nil {
			db = db.Where("sellPrice >= ?", filter.MinSellPrice)
		}
		if filter.MaxSellPrice != nil {
			db = db.Where("sellPrice <= ?", filter.MaxSellPrice)
		}
		if filter.AuthorIds != nil || filter.CategoryIds != nil {
			db = db.Joins("JOIN BookTitle ON Book.booktitleid = BookTitle.id")
			if filter.AuthorIds != nil {
				authorIds := strings.Split(*filter.AuthorIds, "|")
				query := ""
				params := make([]interface{}, 0)
				for i, authorId := range authorIds {
					params = append(params, "%"+authorId+"%")
					if i == 0 {
						query += "BookTitle.authorIds LIKE ?"
					} else {
						query += " OR BookTitle.authorIds LIKE ?"
					}
				}
				db = db.Where(query, params...)
			}

			if filter.CategoryIds != nil {
				categoryIds := strings.Split(*filter.CategoryIds, "|")
				query := ""
				params := make([]interface{}, 0)
				for i, categoryId := range categoryIds {
					params = append(params, "%"+categoryId+"%")
					if i == 0 {
						query += "BookTitle.categoryIds LIKE ?"
					} else {
						query += " OR BookTitle.categoryIds LIKE ?"
					}
				}
				db = db.Where(query, params...)
			}
		}
		if filter.PublisherId != nil {
			db = db.
				Joins("JOIN Publisher ON Book.publisherid = Publisher.id").
				Where("Publisher.id = ?", filter.PublisherId)
		}
	}
}
