package query

import (
	documentE "main/domain/document/entity"
	documentR "main/domain/document/repository"
	tokensF "main/domain/document/valueobject/tokens/factory"
	invertedIndexE "main/domain/invertedindex/entity"
	invertedIndexR "main/domain/invertedindex/repository"
	"sort"

	"main/util"
)

type Searcher interface {
	SearchAnd(searchWord string) (document []documentE.Document, err error)
}
type searcher struct {
	ir invertedIndexR.InvertedIndex
	dr documentR.Document
	tf tokensF.Tokens
}

func NewSearcher(ir invertedIndexR.InvertedIndex, dr documentR.Document, tf tokensF.Tokens) Searcher {
	return &searcher{ir: ir, dr: dr, tf: tf}
}
func (s *searcher) SearchAnd(searchWord string) (document []documentE.Document, err error) {
	searchWords, err := s.tf.CreateTokens(searchWord)
	if err != nil {
		return nil, err
	}
	if len(searchWord) == 0 {
		return nil, nil
	}
	invertedIndexes := []invertedIndexE.InvertedIndex{}
	errors := []error{}
	for _, searchWord := range searchWords {
		invertedIndex, err := s.ir.Get(searchWord)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		invertedIndexes = append(invertedIndexes, invertedIndex)
	}
	if len(errors) != 0 {
		return nil, util.NewInternalServerError()
	}
	if len(invertedIndexes) == 0 {
		return nil, nil
	}
	type DocumentMapValue struct {
		Count int
		Score float64
	}
	documentIdsCount := make(map[int]DocumentMapValue)
	for _, ininvertedIndex := range invertedIndexes {
		invertedIndex := ininvertedIndex.Get().DocumentScore
		for _, documentScore := range invertedIndex {
			countSum := documentIdsCount[documentScore.DocumentId].Count
			scoreSum := documentIdsCount[documentScore.DocumentId].Score
			documentIdsCount[documentScore.DocumentId] = DocumentMapValue{Count: countSum + 1, Score: scoreSum + documentScore.Score}
		}
	}
	documentIdsHit := []int{}
	for documentId, value := range documentIdsCount {
		if value.Count == len(invertedIndexes) {
			documentIdsHit = append(documentIdsHit, documentId)
		}
	}
	sort.Slice(documentIdsHit, func(i, j int) bool {
		return documentIdsCount[documentIdsHit[i]].Score > documentIdsCount[documentIdsHit[j]].Score
	})
	documents := []documentE.Document{}
	errors = []error{}
	for _, documentId := range documentIdsHit {
		document, err := s.dr.Get(documentId)
		if err != nil {
			errors = append(errors, err)
		}
		documents = append(documents, document)
	}
	if len(errors) != 0 {
		return nil, util.NewInternalServerError()
	}

	return documents, nil

}
