package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	rr := NewRegistry()
	assert.Regexp(t, "^[0-9]{100}$", rr.randomString(100, &randomOpts{withNumbers: true}))
	assert.Regexp(t, "^[a-zA-Z]{100}$", rr.randomString(100, &randomOpts{withLetters: true}))
	assert.Regexp(t, "^[a-zA-Z0-9]{100}$", rr.randomString(100, &randomOpts{withLetters: true, withNumbers: true}))
	assert.Regexp(t, "^([a-zA-Z0-9]|[[:ascii:]]){100}$", rr.randomString(100, &randomOpts{withLetters: true, withAscii: true}))
	assert.Regexp(t, "^[42@]{100}$", rr.randomString(100, &randomOpts{withChars: []rune{'4', '2', '@'}}))
}
