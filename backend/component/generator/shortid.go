package generator

import (
	"book-store-management-backend/common"
	"github.com/teris-io/shortid"
)

type shortIdGenerator struct {
}

func NewShortIdGenerator() *shortIdGenerator {
	return &shortIdGenerator{}
}

func (*shortIdGenerator) GenerateId() (string, error) {
	return shortid.Generate()
}
func (g *shortIdGenerator) IdProcess(id *string) (*string, error) {
	if id != nil && len(*id) != 0 {
		if len(*id) > common.MaxLengthIdCanGenerate {
			return nil, common.ErrIdIsTooLong()
		}
		return id, nil
	} else {
		idGenerated, err := g.GenerateId()
		if err != nil {
			return nil, common.ErrInternal(err)
		}

		return &idGenerated, nil
	}
}
