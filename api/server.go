package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pootytang/danielleapi/types"
)

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

	slog.Info("-------------------- STARTING --------------------")
	slog.Info(fmt.Sprintf("---------- ENVIRONMENT = %s ----------", types.GetHost()))
	e.Logger.Fatal(e.Start(":1323"))
}

func openLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}
