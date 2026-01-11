package sprigin_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/sprigin"
)

// TestSprigGet tests both Sprig (dict, key) and Sprout (key, dict) signatures.
func TestSprigGet(t *testing.T) {
	tc := []pesticide.TestCase{
		// Old Sprig signature: get(dict, key)
		{Name: "TestSprigEmpty", Input: `{{get . "a" }}`, ExpectedOutput: ""},
		{Name: "TestSprigWithKey", Input: `{{get . "a" }}`, ExpectedOutput: "1", Data: map[string]any{"a": 1}},
		{Name: "TestSprigWithNestedKeyNotFound", Input: `{{get . "b" }}`, ExpectedOutput: "", Data: map[string]any{"a": 1}},
		// New Sprout signature: get(key, dict) via piping
		{Name: "TestSproutWithPipe", Input: `{{. | get "a" }}`, ExpectedOutput: "1", Data: map[string]any{"a": 1}},
		{Name: "TestSproutWithPipeNotFound", Input: `{{. | get "b" }}`, ExpectedOutput: "", Data: map[string]any{"a": 1}},
		// Edge cases
		{Name: "TestGetDeepNested", Input: `{{get .deep "level1"}}`, ExpectedOutput: "map[level2:map[level3:value]]", Data: map[string]any{"deep": map[string]any{"level1": map[string]any{"level2": map[string]any{"level3": "value"}}}}},
		{Name: "TestGetDeepNestedPiped", Input: `{{.deep | get "level1"}}`, ExpectedOutput: "map[level2:map[level3:value]]", Data: map[string]any{"deep": map[string]any{"level1": map[string]any{"level2": map[string]any{"level3": "value"}}}}},
		{Name: "TestGetFromEmptyMap", Input: `{{get . "key"}}`, ExpectedOutput: "", Data: map[string]any{}},
		{Name: "TestGetNilValue", Input: `{{get . "nil"}}`, ExpectedOutput: "<no value>", Data: map[string]any{"nil": nil}},
		{Name: "TestGetSliceFromMap", Input: `{{get . "list"}}`, ExpectedOutput: "[a b c]", Data: map[string]any{"list": []string{"a", "b", "c"}}},
		{Name: "TestGetEmptyStringKey", Input: `{{get . ""}}`, ExpectedOutput: "empty", Data: map[string]any{"": "empty"}},
		{Name: "TestGetKeyWithSpaces", Input: `{{get . "key with spaces"}}`, ExpectedOutput: "value", Data: map[string]any{"key with spaces": "value"}},
		{Name: "TestGetFromNilMap", Input: `{{get .NilMap "key"}}`, ExpectedOutput: "<no value>", Data: map[string]any{"NilMap": nil}},
	}

	pesticide.RunTestCasesWithFuncs(t, sprigin.FuncMap(), tc)
}

// TestSprigSet tests both Sprig (dict, key, value) and Sprout (key, value, dict) signatures.
func TestSprigSet(t *testing.T) {
	tc := []pesticide.TestCase{
		// Old Sprig signature: set(dict, key, value)
		{Name: "TestSprigWithKey", Input: `{{$d := set . "a" 2}}{{$d}}`, ExpectedOutput: "map[a:2]", Data: map[string]any{"a": 1}},
		{Name: "TestSprigWithNewKey", Input: `{{$d := set . "b" 3}}{{$d}}`, ExpectedOutput: "map[a:1 b:3]", Data: map[string]any{"a": 1}},
		{Name: "TestSprigWithNilValue", Input: `{{$d := set .V "a" .Nil}}{{$d}}`, ExpectedOutput: "map[a:<nil>]", Data: map[string]any{"V": map[string]any{"a": 1}, "Nil": nil}},
		// New Sprout signature: set(key, value, dict) via piping
		{Name: "TestSproutWithPipe", Input: `{{$d := . | set "a" 2}}{{$d}}`, ExpectedOutput: "map[a:2]", Data: map[string]any{"a": 1}},
		{Name: "TestSproutWithPipeNewKey", Input: `{{$d := . | set "b" 3}}{{$d}}`, ExpectedOutput: "map[a:1 b:3]", Data: map[string]any{"a": 1}},
		// Edge cases
		{Name: "TestSetOnEmptyMap", Input: `{{$d := set . "key" "value"}}{{$d}}`, ExpectedOutput: "map[key:value]", Data: map[string]any{}},
		{Name: "TestSetNilValue", Input: `{{$d := set .M "key" .Nil}}{{$d}}`, ExpectedOutput: "map[key:<nil>]", Data: map[string]any{"M": map[string]any{}, "Nil": nil}},
		// Ambiguous case - both first and last are maps (default case)
		{Name: "TestSetAmbiguousBothMaps", Input: `{{$d := set .M1 "key" .M2}}{{$d}}`, ExpectedOutput: "map[a:1 key:map[b:2]]", Data: map[string]any{"M1": map[string]any{"a": 1}, "M2": map[string]any{"b": 2}}},
	}

	pesticide.RunTestCasesWithFuncs(t, sprigin.FuncMap(), tc)
}

