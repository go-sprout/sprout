package slices_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/slices"
)

func TestDeprecatedWithout(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ without .V "a" }}`, ExpectedOutput: "[b c d e]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ without .V "a" }}`, ExpectedOutput: "[b c d e]", Data: map[string]any{"V": []string{"b", "c", "d", "e"}}},
		{Input: `{{ without .V "a" }}`, ExpectedOutput: "[b c d e]", Data: map[string]any{"V": []string{"b", "c", "d", "e", "a"}}},
		{Input: `{{ without .V "a" }}`, ExpectedOutput: "[]", Data: map[string]any{"V": []string{"a"}}},
		{Input: `{{ without .V "a" }}`, ExpectedOutput: "[]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ without .V "a" }}`, Data: map[string]any{"V": nil}, ExpectedErr: "cannot without nil"},
		{Input: `{{ without .V "a" }}`, Data: map[string]any{"V": 1}, ExpectedErr: "last argument must be a slice but got string"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestDeprecatedAppend(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ append .V "a" }}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ append .V "a" }}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []string(nil)}},
		{Input: `{{ append .V "a" }}`, ExpectedOutput: "[x a]", Data: map[string]any{"V": []string{"x"}}},
		{Input: `{{ append .V "a" }}`, ExpectedOutput: "[x a]", Data: map[string]any{"V": [1]string{"x"}}},
		{Input: `{{ append .V "a" }}`, Data: map[string]any{"V": nil}, ExpectedErr: "cannot append to nil"},
		{Input: `{{ append .V "a" }}`, Data: map[string]any{"V": "1"}, ExpectedErr: "cannot append on type string"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestDeprecatedPrepend(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ prepend .V "a" }}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ prepend .V "a" }}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []string(nil)}},
		{Input: `{{ prepend .V "a" }}`, ExpectedOutput: "[a x]", Data: map[string]any{"V": []string{"x"}}},
		{Input: `{{ prepend .V "a" }}`, ExpectedOutput: "[a x]", Data: map[string]any{"V": [1]string{"x"}}},
		{Input: `{{ prepend .V "a" }}`, Data: map[string]any{"V": nil}, ExpectedErr: "cannot prepend to nil"},
		{Input: `{{ prepend .V "a" }}`, Data: map[string]any{"V": "1"}, ExpectedErr: "cannot prepend on type string"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestDeprecatedSlice(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ slice .V }}`, ExpectedOutput: "[a b c d e]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ slice .V 1 }}`, ExpectedOutput: "[b c d e]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ slice .V 1 3 }}`, ExpectedOutput: "[b c]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ slice .V 0 1 }}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ slice .V 0 1  }}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []string{"a"}}},
		{Input: `{{ slice .V 0 1 }}`, ExpectedOutput: "<no value>", Data: map[string]any{"V": []string{}}},
		{Input: `{{ slice .V 0 1 }}`, Data: map[string]any{"V": nil}, ExpectedErr: "cannot slice nil"},
		{Input: `{{ slice .V 0 1 }}`, Data: map[string]any{"V": 1}, ExpectedErr: "last argument must be a slice but got int"},
		{Input: `{{ slice .V -1 1 }}`, Data: map[string]any{"V": []string{"a"}}, ExpectedErr: "start index out of bounds"},
		{Input: `{{ slice .V 0 52 }}`, Data: map[string]any{"V": []string{"a"}}, ExpectedErr: "end index out of bounds"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}
