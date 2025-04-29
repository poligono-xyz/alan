package prompts

import (
	"fmt"

	"github.com/poligono-xyz/alan"
	clientOption "github.com/poligono-xyz/alan/option/client"
)

func main() {
	client, err := alan.NewClient(
		clientOption.WithProvider(alan.ClaudeProvider),
		clientOption.WithModel(alan.ModelClaude3_7SonnetLatest),
		clientOption.WithAPIKey("YOUR_API_KEY"), // Replace with your Anthropic API key
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
