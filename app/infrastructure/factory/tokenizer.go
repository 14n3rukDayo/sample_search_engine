package ifactory

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	TokensVO "main/domain/document/valueobject/tokens"
	TokensF "main/domain/document/valueobject/tokens/factory"
	"main/util"
	"net/http"
	"os"
)

type Tokenizer struct {
}

func NewTokenizer() TokensF.Tokenizer {
	return &Tokenizer{}
}

func (t *Tokenizer) Tokenize(text string) (tokens TokensVO.Tokens, err error) {
	API_URL := os.Getenv("API_URL")
	url := fmt.Sprintf("%s/tokenize", API_URL)
	requestBody := map[string]string{
		"text": text,
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, util.NewInternalServerError()
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, util.NewInternalServerError()
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, util.NewInternalServerError()
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, util.NewInternalServerError()
	}
	if resp.StatusCode != util.OK_RESPONSE_CODE {
		return nil, util.NewInternalServerError()
	}
	var res struct {
		Tokens []string
	}

	err = json.Unmarshal(body, &res)

	if err != nil {
		return nil, util.NewInternalServerError()
	}
	tokens = res.Tokens
	return tokens, nil
}
