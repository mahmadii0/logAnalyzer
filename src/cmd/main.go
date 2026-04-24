package main

import (
	"log"
	"logAnalyzer/api"
	"logAnalyzer/data/db/db_init"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	file, err := os.OpenFile("src/logs/app.log", 
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err !=nil{
		log.Fatal(err)
	}
	log.SetOutput(file)
	database.CreateTables()
	api.InitServer()

}
