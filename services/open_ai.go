package services

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

func OpenAIService() {

	viper.AutomaticEnv()

	viper.SetConfigFile("configs/env.yml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Erro ao ler arquivo de configuração:", err.Error())
		return
	}

	apiKey := viper.GetString("APIKEY")

	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Qual a cidade que tem a estátua da liberdade?",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
