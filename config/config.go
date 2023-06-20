package config

import (
	"github.com/spf13/viper"
)

//keys for database configuration.
const (
	DbName = "db.name"
	DbHost = "db.ip"
	DbPort = "db.port"
	DbUser = "db.user"
	DbPass = "db.pass"

	ServerHost = "server.host"
	ServerPort = "server.port"

	SecretKey   = "auth.secret"
	CookieName  = "auth.cookie"
	TokenExpiry = "auth.expiry"

	KafkaHost  = "kafka.host"
	KafkaPort  = "kafka.port"
	KafkaTopic = "kafka.topic"
)

func init() {
	// env var for db
	_ = viper.BindEnv(DbName, "DB_NAME")
	_ = viper.BindEnv(DbHost, "DB_HOST")
	_ = viper.BindEnv(DbPort, "DB_PORT")
	_ = viper.BindEnv(DbUser, "DB_USER")
	_ = viper.BindEnv(DbPass, "DB_PASS")

	// env var for server
	_ = viper.BindEnv(ServerHost, "SERVER_HOST")
	_ = viper.BindEnv(ServerPort, "SERVER_PORT")

	// env var for auth
	_ = viper.BindEnv(SecretKey, "AUTH_SECRET_KEY")
	_ = viper.BindEnv(CookieName, "AUTH_COOKIE_NAME")
	_ = viper.BindEnv(TokenExpiry, "AUTH_TOKEN_EXPIRY_TIME")

	// env var for kafka
	_ = viper.BindEnv(KafkaHost, "KAFKA_HOST")
	_ = viper.BindEnv(KafkaPort, "KAFKA_PORT")
	_ = viper.BindEnv(KafkaTopic, "KAFKA_TOPIC")

	// defaults
	viper.SetDefault(DbName, "companies_portal")
	viper.SetDefault(DbHost, "localhost")
	viper.SetDefault(DbPort, "27017")
	viper.SetDefault(DbUser, "user")
	viper.SetDefault(DbPass, "password")
	viper.SetDefault(ServerHost, "127.0.0.1")
	viper.SetDefault(ServerPort, "8080")
	viper.SetDefault(CookieName, "jwt")
	viper.SetDefault(TokenExpiry, "2h")
	viper.SetDefault(KafkaHost, "127.0.0.1")
	viper.SetDefault(KafkaPort, "9094")
	viper.SetDefault(KafkaTopic, "company-events")
}
