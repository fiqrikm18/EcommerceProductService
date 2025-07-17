package config

import (
	"ecommerce/constants"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log/slog"
)

type DatabaseConfiguration struct {
	*gorm.DB
}

var DatabaseProvider *DatabaseConfiguration

func InitializeDatabase(appConfig *ApplicationConfiguration, slog *slog.Logger) {
	db, err := gorm.Open(postgres.Open(appConfig.PostgresDsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		slog.Error(err.Error())
	}

	rawDb, err := db.DB()
	if err != nil {
		slog.Error(err.Error())
	}

	err = rawDb.Ping()
	if err != nil {
		slog.Error(err.Error())
	}

	rawDb.SetMaxIdleConns(constants.MaxIdleConn)
	rawDb.SetMaxOpenConns(constants.MaxOpenConn)
	rawDb.SetConnMaxLifetime(constants.ConnMaxLifetime)
	rawDb.SetConnMaxIdleTime(constants.ConnMaxIdleTime)

	DatabaseProvider = &DatabaseConfiguration{db}
}
