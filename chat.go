package alan

import (
	"context"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/openai/openai-go"
	"google.golang.org/genai"
)

type ChatOption func(*ChatConfig)

type Chat interface {
	// Prompt sends a prompt to the chat model and returns the response.
	Prompt(t string) (string, error)
}

type ChatConfig struct {
	Temperature       float32 `json:"temperature,omitempty"`
	TopK              float32 `json:"TopK,omitempty"`
	TopP              float32 `json:"TopP,omitempty"`
	SystemInstruction string  `json:"systemInstruction,omitempty"`
	CandidateCount    int     `json:"candidateCount,omitempty"`
}

type chatImpl struct {
	ctx                              context.Context
	client                           Client
	internalProviderClient           interface{}
	providerClient                   interface{}
	chat                             interface{} // Empty for chatgpt
	chatCompletionMessageParamUnions []openai.ChatCompletionMessageParamUnion
	claudeMessages                   []anthropic.MessageParam
}

func (self *chatImpl) Prompt(t string) (string, error) {
	switch self.client.GetProvider() {
	case GeminiProvider:
		chat, ok := self.chat.(*genai.Chat)
		if !ok {
			return "", nil
		}

		result, err := chat.SendMessage(self.ctx, genai.Part{Text: t})
		if err != nil {
			return "", err
		}
		return result.Text(), nil
	case ChatGPTProvider:
		client, ok := self.providerClient.(*openai.Client)
		if !ok {
			return "", nil
		}

		internalClient, ok := self.internalProviderClient.(chatgptImpl)
		if !ok {
			return "", nil
		}

		self.chatCompletionMessageParamUnions = append(self.chatCompletionMessageParamUnions, openai.UserMessage(t))

		model, err := internalClient.fromAlanModelToOpenAIChatModel(self.client.GetModel())
		if err != nil {
			return "", err
		}

		chatCompletion, err := client.Chat.Completions.New(
			self.ctx, openai.ChatCompletionNewParams{
				Messages: openai.F(self.chatCompletionMessageParamUnions),
				Model:    openai.F(model),
			})
		if err != nil {
			return "", err
		}
		return chatCompletion.Choices[0].Message.Content, nil
	case ClaudeProvider:
		client, ok := self.providerClient.(*anthropic.Client)
		if !ok {
			return "", nil
		}

		self.claudeMessages = append(self.claudeMessages, anthropic.NewUserMessage(anthropic.NewTextBlock(t)))

		internalClient, ok := self.internalProviderClient.(claudeImpl)
		if !ok {
			return "", nil
		}

		model, err := internalClient.fromAlanModelToOpenAIChatModel(self.client.GetModel())
		if err != nil {
			return "", err
		}

		message, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
			Model:     model,
			Messages:  self.claudeMessages,
			MaxTokens: 1024,
		})
		if err != nil {
			return "", err
		}

		self.claudeMessages = append(self.claudeMessages, message.ToParam())
		return message.Content[0].Text, nil
	default:
		return "", nil
	}
}
