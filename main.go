package alan

import (
	"context"
	"errors"

	"github.com/openai/openai-go"
	chatgptOption "github.com/openai/openai-go/option"
	"google.golang.org/genai"
)

type Client interface {
	Prompt(t string) (string, error)
}

func NewClient(ctx context.Context, clientConfig interface{}) (Client, error) {
	switch config := clientConfig.(type) {
	case GeminiConfig:
		client, err := genai.NewClient(context.Background(), &genai.ClientConfig{
			APIKey:  config.ApiKey,
			Backend: genai.BackendGeminiAPI,
		})
		if err != nil {
			return nil, err
		}

		return &geminiImpl{ctx: ctx, client: client, config: config}, nil
	case ChatGPTConfig:
		client := openai.NewClient(chatgptOption.WithAPIKey(config.ApiKey))

		return &chatgptImpl{client: client, config: config}, nil
	default:
		return nil, nil
	}
}

func StringToGeminiModel(model string) (GeminiModel, error) {
	switch model {
	case "gemini-2.0-flash":
		return Gemini20Flash, nil
	case "gemini-2.0-flash-lite-preview-02-05":
		return Gemini20FlashLitePreview0205, nil
	case "gemini-1.5-flash":
		return Gemini15Flash, nil
	case "gemini-1.5-flash-8b":
		return Gemini15Flash8B, nil
	case "gemini-1.5-pro":
		return Gemini15Pro, nil
	case "text-embedding-004":
		return GeminiTextEmbedding004, nil
	default:
		return "", errors.New("unsupported model")
	}
}

func StringToChatModel(model string) (ChatModel, error) {
	switch model {
	// GPT 4.5
	case "gpt-4-5-preview":
		return ChatModelGPT4_5Preview, nil
	case "gpt-4-5-preview-2025-02-27":
		return ChatModelGPT4_5Preview2025_02_27, nil

	// GPT 4
	case "gpt-4":
		return ChatModelGPT4, nil
	case "gpt-4-0125-preview":
		return ChatModelGPT4_0125Preview, nil
	case "gpt-4-0314":
		return ChatModelGPT4_0314, nil
	case "gpt-4-0613":
		return ChatModelGPT4_0613, nil
	case "gpt-4-1106-preview":
		return ChatModelGPT4_1106Preview, nil
	case "gpt-4-32k":
		return ChatModelGPT4_32k, nil
	case "gpt-4-32k-0314":
		return ChatModelGPT4_32k0314, nil
	case "gpt-4-32k-0613":
		return ChatModelGPT4_32k0613, nil

	// GPT 4 Turbo
	case "gpt-4-turbo":
		return ChatModelGPT4Turbo, nil
	case "gpt-4-turbo-preview":
		return ChatModelGPT4TurboPreview, nil

	// GPT 4o
	case "gpt-4o":
		return ChatModelGPT4o, nil
	case "chatgpt-4o-latest":
		return ChatModelChatgpt4oLatest, nil
	case "gpt-4o-2024-05-13":
		return ChatModelGPT4o2024_05_13, nil
	case "gpt-4o-2024-08-06":
		return ChatModelGPT4o2024_08_06, nil
	case "gpt-4o-2024-11-20":
		return ChatModelGPT4o2024_11_20, nil

	// GPT 4o Mini
	case "gpt-4o-mini":
		return ChatModelGPT4oMini, nil
	case "gpt-4o-mini-2024-07-18":
		return ChatModelGPT4oMini2024_07_18, nil

	// GPT 3.5 Turbo
	case "gpt-3.5-turbo":
		return ChatModelGPT3_5Turbo, nil
	case "gpt-3.5-turbo-0125":
		return ChatModelGPT3_5Turbo0125, nil
	case "gpt-3.5-turbo-0613":
		return ChatModelGPT3_5Turbo0613, nil
	case "gpt-3.5-turbo-1106":
		return ChatModelGPT3_5Turbo1106, nil
	case "gpt-3.5-turbo-16k":
		return ChatModelGPT3_5Turbo16k, nil
	case "gpt-3.5-turbo-16k-0613":
		return ChatModelGPT3_5Turbo16k0613, nil

	default:
		return "", errors.New("unsupported model")
	}
}
