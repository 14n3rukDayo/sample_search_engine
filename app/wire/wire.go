//go:build wireinject

package wire

import (
	"main/application/query"
	"main/application/service"
	"main/application/usecase"
	"main/config"
	documentF "main/domain/document/factory"
	tokensF "main/domain/document/valueobject/tokens/factory"
	tokenScoreF "main/domain/document/valueobject/tokenscore/factory"
	invertedIndexF "main/domain/invertedindex/factory"
	ifactory "main/infrastructure/factory"
	irepository "main/infrastructure/repository"
	"main/presentation/handler"
	"main/router"

	wire "github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func InitializeEcho() (*echo.Echo, error) {
	wire.Build(
		config.NewRedisClient,
		config.CreateAliases,
		irepository.NewInvertedIndexRepository,
		irepository.NewDocumentRepository,
		irepository.NewSynonumRepository,
		invertedIndexF.NewInvertedIndexFactory,
		documentF.NewDocumentFactory,
		tokensF.NewTokensFactory,
		ifactory.NewTokenizer,
		tokenScoreF.NewTokenScoreFactory,
		service.NewDocumentService,
		query.NewSearcher,
		usecase.NewSearchDocument,
		usecase.NewAddDocumentUsecase,
		handler.NewSearchHandler,
		handler.NewDocumentHandler,
		router.NewEchoInstance,
	)
	return nil, nil
}
