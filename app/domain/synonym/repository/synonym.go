package repository

type Synonym interface {
	Get(word string) (alias string)
}
