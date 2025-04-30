package alan

import (
	"context"

	"github.com/anthropics/anthropic-sdk-go"
	anthropicOption "github.com/anthropics/anthropic-sdk-go/option"
	"github.com/openai/openai-go"
	chatgptOption "github.com/openai/openai-go/option"
	"google.golang.org/genai"
)

type Client interface {
	GetProvider() Provider
	GetModel() Model
	NewChat(options ...ChatOption) (Chat, error)
	Prompt(t string) (string, error)
}

type ClientOption func(*ClientConfig)

type ClientConfig struct {
	Context  context.Context
	Provider Provider
	APIKey   string
	Model    Model
}

func NewClient(options ...ClientOption) (Client, error) {
	config := &ClientConfig{
		Context: context.Background(),
	}

	for _, option := range options {
		option(config)
	}

	switch config.Provider {
	case GeminiProvider:
		client, err := genai.NewClient(context.Background(), &genai.ClientConfig{
			APIKey:  config.APIKey,
			Backend: genai.BackendGeminiAPI,
		})
		if err != nil {
			return nil, err
		}

		return &geminiImpl{ctx: config.Context, client: client, model: config.Model}, nil
	case ChatGPTProvider:
		client := openai.NewClient(chatgptOption.WithAPIKey(config.APIKey))

		return &chatgptImpl{ctx: config.Context, client: client, model: config.Model}, nil
	case ClaudeProvider:
		client := anthropic.NewClient(anthropicOption.WithAPIKey(config.APIKey))

		return &claudeImpl{ctx: config.Context, client: client, model: config.Model}, nil
	default:
		return nil, nil
	}
}

// hello world
