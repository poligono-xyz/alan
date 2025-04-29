package chats

import (
	"github.com/poligono-xyz/alan"
	chatOption "github.com/poligono-xyz/alan/option/chat"
	clientOption "github.com/poligono-xyz/alan/option/client"
)

func main() {
	client, err := alan.NewClient(
		clientOption.WithProvider(alan.ClaudeProvider),
		clientOption.WithModel(alan.ModelClaude3_7SonnetLatest),
		clientOption.WithAPIKey("YOUR_API_KEY"),
	)

	if err != nil {
		panic(err)
	}

	model := client.GetModel()
	provider := client.GetProvider()
	println("Model: ", model)
	println("Provider: ", provider)

	chat, err := client.NewChat(chatOption.WithTemperature(0.5))
	if err != nil {
		panic(err)
	}

	result, err := chat.Prompt("What's the weather in the Dominican Republic?")
	if err != nil {
		panic(err)
	}
	println(result)

	result, err = chat.Prompt("What's the weather in the Japan?")
	if err != nil {
		panic(err)
	}
	println(result)
}
