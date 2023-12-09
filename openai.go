package main

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type OpenAIConstants struct {
	Key   string
	Model string
}

func FetchOpenAI(content string, obj OpenAIConstants) string {
	client := openai.NewClient(obj.Key)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: obj.Model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)
	CheckError("ChatCompletion error", err)

	answer := resp.Choices[0].Message.Content
	return answer
}
