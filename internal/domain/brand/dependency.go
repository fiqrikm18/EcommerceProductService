package brand

import (
	"context"
	"ecommerce/config"
	"ecommerce/internal/domain/brand/presenter"
	brandReposiotry "ecommerce/internal/domain/brand/repository"
	brandUseCase "ecommerce/internal/domain/brand/usecase"
	"log/slog"
)

func NewBrandDependency(
	ctx context.Context,
	dbProvider *config.DatabaseConfiguration,
	logger *slog.Logger,
) presenter.IBrandPresenter {
	repository := brandReposiotry.NewBrandRepository(ctx, dbProvider, logger)
	useCase := brandUseCase.NewBrandUseCase(repository)
	return presenter.NewBrandPresenter(useCase)
}
