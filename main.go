package main

import (
	"github.com/joaolima7/gopportunities.git/config"
	"github.com/joaolima7/gopportunities.git/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		panic(err)
	}
	router.Initialize()
}
