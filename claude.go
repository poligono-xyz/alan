// package main

// import (
//     "fmt"
//     "github.com/3JoB/anthropic-sdk-go/v2"
//     "github.com/3JoB/anthropic-sdk-go/v2/data"
// )

// func main() {
//     // Replace 'YOUR_API_KEY' with your actual Anthropic API key
//     client := anthropic.NewClient("YOUR_API_KEY")

//     // Create a completion request
//     request := data.CompleteRequest{
//         Model: anthropic.ModelClaude3Opus20240229,
//         Prompt: "Translate the following English text to French: 'Hello, how are you?'",
//         MaxTokens: 100,
//     }

//     // Generate the completion
//     response, err := client.CreateComplete(context.Background(), request)
//     if err != nil {
//         fmt.Println("Error generating completion:", err)
//         return
//     }

//	    fmt.Println("Generated text:", response.Completion)
//	}
package alan
