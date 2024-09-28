package app

import (
	"github.com/oktadafa/gotoko/app/controllers"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {
	var server = controllers.Server{}
	var appConfig = controllers.AppConfig{}

	var dbConfig = controllers.DBConfig{}
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erorr Unloading .env File")
	}
	appConfig.AppName = getEnv("APP_NAME", "gotokoWeb")
	appConfig.AppEnv = getEnv("APP_ENV", "Development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")

	dbConfig.DBName = getEnv("DB_NAME", "gotoko")
	dbConfig.DBUser = getEnv("DB_USER", "postgres")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "postgres")
	dbConfig.DBPort = getEnv("DB_PORT", "5432")
	dbConfig.DBHost = getEnv("DB_HOST", "localhost")

	server.Initialize(appConfig, dbConfig)
	server.Run(":" + appConfig.AppPort)
}
