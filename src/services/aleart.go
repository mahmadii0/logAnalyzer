package services

import (
	"log"
	"logAnalyzer/models"
	"time"
)

func GetAleartStatus(limitNumber int, eg *models.ErrorGroup) bool{
	if limitNumber != -1 {
		if eg.Count >= limitNumber {
			if eg.AISummary==""{
				err := GetSummary(eg)
				if err != nil {
					log.Printf("Error While Getting Summary")
				}
			    eg.Confidence="low"
			}else{
				err:=GetConfidence(eg)
				if err != nil {
					log.Printf("Error While Getting Summary")
				}
			}
			eg.LastSeen = time.Now()
			eg.UpdateErrorG()
			return true
		}
		eg.LastSeen = time.Now()
		eg.UpdateErrorG()
		return false
	}
	
	eg.CalBaselineAVG()
	eg.LastSeen = time.Now()
	eg.UpdateErrorG()
	if eg.Count >= int(eg.BaselineAVG) {
		if eg.AISummary==""{
			err := GetSummary(eg)
			if err != nil {
				log.Printf("Error While Getting Summary")
			}
			eg.Confidence="low"
		}else{
			err:=GetConfidence(eg)
			if err != nil {
				log.Printf("Error While Getting Summary")
			}
		}
		return true
	}
	return false
}
