package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Set_env() {
	env := os.Getenv("APP_ENV")
	err := godotenv.Load("./config/" + env + ".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
