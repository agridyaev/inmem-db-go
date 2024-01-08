package compute

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuery(t *testing.T) {
	query := NewQuery(GetCommandID, []string{"GET", "key"})
	assert.Equal(t, GetCommandID, query.CommandID())
	assert.Equal(t, []string{"GET", "key"}, query.Arguments())
}
