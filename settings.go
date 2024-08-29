package main

import (
	"log"
	"os"

	dotenv "github.com/joho/godotenv"
)

type settings struct {
	PROJECT_NAME string
	DB_USER      string
	DB_PASS      string
}

var AppSettings settings

func init() {
	err := dotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	AppSettings = settings{
		PROJECT_NAME: os.Getenv("PROJECT_NAME"),
		DB_USER:      os.Getenv("DB_USER"),
		DB_PASS:      os.Getenv("DB_PASS"),
	}
}
