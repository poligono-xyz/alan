package alan

import (
	"testing"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFromAlanModelToClaudeChatModel(t *testing.T) {
	// Dummy client needed to call the method
	impl := &claudeImpl{}

	testCases := []struct {
		name        string
		alanModel   Model
		expected    anthropic.Model
		expectError bool
	}{
		{
			name:      "Claude 3.7 Sonnet Latest",
			alanModel: ModelClaude3_7SonnetLatest,
			expected:  anthropic.ModelClaude3_7SonnetLatest,
		},
		{
			name:        "Unsupported Model",
			alanModel:   "unknown-claude-model",
			expectError: true,
		},
		{
			name:        "GPT Model (should fail)",
			alanModel:   ChatModelGPT4o,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			claudeModel, err := impl.fromAlanModelToOpenAIChatModel(tc.alanModel) // Method name seems incorrect, should be toClaudeModel?

			if tc.expectError {
				assert.Error(t, err, "Expected an error for model %s", tc.alanModel)
				assert.Empty(t, claudeModel, "Expected empty model string on error")
			} else {
				require.NoError(t, err, "Did not expect an error for model %s", tc.alanModel)
				assert.Equal(t, tc.expected, claudeModel, "Model mapping mismatch for %s", tc.alanModel)
			}
		})
	}
}

// Add tests for claudeImpl.Prompt and claudeImpl.NewChat using mocks
// for the anthropic.Client interface if needed.
