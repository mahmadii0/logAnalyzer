package services

import (
	"log"
	
	"logAnalyzer/models"
)

func RegisterLog(lg *models.Log) error {
	if err := lg.CreateLog(); err != nil {
		log.Printf("Error while creating log: %v", err)
		return err
	}
	key := lg.Service + "|" + lg.Level + "|" + lg.Message
	RegisterErrorGroup(lg,key)
	return nil
}
