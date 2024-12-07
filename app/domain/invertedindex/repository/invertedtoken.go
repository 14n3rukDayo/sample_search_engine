package repository

import invertedIndexE "main/domain/invertedindex/entity"

type InvertedIndex interface {
	MultiUpsert(invertedIndexes []invertedIndexE.InvertedIndex) error
	Get(word string) (invertedIndex invertedIndexE.InvertedIndex, err error)
	AddAllDL(invertedIndexes []invertedIndexE.InvertedIndex) error
	GetAllDLNum() (allDLNum int, err error)
}
