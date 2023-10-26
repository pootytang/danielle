package initializers

import (
	"fmt"
	"log/slog"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect(config *Config) {
	slog.Info("DBConnect(): Called")
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort,
	)
	slog.Debug(fmt.Sprintf("DBConnect(): dsn set to: %s", dsn))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("Failed to connect to the Database")
		os.Exit(1)
	}

	slog.Info("DBConnect(): Connected Successfully to the Database")
}
