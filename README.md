# Alan

The `alan` library provides a unified interface for interacting with multiple large language models (LLMs), including OpenAI's ChatGPT and Google's Gemini. It simplifies the process of configuring and using these models in your Go applications.

## Features

- **Multi-Model Support**: Seamlessly interact with ChatGPT and Gemini models.
- **Customizable Configurations**: Configure API keys and model versions for each provider.
- **Unified Interface**: Use a single interface to prompt different LLMs.
- **Extensible**: Easily add support for additional models.

## Installation

To use the `alan` library, add it to your Go project:

```bash
go get github.com/poligono-xyz/alan
```

## Usage

Creating a Client

You can create a client for Gemini, ChatGPT, or Claude by providing the appropriate configuration.

### ChatGPT Examples

<details>
<summary>Prompts</summary>

```go
import (
  "fmt"

  "github.com/poligono-xyz/alan"
  clientOption "github.com/poligono-xyz/alan/option/client"
)

func main() {
  client, err := alan.NewClient(
    clientOption.WithProvider(alan.ChatGPTProvider),
    clientOption.WithModel(alan.ChatModelGPT4oMini),
    clientOption.WithAPIKey("YOUR_API_KEY"), // Replace with your OpenAI API key
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
```

</details>

<details>
<summary>Chats</summary>

```go
import (
  "github.com/poligono-xyz/alan"
  chatOption "github.com/poligono-xyz/alan/option/chat"
  clientOption "github.com/poligono-xyz/alan/option/client"
)

func main() {
  client, err := alan.NewClient(
    clientOption.WithProvider(alan.ChatGPTProvider),
    clientOption.WithModel(alan.ChatModelGPT4),
    clientOption.WithAPIKey("YOUR_API_KEY"),
  )

  if err != nil {
    panic(err)
  }

  model := client.GetModel()
  println("Model: ", model)

  provider := client.GetProvider()
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
```

</details>

### Gemini Examples

<details>
<summary>Prompts</summary>

```go
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
```

</details>

<details>
<summary>Chats</summary>

```go
import (
  "github.com/poligono-xyz/alan"
  chatOption "github.com/poligono-xyz/alan/option/chat"
  clientOption "github.com/poligono-xyz/alan/option/client"
)

func main() {
  client, err := alan.NewClient(
    clientOption.WithProvider(alan.GeminiProvider),
    clientOption.WithModel(alan.Gemini15Flash),
    clientOption.WithAPIKey("YOUR_API_KEY"),
  )

  if err != nil {
    panic(err)
  }

  model := client.GetModel()
  println("Model: ", model)

  provider := client.GetProvider()
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

```

</details>

### Claude Examples

<details>
<summary>Prompts</summary>

```go
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

```

</details>

<details>
<summary>Chats</summary>

```go
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
  println("Model: ", model)

  provider := client.GetProvider()
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
```

</details>

## Supported Models

ChatGPT

- GPT-4
- GPT-4 Turbo
- GPT-3.5 Turbo
- GPT-4o and GPT-4o Mini

Gemini

- Gemini 2.0 Flash
- Gemini 1.5 Flash
- Gemini Text Embedding

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.
