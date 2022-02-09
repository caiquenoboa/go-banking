package main

import (
	"github.com/caiquenoboa/go-banking/app"
	"github.com/caiquenoboa/go-banking/logger"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}
