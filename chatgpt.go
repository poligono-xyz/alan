package alan

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
)

type ChatModel string

var (
	// GPT 4.5
	ChatModelGPT4_5Preview           ChatModel = "gpt-4-5-preview"
	ChatModelGPT4_5Preview2025_02_27 ChatModel = "gpt-4-5-preview-2025-02-27"

	// GPT 4
	ChatModelGPT4             ChatModel = "gpt-4"
	ChatModelGPT4_0125Preview ChatModel = "gpt-4-0125-preview"
	ChatModelGPT4_0314        ChatModel = "gpt-4-0314"
	ChatModelGPT4_0613        ChatModel = "gpt-4-0613"
	ChatModelGPT4_1106Preview ChatModel = "gpt-4-1106-preview"
	ChatModelGPT4_32k         ChatModel = "gpt-4-32k"
	ChatModelGPT4_32k0314     ChatModel = "gpt-4-32k-0314"
	ChatModelGPT4_32k0613     ChatModel = "gpt-4-32k-0613"

	// GPT 4 Turbo
	ChatModelGPT4Turbo           ChatModel = "gpt-4-turbo"
	ChatModelGPT4TurboPreview    ChatModel = "gpt-4-turbo-preview"
	ChatModelGPT4Turbo2024_04_09 ChatModel = "gpt-4-turbo-2024-04-09"

	// GPT 4o
	ChatModelGPT4o           ChatModel = "gpt-4o"
	ChatModelChatgpt4oLatest ChatModel = "chatgpt-4o-latest"
	ChatModelGPT4o2024_05_13 ChatModel = "gpt-4o-2024-05-13"
	ChatModelGPT4o2024_08_06 ChatModel = "gpt-4o-2024-08-06"
	ChatModelGPT4o2024_11_20 ChatModel = "gpt-4o-2024-11-20"

	// GPT 4o Mini
	ChatModelGPT4oMini           ChatModel = "gpt-4o-mini"
	ChatModelGPT4oMini2024_07_18 ChatModel = "gpt-4o-mini-2024-07-18"

	// GPT 3.5 Turbo
	ChatModelGPT3_5Turbo        ChatModel = "gpt-3.5-turbo"
	ChatModelGPT3_5Turbo0125    ChatModel = "gpt-3.5-turbo-0125"
	ChatModelGPT3_5Turbo0301    ChatModel = "gpt-3.5-turbo-0301"
	ChatModelGPT3_5Turbo0613    ChatModel = "gpt-3.5-turbo-0613"
	ChatModelGPT3_5Turbo1106    ChatModel = "gpt-3.5-turbo-1106"
	ChatModelGPT3_5Turbo16k     ChatModel = "gpt-3.5-turbo-16k"
	ChatModelGPT3_5Turbo16k0613 ChatModel = "gpt-3.5-turbo-16k-0613"
)

type ChatGPTConfig struct {
	Model  ChatModel
	ApiKey string
}

type chatgptImpl struct {
	ctx    context.Context
	client *openai.Client
	config ChatGPTConfig
}

func (self *chatgptImpl) Prompt(t string) (string, error) {
	model, err := self.fromAlanChatModelTOpenAIChatModel(self.config.Model)
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

func (self *chatgptImpl) fromAlanChatModelTOpenAIChatModel(model ChatModel) (openai.ChatModel, error) {
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

// 	return "", fmt.Errorf("unsupported model: %s", model)

// func (self *chatgptImpl) PromptStream(t string) (string, error) {
// 	chatCompletion, err := self.client.Chat.Completions.New(self.ctx, openai.ChatCompletionNewParams{
// 		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
// 			openai.UserMessage(t),
// 		}),
// 		Model: openai.F(openai.ChatModelGPT4o),
// 	})
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println(chatCompletion.Choices[0].Message.Content)
// 	return t, nil
// }
// func (self *chatgptImpl) PromptStreamWithContext(t string, ctx context.Context) (string, error) {
// 	chatCompletion, err := self.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
// 		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
// 			openai.UserMessage(t),
// 		}),
// 		Model: openai.F(openai.ChatModelGPT4o),
// 	})
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println(chatCompletion.Choices[0].Message.Content)
// 	return t, nil
// }
// func (self *chatgptImpl) PromptStreamWithContextAndParams(t string, ctx context.Context, params openai.ChatCompletionNewParams) (string, error) {
// 	chatCompletion, err := self.client.Chat.Completions.New(ctx, params)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println(chatCompletion.Choices[0].Message.Content)
// 	return t, nil
// }
