package factory

import (
	documentR "main/domain/document/repository"
	tokensVO "main/domain/document/valueobject/tokens"
	tokenScoreVO "main/domain/document/valueobject/tokenscore"
	invertedIndexR "main/domain/invertedindex/repository"
	"math"
)

type TokenScore interface {
	CreateTokenScore(tokens tokensVO.Tokens) (tokenScore []tokenScoreVO.TokenScore, err error)
}
type tokenScore struct {
	dr  documentR.Document
	iir invertedIndexR.InvertedIndex
}

func NewTokenScoreFactory(dr documentR.Document, iir invertedIndexR.InvertedIndex) TokenScore {
	return &tokenScore{dr: dr, iir: iir}
}

func (t *tokenScore) CreateTokenScore(tokens tokensVO.Tokens) (tokenScore []tokenScoreVO.TokenScore, err error) {
	tokenCount := make(map[string]int)
	for _, token := range tokens {
		tokenCount[token]++
	}
	d := float64(len(tokenCount))
	allD, err := t.iir.GetAllDLNum()
	if err != nil {
		return nil, err
	}
	preTotalDoc, err := t.dr.GetTotalNum()
	if err != nil {
		return nil, err
	}
	newTotalDoc := float64(preTotalDoc + 1)
	avgdl := float64(allD) / newTotalDoc
	dl := float64(newTotalDoc / avgdl)
	k := 1.2
	b := 0.75
	for token, count := range tokenCount {
		tf := float64(count) / d

		preInvertedIndex, err := t.iir.Get(token)
		if err != nil {
			return nil, err
		}
		preDocumentNum := preInvertedIndex.GetDocumentNum()
		newDocumentNum := preDocumentNum + 1
		n := (float64(newTotalDoc) - float64(newDocumentNum) + 0.5) / (float64(newDocumentNum) + 0.5)
		idf := math.Log(n + 1)

		numerator := tf * (k + 1)
		denominator := tf + k*(1-b+b*dl)
		score := idf * (numerator / denominator)
		tokenScore = append(tokenScore, tokenScoreVO.TokenScore{Token: token, Score: score})
	}

	return tokenScore, nil
}
