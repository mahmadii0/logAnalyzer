package services

import (
	"log"
	"logAnalyzer/constants"
	"logAnalyzer/models"
	"time"
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

		if err := eg.CreateErrorG(lg.CreatedAt, key); err != nil {
			log.Printf("Error while creating errorGroup: %v", err)
		}
		return err
	}
	if err = eg.IncraeseCount(); err != nil {
		log.Printf("Error while increasing count of errorGroup: %v", err)
		return err
	}

	var ok bool
	if eg.BaselineAVG == 0 {
		eg.BaselineAVG = 1
		eg.UpdateErrorG()
	}
	if eg.BaselineAVG == 1 && eg.BaselineAVG != float64(eg.Count) {
		limitNumber := constants.LimitNumber
		eg.BaselineAVG = float64(eg.Count)
		ok=GetAleartStatus(limitNumber, eg)
	} else if eg.BaselineAVG == 1 && eg.BaselineAVG == float64(eg.Count) {
		eg.LastSeen = time.Now()
		eg.UpdateErrorG()
		return nil
	} else {
		ok=GetAleartStatus(-1, eg)
	}
	if ok{
		
	}
	return nil
}
