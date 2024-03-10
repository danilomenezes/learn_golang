package main

import (
	"os"
	// "rest-api/domain/service"

	"go.uber.org/zap"
)

func init() {
	logger := zap.Must(zap.NewProduction())
	if os.Getenv("APP_ENV") == "development" {
		logger = zap.Must(zap.NewDevelopment())
	}

	zap.ReplaceGlobals(logger)
}

func main() {
	zap.L().Info("Service started...")

	logger := zap.Must(zap.NewProduction())

	defer logger.Sync()

	// result := service.ProcessRule()

	// logger.Info("Result:", zap.String("status http", result.Status))
}
