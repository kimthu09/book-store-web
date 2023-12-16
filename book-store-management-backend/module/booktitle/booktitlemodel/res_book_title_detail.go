package booktitlemodel

//type ResBookTitleDetail struct {
//	ID   string `json:"id"`
//	Name string `json:"name"`
//}

type ResBookTitleDetail struct {
	Data BookTitleDetail `json:"data"`
}

func NewResBookTitleDetail(data BookTitleDetail) ResBookTitleDetail {
	return ResBookTitleDetail{Data: data}
}
