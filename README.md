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

You can create a client for either ChatGPT or Gemini by providing the appropriate configuration.

### ChatGPT Example

```go
import (
  "fmt"

  "github.com/poligono-xyz/alan"
  "github.com/poligono-xyz/alan/option"
)

func main() {
  client, err := alan.NewClient(
    option.WithProvider(alan.GeminiProvider),
    option.WithModel(alan.ChatModelChatgpt4oLatest),
    option.WithAPIKey("api-key"),
  )
  if err != nil {
    panic(err)
  }

  result, err := client.Prompt("Where is the Dominican Republic located?")
  if err != nil {
    panic(err)
  }

  fmt.Println(result)
}
```

### Gemini Example

```go
import (
  "fmt"

  "github.com/poligono-xyz/alan"
  "github.com/poligono-xyz/alan/option"
)

func main() {
  client, err := alan.NewClient(
    option.WithProvider(alan.GeminiProvider),
    option.WithModel(alan.Gemini15Flash),
    option.WithAPIKey("api-key"),
  )
  if err != nil {
    panic(err)
  }

  result, err := client.Prompt("Where is the Dominican Republic located?")
  if err != nil {
    panic(err)
  }

  fmt.Println(result)
}

```

### Claude Example

```go
import (
  "fmt"

  "github.com/poligono-xyz/alan"
  "github.com/poligono-xyz/alan/option"
)

func main() {
  client, err := alan.NewClient(
    option.WithProvider(alan.GeminiProvider),
    option.WithModel(alan.Gemini15Flash),
    option.WithAPIKey("api-key"),
  )
  if err != nil {
    panic(err)
  }

  result, err := client.Prompt("Where is the Dominican Republic located?")
  if err != nil {
    panic(err)
  }

  fmt.Println(result)
}

```

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
