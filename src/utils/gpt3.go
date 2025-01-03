package utils

import (
	"context"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func GetGPT4Response(userMessage string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(apiKey)
	ctx := context.Background()

	messages := []openai.ChatCompletionMessage{
		{
			Role:    "user",
			Content: userMessage,
		},
	}

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT4Turbo,
		Messages:  messages,
		MaxTokens: 150,
	}

	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) > 0 {
		return resp.Choices[0].Message.Content, nil
	}

	return "Sorry, I couldn't understand that.", nil
}
