package main

import (
	"github.com/rodrigocarvalho10/go-opportunities/config"
	"github.com/rodrigocarvalho10/go-opportunities/router"
)

func main() {

	//Initialize Configs
	err error := config.Init()
	if err != nil {
		panic(err)
	}
	// Initialize router
	router.Initialize()
}
