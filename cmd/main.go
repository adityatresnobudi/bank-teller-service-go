package main

import (
	"log"

	"github.com/adityatresnobudi/bank-teller-service-go/config"
	"github.com/adityatresnobudi/bank-teller-service-go/internal/server"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("err loading .env file: %s\n", err.Error())
	}
}

func main() {
	cfg := config.NewConfig()

	s := server.NewServer(cfg)

	s.Run()
}
