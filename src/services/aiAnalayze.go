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



func GetAnalayzeResponse(eg *models.ErrorGroup) string{
	BaseUrl:="http://localhost:1234/v1"
	developerMessage := env.GetEnvString("DEVELOPER_MESSAGE", "")
	groupKey:=strings.Split(eg.GroupKey, "|")
	message:=fmt.Sprintf(
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