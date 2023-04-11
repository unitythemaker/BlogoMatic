package openai

import (
	"context"
	"fmt"
	"os"
)

import (
	openai "github.com/sashabaranov/go-openai"
)

func GenerateArticle(prompt string, systemPrompt string) (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	if resp.Choices[0].FinishReason != "stop" {
		fmt.Printf("ChatCompletion error (FinishReason): %v\n", resp.Choices[0].FinishReason)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
