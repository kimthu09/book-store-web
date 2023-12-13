package common

type Paging struct {
	Page  int64 `json:"page" form:"page" example:"1"`
	Limit int64 `json:"limit" form:"limit" example:"10"`
	Total int64 `json:"total" example:"11" required:"false"`
}

func (p *Paging) Fulfill() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 10
	}
}
