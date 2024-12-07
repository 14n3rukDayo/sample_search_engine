package irepository

import (
	"main/config"
	synonymR "main/domain/synonym/repository"
)

type synonym struct {
	a config.Aliases
}

func NewSynonumRepository(a config.Aliases) synonymR.Synonym {
	return &synonym{a: a}
}

func (s *synonym) Get(word string) (alias string) {
	aliases := s.a
	ailias := aliases[word]
	return ailias
}
