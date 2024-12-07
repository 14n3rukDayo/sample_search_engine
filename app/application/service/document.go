package service

import (
	documentE "main/domain/document/entity"
	documentR "main/domain/document/repository"
	invertedIndexF "main/domain/invertedindex/factory"
	invertedIndexR "main/domain/invertedindex/repository"
)

type Document interface {
	Add(document documentE.Document) error
}

type document struct {
	dr  documentR.Document
	iif invertedIndexF.InvertedIndex
	ir  invertedIndexR.InvertedIndex
}

func NewDocumentService(dr documentR.Document, iif invertedIndexF.InvertedIndex, ir invertedIndexR.InvertedIndex) Document {
	return &document{dr: dr, iif: iif, ir: ir}
}
func (d *document) Add(document documentE.Document) error {
	err := d.dr.Add(document)
	if err != nil {
		return err
	}
	invertedIndexes := d.iif.CreateInvertedIndexes(document)
	err = d.ir.MultiUpsert(invertedIndexes)
	if err != nil {
		return err
	}
	err = d.ir.AddAllDL(invertedIndexes)
	if err != nil {
		return err
	}
	return nil
}
