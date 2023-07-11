package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mrehanabbasi/company-inc/config"
	"github.com/mrehanabbasi/company-inc/database"
	"github.com/mrehanabbasi/company-inc/logger"
	log "github.com/mrehanabbasi/company-inc/logger"
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

	// The package "endless" does not work in a Windows environment
	// _ = endless.ListenAndServe(viper.GetString(config.ServerHost)+":"+viper.GetString(config.ServerPort), server)

	// In case graceful shutdown is not required, server.Run() can be used
	// _ = server.Run(viper.GetString(config.ServerHost) + ":" + viper.GetString(config.ServerPort))

	srv := &http.Server{
		Addr:    viper.GetString(config.ServerHost) + ":" + viper.GetString(config.ServerPort),
		Handler: server,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panic("server listening error: ", err)
			panic(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Panic("server shutdown with error: ", err)
		panic(err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Info("waited 5 seconds for the server to close")
	}
	log.Info("Exiting server")
}
