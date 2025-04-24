package option

import (
	"context"

	"github.com/poligono-xyz/alan"
)

func WithProvider(provider alan.Provider) alan.Option {
	return func(config *alan.Config) {
		config.Provider = provider
	}
}

func WithModel(model alan.Model) alan.Option {
	return func(config *alan.Config) {
		config.Model = model
	}
}
func WithAPIKey(apiKey string) alan.Option {
	return func(config *alan.Config) {
		config.APIKey = apiKey
	}
}

func WithContext(ctx context.Context) alan.Option {
	return func(config *alan.Config) {
		config.Context = ctx
	}
}
