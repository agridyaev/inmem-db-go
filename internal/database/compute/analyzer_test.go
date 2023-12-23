package compute

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
)

func TestAnalyzeQuery(t *testing.T) {
	testCases := []struct {
		name   string
		tokens []string
		query  Query
		err    error
	}{
		{
			name:   "valid GET command",
			tokens: []string{"GET", "key"},
			query:  NewQuery(GetCommandID, []string{"key"}),
			err:    nil,
		},
		{
			name:   "valid SET command",
			tokens: []string{"SET", "key", "value"},
			query:  NewQuery(SetCommandID, []string{"key", "value"}),
			err:    nil,
		},
		{
			name:   "valid DEL command",
			tokens: []string{"DEL", "key"},
			query:  NewQuery(DelCommandID, []string{"key"}),
			err:    nil,
		},
		{
			name:   "empty tokens",
			tokens: []string{},
			err:    errInvalidCommand,
		},
		{
			name:   "invalid command",
			tokens: []string{"TRUNCATE"},
			err:    errInvalidCommand,
		},
		{
			name:   "invalid number arguments for SET command",
			tokens: []string{"SET", "key"},
			err:    errInvalidArguments,
		},
		{
			name:   "invalid number arguments for GET command",
			tokens: []string{"GET", "key", "value"},
			err:    errInvalidArguments,
		},
		{
			name:   "invalid number arguments for DEL command",
			tokens: []string{"DEL", "key", "value"},
			err:    errInvalidArguments,
		},
	}

	ctx := context.WithValue(context.Background(), "tx", int64(555))
	analyzer, err := NewAnalyzer(zap.NewNop())
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			query, err := analyzer.AnalyzeQuery(ctx, tc.tokens)

			assert.Equal(t, tc.query, query)
			assert.Equal(t, tc.err, err)
		})
	}
}

func TestAnalyzeSetQuery(t *testing.T) {
	testCases := []struct {
		name  string
		query Query
		err   error
	}{
		{
			name:  "empty arguments",
			query: NewQuery(SetCommandID, []string{}),
			err:   errInvalidArguments,
		},
		{
			name:  "one argument",
			query: NewQuery(SetCommandID, []string{"key"}),
			err:   errInvalidArguments,
		},
		{
			name:  "valid SET command",
			query: NewQuery(SetCommandID, []string{"key", "value"}),
			err:   nil,
		},
		{
			name:  "three argumens",
			query: NewQuery(SetCommandID, []string{"key", "value", "wrong"}),
			err:   errInvalidArguments,
		},
	}

	ctx := context.WithValue(context.Background(), "tx", int64(555))
	analyzer, err := NewAnalyzer(zap.NewNop())
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := analyzer.analyzeSetQuery(ctx, tc.query)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestAnalyzeGetQuery(t *testing.T) {
	testCases := []struct {
		name  string
		query Query
		err   error
	}{
		{
			name:  "empty arguments",
			query: NewQuery(GetCommandID, []string{}),
			err:   errInvalidArguments,
		},
		{
			name:  "valid GET command",
			query: NewQuery(GetCommandID, []string{"key"}),
			err:   nil,
		},
		{
			name:  "two arguments",
			query: NewQuery(GetCommandID, []string{"key", "value"}),
			err:   errInvalidArguments,
		},
	}

	ctx := context.WithValue(context.Background(), "tx", int64(555))
	analyzer, err := NewAnalyzer(zap.NewNop())
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := analyzer.analyzeGetQuery(ctx, tc.query)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestAnalyzeDelQuery(t *testing.T) {
	testCases := []struct {
		name  string
		query Query
		err   error
	}{
		{
			name:  "empty arguments",
			query: NewQuery(DelCommandID, []string{}),
			err:   errInvalidArguments,
		},
		{
			name:  "valid GET command",
			query: NewQuery(DelCommandID, []string{"key"}),
			err:   nil,
		},
		{
			name:  "two arguments",
			query: NewQuery(DelCommandID, []string{"key", "value"}),
			err:   errInvalidArguments,
		},
	}

	ctx := context.WithValue(context.Background(), "tx", int64(555))
	analyzer, err := NewAnalyzer(zap.NewNop())
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := analyzer.analyzeDelQuery(ctx, tc.query)

			assert.Equal(t, tc.err, err)
		})
	}
}
