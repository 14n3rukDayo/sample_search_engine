package usecase

import (
	"context"
	"main/application/service"
	documentF "main/domain/document/factory"
)

type AddDocument interface {
	Execute(description string) error
}

type addDocument struct {
	ds service.Document
	df documentF.Document
}

func NewAddDocumentUsecase(ds service.Document, df documentF.Document) AddDocument {
	return &addDocument{ds: ds, df: df}
}

func (du *addDocument) Execute(description string) error {
	document, err := du.df.NewDocument(context.Background(), description)
	if err != nil {
		return err
	}
	err = du.ds.Add(document)
	if err != nil {
		return err
	}
	return nil
}
