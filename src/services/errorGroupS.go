package services

import(
	"logAnalyzer/models"
	"logAnalyzer/constants"
	"time"
	"log"
)

func RegisterErrorGroup(lg *models.Log,key string){
	eg, err := models.GetErrorGbyKey(key)
	if err != nil {
		log.Printf("Error while getting errorGroup: %v", err)
	} else if eg.ID == 0 {

		if err := eg.CreateErrorG(lg.CreatedAt, key); err != nil {
			log.Printf("Error while creating errorGroup: %v", err)
		}
	}
	if err = eg.IncraeseCount(); err != nil {
		log.Printf("Error while increasing count of errorGroup: %v", err)
	}

	var ok bool = false
	if eg.BaselineAVG == 0 {
		eg.BaselineAVG = 1
		eg.UpdateErrorG()
	}
	if eg.BaselineAVG == 1 && eg.BaselineAVG != float64(eg.Count) {
		limitNumber := constants.LimitNumber
		eg.BaselineAVG = float64(eg.Count)
		ok=GetAlertStatus(limitNumber, eg)
	} else if eg.BaselineAVG == 1 && eg.BaselineAVG == float64(eg.Count) {
		eg.LastSeen = time.Now()
		eg.UpdateErrorG()
	} else {
		ok=GetAlertStatus(-1, eg)
	}
	if ok{
		SendAlert(eg)
	}
}
