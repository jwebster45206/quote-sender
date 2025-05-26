package ai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

// OpenAIProvider implements the Provider interface using OpenAI's API
type OpenAIProvider struct {
	client *openai.Client
	model  string
}

type OpenAIConfig struct {
	APIKey string
	Model  string
}

func NewOpenAIProvider(config OpenAIConfig) Provider {
	client := openai.NewClient(config.APIKey)
	return &OpenAIProvider{
		client: client,
		model:  config.Model,
	}
}

func (p *OpenAIProvider) GenerateQuote(ctx context.Context) (string, error) {
	resp, err := p.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: p.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: "You are Uncle Iroh from Avatar: The Last Airbender. " +
						"Generate a wise, thoughtful quote in his style about life, " +
						"wisdom, tea, balance, remorse, personal growth, and redemption. " +
						"The quote should be brief (1-2 sentences) and profound.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Generate an Uncle Iroh quote.",
				},
			},
			Temperature: 0.7,
			MaxTokens:   100,
		},
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate quote: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no quote generated")
	}

	return resp.Choices[0].Message.Content, nil
}
