package services

import (
	"context"
	"fmt"
	"logAnalyzer/models"
	"strings"

	"logAnalyzer/env"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func GetSummary(eg *models.ErrorGroup) error {
	summary := GetAnalayzeResponse(eg)
	if summary != "" {
		eg.AISummary = summary
		eg.UpdateErrorG()
	}
	return nil
}

func GetConfidence(eg *models.ErrorGroup) error {
	var confidence string = "Null"
	response := GetConfidenceResponse(eg)
	response=strings.ToLower(response)
	
	if ok:=strings.Contains(response,"low");ok{
		confidence="low"
	}

	if ok:=strings.Contains(response,"medium");ok{
		if confidence != "low"{
			confidence="medium"
		}
	}

	if ok:=strings.Contains(response,"high");ok{
		if confidence != "medium"{
			confidence="high"
		}
	}
	eg.Confidence = confidence
	eg.UpdateErrorG()
	return nil
}

func GetAnalayzeResponse(eg *models.ErrorGroup) string {
	BaseUrl := "http://localhost:1234/v1"
	developerMessage := env.GetEnvString("SUMMARY_DEVELOPER_MESSAGE", "")
	groupKey := strings.Split(eg.GroupKey, "|")
	message := fmt.Sprintf(
		"What is the most likely reason for this error?\nInformation\nService: %s\nLevel: %s\nMessage: %s\nCount of log: %d",
		groupKey[0], groupKey[1], groupKey[2], eg.Count,
	)

	client := openai.NewClient(
		option.WithAPIKey("lm-studio-key"),
		option.WithBaseURL(BaseUrl),
	)

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.DeveloperMessage(developerMessage),
			openai.UserMessage(message),
		},
		Model: "gemma-3n-e4b-it-text",
	})
	if err != nil {
		panic(err)
	}

	return chatCompletion.Choices[0].Message.Content

}

func GetConfidenceResponse(eg *models.ErrorGroup) string {
	BaseUrl := "http://localhost:1234/v1"
	developerMessage := env.GetEnvString("CONFIDENCE_DEVELOPER_MESSAGE", "")
	groupKey := strings.Split(eg.GroupKey, "|")
	message := fmt.Sprintf(
		"Explain this repeated error briefly and return just a confidence:\nInformation\nService: %s\nLevel: %s\nMessage: %s\nCount of log: %d",
		groupKey[0], groupKey[1], groupKey[2], eg.Count,
	)

	client := openai.NewClient(
		option.WithAPIKey("lm-studio-key"),
		option.WithBaseURL(BaseUrl),
	)

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.DeveloperMessage(developerMessage),
			openai.UserMessage(message),
		},
		Model: "gemma-3n-e4b-it-text",
	})
	if err != nil {
		panic(err)
	}

	return chatCompletion.Choices[0].Message.Content

}
