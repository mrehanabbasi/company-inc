package main

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"github.com/mrehanabbasi/company-inc/config"
	"github.com/mrehanabbasi/company-inc/database"
	"github.com/mrehanabbasi/company-inc/logger"
	"github.com/mrehanabbasi/company-inc/routes"
	"github.com/spf13/viper"
)

func main() {
	// Initializing logger
	logger.TextLogInit()

	// Convert fe.Field() from StructField to json field for custom validation messages
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	// Initializing database
	dbClient := database.InitDB()

	// Register all the routes
	server := routes.NewRouter(dbClient)

	_ = server.Run(viper.GetString(config.ServerHost) + ":" + viper.GetString(config.ServerPort))
}
