package main

import (
	"log"
	"tender-service/internal/repository"
	"tender-service/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	psqlDB, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	defer psqlDB.Close()

	server.Run(psqlDB)
}
