package importnotebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/importnote/importnotemodel"
	"context"
)

type SeeDetailImportNoteRepo interface {
	SeeDetailImportNote(
		ctx context.Context,
		importNoteId string,
		paging *common.Paging,
	) (*importnotemodel.ImportNote, error)
}

type seeDetailImportNoteBiz struct {
	repo      SeeDetailImportNoteRepo
	requester middleware.Requester
}

func NewSeeDetailImportNoteBiz(
	repo SeeDetailImportNoteRepo,
	requester middleware.Requester) *seeDetailImportNoteBiz {
	return &seeDetailImportNoteBiz{repo: repo, requester: requester}
}

func (biz *seeDetailImportNoteBiz) SeeDetailImportNote(
	ctx context.Context,
	importNoteId string,
	paging *common.Paging) (*importnotemodel.ImportNote, error) {
	if !biz.requester.IsHasFeature(common.ImportNoteViewFeatureCode) {
		return nil, importnotemodel.ErrImportNoteViewNoPermission
	}

	importNote, err := biz.repo.SeeDetailImportNote(
		ctx,
		importNoteId,
		paging)

	if err != nil {
		return nil, err
	}

	return importNote, nil
}
