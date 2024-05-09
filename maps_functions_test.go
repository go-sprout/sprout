package sprout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDict(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{dict}}`, "map[]", nil},
		{"TestWithEvenKeyPair", `{{dict "a" 1 "b" 2}}`, "map[a:1 b:2]", nil},
		{"TestWithOddKeyPair", `{{dict "a" 1 "b" 2 "c" 3 "d"}}`, "map[a:1 b:2 c:3 d:]", nil},
		{"TestWithANilKey", `{{dict "a" 1 "b" 2 "c" 3 .Nil 4}}`, "map[<nil>:4 a:1 b:2 c:3]", map[string]any{"Nil": nil}},
	}

	runTestCases(t, tests)
}

func TestGet(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{get . "a"}}`, "", nil},
		{"TestWithKey", `{{get . "a"}}`, "1", map[string]any{"a": 1}},
		{"TestWithNestedKeyNotFound", `{{get . "b"}}`, "", map[string]any{"a": 1}},
	}

	runTestCases(t, tests)
}

func TestSet(t *testing.T) {
	var tests = testCases{
		{"TestWithKey", `{{$d := set . "a" 2}}{{$d}}`, "map[a:2]", map[string]any{"a": 1}},
		{"TestWithNewKey", `{{$d := set . "b" 3}}{{$d}}`, "map[a:1 b:3]", map[string]any{"a": 1}},
		{"TestWithNilValue", `{{$d := set .V "a" .Nil}}{{$d}}`, "map[a:<nil>]", map[string]any{"V": map[string]any{"a": 1}, "Nil": nil}},
	}

	runTestCases(t, tests)
}

func TestUnset(t *testing.T) {
	var tests = testCases{
		{"TestWithKey", `{{$d := unset . "a"}}{{$d}}`, "map[]", map[string]any{"a": 1}},
		{"TestWithNestedKeyNotFound", `{{$d := unset . "b"}}{{$d}}`, "map[a:1]", map[string]any{"a": 1}},
	}

	runTestCases(t, tests)
}

func TestKeys(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{keys .}}`, "[]", nil},
		{"TestWithKeys", `{{keys .}}`, `[a b]`, map[string]any{"a": 1, "b": 2}},
		{"TestWithMultiplesMaps", `{{keys .A .B}}`, `[a b c d]`, map[string]any{"A": map[string]any{"a": 1, "b": 2}, "B": map[string]any{"c": 3, "d": 4}}},
	}

	runTestCases(t, tests)
}

func TestValues(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{values .}}`, "[]", nil},
		{"TestWithValues", `{{values .}}`, "[1 foo]", map[string]any{"a": 1, "b": "foo"}},
	}

	runTestCases(t, tests)
}

func TestPluck(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{pluck "a" .}}`, "[]", nil},
		{"TestWithOneMap", `{{. | pluck "a"}}`, "[1]", map[string]any{"a": 1, "b": 2}},
		{"TestWithTwoMaps", `{{pluck "a" .A .B }}`, "[1 3]", map[string]any{"A": map[string]any{"a": 1, "b": 2}, "B": map[string]any{"a": 3, "b": 4}}},
	}

	runTestCases(t, tests)
}

func TestPick(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{pick . "a"}}`, "map[]", nil},
		{"TestWithOneValue", `{{pick . "a"}}`, "map[a:1]", map[string]any{"a": 1, "b": 2}},
		{"TestWithTwoValues", `{{pick . "a" "b"}}`, "map[a:1 b:2]", map[string]any{"a": 1, "b": 2}},
		{"TestWithNestedKeyNotFound", `{{pick . "nope"}}`, "map[]", map[string]any{"a": 1}},
	}

	runTestCases(t, tests)
}

