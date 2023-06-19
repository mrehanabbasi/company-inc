package main

import (
	"os"

	"github.com/mrehanabbasi/company-inc/logger"
	"github.com/mrehanabbasi/company-inc/routes"
)

func main() {
	// Initializing logger
	logger.TextLogInit()

	// Register all the routes
	server := routes.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	_ = server.Run(":" + port)
}
