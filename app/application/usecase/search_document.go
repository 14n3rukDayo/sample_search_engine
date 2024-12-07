package usecase

import (
	"main/application/query"
	documentE "main/domain/document/entity"
)

type SearchDocument interface {
	Execute(searchWord string) (documents []documentE.Document, err error)
}

type searchDocument struct {
	sq query.Searcher
}

func NewSearchDocument(sq query.Searcher) SearchDocument {
	return &searchDocument{sq: sq}
}

func (s *searchDocument) Execute(searchWord string) (documents []documentE.Document, err error) {
	documents, err = s.sq.SearchAnd(searchWord)
	if err != nil {
		return nil, err
	}
	if documents == nil {
		return []documentE.Document{}, nil
	}
	return documents, nil
}
