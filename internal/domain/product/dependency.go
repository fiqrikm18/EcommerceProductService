package product

import (
	"context"
	"ecommerce/config"
	BrandRepository "ecommerce/internal/domain/brand/repository"
	"ecommerce/internal/domain/product/presenter"
	ProductRepository "ecommerce/internal/domain/product/repository"
	"ecommerce/internal/domain/product/usecase"
	"log/slog"
)

func NewProductDependency(
	ctx context.Context,
	dbProvider *config.DatabaseConfiguration,
	logger *slog.Logger,
) presenter.IProductPresenter {
	productRepository := ProductRepository.NewProductRepository(ctx, dbProvider, logger)
	brandRepository := BrandRepository.NewBrandRepository(ctx, dbProvider, logger)
	useCase := usecase.NewProductUseCase(productRepository, brandRepository)
	return presenter.NewProductPresenter(useCase)
}
