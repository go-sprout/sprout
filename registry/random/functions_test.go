package random_test

import (
	"encoding/base64"
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/random"
	"github.com/stretchr/testify/assert"
)

func TestRandAlphaNumeric(t *testing.T) {
	var tc = []pesticide.RegexpTestCase{
		{Name: "TestRandAlphaNumWithNegativeValue", Template: `{{ randAlphaNum -1 }}`, Regexp: "", Length: 0},
		{Name: "TestRandAlphaNumWithZero", Template: `{{ randAlphaNum 0 }}`, Regexp: "", Length: 0},
		{Name: "TestRandAlphaNum", Template: `{{ randAlphaNum 100 }}`, Regexp: `^[a-zA-Z0-9]{100}$`, Length: 100},
	}

	pesticide.RunRegexpTestCases(t, random.NewRegistry(), tc)
}

func TestRandAlpha(t *testing.T) {
	var tc = []pesticide.RegexpTestCase{
		{Name: "TestRandAlphaWithNegativeValue", Template: `{{ randAlpha -1 }}`, Regexp: "", Length: 0},
		{Name: "TestRandAlphaWithZero", Template: `{{ randAlpha 0 }}`, Regexp: "", Length: 0},
		{Name: "TestRandAlpha", Template: `{{ randAlpha 100 }}`, Regexp: `^[a-zA-Z]{100}$`, Length: 100},
	}

	pesticide.RunRegexpTestCases(t, random.NewRegistry(), tc)
}

func TestRandAscii(t *testing.T) {
	var tc = []pesticide.RegexpTestCase{
		{Name: "TestRandAsciiWithNegativeValue", Template: `{{ randAscii -1 }}`, Regexp: "", Length: 0},
		{Name: "TestRandAsciiWithZero", Template: `{{ randAscii 0 }}`, Regexp: "", Length: 0},
		{Name: "TestRandAscii", Template: `{{ randAscii 100 }}`, Regexp: "^[[:ascii:]]{100}$", Length: 100},
	}

	pesticide.RunRegexpTestCases(t, random.NewRegistry(), tc)
}

func TestRandNumeric(t *testing.T) {
	var tc = []pesticide.RegexpTestCase{
		{Name: "TestRandNumericWithNegativeValue", Template: `{{ randNumeric -1 }}`, Regexp: "", Length: 0},
		{Name: "TestRandNumericWithZero", Template: `{{ randNumeric 0 }}`, Regexp: "", Length: 0},
		{Name: "TestRandNumeric", Template: `{{ randNumeric 100 }}`, Regexp: `^[0-9]{100}$`, Length: 100},
	}

	pesticide.RunRegexpTestCases(t, random.NewRegistry(), tc)
}

func TestRandBytes(t *testing.T) {
	var tests = []pesticide.RegexpTestCase{
		{Name: "TestRandBytesWithNegativeValue", Template: `{{ randBytes -1 }}`, Regexp: "", Length: 0},
		{Name: "TestRandBytesWithZero", Template: `{{ randBytes 0 }}`, Regexp: "", Length: 0},
		{Name: "TestRandBytes", Template: `{{ randBytes 100 }}`, Regexp: "", Length: 100},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result, err := pesticide.TestTemplate(t, random.NewRegistry(), test.Template, nil)
			assert.NoError(t, err)

			assert.Regexp(t, test.Regexp, result)

			b, err := base64.StdEncoding.DecodeString(result)
			assert.NoError(t, err)
			assert.Len(t, b, test.Length)
		})
	}
}

func TestRandInt(t *testing.T) {
	var tc = []pesticide.RegexpTestCase{
		{Name: "TestRandIntWithNegativeValue", Template: `{{ randInt -1 10 }}`, Regexp: "", Length: -1},
		{Name: "BetweenZeroAndTen", Template: `{{ randInt 0 10 }}`, Regexp: `^[0-9]{1,2}$`, Length: -1},
		{Name: "BetweenTenAndTwenty", Template: `{{ randInt 10 20 }}`, Regexp: `^[0-9]{1,2}$`, Length: -1},
		{Name: "BetweenNegativeTenAndTwenty", Template: `{{ randInt -10 20 }}`, Regexp: `^-?[0-9]{1,2}$`, Length: -1},
	}

	pesticide.RunRegexpTestCases(t, random.NewRegistry(), tc)
}
