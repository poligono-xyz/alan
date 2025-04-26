package option

import (
	"context"

	"github.com/poligono-xyz/alan"
)

func WithProvider(provider alan.Provider) alan.ClientOption {
	return func(config *alan.ClientConfig) {
		config.Provider = provider
	}
}

func WithModel(model alan.Model) alan.ClientOption {
	return func(config *alan.ClientConfig) {
		config.Model = model
	}
}
func WithAPIKey(apiKey string) alan.ClientOption {
	return func(config *alan.ClientConfig) {
		config.APIKey = apiKey
	}
}

func WithContext(ctx context.Context) alan.ClientOption {
	return func(config *alan.ClientConfig) {
		config.Context = ctx
	}
}
