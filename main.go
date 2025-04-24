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
	Prompt(t string) (string, error)
}

type Option func(*Config)

type Model string

type Provider string

const (
	ChatGPTProvider Provider = "chatgpt"
	GeminiProvider  Provider = "gemini"
	ClaudeProvider  Provider = "claude"
)

const (
	// ChatGPT Models
	// GPT 4.5
	ChatModelGPT4_5Preview           Model = "gpt-4-5-preview"
	ChatModelGPT4_5Preview2025_02_27 Model = "gpt-4-5-preview-2025-02-27"
	// GPT 4
	ChatModelGPT4             Model = "gpt-4"
	ChatModelGPT4_0125Preview Model = "gpt-4-0125-preview"
	ChatModelGPT4_0314        Model = "gpt-4-0314"
	ChatModelGPT4_0613        Model = "gpt-4-0613"
	ChatModelGPT4_1106Preview Model = "gpt-4-1106-preview"
	ChatModelGPT4_32k         Model = "gpt-4-32k"
	ChatModelGPT4_32k0314     Model = "gpt-4-32k-0314"
	ChatModelGPT4_32k0613     Model = "gpt-4-32k-0613"
	// GPT 4 Turbo
	ChatModelGPT4Turbo           Model = "gpt-4-turbo"
	ChatModelGPT4TurboPreview    Model = "gpt-4-turbo-preview"
	ChatModelGPT4Turbo2024_04_09 Model = "gpt-4-turbo-2024-04-09"
	// GPT 4o
	ChatModelGPT4o           Model = "gpt-4o"
	ChatModelChatgpt4oLatest Model = "chatgpt-4o-latest"
	ChatModelGPT4o2024_05_13 Model = "gpt-4o-2024-05-13"
	ChatModelGPT4o2024_08_06 Model = "gpt-4o-2024-08-06"
	ChatModelGPT4o2024_11_20 Model = "gpt-4o-2024-11-20"
	// GPT 4o Mini
	ChatModelGPT4oMini           Model = "gpt-4o-mini"
	ChatModelGPT4oMini2024_07_18 Model = "gpt-4o-mini-2024-07-18"
	// GPT 3.5 Turbo
	ChatModelGPT3_5Turbo        Model = "gpt-3.5-turbo"
	ChatModelGPT3_5Turbo0125    Model = "gpt-3.5-turbo-0125"
	ChatModelGPT3_5Turbo0301    Model = "gpt-3.5-turbo-0301"
	ChatModelGPT3_5Turbo0613    Model = "gpt-3.5-turbo-0613"
	ChatModelGPT3_5Turbo1106    Model = "gpt-3.5-turbo-1106"
	ChatModelGPT3_5Turbo16k     Model = "gpt-3.5-turbo-16k"
	ChatModelGPT3_5Turbo16k0613 Model = "gpt-3.5-turbo-16k-0613"

	// Gemini Models
	Gemini20Flash                Model = "gemini-2.0-flash"
	Gemini20FlashLitePreview0205 Model = "gemini-2.0-flash-lite-preview-02-05"
	Gemini15Flash                Model = "gemini-1.5-flash"
	Gemini15Flash8B              Model = "gemini-1.5-flash-8b"
	Gemini15Pro                  Model = "gemini-1.5-pro"
	GeminiTextEmbedding004       Model = "text-embedding-004"

	// Claude Models
	ModelClaude3_7SonnetLatest Model = "claude-3-7-sonnet-latest"
)

var (
	ChatGPT Provider = "ChatGPT"
	Gemini  Provider = "Gemini"
	Claude  Provider = "Claude"
)

type Config struct {
	Context  context.Context
	Provider Provider
	APIKey   string
	Model    Model
}

func NewClient(options ...Option) (Client, error) {
	config := &Config{
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
	case Claude:
		client := anthropic.NewClient(anthropicOption.WithAPIKey(config.APIKey))

		return &claudeImpl{ctx: config.Context, client: client, model: config.Model}, nil
	default:
		return nil, nil
	}
}

// func StringToGeminiModel(model string) (Model, error) {
// 	switch model {
// 	case "gemini-2.0-flash":
// 		return Gemini20Flash, nil
// 	case "gemini-2.0-flash-lite-preview-02-05":
// 		return Gemini20FlashLitePreview0205, nil
// 	case "gemini-1.5-flash":
// 		return Gemini15Flash, nil
// 	case "gemini-1.5-flash-8b":
// 		return Gemini15Flash8B, nil
// 	case "gemini-1.5-pro":
// 		return Gemini15Pro, nil
// 	case "text-embedding-004":
// 		return GeminiTextEmbedding004, nil
// 	default:
// 		return "", errors.New("unsupported model")
// 	}
// }

// func StringToChatModel(model string) (ChatModel, error) {
// 	switch model {
// 	// GPT 4.5
// 	case "gpt-4-5-preview":
// 		return ChatModelGPT4_5Preview, nil
// 	case "gpt-4-5-preview-2025-02-27":
// 		return ChatModelGPT4_5Preview2025_02_27, nil

// 	// GPT 4
// 	case "gpt-4":
// 		return ChatModelGPT4, nil
// 	case "gpt-4-0125-preview":
// 		return ChatModelGPT4_0125Preview, nil
// 	case "gpt-4-0314":
// 		return ChatModelGPT4_0314, nil
// 	case "gpt-4-0613":
// 		return ChatModelGPT4_0613, nil
// 	case "gpt-4-1106-preview":
// 		return ChatModelGPT4_1106Preview, nil
// 	case "gpt-4-32k":
// 		return ChatModelGPT4_32k, nil
// 	case "gpt-4-32k-0314":
// 		return ChatModelGPT4_32k0314, nil
// 	case "gpt-4-32k-0613":
// 		return ChatModelGPT4_32k0613, nil

// 	// GPT 4 Turbo
// 	case "gpt-4-turbo":
// 		return ChatModelGPT4Turbo, nil
// 	case "gpt-4-turbo-preview":
// 		return ChatModelGPT4TurboPreview, nil

// 	// GPT 4o
// 	case "gpt-4o":
// 		return ChatModelGPT4o, nil
// 	case "chatgpt-4o-latest":
// 		return ChatModelChatgpt4oLatest, nil
// 	case "gpt-4o-2024-05-13":
// 		return ChatModelGPT4o2024_05_13, nil
// 	case "gpt-4o-2024-08-06":
// 		return ChatModelGPT4o2024_08_06, nil
// 	case "gpt-4o-2024-11-20":
// 		return ChatModelGPT4o2024_11_20, nil

// 	// GPT 4o Mini
// 	case "gpt-4o-mini":
// 		return ChatModelGPT4oMini, nil
// 	case "gpt-4o-mini-2024-07-18":
// 		return ChatModelGPT4oMini2024_07_18, nil

// 	// GPT 3.5 Turbo
// 	case "gpt-3.5-turbo":
// 		return ChatModelGPT3_5Turbo, nil
// 	case "gpt-3.5-turbo-0125":
// 		return ChatModelGPT3_5Turbo0125, nil
// 	case "gpt-3.5-turbo-0613":
// 		return ChatModelGPT3_5Turbo0613, nil
// 	case "gpt-3.5-turbo-1106":
// 		return ChatModelGPT3_5Turbo1106, nil
// 	case "gpt-3.5-turbo-16k":
// 		return ChatModelGPT3_5Turbo16k, nil
// 	case "gpt-3.5-turbo-16k-0613":
// 		return ChatModelGPT3_5Turbo16k0613, nil

// 	default:
// 		return "", errors.New("unsupported model")
// 	}
// }
