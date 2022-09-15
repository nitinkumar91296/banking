package main

import (
	"github.com/nitinkumar91296/banking/app"
	"github.com/nitinkumar91296/banking/logger"
)

func main() {
	// log.Println("starting our application...")
	logger.Info("Starting the application...")
	app.Start()
}
