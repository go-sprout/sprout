package sprout

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

type randTestCase struct {
	name     string
	template string
	regexp   string
	length   int
}

func testRandHelper(t *testing.T, tcs []randTestCase) {
	t.Helper()

	for _, test := range tcs {
		t.Run(test.name, func(t *testing.T) {
			t.Helper()

			result, err := runTemplate(t, NewFunctionHandler(), test.template, nil)
			assert.NoError(t, err)

			assert.Regexp(t, test.regexp, result)
			assert.Len(t, result, test.length)
		})
	}
}

func TestRandAlphaNumeric(t *testing.T) {
	var tests = []randTestCase{
		{"TestRandAlphaNumWithNegativeValue", `{{ randAlphaNum -1 }}`, "", 0},
		{"TestRandAlphaNumWithZero", `{{ randAlphaNum 0 }}`, "", 0},
		{"TestRandAlphaNum", `{{ randAlphaNum 100 }}`, `^[a-zA-Z0-9]{100}$`, 100},
	}

	testRandHelper(t, tests)
}

func TestRandAlpha(t *testing.T) {
	var tests = []randTestCase{
		{"TestRandAlphaWithNegativeValue", `{{ randAlpha -1 }}`, "", 0},
		{"TestRandAlphaWithZero", `{{ randAlpha 0 }}`, "", 0},
		{"TestRandAlpha", `{{ randAlpha 100 }}`, `^[a-zA-Z]{100}$`, 100},
	}

	testRandHelper(t, tests)
}

func TestRandAscii(t *testing.T) {
	var tests = []randTestCase{
		{"TestRandAsciiWithNegativeValue", `{{ randAscii -1 }}`, "", 0},
		{"TestRandAsciiWithZero", `{{ randAscii 0 }}`, "", 0},
		{"TestRandAscii", `{{ randAscii 100 }}`, "^[[:ascii:]]{100}$", 100},
	}

	testRandHelper(t, tests)
}

func TestRandNumeric(t *testing.T) {
	var tests = []randTestCase{
		{"TestRandNumericWithNegativeValue", `{{ randNumeric -1 }}`, "", 0},
		{"TestRandNumericWithZero", `{{ randNumeric 0 }}`, "", 0},
		{"TestRandNumeric", `{{ randNumeric 100 }}`, `^[0-9]{100}$`, 100},
	}

	testRandHelper(t, tests)
}

func TestRandBytes(t *testing.T) {
	var tests = []randTestCase{
		{"TestRandBytesWithNegativeValue", `{{ randBytes -1 }}`, "", 0},
		{"TestRandBytesWithZero", `{{ randBytes 0 }}`, "", 0},
		{"TestRandBytes", `{{ randBytes 100 }}`, "", 100},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := runTemplate(t, NewFunctionHandler(), test.template, nil)
			assert.NoError(t, err)

			assert.Regexp(t, test.regexp, result)

			b, err := base64.StdEncoding.DecodeString(result)
			assert.NoError(t, err)
			assert.Len(t, b, test.length)
		})
	}
}

func TestRandomString(t *testing.T) {
	fh := NewFunctionHandler()
	assert.Regexp(t, "^[0-9]{100}$", fh.randomString(100, &randomOpts{withNumbers: true}))
	assert.Regexp(t, "^[a-zA-Z]{100}$", fh.randomString(100, &randomOpts{withLetters: true}))
	assert.Regexp(t, "^[a-zA-Z0-9]{100}$", fh.randomString(100, &randomOpts{withLetters: true, withNumbers: true}))
	assert.Regexp(t, "^([a-zA-Z0-9]|[[:ascii:]]){100}$", fh.randomString(100, &randomOpts{withLetters: true, withAscii: true}))
	assert.Regexp(t, "^[42@]{100}$", fh.randomString(100, &randomOpts{withChars: []rune{'4', '2', '@'}}))
}
