package main

import (
	"context"
	"fmt"
	"log/slog"
	"mikhaylovilya/map-sorcerer/internal/config"
	"mikhaylovilya/map-sorcerer/internal/infra/postgres"
	"mikhaylovilya/map-sorcerer/internal/logger"
	"os"
)

func main() {
	config := config.Parse()

	fmt.Println(config)

	logger := logger.NewLogger(config.Env)

	logger.Info("logger initialized", slog.String("env", config.Env))
	logger.Debug("debug messages are enabled")

	dbpool, err := postgres.NewDBPool(context.Background(), config)
	if err != nil {
		logger.Error(err.Error(), err)
		os.Exit(1)
	}
	defer dbpool.Close()
	// server init
	// server start
}
