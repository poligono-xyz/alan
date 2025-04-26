package alan

type Model string

// ChatGPT Models
const (
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
)

// Gemini Models
const (
	Gemini20Flash                Model = "gemini-2.0-flash"
	Gemini20FlashLitePreview0205 Model = "gemini-2.0-flash-lite-preview-02-05"
	Gemini15Flash                Model = "gemini-1.5-flash"
	Gemini15Flash8B              Model = "gemini-1.5-flash-8b"
	Gemini15Pro                  Model = "gemini-1.5-pro"
	GeminiTextEmbedding004       Model = "text-embedding-004"
)

// Claude Models
const (
	ModelClaude3_7SonnetLatest Model = "claude-3-7-sonnet-latest"
)
