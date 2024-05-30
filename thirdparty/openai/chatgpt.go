package openai

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"os"
)

var Client = openai.NewClient(requireOpenAiAuthToken())

func requireOpenAiAuthToken() string {
	token, ok := os.LookupEnv("OPENAI_AUTH_TOKEN")
	if !ok {
		panic("require a OPENAI_AUTH_TOKEN environment variable.")
	}
	return token
}

func AiMagic(content string) (string, error) {
	resp, err := Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