func TestOmit(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{omit . "a"}}`, "map[]", nil},
		{"TestWithOneValue", `{{omit . "a"}}`, "map[b:2]", map[string]any{"a": 1, "b": 2}},
		{"TestWithTwoValues", `{{omit . "a" "b"}}`, "map[]", map[string]any{"a": 1, "b": 2}},
		{"TestWithNestedKeyNotFound", `{{omit . "nope"}}`, "map[a:1]", map[string]any{"a": 1}},
	}

	runTestCases(t, tests)
}

func TestDig(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"TestEmpty", `{{dig "a" .}}`, "<no value>", nil}, ""},
		{testCase{"TestWithOneValue", `{{dig "a" .}}`, "1", map[string]any{"a": 1, "b": 2}}, ""},
		{testCase{"TestWithNestedKey", `{{dig "a" "b" .}}`, "2", map[string]any{"a": map[string]any{"b": 2}}}, ""},
		{testCase{"TestWithNestedKeyNotFound", `{{dig "b" .}}`, "<no value>", map[string]any{"a": 1}}, ""},
		{testCase{"TestWithNotEnoughArgs", `{{dig "a"}}`, "", nil}, "dig requires at least two arguments"},
		{testCase{"TestWithInvalidKey", `{{dig 1 .}}`, "", nil}, "all keys must be strings, got int at position 0"},
		{testCase{"TestWithInvalidMap", `{{dig "a" 1}}`, "", nil}, "last argument must be a map[string]any"},
		{testCase{"TestToAccessNotMapNestedKey", `{{dig "a" "b" .}}`, "", map[string]any{"a": 1}}, "value at key \"a\" is not a nested dictionary but int"},
	}

	runMustTestCases(t, tests)
}

func TestDigIntoDictWithNoKeys(t *testing.T) {
	_, err := NewFunctionHandler().digIntoDict(map[string]any{}, []string{})
	assert.ErrorContains(t, err, "unexpected termination of key traversal")
}

func TestHasKey(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{hasKey . "a"}}`, "false", nil},
		{"TestWithKey", `{{hasKey . "a"}}`, "true", map[string]any{"a": 1}},
		{"TestWithNestedKeyNotFound", `{{hasKey . "b"}}`, "false", map[string]any{"a": 1}},
	}

	runTestCases(t, tests)
}

func TestMerge(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{merge .}}`, "map[]", nil},
		{"TestWithOneMap", `{{merge .}}`, "map[a:1 b:2]", map[string]any{"a": 1, "b": 2}},
		{"TestWithTwoMaps", `{{merge .A .B}}`, "map[a:1 b:2 c:3 d:4]", map[string]any{"A": map[string]any{"a": 1, "b": 2}, "B": map[string]any{"c": 3, "d": 4}}},
	}

	runTestCases(t, tests)
}

func TestMergeOverwrite(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{mergeOverwrite .}}`, "map[]", nil},
		{"TestWithOneMap", `{{mergeOverwrite .}}`, "map[a:1 b:2]", map[string]any{"a": 1, "b": 2}},
		{"TestWithTwoMaps", `{{mergeOverwrite .A .B}}`, "map[a:1 b:2 c:3 d:4]", map[string]any{"A": map[string]any{"a": 1, "b": 2}, "B": map[string]any{"c": 3, "d": 4}}},
		{"TestWithOverwrite", `{{mergeOverwrite .A .B}}`, "map[a:3 b:2 d:4]", map[string]any{"A": map[string]any{"a": 1, "b": 2}, "B": map[string]any{"a": 3, "d": 4}}},
	}

	runTestCases(t, tests)
}

func TestMustMerge(t *testing.T) {
	var dest map[string]any
	var tests = mustTestCases{
		{testCase{"TestWithNotEnoughArgs", `{{mustMerge .}}`, "map[a:1]", map[string]any{"a": 1}}, ""},
		{testCase{"TestWithDestNonInitialized", `{{mustMerge .A .B}}`, "map[b:2]", map[string]any{"A": dest, "B": map[string]any{"b": 2}}}, ""},
		{testCase{"TestWithDestNotMap", `{{mustMerge .A .B}}`, "", map[string]any{"A": 1, "B": map[string]any{"b": 2}}}, "wrong type for value"},
	}

	runMustTestCases(t, tests)
}

func TestMustMergeOverwrite(t *testing.T) {
	var dest map[string]any
	var tests = mustTestCases{
		{testCase{"TestWithNotEnoughArgs", `{{mustMergeOverwrite .}}`, "map[a:1]", map[string]any{"a": 1}}, ""},
		{testCase{"TestWithDestNonInitialized", `{{mustMergeOverwrite .A .B}}`, "map[b:2]", map[string]any{"A": dest, "B": map[string]any{"b": 2}}}, ""},
		{testCase{"TestWithDestNotMap", `{{mustMergeOverwrite .A .B}}`, "", map[string]any{"A": 1, "B": map[string]any{"b": 2}}}, "wrong type for value"},
	}

	runMustTestCases(t, tests)
}
