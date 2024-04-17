package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	DBName   string
	Username string
	Passwod  string
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func GetDBConfig() *DBConfig {
	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USERNAME"),
		Passwod:  os.Getenv("DB_PASSWORD"),
	}
}
