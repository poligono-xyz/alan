package alan

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
)

type chatgptImpl struct {
	ctx    context.Context
	client *openai.Client
	model  Model
}

func (self *chatgptImpl) Prompt(t string) (string, error) {
	model, err := self.fromAlanModelToOpenAIChatModel(self.model)
	if err != nil {
		return "", err
	}

	chatCompletion, err := self.client.Chat.Completions.New(self.ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(t),
		}),
		Model: openai.F(model),
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(chatCompletion.Choices[0].Message.Content)
	return t, nil
}

func (self *chatgptImpl) GetModel() Model {
	return self.model
}

func (self *chatgptImpl) GetProvider() Provider {
	return ChatGPTProvider
}

func (self *chatgptImpl) SetModel(model Model) {
	self.model = model
}

func (self *chatgptImpl) NewChat(options ...ChatOption) (Chat, error) {
	config := &ChatConfig{
		Temperature: 0.5,
	}

	for _, option := range options {
		option(config)
	}
	chat := &chatImpl{}
	return chat, nil
}

func (self *chatgptImpl) fromAlanModelToOpenAIChatModel(model Model) (openai.ChatModel, error) {
	switch model {
	// GPT 4.5
	case ChatModelGPT4_5Preview:
		return openai.ChatModelGPT4_5Preview, nil
	case ChatModelGPT4_5Preview2025_02_27:
		return openai.ChatModelGPT4_5Preview2025_02_27, nil

	//GPT 4
	case ChatModelGPT4:
		return openai.ChatModelGPT4, nil
	case ChatModelGPT4_0125Preview:
		return openai.ChatModelGPT4_0125Preview, nil
	case ChatModelGPT4_0314:
		return openai.ChatModelGPT4_0314, nil
	case ChatModelGPT4_0613:
		return openai.ChatModelGPT4_0613, nil
	case ChatModelGPT4_1106Preview:
		return openai.ChatModelGPT4_1106Preview, nil
	case ChatModelGPT4_32k:
		return openai.ChatModelGPT4_32k, nil
	case ChatModelGPT4_32k0314:
		return openai.ChatModelGPT4_32k0314, nil
	case ChatModelGPT4_32k0613:
		return openai.ChatModelGPT4_32k0613, nil

	// GPT 4 Turbo
	case ChatModelGPT4Turbo:
		return openai.ChatModelGPT4Turbo, nil
	case ChatModelGPT4TurboPreview:
		return openai.ChatModelGPT4TurboPreview, nil
	case ChatModelGPT4Turbo2024_04_09:
		return openai.ChatModelGPT4Turbo2024_04_09, nil

	// GPT 4o
	case ChatModelGPT4o:
		return openai.ChatModelGPT4o, nil
	case ChatModelChatgpt4oLatest:
		return openai.ChatModelChatgpt4oLatest, nil
	case ChatModelGPT4o2024_05_13:
		return openai.ChatModelGPT4o2024_05_13, nil
	case ChatModelGPT4o2024_08_06:
		return openai.ChatModelGPT4o2024_08_06, nil
	case ChatModelGPT4o2024_11_20:
		return openai.ChatModelGPT4o2024_11_20, nil

	// GPT 4o Mini
	case ChatModelGPT4oMini:
		return openai.ChatModelGPT4oMini, nil
	case ChatModelGPT4oMini2024_07_18:
		return openai.ChatModelGPT4oMini2024_07_18, nil

	// GPT 3.5 Turbo
	case ChatModelGPT3_5Turbo:
		return openai.ChatModelGPT3_5Turbo, nil
	case ChatModelGPT3_5Turbo0125:
		return openai.ChatModelGPT3_5Turbo0125, nil
	case ChatModelGPT3_5Turbo0301:
		return openai.ChatModelGPT3_5Turbo0301, nil
	case ChatModelGPT3_5Turbo0613:
		return openai.ChatModelGPT3_5Turbo0613, nil
	case ChatModelGPT3_5Turbo1106:
		return openai.ChatModelGPT3_5Turbo1106, nil
	case ChatModelGPT3_5Turbo16k:
		return openai.ChatModelGPT3_5Turbo16k, nil
	case ChatModelGPT3_5Turbo16k0613:
		return openai.ChatModelGPT3_5Turbo16k0613, nil

	default:
		return "", fmt.Errorf("unsupported model: %s", model)
	}
}
