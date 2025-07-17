package common

import (
	"context"
	"ecommerce/config"
	"github.com/labstack/echo/v4"
	"log/slog"

	BrandDeps "ecommerce/internal/domain/brand"
	Brand "ecommerce/internal/domain/brand/presenter"

	ProductDeps "ecommerce/internal/domain/product"
	Product "ecommerce/internal/domain/product/presenter"
)

var (
	brandPresenter   Brand.IBrandPresenter
	productPresenter Product.IProductPresenter
)

func RegisterRoute(c *echo.Echo, ctx context.Context, databaseProvider *config.DatabaseConfiguration, logger *slog.Logger) {
	initializePresenter(ctx, databaseProvider, logger)
	api := c.Group("api/v1")

	brandRoute := api.Group("/brands")
	brandRoute.GET("", brandPresenter.GetAll)
	brandRoute.GET("/:id", brandPresenter.Get)
	brandRoute.POST("", brandPresenter.Create)
	brandRoute.PATCH("/:id", brandPresenter.Update)
	brandRoute.DELETE("/:id", brandPresenter.Delete)

	productRoute := api.Group("/products")
	productRoute.GET("", productPresenter.GetAll)
	productRoute.GET("/:id", productPresenter.Get)
	productRoute.POST("", productPresenter.Create)
	productRoute.PATCH("/:id", productPresenter.Update)
	productRoute.DELETE("/:id", productPresenter.Delete)
}

func initializePresenter(ctx context.Context, databaseProvider *config.DatabaseConfiguration, logger *slog.Logger) {
	brandPresenter = BrandDeps.NewBrandDependency(ctx, databaseProvider, logger)
	productPresenter = ProductDeps.NewProductDependency(ctx, databaseProvider, logger)
}
