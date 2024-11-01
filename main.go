package main

import (
	"github.com/rodrigocarvalho10/go-opportunities/config"
	"github.com/rodrigocarvalho10/go-opportunities/router"
)

var (
	logger config.Logger
)

//	@title			Example Api Go
//	@version		1.0
//	@description	This is a sample server go
//	@termsOfService	http://swagger.io/terms/

func main() {
	logger = *config.GetLogger("main")
	//Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	// Initialize router
	router.Initialize()
}
