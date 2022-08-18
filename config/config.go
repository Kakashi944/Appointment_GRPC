package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var AppConfig Config

type Config struct {
	Port        int
	Environment string
	Debug       bool

	DBHost     string
	DBPort     int
	DBDatabase string
	DBUsername string
	DBPassword string
}

func InitializeAppConfig() {
	viper.SetConfigName(".env") // allow directly reading from .env file
	viper.SetConfigType("env")
	viper.AddConfigPath("./../")
	_ = viper.ReadInConfig()

	AppConfig.Port = viper.GetInt("PORT")
	AppConfig.Environment = viper.GetString("ENVIRONMENT")
	AppConfig.Debug = viper.GetBool("DEBUG")

	AppConfig.DBHost = viper.GetString("DB_HOST")
	AppConfig.DBPort = viper.GetInt("DB_PORT")
	AppConfig.DBDatabase = viper.GetString("DB_DATABASE")
	AppConfig.DBUsername = viper.GetString("DB_USERNAME")
	AppConfig.DBPassword = viper.GetString("DB_PASSWORD")
	fmt.Println("[INIT] configuration loaded")
}
