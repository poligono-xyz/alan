package alan

import (
	"context"

	"google.golang.org/genai"
)

type GeminiModel string

var (
	Gemini20Flash                GeminiModel = "gemini-2.0-flash"
	Gemini20FlashLitePreview0205 GeminiModel = "gemini-2.0-flash-lite-preview-02-05"
	Gemini15Flash                GeminiModel = "gemini-1.5-flash"
	Gemini15Flash8B              GeminiModel = "gemini-1.5-flash-8b"
	Gemini15Pro                  GeminiModel = "gemini-1.5-pro"
	GeminiTextEmbedding004       GeminiModel = "text-embedding-004"
)

type GeminiConfig struct {
	Model  GeminiModel
	ApiKey string
}

type geminiImpl struct {
	ctx    context.Context
	client *genai.Client
	config GeminiConfig
}

func (self *geminiImpl) Prompt(t string) (string, error) {
	result, err := self.client.Models.GenerateContent(context.Background(), string(self.config.Model), genai.Text(t), nil)
	if err != nil {
		return t, err
	}

	return result.Text(), nil
}
