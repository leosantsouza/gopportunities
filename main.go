package main

import (
	"github.com/leosantsouza/gopportunities.git/config"
	"github.com/leosantsouza/gopportunities.git/router"
)

func main() {
  
	// Initialize Configs
	err  := config.Init()
	if err != nil{
		println(err)
		return
	}

  // Initialize router
	router.Initialize()

}
