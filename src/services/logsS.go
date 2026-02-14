package services

import (
	"log"
	"logAnalyzer/constants"
	"logAnalyzer/models"
)

func RegisterLog(lg *models.Log) {
	if err := lg.CreateLog(); err != nil {
		log.Printf("Error while creating log: %v", err)
		return
	}
	key := lg.Service + "|" + lg.Level + "|" + lg.Message
	eg, err := models.GetErrorGbyKey(key)
	if err != nil {
		log.Printf("Error while getting errorGroup: %v", err)
	} else if eg.ID == 0 {
		err := eg.CreateErrorG(lg.CreatedAt, key)
		if err != nil {
			log.Printf("Error while creating errorGroup: %v", err)
		}
		return
	}
	if err = eg.IncraeseCount(); err != nil {
		log.Printf("Error while increasing count of errorGroup: %v", err)
	}
	if eg.Count > 5 {
		log.Printf("%sErrorGroup Reached The Limit%s", constants.Purple,constants.Gray)
	}
}
