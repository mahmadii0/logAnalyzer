package main

import (
	"log"
	"logAnalyzer/api"
	"logAnalyzer/data/db/db_init"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.CreateTables()
	api.InitServer()

}
