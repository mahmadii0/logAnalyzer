package services

import (
	"log"
	"logAnalyzer/env"
	"logAnalyzer/models"
	"logAnalyzer/pkg"
	"strconv"
	"time"
)

func GetAlertStatus(limitNumber int, eg *models.ErrorGroup) bool {
	if limitNumber != -1 {
		if eg.Count >= limitNumber {
			if eg.AISummary == "" {
				err := GetSummary(eg)
				if err != nil {
					log.Printf("Error While Getting Summary")
				}
				eg.Confidence = "low"
			} else {
				err := GetConfidence(eg)
				if err != nil {
					log.Printf("Error While Getting Confidence")
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
	if eg.Count == 1 {
		eg.CalBaselineAVG()
	}
	eg.LastSeen = time.Now()
	eg.UpdateErrorG()
	if eg.Count >= int(eg.BaselineAVG) {
		if eg.AISummary == "" {
			err := GetSummary(eg)
			if err != nil {
				log.Printf("Error While Getting Summary")
			}
			eg.Confidence = "low"
		} else {
			err := GetConfidence(eg)
			if err != nil {
				log.Printf("Error While Getting Confidence")
			}
		}
		return true
	}
	return false
}

func SendAlert(eg *models.ErrorGroup) {
	telegram := env.GetEnvInt("TELEGRAM_STATUS", 0)
	slack := env.GetEnvInt("SLACK_STATUS", 0)
	whatsApp := env.GetEnvInt("WHATSAPP_STATUS", 0)
	if telegram == 0 && slack == 0 && whatsApp == 0 {
		log.Printf("All of messengers are OFF")
		return
	}

	//Message Text
	groupKey := pkg.DecodeGroupKey(eg.GroupKey)
	if groupKey == nil {
		log.Printf("Error While Decoding groupKey")
		return
	}
	count := strconv.Itoa(eg.Count)
	id := strconv.Itoa(int(eg.ID))
	txt := "🔥 HIGH SEVERITY ERROR\n\nService: " + groupKey["service"] +
		"\nID: " + id + "\nConfidence: " + eg.Confidence + "\nError Message: " +
		groupKey["message"] + "\nCount: " + count + "\n\nAI Summary: " +
		eg.AISummary

	if telegram == 1 {
		token := env.GetEnvString("TELEGRAM_BOT_TOKEN", "")
		cId := env.GetEnvString("CHAT_ID", "")
		chatId, err := strconv.Atoi(cId)
		if err != nil || cId == "" {
			log.Printf("Erorr While Getting Enviorments")
		}
		SendTelegramMessage(int64(chatId), txt, token)
	}
	if slack == 1 {
		token := env.GetEnvString("SLACK_BOT_TOKEN", "")
		channelId := env.GetEnvString("CHANNEL_ID", "")
		if channelId == "" {
			log.Printf("Erorr While Getting Enviorments")
		}
		sendSlackMessage(channelId, txt, token)
	}
	if whatsApp == 1 {
		token := env.GetEnvString("WHATSAPP_BOT_TOKEN", "")
		phoneId := env.GetEnvString("PHONE_ID", "")
		to:=env.GetEnvString("TO","")
		if token == "" || phoneId == "" || to =="" {
			log.Printf("Erorr While Getting Enviorments")
		}
		sendWhatsAppMessage(token,phoneId,to, txt)
	}

}
