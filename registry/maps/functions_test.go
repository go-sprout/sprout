package maps_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/maps"
)

func TestDict(t *testing.T) {
	var tc = []pesticide.TestCase{
		{"TestEmpty", `{{dict}}`, "map[]", nil},
		{"TestWithEvenKeyPair", `{{dict "a" 1 "b" 2}}`, "map[a:1 b:2]", nil},
		{"TestWithOddKeyPair", `{{dict "a" 1 "b" 2 "c" 3 "d"}}`, "map[a:1 b:2 c:3 d:]", nil},
		{"TestWithANilKey", `{{dict "a" 1 "b" 2 "c" 3 .Nil 4}}`, "map[<nil>:4 a:1 b:2 c:3]", map[string]any{"Nil": nil}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestGet(t *testing.T) {
	var tc = []pesticide.TestCase{
		{"TestEmpty", `{{get . "a"}}`, "", nil},
		{"TestWithKey", `{{get . "a"}}`, "1", map[string]any{"a": 1}},
		{"TestWithNestedKeyNotFound", `{{get . "b"}}`, "", map[string]any{"a": 1}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestSet(t *testing.T) {
	var tc = []pesticide.TestCase{
		{"TestWithKey", `{{$d := set . "a" 2}}{{$d}}`, "map[a:2]", map[string]any{"a": 1}},
		{"TestWithNewKey", `{{$d := set . "b" 3}}{{$d}}`, "map[a:1 b:3]", map[string]any{"a": 1}},
		{"TestWithNilValue", `{{$d := set .V "a" .Nil}}{{$d}}`, "map[a:<nil>]", map[string]any{"V": map[string]any{"a": 1}, "Nil": nil}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestUnset(t *testing.T) {
	var tc = []pesticide.TestCase{
		{"TestWithKey", `{{$d := unset . "a"}}{{$d}}`, "map[]", map[string]any{"a": 1}},
		{"TestWithNestedKeyNotFound", `{{$d := unset . "b"}}{{$d}}`, "map[a:1]", map[string]any{"a": 1}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestKeys(t *testing.T) {
	var tc = []pesticide.TestCase{
		{"TestEmpty", `{{keys .}}`, "[]", nil},
		{"TestWithKeys", `{{keys . | sortAlpha}}`, `[a b]`, map[string]any{"a": 1, "b": 2}},
		{"TestWithMultiplesMaps", `{{keys .A .B | sortAlpha}}`, `[a b c d]`, map[string]any{"A": map[string]any{"a": 1, "b": 2}, "B": map[string]any{"c": 3, "d": 4}}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestValues(t *testing.T) {
	var tc = []pesticide.TestCase{
		{"TestEmpty", `{{values .}}`, "[]", nil},
		{"TestWithValues", `{{values . | sortAlpha}}`, "[1 foo]", map[string]any{"a": 1, "b": "foo"}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestPluck(t *testing.T) {
	var tc = []pesticide.TestCase{
		{"TestEmpty", `{{pluck "a" .}}`, "[]", nil},
		{"TestWithOneMap", `{{. | pluck "a"}}`, "[1]", map[string]any{"a": 1, "b": 2}},
		{"TestWithTwoMaps", `{{pluck "a" .A .B }}`, "[1 3]", map[string]any{"A": map[string]any{"a": 1, "b": 2}, "B": map[string]any{"a": 3, "b": 4}}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestPick(t *testing.T) {
	var tc = []pesticide.TestCase{
		{"TestEmpty", `{{pick . "a"}}`, "map[]", nil},
		{"TestWithOneValue", `{{pick . "a"}}`, "map[a:1]", map[string]any{"a": 1, "b": 2}},
		{"TestWithTwoValues", `{{pick . "a" "b"}}`, "map[a:1 b:2]", map[string]any{"a": 1, "b": 2}},
		{"TestWithNestedKeyNotFound", `{{pick . "nope"}}`, "map[]", map[string]any{"a": 1}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestOmit(t *testing.T) {
	var tc = []pesticide.TestCase{
		{"TestEmpty", `{{omit . "a"}}`, "map[]", nil},
		{"TestWithOneValue", `{{omit . "a"}}`, "map[b:2]", map[string]any{"a": 1, "b": 2}},
		{"TestWithTwoValues", `{{omit . "a" "b"}}`, "map[]", map[string]any{"a": 1, "b": 2}},
		{"TestWithNestedKeyNotFound", `{{omit . "nope"}}`, "map[a:1]", map[string]any{"a": 1}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestDig(t *testing.T) {
	var tc = []pesticide.MustTestCase{
		{pesticide.TestCase{"TestEmpty", `{{dig "a" .}}`, "<no value>", nil}, ""},
		{pesticide.TestCase{"TestWithOneValue", `{{dig "a" .}}`, "1", map[string]any{"a": 1, "b": 2}}, ""},
		{pesticide.TestCase{"TestWithNestedKey", `{{dig "a" "b" .}}`, "2", map[string]any{"a": map[string]any{"b": 2}}}, ""},
		{pesticide.TestCase{"TestWithNestedKeyNotFound", `{{dig "b" .}}`, "<no value>", map[string]any{"a": 1}}, ""},
		{pesticide.TestCase{"TestWithNotEnoughArgs", `{{dig "a"}}`, "", nil}, "dig requires at least two arguments"},
		{pesticide.TestCase{"TestWithInvalidKey", `{{dig 1 .}}`, "", nil}, "all keys must be strings, got int at position 0"},
		{pesticide.TestCase{"TestWithInvalidMap", `{{dig "a" 1}}`, "", nil}, "last argument must be a map[string]any"},
		{pesticide.TestCase{"TestToAccessNotMapNestedKey", `{{dig "a" "b" .}}`, "", map[string]any{"a": 1}}, "value at key \"a\" is not a nested dictionary but int"},
	}

	pesticide.RunMustTestCases(t, maps.NewRegistry(), tc)
}

func TestHasKey(t *testing.T) {
	var tc = []pesticide.TestCase{
		{"TestEmpty", `{{hasKey . "a"}}`, "false", nil},
		{"TestWithKey", `{{hasKey . "a"}}`, "true", map[string]any{"a": 1}},
		{"TestWithNestedKeyNotFound", `{{hasKey . "b"}}`, "false", map[string]any{"a": 1}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestMerge(t *testing.T) {
	var tc = []pesticide.TestCase{
		{"TestEmpty", `{{merge .}}`, "map[]", nil},
		{"TestWithOneMap", `{{merge .}}`, "map[a:1 b:2]", map[string]any{"a": 1, "b": 2}},
		{"TestWithTwoMaps", `{{merge .Dest .Src1}}`, "map[a:1 b:2 c:3 d:4]", map[string]any{"Dest": map[string]any{"a": 1, "b": 2}, "Src1": map[string]any{"c": 3, "d": 4}}},
		{"TestWithZeroValues", `{{merge .Dest .Src1}}`, "map[a:0 b:false c:3 d:4]", map[string]any{"Dest": map[string]any{"a": 0, "b": false}, "Src1": map[string]any{"a": 2, "b": true, "c": 3, "d": 4}}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestMergeOverwrite(t *testing.T) {
	var tc = []pesticide.TestCase{
		{"TestEmpty", `{{mergeOverwrite .}}`, "map[]", nil},
		{"TestWithOneMap", `{{mergeOverwrite .}}`, "map[a:1 b:2]", map[string]any{"a": 1, "b": 2}},
		{"TestWithTwoMaps", `{{mergeOverwrite .Dest .Src1}}`, "map[a:1 b:2 c:3 d:4]", map[string]any{"Dest": map[string]any{"a": 1, "b": 2}, "Src1": map[string]any{"c": 3, "d": 4}}},
		{"TestWithOverwrite", `{{mergeOverwrite .Dest .Src1}}`, "map[a:3 b:2 d:4]", map[string]any{"Dest": map[string]any{"a": 1, "b": 2}, "Src1": map[string]any{"a": 3, "d": 4}}},
		{"TestWithZeroValues", `{{mergeOverwrite .Dest .Src1}}`, "map[a:2 b:true c:3 d:4]", map[string]any{"Dest": map[string]any{"a": 0, "b": false}, "Src1": map[string]any{"a": 2, "b": true, "c": 3, "d": 4}}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestMustMerge(t *testing.T) {
	var dest map[string]any
	var tc = []pesticide.MustTestCase{
		{pesticide.TestCase{"TestWithNotEnoughArgs", `{{mustMerge .}}`, "map[a:1]", map[string]any{"a": 1}}, ""},
		{pesticide.TestCase{"TestWithDestNonInitialized", `{{mustMerge .Dest .Src1}}`, "map[b:2]", map[string]any{"Dest": dest, "Src1": map[string]any{"b": 2}}}, ""},
		{pesticide.TestCase{"TestWithDestNotMap", `{{mustMerge .Dest .Src1}}`, "", map[string]any{"Dest": 1, "Src1": map[string]any{"b": 2}}}, "wrong type for value"},
	}

	pesticide.RunMustTestCases(t, maps.NewRegistry(), tc)
}

func TestMustMergeOverwrite(t *testing.T) {
	var dest map[string]any
	var tc = []pesticide.MustTestCase{
		{pesticide.TestCase{"TestWithNotEnoughArgs", `{{mustMergeOverwrite .}}`, "map[a:1]", map[string]any{"a": 1}}, ""},
		{pesticide.TestCase{"TestWithDestNonInitialized", `{{mustMergeOverwrite .A .B}}`, "map[b:2]", map[string]any{"A": dest, "B": map[string]any{"b": 2}}}, ""},
		{pesticide.TestCase{"TestWithDestNotMap", `{{mustMergeOverwrite .A .B}}`, "", map[string]any{"A": 1, "B": map[string]any{"b": 2}}}, "wrong type for value"},
	}

	pesticide.RunMustTestCases(t, maps.NewRegistry(), tc)
}
