package alan

import (
	"context"
	"encoding/json"

	"google.golang.org/genai"
)

type geminiImpl struct {
	ctx    context.Context
	client *genai.Client
	model  Model
}

func (self *geminiImpl) Prompt(t string) (string, error) {
	result, err := self.client.Models.GenerateContent(context.Background(), string(self.model), genai.Text(t), nil)
	if err != nil {
		return t, err
	}

	return result.Text(), nil
}

func (self *geminiImpl) GetModel() Model {
	return self.model
}

func (self *geminiImpl) SetModel(model Model) {
	self.model = model
}

func (self *geminiImpl) GetProvider() Provider {
	return GeminiProvider
}

func (self *geminiImpl) NewChat(options ...ChatOption) (Chat, error) {
	config := &ChatConfig{}

	for _, option := range options {
		option(config)
	}

	jsonConfig, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	var geminiChatConfig genai.GenerateContentConfig
	err = json.Unmarshal(jsonConfig, &geminiChatConfig)
	if err != nil {
		return nil, err
	}
	// var geminiChatConfig *genai.GenerateContentConfig = &genai.GenerateContentConfig{Temperature: genai.Ptr(config.Temperature), TopK: genai.Ptr(config.TopK), TopP: genai.Ptr(config.TopP), SystemInstruction: config.SystemInstruction, CandidateCount: config.CandidateCount}

	geminiChat, err := self.client.Chats.Create(self.ctx, string(self.model), &geminiChatConfig, nil)
	if err != nil {
		return nil, err
	}

	chat := &chatImpl{
		ctx:    self.ctx,
		client: self,
		chat:   geminiChat,
	}
	return chat, nil
}
