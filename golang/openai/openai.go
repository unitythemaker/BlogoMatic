package openai

import (
	"context"
	"fmt"
	"os"
)

import (
	openai "github.com/sashabaranov/go-openai"
)

func GenerateArticle(prompt string, systemPrompt string) (article string, err error) {
	article, err = generate(prompt, systemPrompt)
	return
}

func GenerateTitle(article string, prompt string, systemPrompt string) (title string, err error) {
	title, err = generate(prompt, systemPrompt, article)
	return
}

func generate(prompt string, systemPrompt string, article ...string) (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	var messages []openai.ChatCompletionMessage
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: systemPrompt,
	})
	if len(article) > 0 {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: article[0],
		})
	}
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: prompt,
	})
	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: messages,
	})

	/*resp, err := client.CreateChatCompletion(
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
	)*/

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
