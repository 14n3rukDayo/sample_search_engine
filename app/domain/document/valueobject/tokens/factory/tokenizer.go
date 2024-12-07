package factory

import tokensVO "main/domain/document/valueobject/tokens"

type Tokenizer interface {
	Tokenize(text string) (tokens tokensVO.Tokens, err error)
}
