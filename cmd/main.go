package main

import (
	"hackaton/internal/server"
	"log"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	cfg, err := server.NewConfig("config.yml")
	if err != nil {
		logger.Sugar().Error(err)
	}

	log.Fatalln(server.Start(cfg, logger.Sugar()))
}
