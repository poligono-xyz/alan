package alan

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- Unit Tests for NewClient ---

func TestNewClient_Gemini(t *testing.T) {
	// Note: This test doesn't actually call the genai.NewClient
	// because it requires credentials or mocking the underlying HTTP transport.
	// We are primarily testing the logic within *our* NewClient function.
	// A more thorough test would involve mocking genai.NewClient itself.

	ctx := context.Background()
	apiKey := "test-gemini-key"
	model := Gemini15Flash

	client, err := NewClient(
		WithContext(ctx),
		WithProvider(GeminiProvider),
		WithAPIKey(apiKey),
		WithModel(model),
	)

	require.NoError(t, err, "NewClient should not return an error for Gemini")
	require.NotNil(t, client, "Client should not be nil")

	assert.Equal(t, GeminiProvider, client.GetProvider(), "Provider should be Gemini")
	assert.Equal(t, model, client.GetModel(), "Model should match input")

	// Check if the underlying type is geminiImpl (optional, but good for verification)
	_, ok := client.(*geminiImpl)
	assert.True(t, ok, "Client should be of type *geminiImpl")

	// TODO: Add mock for genai.NewClient to verify API key usage if needed.
}

func TestNewClient_ChatGPT(t *testing.T) {
	ctx := context.Background()
	apiKey := "test-chatgpt-key"
	model := ChatModelGPT4o

	client, err := NewClient(
		WithContext(ctx),
		WithProvider(ChatGPTProvider),
		WithAPIKey(apiKey),
		WithModel(model),
	)

	require.NoError(t, err, "NewClient should not return an error for ChatGPT")
	require.NotNil(t, client, "Client should not be nil")

	assert.Equal(t, ChatGPTProvider, client.GetProvider(), "Provider should be ChatGPT")
	assert.Equal(t, model, client.GetModel(), "Model should match input")

	// Check if the underlying type is chatgptImpl
	impl, ok := client.(*chatgptImpl)
	assert.True(t, ok, "Client should be of type *chatgptImpl")
	assert.NotNil(t, impl.client, "Internal openai client should be initialized")
	// Note: Verifying the API key requires inspecting the openai client's transport,
	// which is more involved. We trust the openai SDK handles this via options.
}

func TestNewClient_Claude(t *testing.T) {
	ctx := context.Background()
	apiKey := "test-claude-key"
	model := ModelClaude3_7SonnetLatest

	client, err := NewClient(
		WithContext(ctx),
		WithProvider(ClaudeProvider),
		WithAPIKey(apiKey),
		WithModel(model),
	)

	require.NoError(t, err, "NewClient should not return an error for Claude")
	require.NotNil(t, client, "Client should not be nil")

	assert.Equal(t, ClaudeProvider, client.GetProvider(), "Provider should be Claude")
	assert.Equal(t, model, client.GetModel(), "Model should match input")

	// Check if the underlying type is claudeImpl
	impl, ok := client.(*claudeImpl)
	assert.True(t, ok, "Client should be of type *claudeImpl")
	assert.NotNil(t, impl.client, "Internal anthropic client should be initialized")
	// Note: Verifying the API key requires inspecting the anthropic client's transport.
}

func TestNewClient_InvalidProvider(t *testing.T) {
	client, err := NewClient(
		WithProvider("unknown-provider"),
		WithAPIKey("some-key"),
		WithModel("some-model"),
	)

	// Expecting nil client and nil error based on current implementation
	// Consider returning an error for unknown providers for robustness.
	assert.Nil(t, err, "Error should be nil for unknown provider (current behavior)")
	assert.Nil(t, client, "Client should be nil for unknown provider")

	// If you change NewClient to return an error:
	// require.Error(t, err, "NewClient should return an error for unknown provider")
	// assert.Nil(t, client, "Client should be nil for unknown provider")
}

// func TestNewClient_Defaults(t *testing.T) {
// 	// Test default context if no context option is provided
// 	// This test implicitly runs when creating clients above,
// 	// but we can be explicit if needed.
// 	config := &ClientConfig{}
// 	// Apply options (or none to test defaults)
// 	assert.Equal(t, context.Background(), config.Context, "Default context should be context.Background()")

// 	// Test creating a client with minimal options (though APIKey/Provider are needed)
// 	_, err := NewClient(WithProvider(GeminiProvider), WithAPIKey("k")) // Need provider and key
// 	assert.NoError(t, err)
// }

// --- Helper Functions for Options ---

func WithContext(ctx context.Context) ClientOption {
	return func(c *ClientConfig) {
		c.Context = ctx
	}
}

func WithProvider(p Provider) ClientOption {
	return func(c *ClientConfig) {
		c.Provider = p
	}
}

func WithAPIKey(key string) ClientOption {
	return func(c *ClientConfig) {
		c.APIKey = key
	}
}

func WithModel(m Model) ClientOption {
	return func(c *ClientConfig) {
		c.Model = m
	}
}
