package config

import (
	"github.com/joho/godotenv"
	"log"
)

func ConfigureProject() {
	if err := godotenv.Load("config/.env"); err != nil {
		log.Fatal("env file not loaded")
	}
}
