package maps_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/maps"
)

func TestDict(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{dict}}`, ExpectedOutput: "map[]"},
		{Name: "TestWithEvenKeyPair", Input: `{{dict "a" 1 "b" 2}}`, ExpectedOutput: "map[a:1 b:2]"},
		{Name: "TestWithOddKeyPair", Input: `{{dict "a" 1 "b" 2 "c" 3 "d"}}`, ExpectedErr: "number of values must be even"},
		{Name: "TestWithANilKey", Input: `{{dict "a" 1 "b" 2 "c" 3 .Nil 4}}`, ExpectedOutput: "map[<nil>:4 a:1 b:2 c:3]", Data: map[string]any{"Nil": nil}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestGet(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{get "a"  . }}`, ExpectedOutput: ""},
		{Name: "TestWithKey", Input: `{{get "a"  . }}`, ExpectedOutput: "1", Data: map[string]any{"a": 1}},
		{Name: "TestWithNestedKeyNotFound", Input: `{{get "b"  . }}`, ExpectedOutput: "", Data: map[string]any{"a": 1}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestSet(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestWithKey", Input: `{{$d := set "a" 2 .}}{{$d}}`, ExpectedOutput: "map[a:2]", Data: map[string]any{"a": 1}},
		{Name: "TestWithNewKey", Input: `{{$d := set "b" 3 .}}{{$d}}`, ExpectedOutput: "map[a:1 b:3]", Data: map[string]any{"a": 1}},
		{Name: "TestWithNilValue", Input: `{{$d := .V | set "a" .Nil}}{{$d}}`, ExpectedOutput: "map[a:<nil>]", Data: map[string]any{"V": map[string]any{"a": 1}, "Nil": nil}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestUnset(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestWithKey", Input: `{{$d := unset "a" . }}{{$d}}`, ExpectedOutput: "map[]", Data: map[string]any{"a": 1}},
		{Name: "TestWithNestedKeyNotFound", Input: `{{$d := unset "b" . }}{{$d}}`, ExpectedOutput: "map[a:1]", Data: map[string]any{"a": 1}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestKeys(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{keys .}}`, ExpectedOutput: "[]"},
		{Name: "TestWithKeys", Input: `{{keys . | sortAlpha}}`, ExpectedOutput: `[a b]`, Data: map[string]any{"a": 1, "b": 2}},
		{Name: "TestWithMultiplesMaps", Input: `{{keys .A .B | sortAlpha}}`, ExpectedOutput: `[a b c d]`, Data: map[string]any{"A": map[string]any{"a": 1, "b": 2}, "B": map[string]any{"c": 3, "d": 4}}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestValues(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{values .}}`, ExpectedOutput: "[]"},
		{Name: "TestWithValues", Input: `{{values . | sortAlpha}}`, ExpectedOutput: "[1 foo]", Data: map[string]any{"a": 1, "b": "foo"}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestPluck(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{pluck "a" .}}`, ExpectedOutput: "[]"},
		{Name: "TestWithOneMap", Input: `{{. | pluck "a"}}`, ExpectedOutput: "[1]", Data: map[string]any{"a": 1, "b": 2}},
		{Name: "TestWithTwoMaps", Input: `{{pluck "a" .A .B }}`, ExpectedOutput: "[1 3]", Data: map[string]any{"A": map[string]any{"a": 1, "b": 2}, "B": map[string]any{"a": 3, "b": 4}}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestPick(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ . | pick "a"}}`, ExpectedOutput: "map[]"},
		{Name: "TestWithOneValue", Input: `{{ . | pick "a"}}`, ExpectedOutput: "map[a:1]", Data: map[string]any{"a": 1, "b": 2}},
		{Name: "TestWithTwoValues", Input: `{{ . | pick "a" "b"}}`, ExpectedOutput: "map[a:1 b:2]", Data: map[string]any{"a": 1, "b": 2}},
		{Name: "TestWithNestedKeyNotFound", Input: `{{ . | pick "nope"}}`, ExpectedOutput: "map[]", Data: map[string]any{"a": 1}},
		{Name: "TestWithNoMapGivenAsLastArg", Input: `{{ .a | pick "a"}}`, ExpectedErr: "last argument must be a map[string]any", Data: map[string]any{"a": 1}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestOmit(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ . | omit "a"}}`, ExpectedOutput: "map[]"},
		{Name: "TestWithOneValue", Input: `{{ . | omit "a"}}`, ExpectedOutput: "map[b:2]", Data: map[string]any{"a": 1, "b": 2}},
		{Name: "TestWithTwoValues", Input: `{{ . | omit "a" "b"}}`, ExpectedOutput: "map[]", Data: map[string]any{"a": 1, "b": 2}},
		{Name: "TestWithNestedKeyNotFound", Input: `{{ . | omit "nope"}}`, ExpectedOutput: "map[a:1]", Data: map[string]any{"a": 1}},
		{Name: "TestWithInvalidKeys", Input: `{{ . | omit "a" 1}}`, ExpectedErr: "all keys must be strings", Data: map[string]any{"a": 1}},
		{Name: "TestInvaidArgumentsType", Input: `{{ . | omit 1 }}`, ExpectedErr: "expected map or string, got int", Data: map[string]any{"a": 1}},
		{Name: "TestWithNoMapGivenAsLastArg", Input: `{{ .a | omit "a"}}`, ExpectedErr: "last argument must be a map[string]any", Data: map[string]any{"a": 1}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestDig(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{dig "a" .}}`, ExpectedOutput: "<no value>"},
		{Name: "TestWithOneValue", Input: `{{dig "a" .}}`, ExpectedOutput: "1", Data: map[string]any{"a": 1, "b": 2}},
		{Name: "TestWithNestedKey", Input: `{{dig "a" "b" .}}`, ExpectedOutput: "2", Data: map[string]any{"a": map[string]any{"b": 2}}},
		{Name: "TestWithNestedKeyNotFound", Input: `{{dig "b" .}}`, ExpectedOutput: "<no value>", Data: map[string]any{"a": 1}},
		{Name: "TestWithNotEnoughArgs", Input: `{{dig "a"}}`, ExpectedErr: "dig requires at least two arguments"},
		{Name: "TestWithInvalidKey", Input: `{{dig 1 .}}`, ExpectedErr: "all keys must be strings, got int at position 0"},
		{Name: "TestWithInvalidMap", Input: `{{dig "a" 1}}`, ExpectedErr: "last argument must be a map[string]any"},
		{Name: "TestToAccessNotMapNestedKey", Input: `{{dig "a" "b" .}}`, Data: map[string]any{"a": 1}, ExpectedErr: "value at key \"a\" is not a nested dictionary but int"},
		{Name: "TestWithDotSyntax", Input: `{{ dig "a.b" . }}`, ExpectedOutput: "2", Data: map[string]any{"a": map[string]any{"b": 2}}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestHasKey(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{. | hasKey "a"}}`, ExpectedOutput: "false"},
		{Name: "TestWithKey", Input: `{{. | hasKey "a"}}`, ExpectedOutput: "true", Data: map[string]any{"a": 1}},
		{Name: "TestWithNestedKeyNotFound", Input: `{{. | hasKey "b"}}`, ExpectedOutput: "false", Data: map[string]any{"a": 1}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestMerge(t *testing.T) {
	var dest map[string]any

	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{merge .}}`, ExpectedOutput: "map[]"},
		{Name: "TestWithOneMap", Input: `{{merge .}}`, ExpectedOutput: "map[a:1 b:2]", Data: map[string]any{"a": 1, "b": 2}},
		{Name: "TestWithTwoMaps", Input: `{{merge .Dest .Src1}}`, ExpectedOutput: "map[a:1 b:2 c:3 d:4]", Data: map[string]any{"Dest": map[string]any{"a": 1, "b": 2}, "Src1": map[string]any{"c": 3, "d": 4}}},
		{Name: "TestWithZeroValues", Input: `{{merge .Dest .Src1}}`, ExpectedOutput: "map[a:0 b:false c:3 d:4]", Data: map[string]any{"Dest": map[string]any{"a": 0, "b": false}, "Src1": map[string]any{"a": 2, "b": true, "c": 3, "d": 4}}},
		{Name: "TestWithNotEnoughArgs", Input: `{{merge .}}`, ExpectedOutput: "map[a:1]", Data: map[string]any{"a": 1}},
		{Name: "TestWithDestNonInitialized", Input: `{{merge .Dest .Src1}}`, ExpectedOutput: "map[b:2]", Data: map[string]any{"Dest": dest, "Src1": map[string]any{"b": 2}}},
		{Name: "TestWithDestNotMap", Input: `{{merge .Dest .Src1}}`, Data: map[string]any{"Dest": 1, "Src1": map[string]any{"b": 2}}, ExpectedErr: "wrong type for value"},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestMergeOverwrite(t *testing.T) {
	var dest map[string]any

	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{mergeOverwrite .}}`, ExpectedOutput: "map[]"},
		{Name: "TestWithOneMap", Input: `{{mergeOverwrite .}}`, ExpectedOutput: "map[a:1 b:2]", Data: map[string]any{"a": 1, "b": 2}},
		{Name: "TestWithTwoMaps", Input: `{{mergeOverwrite .Dest .Src1}}`, ExpectedOutput: "map[a:1 b:2 c:3 d:4]", Data: map[string]any{"Dest": map[string]any{"a": 1, "b": 2}, "Src1": map[string]any{"c": 3, "d": 4}}},
		{Name: "TestWithOverwrite", Input: `{{mergeOverwrite .Dest .Src1}}`, ExpectedOutput: "map[a:3 b:2 d:4]", Data: map[string]any{"Dest": map[string]any{"a": 1, "b": 2}, "Src1": map[string]any{"a": 3, "d": 4}}},
		{Name: "TestWithZeroValues", Input: `{{mergeOverwrite .Dest .Src1}}`, ExpectedOutput: "map[a:2 b:true c:3 d:4]", Data: map[string]any{"Dest": map[string]any{"a": 0, "b": false}, "Src1": map[string]any{"a": 2, "b": true, "c": 3, "d": 4}}},
		{Name: "TestWithNotEnoughArgs", Input: `{{mergeOverwrite .}}`, ExpectedOutput: "map[a:1]", Data: map[string]any{"a": 1}},
		{Name: "TestWithDestNonInitialized", Input: `{{mergeOverwrite .A .B}}`, ExpectedOutput: "map[b:2]", Data: map[string]any{"A": dest, "B": map[string]any{"b": 2}}},
		{Name: "TestWithDestNotMap", Input: `{{mergeOverwrite .A .B}}`, Data: map[string]any{"A": 1, "B": map[string]any{"b": 2}}, ExpectedErr: "wrong type for value"},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}
