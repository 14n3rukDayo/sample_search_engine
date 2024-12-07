package tokenizer

import (
	"strings"
	"unicode"
)

func Tokenize(text string) []string {
	var tokens []string
	var sb strings.Builder

	for _, r := range text {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			sb.WriteRune(unicode.ToLower(r))
		} else {
			if sb.Len() > 0 {
				tokens = append(tokens, sb.String())
				sb.Reset()
			}
		}
	}

	if sb.Len() > 0 {
		tokens = append(tokens, sb.String())
	}

	return tokens
}
