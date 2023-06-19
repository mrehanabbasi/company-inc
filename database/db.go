package database

import (
	"context"
	"fmt"

	"github.com/mrehanabbasi/company-inc/config"
	log "github.com/mrehanabbasi/company-inc/logger"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const companyCollectionName = "Company"

type client struct {
	conn *mongo.Client
}

func InitDB() *client {
	uri := fmt.Sprintf("mongodb://%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	log.Info("Initializing MongoDB: ", uri)

	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Panic(err.Error())
		panic(err)
	}

	return &client{conn: cli}
}

func (m *client) GetMongoDatabase() *mongo.Database {
	return m.conn.Database(viper.GetString(config.DbName))
}

func (m *client) GetMongoCompanyCollection() *mongo.Collection {
	return m.GetMongoDatabase().Collection(companyCollectionName)
}
