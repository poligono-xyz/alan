package main

import (
	"github.com/poligono-xyz/alan"
	chatOption "github.com/poligono-xyz/alan/option/chat"
	clientOption "github.com/poligono-xyz/alan/option/client"
)

func main() {
	gemini, err := alan.NewClient(
		clientOption.WithProvider(alan.GeminiProvider),
		clientOption.WithModel(alan.Gemini15Flash),
		clientOption.WithAPIKey("YOUR_API_KEY"),
	)

	if err != nil {
		panic(err)
	}

	model := gemini.GetModel()
	provider := gemini.GetProvider()
	println("Model: ", model)
	println("Provider: ", provider)

	chat, err := gemini.NewChat(chatOption.WithTemperature(0.5))
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
