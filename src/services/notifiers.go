package services

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/slack-go/slack"
	"bytes"
	"fmt"
	"net/http"
)

//Telegram Bot

var bot *tgbotapi.BotAPI

func SendTelegramMessage(chatId int64, text string, token string) {
	bot, err := tgbotapi.NewBotAPI(token)
    if err != nil {
        log.Printf("Error While Connecting To Telegram Bot")
		return
    }
    msg := tgbotapi.NewMessage(chatId, text)
    _, err = bot.Send(msg)
	if err !=nil{
		log.Printf("Error While Sending Message")
	}
}

//Slack Bot

func sendSlackMessage(channelId string, text string, token string) {
	client := slack.New(token)
	_, _, err := client.PostMessage(channelId, slack.MsgOptionText(text, false))
	log.Printf("%v",err)
}

//WhatsApp Bot

func sendWhatsAppMessage(token, phoneID, to, text string) {
	body := fmt.Sprintf(`{
		"messaging_product": "whatsapp",
		"to": "%s",
		"type": "text",
		"text": {"body": "%s"}
	}`, to, text)

	req, _ := http.NewRequest("POST",
		"https://graph.facebook.com/v19.0/"+phoneID+"/messages",
		bytes.NewBufferString(body))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("%v",err)
		return 
	}
	defer resp.Body.Close()
}