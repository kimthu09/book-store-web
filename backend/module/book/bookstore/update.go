package bookstore

import "context"

func (s *sqlStore) UpdateBook(ctx context.Context, id string, data *BookDBModel) error {
	data.ID = nil
	db := s.db.Table(data.TableName()).Where("id = ? and isActive = ?", id, "1").Updates(data)
	if err := db.Error; err != nil {
		return err
	}
	return nil
}
