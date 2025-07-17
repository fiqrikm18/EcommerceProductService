// no lint
package cmd

import (
	"context"
	"ecommerce/common"
	"ecommerce/config"
	"ecommerce/docs"
	ValidatorUtils "ecommerce/pkg/validator"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	_ "ecommerce/docs"
	"github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/files"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Ecommerce is a CLI tool for managing Ecommerce products",
	Long:  "Ecommerce is a CLI tool for managing Ecommerce products.",
	Run:   runRootCommand,
}

var port string

func init() {
	config.InitializeCConfiguration()
}

func runRootCommand(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	config.InitializeDatabase(config.AppConfig, logger)
	e := echo.New()
	e.Validator = &ValidatorUtils.RequestValidator{Validator: validator.New()}

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Request().URL.Path, "docs") {
				return true
			}
			return false
		},
	}))

	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Host = config.AppConfig.ApplicationHost + ":" + config.AppConfig.ApplicationPort

	e.GET("/docs/*", echoSwagger.WrapHandler)
	common.RegisterRoute(e, ctx, config.DatabaseProvider, logger)

	go func() {
		if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal("❌ Echo start error:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("🔄 Shutdown...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal("❌ Shutdown error:", err)
	}

	fmt.Println("✅ Server stopped cleanly")
}

func Execute() {
	rootCmd.PersistentFlags().StringVar(&port, "port", "8080", "The port the server will listen on")
	rootCmd.MarkFlagsRequiredTogether("port")
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
