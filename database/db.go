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

const (
	companyCollectionName = "Company"
	userCollectionName    = "User"

	companyTestCollectionName = "Company_Test"
	userTestCollectionName    = "User_Test"
)

type Client struct {
	Conn *mongo.Client
}

func InitDB(ctx context.Context) *Client {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		viper.GetString(config.DbUser),
		viper.GetString(config.DbPass),
		viper.GetString(config.DbHost),
		viper.GetString(config.DbPort),
	)
	log.Info("Initializing MongoDB: ", fmt.Sprintf("mongodb://****:****@%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort)))

	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Panic(err.Error())
		panic(err)
	}

	return &Client{Conn: cli}
}

func InitTestDB(ctx context.Context) *Client {
	uri := "mongodb://user:password@localhost:27017"
	log.Info("Initializing MongoDB: ", "mongodb://****:****@localhost:27017")

	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Panic(err.Error())
		panic(err)
	}

	return &Client{Conn: cli}
}

func (m *Client) GetMongoDatabase() *mongo.Database {
	return m.Conn.Database(viper.GetString(config.DbName))
}

func (m *Client) GetMongoTestDatabase() *mongo.Database {
	return m.Conn.Database("test_company_db")
}

func (m *Client) GetMongoCompanyCollection() *mongo.Collection {
	return m.GetMongoDatabase().Collection(companyCollectionName)
}

func (m *Client) GetMongoTestCompanyCollection() *mongo.Collection {
	return m.GetMongoTestDatabase().Collection(companyTestCollectionName)
}

func (m *Client) GetMongoUserCollection() *mongo.Collection {
	return m.GetMongoDatabase().Collection(userCollectionName)
}

func (m *Client) GetMongoTestUserCollection() *mongo.Collection {
	return m.GetMongoTestDatabase().Collection(userTestCollectionName)
}

func (m *Client) InitIndices() error {
	// Create a unique index on the "company_name" field
	coNameIndexModel := mongo.IndexModel{
		Keys:    bson.M{"company_name": 1},
		Options: options.Index().SetUnique(true),
	}

	// Create a unique index on the "email" field
	userEmailIndexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	// Create the company name index
	_, err := m.GetMongoCompanyCollection().Indexes().CreateOne(context.TODO(), coNameIndexModel)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Create the user email index
	_, err = m.GetMongoUserCollection().Indexes().CreateOne(context.TODO(), userEmailIndexModel)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
