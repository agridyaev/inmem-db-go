package compute

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
)

func TestParser(t *testing.T) {
	testCases := []struct {
		name           string
		query          string
		expectedError  error
		expectedTokens []string
	}{
		{name: "valid SET command", query: "SET a b",
			expectedError: nil, expectedTokens: []string{"SET", "a", "b"}},
		{name: "whitespace before SET command", query: " SET a b",
			expectedError: nil, expectedTokens: []string{"SET", "a", "b"}},
		{name: "whitespace after SET command", query: "SET a b ",
			expectedError: nil, expectedTokens: []string{"SET", "a", "b"}},
		{name: "newline after SET command", query: "SET a b\n",
			expectedError: nil, expectedTokens: []string{"SET", "a", "b"}},
		{name: "tab separator", query: "SET\ta\tb\t",
			expectedError: nil, expectedTokens: []string{"SET", "a", "b"}},
		{name: "new line separator", query: "SET\na\nb\n",
			expectedError: nil, expectedTokens: []string{"SET", "a", "b"}},
		{name: "invalid first argument SET", query: "SET % b",
			expectedError: errInvalidSymbol, expectedTokens: nil},
		{name: "invalid second argument SET", query: "SET a %",
			expectedError: errInvalidSymbol, expectedTokens: nil},
		{name: "valid GET command", query: "GET a",
			expectedError: nil, expectedTokens: []string{"GET", "a"}},
		{name: "invalid argument GET", query: "GET %",
			expectedError: errInvalidSymbol, expectedTokens: nil},
		{name: "valid DEL command", query: "DEL a",
			expectedError: nil, expectedTokens: []string{"DEL", "a"}},
		{name: "invalid argument DEL", query: "DEL %",
			expectedError: errInvalidSymbol, expectedTokens: nil},
		{name: "invalid command", query: "Б",
			expectedError: errInvalidSymbol, expectedTokens: nil},
		{name: "empty command", query: "",
			expectedError: nil, expectedTokens: nil},
		{name: "empty tokens command", query: " ",
			expectedError: nil, expectedTokens: nil},
	}
	ctx := context.WithValue(context.Background(), "tx", int64(555))
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			parser, err := NewParser(zap.NewNop())
			require.NoError(t, err)

			tokens, err := parser.ParseQuery(ctx, tc.query)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedTokens, tokens)
		})
	}
}
