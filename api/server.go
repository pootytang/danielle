package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pootytang/danielleapi/initializers"
	"github.com/pootytang/danielleapi/types"
)

func init() {
	config, err := initializers.LoadConfig("./")
	if err != nil {
		log.Fatal("init(): Could not load environment variables: ", err)
	}

	initializers.DBConnect(&config)
}

func main() {
	file, err := openLogFile(types.LOG_FILE)
	if err != nil {
		slog.Error(err.Error())
	}
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handler := slog.NewTextHandler(file, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	e := echo.New()
	setupRoutes(e)

	_, hostname := types.GetHost()
	slog.Info(fmt.Sprintf("---------- STARTING IN ENVIRONMENT = %s ----------", hostname))
	e.Logger.Fatal(e.Start(":1323"))
}

func openLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}
