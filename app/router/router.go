package router

import (
	"main/presentation/handler"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, documentHandler handler.Document, searchHandler handler.Searcher) {
	e.POST("/document", documentHandler.Add)
	e.POST("/search", searchHandler.Search)
}

func NewEchoInstance(documentHandler handler.Document, searchHandler handler.Searcher) *echo.Echo {
	e := echo.New()
	RegisterRoutes(e, documentHandler, searchHandler)
	return e
}
