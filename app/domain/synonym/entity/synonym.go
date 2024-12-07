package entity

type Synonym interface {
}

type synonym struct {
	Alias    string
	Synonyms []string
}

func NewSynonum(Alias string, synonyms []string) Synonym {
	return &synonym{Alias: Alias, Synonyms: synonyms}
}
