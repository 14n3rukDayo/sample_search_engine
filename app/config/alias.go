package config

import (
	"encoding/json"
	"io"
	"os"
)

type Aliases map[string]string

func CreateAliases() (aliases Aliases) {

	file, err := os.Open("./config/alias.json")
	if err != nil {
		panic("not opened")
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic("not read")
	}
	aliasesList := make(map[string][]string)
	err = json.Unmarshal(bytes, &aliasesList)
	if err != nil {
		panic("not unmarshal")
	}
	aliases = make(Aliases)
	for key, values := range aliasesList {
		for _, value := range values {
			aliases[value] = key
		}
	}
	return aliases
}
