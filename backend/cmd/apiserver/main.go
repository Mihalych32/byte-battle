package main

import (
	"byte-battle_backend/config"
	"byte-battle_backend/internal/app"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Could not load .env file")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config err: %s", err)
	}
	app.Run(cfg)
}
