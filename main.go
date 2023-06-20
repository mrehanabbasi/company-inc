package main

import (
	"context"

	"github.com/mrehanabbasi/company-inc/config"
	"github.com/mrehanabbasi/company-inc/database"
	"github.com/mrehanabbasi/company-inc/logger"
	"github.com/mrehanabbasi/company-inc/models"
	"github.com/mrehanabbasi/company-inc/msq"
	"github.com/mrehanabbasi/company-inc/routes"
	"github.com/spf13/viper"
)

func main() {
	// Initializing logger
	logger.TextLogInit()

	models.InitCustomValidation()

	ctx := context.Background()

	// Initializing database
	dbClient := database.InitDB(ctx)
	_ = dbClient.InitIndices()
	defer func() {
		_ = dbClient.Conn.Disconnect(ctx)
	}()

	// Initializing kafka writer
	kConn := msq.InitKafka()
	defer func() {
		_ = kConn.Conn.Close()
	}()

	// Register all the routes
	server := routes.NewRouter(dbClient, kConn)

	_ = server.Run(viper.GetString(config.ServerHost) + ":" + viper.GetString(config.ServerPort))
}
