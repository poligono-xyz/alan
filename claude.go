package alan

import (
	"context"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go"
)

type ClaudeModel string

var (
	Claude1_3Opus20240229   ClaudeModel = "claude-1.3-opus-2024-02-29"
	Claude1_3Opus20240229V2 ClaudeModel = "claude-1.3-opus-2024-02-29-v2"
	Claude1_3Opus20240229V3 ClaudeModel = "claude-1.3-opus-2024-02-29-v3"
	Claude1_3Opus20240229V4 ClaudeModel = "claude-1.3-opus-2024-02-29-v4"

	Claude1_3Opus20240229V5 ClaudeModel = "claude-1.3-opus-2024-02-29-v5"
	Claude1_3Opus20240229V6 ClaudeModel = "claude-1.3-opus-2024-02-29-v6"
	Claude1_3Opus20240229V7 ClaudeModel = "claude-1.3-opus-2024-02-29-v7"
	Claude1_3Opus20240229V8 ClaudeModel = "claude-1.3-opus-2024-02-29-v8"
)

type claudeImpl struct {
	ctx    context.Context
	client anthropic.Client
	model  Model
}

func (self *claudeImpl) Prompt(t string) (string, error) {
	model, err := self.fromAlanModelToOpenAIChatModel(self.model)
	if err != nil {
		return "", err
	}

	message, err := self.client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		MaxTokens: 1024,
		Messages: []anthropic.MessageParam{{
			Role: anthropic.MessageParamRoleUser,
			Content: []anthropic.ContentBlockParamUnion{{
				OfRequestTextBlock: &anthropic.TextBlockParam{Text: t},
			}},
		}},
		Model: model,
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v\n", message.Content)
	return "", nil
}

func (self *claudeImpl) fromAlanModelToOpenAIChatModel(model Model) (anthropic.Model, error) {
	switch model {
	case ModelClaude3_7SonnetLatest:
		return anthropic.ModelClaude3_7SonnetLatest, nil
	default:
		return "", fmt.Errorf("unsupported model: %s", model)
	}
}
