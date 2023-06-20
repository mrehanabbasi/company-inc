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
	TokenExpiry = "token.expiry"
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
	_ = viper.BindEnv(SecretKey, "SECRET_KEY")
	_ = viper.BindEnv(TokenExpiry, "TOKEN_EXPIRY_TIME")

	// defaults
	viper.SetDefault(DbName, "companies_portal")
	viper.SetDefault(DbHost, "localhost")
	viper.SetDefault(DbPort, "27017")
	viper.SetDefault(DbUser, "user")
	viper.SetDefault(DbPass, "password")
	viper.SetDefault(ServerHost, "127.0.0.1")
	viper.SetDefault(ServerPort, "8080")
	viper.SetDefault(TokenExpiry, "2h")
}
