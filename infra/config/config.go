package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type appConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBDialect  string
	Port       string
}

func LoadAppConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env file")
	}
}

func GetAppConfig() appConfig {
	return appConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBDialect:  os.Getenv("DB_DIALECT"),
		Port:       os.Getenv("PORT"),
	}
}