// TestSprigUnset tests both Sprig (dict, key) and Sprout (key, dict) signatures.
func TestSprigUnset(t *testing.T) {
	tc := []pesticide.TestCase{
		// Old Sprig signature: unset(dict, key)
		{Name: "TestSprigWithKey", Input: `{{$d := unset . "a"}}{{$d}}`, ExpectedOutput: "map[]", Data: map[string]any{"a": 1}},
		{Name: "TestSprigWithNestedKeyNotFound", Input: `{{$d := unset . "b"}}{{$d}}`, ExpectedOutput: "map[a:1]", Data: map[string]any{"a": 1}},
		// New Sprout signature: unset(key, dict) via piping
		{Name: "TestSproutWithPipe", Input: `{{$d := . | unset "a"}}{{$d}}`, ExpectedOutput: "map[]", Data: map[string]any{"a": 1}},
		{Name: "TestSproutWithPipeNotFound", Input: `{{$d := . | unset "b"}}{{$d}}`, ExpectedOutput: "map[a:1]", Data: map[string]any{"a": 1}},
	}

	pesticide.RunTestCasesWithFuncs(t, sprigin.FuncMap(), tc)
}

// TestSprigHasKey tests both Sprig (dict, key) and Sprout (key, dict) signatures.
func TestSprigHasKey(t *testing.T) {
	tc := []pesticide.TestCase{
		// Old Sprig signature: hasKey(dict, key)
		{Name: "TestSprigEmpty", Input: `{{hasKey . "a"}}`, ExpectedOutput: "false"},
		{Name: "TestSprigWithKey", Input: `{{hasKey . "a"}}`, ExpectedOutput: "true", Data: map[string]any{"a": 1}},
		{Name: "TestSprigWithNestedKeyNotFound", Input: `{{hasKey . "b"}}`, ExpectedOutput: "false", Data: map[string]any{"a": 1}},
		// New Sprout signature: hasKey(key, dict) via piping
		{Name: "TestSproutWithPipe", Input: `{{. | hasKey "a"}}`, ExpectedOutput: "true", Data: map[string]any{"a": 1}},
		{Name: "TestSproutWithPipeNotFound", Input: `{{. | hasKey "b"}}`, ExpectedOutput: "false", Data: map[string]any{"a": 1}},
		// Edge cases
		{Name: "TestHasKeyOnEmptyMap", Input: `{{hasKey . "key"}}`, ExpectedOutput: "false", Data: map[string]any{}},
		{Name: "TestHasKeyWithNilValue", Input: `{{hasKey . "nil"}}`, ExpectedOutput: "true", Data: map[string]any{"nil": nil}},
	}

	pesticide.RunTestCasesWithFuncs(t, sprigin.FuncMap(), tc)
}

// TestSprigPick tests both Sprig (dict, keys...) and Sprout (keys..., dict) signatures.
func TestSprigPick(t *testing.T) {
	tc := []pesticide.TestCase{
		// Old Sprig signature: pick(dict, keys...)
		{Name: "TestSprigEmpty", Input: `{{pick . "a" "b"}}`, ExpectedOutput: "map[]"},
		{Name: "TestSprigWithKeys", Input: `{{pick . "a" "b"}}`, ExpectedOutput: "map[a:1 b:2]", Data: map[string]any{"a": 1, "b": 2}},
		{Name: "TestSprigWithNestedKeyNotFound", Input: `{{pick . "a" "b"}}`, ExpectedOutput: "map[a:1]", Data: map[string]any{"a": 1}},
		// New Sprout signature: pick(keys..., dict) via piping
		{Name: "TestSproutWithPipe", Input: `{{. | pick "a" "b"}}`, ExpectedOutput: "map[a:1 b:2]", Data: map[string]any{"a": 1, "b": 2}},
		{Name: "TestSproutWithPipeNotFound", Input: `{{. | pick "a" "b"}}`, ExpectedOutput: "map[a:1]", Data: map[string]any{"a": 1}},
		// Edge cases
		{Name: "TestPickFromEmptyMap", Input: `{{pick . "a" "b"}}`, ExpectedOutput: "map[]", Data: map[string]any{}},
	}

	pesticide.RunTestCasesWithFuncs(t, sprigin.FuncMap(), tc)
}

