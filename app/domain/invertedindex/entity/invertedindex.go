package entity

import documentScoreVO "main/domain/invertedindex/valueobject"

type InvertedIndex interface {
	Get() *invertedIndex
	GetDocumentNum() int
}
type invertedIndex struct {
	Token         string
	DocumentScore []documentScoreVO.DocumentScore
}

func NewInvertedIndex(token string, documentIds []documentScoreVO.DocumentScore) InvertedIndex {
	return &invertedIndex{Token: token, DocumentScore: documentIds}
}

func (ii *invertedIndex) Get() *invertedIndex {
	return ii
}

func (ii *invertedIndex) GetDocumentNum() int {
	num := len(ii.DocumentScore)
	return num
}
