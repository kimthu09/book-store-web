package importnotemodel

import "book-store-management-backend/common"

type ReqUpdateImportNote struct {
	CloseBy    string            `json:"-" gorm:"column:closeBy;"`
	Id         string            `json:"-" gorm:"-"`
	SupplierId string            `json:"-" gorm:"-"`
	TotalPrice float32           `json:"-" gorm:"-"`
	Status     *ImportNoteStatus `json:"status" gorm:"column:status;" example:"Done"`
}

func (*ReqUpdateImportNote) TableName() string {
	return common.TableImportNote
}

func (data *ReqUpdateImportNote) Validate() *common.AppError {
	if data.Status == nil {
		return ErrImportNoteStatusEmpty
	}
	if *data.Status == InProgress {
		return ErrImportNoteStatusInvalid
	}
	return nil
}
