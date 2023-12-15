package booktitlemodel

import "book-store-management-backend/common"

type ResListBookTitle struct {
	Data   []BookTitleDetail `json:"data"`
	Paging common.Paging     `json:"paging,omitempty"`
	Filter Filter            `json:"filter,omitempty"`
}

func NewResListBookTitle(data []BookTitleDetail, paging common.Paging, filter Filter) *ResListBookTitle {
	return &ResListBookTitle{Data: data, Paging: paging, Filter: filter}
}
