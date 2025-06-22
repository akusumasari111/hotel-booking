package main

import (
	"hotel-booking/config"
	"hotel-booking/internal/api"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cfg, err := config.SetupEnv()
	if err != nil {
		log.Printf("config file is not loaded properly %v\n", err)
	}

	api.StartServer(cfg)

}
