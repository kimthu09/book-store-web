package booktitlestore

import "context"

func (store *sqlStore) DetailBookTitle(ctx context.Context, id string) (*BookTitleDBModel, error) {
	var result BookTitleDBModel
	db := store.db.Table(result.TableName()).Where("id = ?", id).First(&result)
	if db.Error != nil {
		return nil, db.Error
	}
	return &result, nil
}
