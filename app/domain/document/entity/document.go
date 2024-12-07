package entity

import (
	tokenscoreVO "main/domain/document/valueobject/tokenscore"
)

type Document interface {
	Get() *document
	AddTokens(tokens []tokenscoreVO.TokenScore)
}
type document struct {
	DocumentId  int
	Description string
	Tokens      []tokenscoreVO.TokenScore
}

func NewDocumentEntity(id int, description string) Document {
	return &document{
		DocumentId:  id,
		Description: description,
	}
}
func (d *document) Get() *document {
	return d
}

func (d *document) AddTokens(tokens []tokenscoreVO.TokenScore) {
	d.Tokens = append(d.Tokens, tokens...)
}
