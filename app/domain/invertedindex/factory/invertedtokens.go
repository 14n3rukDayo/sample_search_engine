package factory

import (
	documentE "main/domain/document/entity"
	invertedIndexE "main/domain/invertedindex/entity"
	documentScoreVO "main/domain/invertedindex/valueobject"
)

type InvertedIndex interface {
	CreateInvertedIndexes(document documentE.Document) []invertedIndexE.InvertedIndex
}

type invertedIndex struct {
}

func NewInvertedIndexFactory() InvertedIndex {
	return &invertedIndex{}
}

func (ii *invertedIndex) CreateInvertedIndexes(document documentE.Document) []invertedIndexE.InvertedIndex {
	tokens := document.Get().Tokens
	documentId := document.Get().DocumentId
	var invertedIndexes []invertedIndexE.InvertedIndex
	for _, token := range tokens {
		var documentIds []documentScoreVO.DocumentScore
		documentIds = append(documentIds, documentScoreVO.DocumentScore{DocumentId: documentId, Score: token.Score})
		invertedIndex := invertedIndexE.NewInvertedIndex(token.Token, documentIds)
		invertedIndexes = append(invertedIndexes, invertedIndex)
	}
	return invertedIndexes
}
