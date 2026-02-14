package database

import (
	"log"
	"logAnalyzer/env"
	"logAnalyzer/models"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateTables() {
	source := env.GetEnvString("DATABASE_SOURCE", "")
	db, err := gorm.Open(mysql.Open(source), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(models.Log{}, models.ErrorGroup{})

	db, err = gorm.Open(mysql.Open(source), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(models.Log{}, models.ErrorGroup{})
}