// TestSprigOmit tests both Sprig (dict, keys...) and Sprout (keys..., dict) signatures.
func TestSprigOmit(t *testing.T) {
	tc := []pesticide.TestCase{
		// Old Sprig signature: omit(dict, keys...)
		{Name: "TestSprigEmpty", Input: `{{omit . "a" "b"}}`, ExpectedOutput: "map[]"},
		{Name: "TestSprigWithKeys", Input: `{{omit . "a" "b"}}`, ExpectedOutput: "map[]", Data: map[string]any{"a": 1, "b": 2}},
		{Name: "TestSprigWithNestedKeyNotFound", Input: `{{omit . "b"}}`, ExpectedOutput: "map[a:1]", Data: map[string]any{"a": 1}},
		// New Sprout signature: omit(keys..., dict) via piping
		{Name: "TestSproutWithPipe", Input: `{{. | omit "a" "b"}}`, ExpectedOutput: "map[]", Data: map[string]any{"a": 1, "b": 2}},
		{Name: "TestSproutWithPipeNotFound", Input: `{{. | omit "b"}}`, ExpectedOutput: "map[a:1]", Data: map[string]any{"a": 1}},
		// Edge cases
		{Name: "TestOmitFromEmptyMap", Input: `{{omit . "a" "b"}}`, ExpectedOutput: "map[]", Data: map[string]any{}},
	}

	pesticide.RunTestCasesWithFuncs(t, sprigin.FuncMap(), tc)
}

// TestSprigDig tests the Sprig-compatible dig signature (keys + default + dict).
func TestSprigDig(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestSingleKeyFound", Input: `{{dig "a" "default" .}}`, ExpectedOutput: "found", Data: map[string]any{"a": "found"}},
		{Name: "TestSingleKeyNotFound", Input: `{{dig "b" "default" .}}`, ExpectedOutput: "default", Data: map[string]any{"a": "found"}},
		{Name: "TestNestedKeysFound", Input: `{{dig "a" "b" "default" .}}`, ExpectedOutput: "nested", Data: map[string]any{"a": map[string]any{"b": "nested"}}},
		{Name: "TestNestedKeysNotFound", Input: `{{dig "a" "c" "default" .}}`, ExpectedOutput: "default", Data: map[string]any{"a": map[string]any{"b": "nested"}}},
		{Name: "TestDeeplyNested", Input: `{{dig "a" "b" "c" "default" .}}`, ExpectedOutput: "deep", Data: map[string]any{"a": map[string]any{"b": map[string]any{"c": "deep"}}}},
		{Name: "TestDeeplyNestedNotFound", Input: `{{dig "a" "b" "x" "default" .}}`, ExpectedOutput: "default", Data: map[string]any{"a": map[string]any{"b": map[string]any{"c": "deep"}}}},
		{Name: "TestEmptyDict", Input: `{{dig "a" "default" .}}`, ExpectedOutput: "default", Data: map[string]any{}},
		{Name: "TestIntegerValue", Input: `{{dig "count" "0" .}}`, ExpectedOutput: "42", Data: map[string]any{"count": 42}},
		{Name: "TestKeysWithDots", Input: `{{dig "has.dot" "default" .}}`, ExpectedOutput: "value", Data: map[string]any{"has.dot": "value"}},
		// Test digging into a non-dict value mid-path (returns default due to error)
		{Name: "TestDigNonDictInPath", Input: `{{dig "a" "b" "default" .}}`, ExpectedOutput: "default", Data: map[string]any{"a": "notadict"}},
	}

	pesticide.RunTestCasesWithFuncs(t, sprigin.FuncMap(), tc)
}

