package main

import (
	"fmt"
	"mikhaylovilya/map-sorcer/internal/config"
	"mikhaylovilya/map-sorcer/internal/logger"
)

func main() {
	config := config.Parse()

	fmt.Println(config)

	logger := logger.NewLogger(config.Env)

	logger.Debug("Logger initialized")

	// server init
	// server start
}
