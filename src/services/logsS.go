package services

import (
	"log"
	"logAnalyzer/constants"
	"logAnalyzer/models"
)

func RegisterLog(lg *models.Log) error {
	if err := lg.CreateLog(); err != nil {
		log.Printf("Error while creating log: %v", err)
		return err
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
		return err
	}
	if err = eg.IncraeseCount(); err != nil {
		log.Printf("Error while increasing count of errorGroup: %v", err)
		return err
	}
	if eg.Count > constants.LimitNumber {
		summary:=GetAnalayzeResponse(eg)
		if summary != ""{
			eg.AISummary=summary
			if err = eg.SetSummary(); err !=nil{
				log.Printf("Error while setting ai summary errorGroup: %v", err)
		return err
			}
		}
		
	}
	return nil
}