// TestSprigAppend tests both Sprig (list, v) and Sprout (v, list) signatures.
func TestSprigAppend(t *testing.T) {
	tc := []pesticide.TestCase{
		// Old Sprig signature: append(list, value)
		{Name: "TestSprigAppend", Input: `{{append .V "a"}}`, ExpectedOutput: "[x a]", Data: map[string]any{"V": []string{"x"}}},
		{Name: "TestSprigAppendToEmpty", Input: `{{append .V "a"}}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []string{}}},
		// New Sprout signature: append(value, list) via piping
		{Name: "TestSproutWithPipe", Input: `{{.V | append "a"}}`, ExpectedOutput: "[x a]", Data: map[string]any{"V": []string{"x"}}},
		{Name: "TestSproutWithPipeToEmpty", Input: `{{.V | append "a"}}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []string{}}},
		// Edge cases
		{Name: "TestAppendToEmptySlice", Input: `{{append .V "a"}}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []any{}}},
		{Name: "TestAppendNilValue", Input: `{{append .V .Nil}}`, ExpectedOutput: "[a <nil>]", Data: map[string]any{"V": []any{"a"}, "Nil": nil}},
		{Name: "TestAppendNestedSlice", Input: `{{append .V .Nested}}`, ExpectedOutput: "[a [[b [c]]]]", Data: map[string]any{"V": []any{"a"}, "Nested": []any{[]any{"b", []any{"c"}}}}},
		{Name: "TestAppendMapToSlice", Input: `{{append .V .Map}}`, ExpectedOutput: "[a map[key:value]]", Data: map[string]any{"V": []any{"a"}, "Map": map[string]any{"key": "value"}}},
		{Name: "TestAppendMapToSlicePiped", Input: `{{.V | append .Map}}`, ExpectedOutput: "[a map[key:value]]", Data: map[string]any{"V": []any{"a"}, "Map": map[string]any{"key": "value"}}},
		{Name: "TestChainedAppends", Input: `{{.V | append "b" | append "c"}}`, ExpectedOutput: "[a b c]", Data: map[string]any{"V": []any{"a"}}},
		// Issue #162: Stack overflow when both args are slices
		{Name: "TestSliceToSlice#Issue162", Input: `{{append .List .Item}}`, ExpectedOutput: "[[a] [[b]]]", Data: map[string]any{"List": []any{[]string{"a"}}, "Item": []any{[]string{"b"}}}},
		{Name: "TestSliceToSlicePiped#Issue162", Input: `{{.List | append .Item}}`, ExpectedOutput: "[[b] [[a]]]", Data: map[string]any{"List": []any{[]string{"a"}}, "Item": []any{[]string{"b"}}}},
		{Name: "TestAppendSliceToSlice#Issue162", Input: `{{append .A .B}}`, ExpectedOutput: "[[1 2] [[3 4]]]", Data: map[string]any{"A": []any{[]int{1, 2}}, "B": []any{[]int{3, 4}}}},
		{Name: "TestDeeplyNestedSlice#Issue162", Input: `{{append .Deep "new"}}`, ExpectedOutput: "[[[[inner]]] new]", Data: map[string]any{"Deep": []any{[]any{[]any{[]string{"inner"}}}}}},
		{Name: "TestRapidAppends#Issue162", Input: `{{.V | append "1" | append "2" | append "3" | append "4" | append "5"}}`, ExpectedOutput: "[a 1 2 3 4 5]", Data: map[string]any{"V": []any{"a"}}},
		// Same type ambiguity tests
		{Name: "TestAppendTwoSlices#Issue162", Input: `{{append .List1 .List2}}`, ExpectedOutput: "[[a] [[b]]]", Data: map[string]any{"List1": []any{[]string{"a"}}, "List2": []any{[]string{"b"}}}},
		{Name: "TestAppendEmptyToEmpty#Issue162", Input: `{{append .E1 .E2}}`, ExpectedOutput: "[[]]", Data: map[string]any{"E1": []any{}, "E2": []any{}}},
		{Name: "TestAppendSliceOfSlices#Issue162", Input: `{{append .SoS1 .SoS2}}`, ExpectedOutput: "[[a] [b] [[c] [d]]]", Data: map[string]any{"SoS1": []any{[]string{"a"}, []string{"b"}}, "SoS2": []any{[]string{"c"}, []string{"d"}}}},
	}

	pesticide.RunTestCasesWithFuncs(t, sprigin.FuncMap(), tc)
}

// TestSprigPrepend tests both Sprig (list, v) and Sprout (v, list) signatures.
func TestSprigPrepend(t *testing.T) {
	tc := []pesticide.TestCase{
		// Old Sprig signature: prepend(list, value)
		{Name: "TestSprigPrepend", Input: `{{prepend .V "a"}}`, ExpectedOutput: "[a x]", Data: map[string]any{"V": []string{"x"}}},
		{Name: "TestSprigPrependToEmpty", Input: `{{prepend .V "a"}}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []string{}}},
		// New Sprout signature: prepend(value, list) via piping
		{Name: "TestSproutWithPipe", Input: `{{.V | prepend "a"}}`, ExpectedOutput: "[a x]", Data: map[string]any{"V": []string{"x"}}},
		{Name: "TestSproutWithPipeToEmpty", Input: `{{.V | prepend "a"}}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []string{}}},
		// Edge cases
		{Name: "TestPrependToEmptySlice", Input: `{{prepend .V "a"}}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []any{}}},
		{Name: "TestPrependNilValue", Input: `{{prepend .V .Nil}}`, ExpectedOutput: "[<nil> a]", Data: map[string]any{"V": []any{"a"}, "Nil": nil}},
		{Name: "TestChainedPrepends", Input: `{{.V | prepend "b" | prepend "c"}}`, ExpectedOutput: "[c b a]", Data: map[string]any{"V": []any{"a"}}},
		// Issue #162: Stack overflow when both args are slices
		{Name: "TestSliceToSlice#Issue162", Input: `{{prepend .List .Item}}`, ExpectedOutput: "[[[b]] [a]]", Data: map[string]any{"List": []any{[]string{"a"}}, "Item": []any{[]string{"b"}}}},
		{Name: "TestSliceToSlicePiped#Issue162", Input: `{{.List | prepend .Item}}`, ExpectedOutput: "[[[a]] [b]]", Data: map[string]any{"List": []any{[]string{"a"}}, "Item": []any{[]string{"b"}}}},
		{Name: "TestPrependTwoSlices#Issue162", Input: `{{prepend .List1 .List2}}`, ExpectedOutput: "[[[b]] [a]]", Data: map[string]any{"List1": []any{[]string{"a"}}, "List2": []any{[]string{"b"}}}},
	}

	pesticide.RunTestCasesWithFuncs(t, sprigin.FuncMap(), tc)
}

