package factory

import (
	documentR "main/domain/document/repository"
	tokensVO "main/domain/document/valueobject/tokens"
	synonymR "main/domain/synonym/repository"
	"main/util"
	"unicode"
	"unicode/utf8"
)

type Tokens interface {
	CreateTokens(text string) (tokens tokensVO.Tokens, err error)
}

type tokens struct {
	tnf Tokenizer
	dr  documentR.Document
	sr  synonymR.Synonym
}

func NewTokensFactory(tnf Tokenizer, dr documentR.Document, sr synonymR.Synonym) Tokens {
	return &tokens{tnf: tnf, dr: dr, sr: sr}
}


func isJapanese(r rune) bool {
	return unicode.In(r, unicode.Hiragana, unicode.Katakana, unicode.Han)
}
func isAlphabet(r rune) bool {
	return unicode.IsLetter(r)
}
func (t *tokens) filterTokens(tokens []string) []string {
	var filteredTokens []string
	for _, token := range tokens {
		if len(token) > 4 && token[:4] == t.dr.GetPrefixDocumentIdReserved() {
			continue
		}
		var filteredToken []rune
		for _, r := range token {
			if isJapanese(r) {
				filteredToken = append(filteredToken, r)
				continue
			}
			if isAlphabet(r) {
				filteredToken = append(filteredToken, unicode.ToLower(r))
				continue
			}
		}
		if len(filteredToken) == utf8.RuneCountInString(token) {
			addedToken := string(filteredToken)

			alias := t.sr.Get(addedToken)
			if len(alias) > 0 {
				addedToken = alias
			}

			filteredTokens = append(filteredTokens, addedToken)
		}
	}

	return filteredTokens
}
func (t *tokens) CreateTokens(text string) (tokens tokensVO.Tokens, err error) {
	tokensOrigin, err := t.tnf.Tokenize(text)
	if err != nil {
		return nil, util.NewInternalServerError()
	}
	tokens = t.filterTokens(tokensOrigin)
	return tokens, nil
}
