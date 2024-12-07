package factory

import (
	"context"
	documentE "main/domain/document/entity"
	documentR "main/domain/document/repository"
	tokensF "main/domain/document/valueobject/tokens/factory"
	tokenScoreF "main/domain/document/valueobject/tokenscore/factory"
)

type Document interface {
	NewDocument(ctx context.Context, description string) (document documentE.Document, err error)
}
type document struct {
	dr  documentR.Document
	tf  tokensF.Tokens
	tsf tokenScoreF.TokenScore
}

func NewDocumentFactory(dr documentR.Document, tf tokensF.Tokens, tsf tokenScoreF.TokenScore) Document {
	return &document{dr: dr, tf: tf, tsf: tsf}
}

func (f *document) NewDocument(ctx context.Context, description string) (document documentE.Document, err error) {

	documentId, err := f.dr.GenerateID()
	if err != nil {
		return nil, err
	}
	document = documentE.NewDocumentEntity(documentId, description)
	tokens, err := f.tf.CreateTokens(description)
	if err != nil {
		return nil, err
	}
	tokenScores, err := f.tsf.CreateTokenScore(tokens)
	if err != nil {
		return nil, err
	}
	document.AddTokens(tokenScores)
	return document, nil
}
