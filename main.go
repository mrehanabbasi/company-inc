package main

import (
	"github.com/mrehanabbasi/company-inc/config"
	"github.com/mrehanabbasi/company-inc/database"
	"github.com/mrehanabbasi/company-inc/logger"
	"github.com/mrehanabbasi/company-inc/routes"
	"github.com/spf13/viper"
)

func main() {
	// Initializing logger
	logger.TextLogInit()

	// Initializing database
	db := database.InitDB()
	db.GetMongoCompanyCollection()

	// Register all the routes
	server := routes.NewRouter()

	_ = server.Run(viper.GetString(config.ServerHost) + ":" + viper.GetString(config.ServerPort))
}
