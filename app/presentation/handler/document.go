package handler

import (
	"fmt"
	"main/application/usecase"
	"main/presentation/schema"
	"main/util"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type Document interface {
	Add(c echo.Context) error
}
type document struct {
	adu usecase.AddDocument
}

func NewDocumentHandler(adu usecase.AddDocument) Document {
	return &document{adu: adu}
}

func (d *document) Add(c echo.Context) error {
	document := new(schema.AddDocumentRequest)
	if err := c.Bind(document); err != nil {
		return util.NewUnprocessableEntityError([]string{util.INVALID_VALIDATION})
	}
	validate := validator.New()
	err := validate.Struct(document)
	if err != nil {
		validationErrors := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s: %s", err.Field(), err.Tag()))
		}
		return util.NewUnprocessableEntityError(validationErrors)
	}

	err = d.adu.Execute(document.Description)
	if err != nil {
		return err
	}
	return util.NewOKResponse(schema.AddDocumentResponse{Message: util.OK})
}
