package common

import (
	"context"
	"ecommerce/config"
	"github.com/labstack/echo/v4"
	"log/slog"

	BrandDeps "ecommerce/internal/domain/brand"
	Brand "ecommerce/internal/domain/brand/presenter"
)

var (
	brandPresenter Brand.IBrandPresenter
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
}

func initializePresenter(ctx context.Context, databaseProvider *config.DatabaseConfiguration, logger *slog.Logger) {
	brandPresenter = BrandDeps.NewBrandDependency(ctx, databaseProvider, logger)
}
