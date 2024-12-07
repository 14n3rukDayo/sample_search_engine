package handler

import (
	"main/application/usecase"
	"main/presentation/schema"
	"main/util"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type Searcher interface {
	Search(c echo.Context) error
}

type searcher struct {
	su usecase.SearchDocument
}

func NewSearchHandler(su usecase.SearchDocument) Searcher {
	return &searcher{su: su}
}

func (s *searcher) Search(c echo.Context) error {
	searcherWord := new(schema.SearcherRequest)
	if err := c.Bind(searcherWord); err != nil {
		return util.NewUnprocessableEntityError([]string{util.INVALID_VALIDATION})
	}
	validate := validator.New()
	err := validate.Struct(searcherWord)
	if err != nil {
		errs := util.FormatValidationErrors(err)
		return util.NewUnprocessableEntityError(errs)
	}
	documents, err := s.su.Execute(searcherWord.SearchWords)
	if err != nil {
		return util.NewInternalServerError()
	}
	var res schema.SearchResponse
	for _, document := range documents {
		res.Documents = append(res.Documents, schema.DocumentResponse{DocumentId: document.Get().DocumentId, Description: document.Get().Description})
	}
	return util.NewOKResponse(res)
}
