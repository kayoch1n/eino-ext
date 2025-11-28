package openai

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFunctionDefinition_Marshal_OmitEmptyParameters(t *testing.T) {
	fd := &functionDefinition{
		Name:        "weather",
		Description: "Get current weather",
	}
	data, err := json.Marshal(fd)
	require.NoError(t, err)
	assert.NotContains(t, string(data), "parameters")
}
