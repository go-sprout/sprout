package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRandomString(t *testing.T) {
	rr := NewRegistry()

	var tc = []struct {
		opts         *randomOpts
		regexpString string
		length       int
	}{
		{&randomOpts{withLetters: true, withNumbers: true}, "^[a-zA-Z0-9]{100}$", 100},
		{&randomOpts{withLetters: true, withNumbers: true, withChars: []rune{'4', '2', '@'}}, "^[42@]{100}$", 100},
		{&randomOpts{withLetters: true}, "^[a-zA-Z]{100}$", 100},
		{&randomOpts{withNumbers: true}, "^[0-9]{100}$", 100},
		{&randomOpts{withAscii: true}, "^([a-zA-Z0-9]|[[:ascii:]]){100}$", 100},
	}

	for _, c := range tc {

		result, err := rr.randomString(c.length, c.opts)
		require.NoError(t, err)
		assert.Regexp(t, c.regexpString, result)
		assert.Len(t, result, c.length)
	}
}
