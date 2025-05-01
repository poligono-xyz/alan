package alan

import (
	"testing"

	"github.com/openai/openai-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFromAlanModelToOpenAIChatModel(t *testing.T) {
	// Dummy client needed to call the method
	impl := &chatgptImpl{}

	testCases := []struct {
		name        string
		alanModel   Model
		expected    openai.ChatModel
		expectError bool
	}{
		{
			name:      "GPT-4o",
			alanModel: ChatModelGPT4o,
			expected:  openai.ChatModelGPT4o,
		},
		{
			name:      "GPT-3.5 Turbo",
			alanModel: ChatModelGPT3_5Turbo,
			expected:  openai.ChatModelGPT3_5Turbo,
		},
		{
			name:      "GPT-4 Turbo Preview",
			alanModel: ChatModelGPT4TurboPreview,
			expected:  openai.ChatModelGPT4TurboPreview,
		},
		{
			name:        "Unsupported Model",
			alanModel:   "unknown-gpt-model",
			expectError: true,
		},
		{
			name:        "Gemini Model (should fail)",
			alanModel:   Gemini15Flash,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			openaiModel, err := impl.fromAlanModelToOpenAIChatModel(tc.alanModel)

			if tc.expectError {
				assert.Error(t, err, "Expected an error for model %s", tc.alanModel)
				assert.Empty(t, openaiModel, "Expected empty model string on error")
			} else {
				require.NoError(t, err, "Did not expect an error for model %s", tc.alanModel)
				assert.Equal(t, tc.expected, openaiModel, "Model mapping mismatch for %s", tc.alanModel)
			}
		})
	}
}

// Add tests for chatgptImpl.Prompt and chatgptImpl.NewChat using mocks
// for the openai.Client interface if needed.