// TestSprigSlice tests both Sprig (list, indices...) and Sprout (indices..., list) signatures.
func TestSprigSlice(t *testing.T) {
	tc := []pesticide.TestCase{
		// Old Sprig signature: slice(list, indices...)
		{Name: "TestSprigSlice", Input: `{{slice .V 1 3}}`, ExpectedOutput: "[b c]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Name: "TestSprigSliceFromStart", Input: `{{slice .V 0 2}}`, ExpectedOutput: "[a b]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		// New Sprout signature: slice(indices..., list) via piping
		{Name: "TestSproutWithPipe", Input: `{{.V | slice 1 3}}`, ExpectedOutput: "[b c]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Name: "TestSproutWithPipeFromStart", Input: `{{.V | slice 0 2}}`, ExpectedOutput: "[a b]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		// Edge cases
		{Name: "TestSliceSingleElement", Input: `{{slice .V 0 1}}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []string{"a", "b", "c"}}},
		{Name: "TestSliceSingleElementPiped", Input: `{{.V | slice 0 1}}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []string{"a", "b", "c"}}},
		{Name: "TestSliceOnlyList", Input: `{{slice .V}}`, ExpectedOutput: "[a b c]", Data: map[string]any{"V": []string{"a", "b", "c"}}},
	}

	pesticide.RunTestCasesWithFuncs(t, sprigin.FuncMap(), tc)
}

// TestSprigWithout tests both Sprig (list, omit...) and Sprout (omit..., list) signatures.
func TestSprigWithout(t *testing.T) {
	tc := []pesticide.TestCase{
		// Old Sprig signature: without(list, omit...)
		{Name: "TestSprigWithout", Input: `{{without .V "a"}}`, ExpectedOutput: "[b c d e]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Name: "TestSprigWithoutMultiple", Input: `{{without .V "a" "c"}}`, ExpectedOutput: "[b d e]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		// New Sprout signature: without(omit..., list) via piping
		{Name: "TestSproutWithPipe", Input: `{{.V | without "a"}}`, ExpectedOutput: "[b c d e]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Name: "TestSproutWithPipeMultiple", Input: `{{.V | without "a" "c"}}`, ExpectedOutput: "[b d e]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		// Edge cases
		{Name: "TestWithoutFromEmptySlice", Input: `{{without .V "a"}}`, ExpectedOutput: "[]", Data: map[string]any{"V": []any{}}},
	}

	pesticide.RunTestCasesWithFuncs(t, sprigin.FuncMap(), tc)
}
