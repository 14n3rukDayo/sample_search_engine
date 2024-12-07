package util

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

const INVALID_VALIDATION = "invalid validation"
const OK = "ok"

type HTTPError struct {
	StatusCode int      `json:"status_code"`
	Message    string   `json:"message"`
	Errors     []string `json:"errors,omitempty"`
}

func NewOKResponse[T any](data T) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusOK, data)
}

func NewUnprocessableEntityError(errors []string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusUnprocessableEntity, HTTPError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    "Validation failed",
		Errors:     errors,
	})
}
func NewInternalServerError() *echo.HTTPError {
	return echo.NewHTTPError(http.StatusInternalServerError, HTTPError{
		StatusCode: http.StatusInternalServerError,
		Message:    "internal server",
	})
}

func NewAddOperationError(errors []string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusBadRequest, HTTPError{
		StatusCode: http.StatusBadRequest,
		Message:    "Failed to add the requested item",
		Errors:     errors,
	})
}

func FormatValidationErrors(err error) []string {
	var validationErrors []string
	for _, ve := range err.(validator.ValidationErrors) {
		validationErrors = append(validationErrors, fmt.Sprintf("%s: %s", ve.Field(), ve.Tag()))
	}
	return validationErrors
}
