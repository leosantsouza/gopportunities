package main

import (
	"github.com/leosantsouza/gopportunities.git/config"
	"github.com/leosantsouza/gopportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {
  logger = config.GetLogger("main")
	// Initialize Configs
	err  := config.Init()
	if err != nil{
		logger.Errorf("config initialization error: %v", err)
		return
	}

  // Initialize router
	router.Initialize()

}
