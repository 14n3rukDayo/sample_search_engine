package repository

import documentE "main/domain/document/entity"

type Document interface {
	Add(document documentE.Document) error
	Get(documentId int) (document documentE.Document, err error)
	GenerateID() (documentId int, err error)
	GetTotalNum() (total int, err error)
	GetPrefixDocumentIdReserved() string
}
