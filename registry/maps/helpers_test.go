package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigIntoDictWithNoKeys(t *testing.T) {
	_, err := NewRegistry().digIntoDict(map[string]any{}, []string{})
	assert.ErrorContains(t, err, "unexpected termination of key traversal")
}
