package bookstore

//func (s *sqlStore) UpdatePriceBook(
//	ctx context.Context,
//	id string,
//	data *bookmodel.BookUpdatePrice) error {
//	db := s.db
//
//	if err := db.Table(common.TableBook).
//		Where("id = ?", id).
//		Updates(data).
//		Error; err != nil {
//		return common.ErrDB(err)
//	}
//
//	return nil
//}
