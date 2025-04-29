package prompts

import (
	"fmt"

	"github.com/poligono-xyz/alan"
	clientOption "github.com/poligono-xyz/alan/option/client"
)

func main() {
	client, err := alan.NewClient(
		clientOption.WithProvider(alan.GeminiProvider),
		clientOption.WithModel(alan.Gemini15Flash),
		clientOption.WithAPIKey("YOUR_API_KEY"), // Replace with your Gemini API key
	)
	if err != nil {
		panic(err)
	}

	result, err := client.Prompt("What's the weather in the Dominican Republic?")
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
