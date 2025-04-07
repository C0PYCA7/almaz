package main

import (
	"CartridgeServer/internal/config"
	"CartridgeServer/internal/logger"
)

func main() {
	log := logger.New()
	cfg := config.New()
	log.Debug("cfg", cfg)
}
