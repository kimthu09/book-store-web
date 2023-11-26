package common

type Paging struct {
	Page  int64 `json:"page" form:"page"`
	Limit int64 `json:"limit" form:"-"`
	Total int64 `json:"total" form:"-"`
}

func (p *Paging) Fulfill() {
	if p.Page <= 0 {
		p.Page = 1
	}

	p.Limit = 10
}
