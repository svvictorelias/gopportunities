package main

import (
	"github.com/joho/godotenv"
	"github.com/svvictorelias/gopportunities/config"
	"github.com/svvictorelias/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Errorf("error loading .env file: %v", err)
		return
	}

	logger = config.GetLogger("main")

	err = config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
