package alan

import (
	"context"

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
	ctx    context.Context
	client Client
	chat   interface{}
}

func (self *chatImpl) Prompt(t string) (string, error) {
	switch self.chat.(type) {
	case *genai.Chat:
		chat, ok := self.chat.(*genai.Chat)
		if !ok {
			return "", nil
		}

		result, err := chat.SendMessage(self.ctx, genai.Part{Text: t})
		if err != nil {
			return "", err
		}
		return result.Text(), nil
	default:
		return "", nil
	}
}
