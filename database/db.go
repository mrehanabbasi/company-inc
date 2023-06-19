package database

import (
	"context"
	"fmt"

	"github.com/mrehanabbasi/company-inc/config"
	log "github.com/mrehanabbasi/company-inc/logger"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const companyCollectionName = "Company"

type Client struct {
	Conn *mongo.Client
}

func InitDB() *Client {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		viper.GetString(config.DbUser),
		viper.GetString(config.DbPass),
		viper.GetString(config.DbHost),
		viper.GetString(config.DbPort),
	)
	log.Info("Initializing MongoDB: ", uri)

	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Panic(err.Error())
		panic(err)
	}

	return &Client{Conn: cli}
}

func (m *Client) GetMongoDatabase() *mongo.Database {
	return m.Conn.Database(viper.GetString(config.DbName))
}

func (m *Client) GetMongoCompanyCollection() *mongo.Collection {
	return m.GetMongoDatabase().Collection(companyCollectionName)
}

func (m *Client) InitIndices() error {
	// Create a unique index on the "company_name" field
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"company_name": 1},
		Options: options.Index().SetUnique(true),
	}

	// Create the index
	_, err := m.GetMongoCompanyCollection().Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
