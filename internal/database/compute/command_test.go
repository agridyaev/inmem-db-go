package compute

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNameToID(t *testing.T) {
	testCases := []struct {
		name        string
		commandId   int
		commandName string
	}{
		{"set command", SetCommandID, "SET"},
		{"get command", GetCommandID, "GET"},
		{"del command", DelCommandID, "DEL"},
		{"unknown command", UnknownCommandID, "DROP"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.commandId, CommandNameToCommandID(tc.commandName))
		})
	}
}
