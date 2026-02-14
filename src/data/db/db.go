package database

import (
	"context"
	"log"
	"logAnalyzer/env"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	ctx context.Context
)

func Connect() {
	source := env.GetEnvString("DATABASE_SOURCE", "")
	d, err := gorm.Open(mysql.Open(source), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	ctx = context.Background()
	db = d
}

func GetDBctx() (*gorm.DB, context.Context) {
	Connect()
	return db, ctx
}
