package sprout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHello asserts the Hello method returns the expected greeting.
func TestHello(t *testing.T) {
	handler := NewFunctionHandler()
	expected := "Hello!"

	assert.Equal(t, expected, handler.Hello())

	tmplResponse, err := runTemplate(t, handler, `{{hello}}`, nil)
	assert.NoError(t, err)
	assert.Equal(t, expected, tmplResponse)
}

func TestDefault(t *testing.T) {
	var tests = testCases{
		{"TestDefaultEmptyInput", `{{default "default" ""}}`, "default", nil},
		{"TestDefaultGivenInput", `{{default "default" "given"}}`, "given", nil},
		{"TestDefaultIntInput", `{{default "default" 42}}`, "42", nil},
		{"TestDefaultFloatInput", `{{default "default" 2.42}}`, "2.42", nil},
		{"TestDefaultTrueInput", `{{default "default" true}}`, "true", nil},
		{"TestDefaultFalseInput", `{{default "default" false}}`, "default", nil},
		{"TestDefaultNilInput", `{{default "default" nil}}`, "default", nil},
		{"TestDefaultNothingInput", `{{default "default" .Nothing}}`, "default", nil},
		{"TestDefaultMultipleNothingInput", `{{default "default" .Nothing}}`, "default", nil},
		{"TestDefaultMultipleArgument", `{{"first" | default "default" "second"}}`, "second", nil},
	}

	runTestCases(t, tests)
}

func TestEmpty(t *testing.T) {
	var tests = testCases{
		{"TestEmptyEmptyInput", `{{if empty ""}}1{{else}}0{{end}}`, "1", nil},
		{"TestEmptyGivenInput", `{{if empty "given"}}1{{else}}0{{end}}`, "0", nil},
		{"TestEmptyIntInput", `{{if empty 42}}1{{else}}0{{end}}`, "0", nil},
		{"TestEmptyUintInput", `{{if empty .i}}1{{else}}0{{end}}`, "0", map[string]any{"i": uint(42)}},
		{"TestEmptyComplexInput", `{{if empty .c}}1{{else}}0{{end}}`, "0", map[string]any{"c": complex(42, 42)}},
		{"TestEmptyFloatInput", `{{if empty 2.42}}1{{else}}0{{end}}`, "0", nil},
		{"TestEmptyTrueInput", `{{if empty true}}1{{else}}0{{end}}`, "0", nil},
		{"TestEmptyFalseInput", `{{if empty false}}1{{else}}0{{end}}`, "1", nil},
		{"TestEmptyStructInput", `{{if empty .s}}1{{else}}0{{end}}`, "0", map[string]any{"s": struct{}{}}},
		{"TestEmptyNilInput", `{{if empty nil}}1{{else}}0{{end}}`, "1", nil},
		{"TestEmptyNothingInput", `{{if empty .Nothing}}1{{else}}0{{end}}`, "1", nil},
		{"TestEmptyNestedInput", `{{if empty .top.NoSuchThing}}1{{else}}0{{end}}`, "1", map[string]any{"top": map[string]interface{}{}}},
		{"TestEmptyNestedNoDataInput", `{{if empty .bottom.NoSuchThing}}1{{else}}0{{end}}`, "1", nil},
		{"TestEmptyNimPointerInput", `{{if empty .nilPtr}}1{{else}}0{{end}}`, "1", map[string]any{"nilPtr": (*int)(nil)}},
	}

	runTestCases(t, tests)
}

func TestAll(t *testing.T) {
	var tests = testCases{
		{"TestAllEmptyInput", `{{if all ""}}1{{else}}0{{end}}`, "0", nil},
		{"TestAllGivenInput", `{{if all "given"}}1{{else}}0{{end}}`, "1", nil},
		{"TestAllIntInput", `{{if all 42 0 1}}1{{else}}0{{end}}`, "0", nil},
		{"TestAllVariableInput1", `{{ $two := 2 }}{{if all "" 0 nil $two }}1{{else}}0{{end}}`, "0", nil},
		{"TestAllVariableInput2", `{{ $two := 2 }}{{if all "" $two 0 0 0 }}1{{else}}0{{end}}`, "0", nil},
		{"TestAllVariableInput3", `{{ $two := 2 }}{{if all "" $two 3 4 5 }}1{{else}}0{{end}}`, "0", nil},
		{"TestAllNoInput", `{{if all }}1{{else}}0{{end}}`, "1", nil},
	}

	runTestCases(t, tests)
}

func TestAny(t *testing.T) {
	var tests = testCases{
		{"TestAnyEmptyInput", `{{if any ""}}1{{else}}0{{end}}`, "0", nil},
		{"TestAnyGivenInput", `{{if any "given"}}1{{else}}0{{end}}`, "1", nil},
		{"TestAnyIntInput", `{{if any 42 0 1}}1{{else}}0{{end}}`, "1", nil},
		{"TestAnyVariableInput1", `{{ $two := 2 }}{{if any "" 0 nil $two }}1{{else}}0{{end}}`, "1", nil},
		{"TestAnyVariableInput2", `{{ $two := 2 }}{{if any "" $two 3 4 5 }}1{{else}}0{{end}}`, "1", nil},
		{"TestAnyVariableInput3", `{{ $zero := 0 }}{{if any "" $zero 0 0 0 }}1{{else}}0{{end}}`, "0", nil},
		{"TestAnyNoInput", `{{if any }}1{{else}}0{{end}}`, "0", nil},
	}

	runTestCases(t, tests)
}

func TestCoalesce(t *testing.T) {
	var tests = testCases{
		{"TestCoalesceEmptyInput", `{{coalesce ""}}`, "<no value>", nil},
		{"TestCoalesceGivenInput", `{{coalesce "given"}}`, "given", nil},
		{"TestCoalesceIntInput", `{{ coalesce "" 0 nil 42 }}`, "42", nil},
		{"TestCoalesceVariableInput1", `{{ $two := 2 }}{{ coalesce "" 0 nil $two }}`, "2", nil},
		{"TestCoalesceVariableInput2", `{{ $two := 2 }}{{ coalesce "" $two 0 0 0 }}`, "2", nil},
		{"TestCoalesceVariableInput3", `{{ $two := 2 }}{{ coalesce "" $two 3 4 5 }}`, "2", nil},
		{"TestCoalesceNoInput", `{{ coalesce }}`, "<no value>", nil},
	}

	runTestCases(t, tests)
}

func TestTernary(t *testing.T) {
	var tests = testCases{
		{"", `{{true | ternary "foo" "bar"}}`, "foo", nil},
		{"", `{{ternary "foo" "bar" true}}`, "foo", nil},
		{"", `{{false | ternary "foo" "bar"}}`, "bar", nil},
		{"", `{{ternary "foo" "bar" false}}`, "bar", nil},
	}

	runTestCases(t, tests)
}

func TestUuidv4(t *testing.T) {
	handler := NewFunctionHandler()

	tmplResponse, err := runTemplate(t, handler, `{{uuidv4}}`, nil)
	assert.NoError(t, err)
	assert.Regexp(t, `^[\da-f]{8}-[\da-f]{4}-4[\da-f]{3}-[\da-f]{4}-[\da-f]{12}$`, tmplResponse)
}

func TestCat(t *testing.T) {
	var tests = testCases{
		{"TestCatEmptyInput", `{{cat ""}}`, "", nil},
		{"TestCatGivenInput", `{{cat "given"}}`, "given", nil},
		{"TestCatIntInput", `{{cat 42}}`, "42", nil},
		{"TestCatFloatInput", `{{cat 2.42}}`, "2.42", nil},
		{"TestCatTrueInput", `{{cat true}}`, "true", nil},
		{"TestCatFalseInput", `{{cat false}}`, "false", nil},
		{"TestCatNilInput", `{{cat nil}}`, "", nil},
		{"TestCatNothingInput", `{{cat .Nothing}}`, "", nil},
		{"TestCatMultipleInput", `{{cat "first" "second"}}`, "first second", nil},
		{"TestCatMultipleArgument", `{{"first" | cat "second"}}`, "second first", nil},
		{"TestCatVariableInput", `{{$b := "b"}}{{"c" | cat "a" $b}}`, "a b c", nil},
		{"TestCatDataInput", `{{.text | cat "a" "b"}}`, "a b cd", map[string]any{"text": "cd"}},
	}

	runTestCases(t, tests)
}

func TestUntil(t *testing.T) {
	var tests = testCases{
		{"", `{{range $i, $e := until 5}}({{$i}}{{$e}}){{end}}`, "(00)(11)(22)(33)(44)", nil},
		{"", `{{range $i, $e := until -5}}({{$i}}{{$e}}){{end}}`, "(00)(1-1)(2-2)(3-3)(4-4)", nil},
	}

	runTestCases(t, tests)
}

func TestUntilStep(t *testing.T) {
	var tests = testCases{
		{"", `{{range $i, $e := untilStep 0 5 1}}({{$i}}{{$e}}){{end}}`, "(00)(11)(22)(33)(44)", nil},
		{"", `{{range $i, $e := untilStep 3 6 1}}({{$i}}{{$e}}){{end}}`, "(03)(14)(25)", nil},
		{"", `{{range $i, $e := untilStep 0 -10 -2}}({{$i}}{{$e}}){{end}}`, "(00)(1-2)(2-4)(3-6)(4-8)", nil},
		{"", `{{range $i, $e := untilStep 3 0 1}}({{$i}}{{$e}}){{end}}`, "", nil},
		{"", `{{range $i, $e := untilStep 3 99 0}}({{$i}}{{$e}}){{end}}`, "", nil},
		{"", `{{range $i, $e := untilStep 3 99 -1}}({{$i}}{{$e}}){{end}}`, "", nil},
		{"", `{{range $i, $e := untilStep 3 0 0}}({{$i}}{{$e}}){{end}}`, "", nil},
	}

	runTestCases(t, tests)
}

func TestTypeIs(t *testing.T) {
	type testStruct struct{}

	var tests = testCases{
		{"TestTypeIsInt", `{{typeIs "int" 42}}`, "true", nil},
		{"TestTypeIsString", `{{42 | typeIs "string"}}`, "false", nil},
		{"TestTypeIsVariable", `{{$var := 42}}{{typeIs "string" $var}}`, "false", nil},
		{"TestTypeIsStruct", `{{.var | typeIs "*sprout.testStruct"}}`, "true", map[string]any{"var": &testStruct{}}},
	}

	runTestCases(t, tests)
}

func TestTypeIsLike(t *testing.T) {
	type testStruct struct{}

	var tests = testCases{
		{"TestTypeIsLikeInt", `{{typeIsLike "int" 42}}`, "true", nil},
		{"TestTypeIsLikeString", `{{42 | typeIsLike "string"}}`, "false", nil},
		{"TestTypeIsLikeVariable", `{{$var := 42}}{{typeIsLike "string" $var}}`, "false", nil},
		{"TestTypeIsLikeStruct", `{{.var | typeIsLike "*sprout.testStruct"}}`, "true", map[string]any{"var": &testStruct{}}},
		{"TestTypeIsLikeStructWithoutPointerMark", `{{.var | typeIsLike "sprout.testStruct"}}`, "true", map[string]any{"var": &testStruct{}}},
	}

	runTestCases(t, tests)
}

func TestTypeOf(t *testing.T) {
	type testStruct struct{}

	var tests = testCases{
		{"TestTypeOfInt", `{{typeOf 42}}`, "int", nil},
		{"TestTypeOfString", `{{typeOf "42"}}`, "string", nil},
		{"TestTypeOfVariable", `{{$var := 42}}{{typeOf $var}}`, "int", nil},
		{"TestTypeOfStruct", `{{typeOf .var}}`, "*sprout.testStruct", map[string]any{"var": &testStruct{}}},
	}

	runTestCases(t, tests)
}

func TestKindIs(t *testing.T) {
	type testStruct struct{}

	var tests = testCases{
		{"TestKindIsInt", `{{kindIs "int" 42}}`, "true", nil},
		{"TestKindIsString", `{{42 | kindIs "string"}}`, "false", nil},
		{"TestKindIsVariable", `{{$var := 42}}{{kindIs "string" $var}}`, "false", nil},
		{"TestKindIsStruct", `{{.var | kindIs "ptr"}}`, "true", map[string]any{"var": &testStruct{}}},
	}

	runTestCases(t, tests)
}

func TestKindOf(t *testing.T) {
	type testStruct struct{}

	var tests = testCases{
		{"TestKindOfInt", `{{kindOf 42}}`, "int", nil},
		{"TestKindOfString", `{{kindOf "42"}}`, "string", nil},
		{"TestKindOfSlice", `{{kindOf .var}}`, "slice", map[string]any{"var": []int{}}},
		{"TestKindOfVariable", `{{$var := 42}}{{kindOf $var}}`, "int", nil},
		{"TestKindOfStruct", `{{kindOf .var}}`, "ptr", map[string]any{"var": &testStruct{}}},
		{"TestKindOfStructWithoutPointerMark", `{{kindOf .var}}`, "struct", map[string]any{"var": testStruct{}}},
	}

	runTestCases(t, tests)
}

func TestDeepEqual(t *testing.T) {
	var tests = testCases{
		{"TestDeepEqualInt", `{{deepEqual 42 42}}`, "true", nil},
		{"TestDeepEqualString", `{{deepEqual "42" "42"}}`, "true", nil},
		{"TestDeepEqualSlice", `{{deepEqual .a .b}}`, "true", map[string]any{"a": []int{1, 2, 3}, "b": []int{1, 2, 3}}},
		{"TestDeepEqualMap", `{{deepEqual .a .b}}`, "true", map[string]any{"a": map[string]int{"a": 1, "b": 2}, "b": map[string]int{"a": 1, "b": 2}}},
		{"TestDeepEqualStruct", `{{deepEqual .a .b}}`, "true", map[string]any{"a": struct{ A int }{A: 1}, "b": struct{ A int }{A: 1}}},
		{"TestDeepEqualVariable", `{{$a := 42}}{{$b := 42}}{{deepEqual $a $b}}`, "true", nil},
		{"TestDeepEqualDifferent", `{{deepEqual 42 32}}`, "false", nil},
		{"TestDeepEqualDifferentType", `{{deepEqual 42 "42"}}`, "false", nil},
	}

	runTestCases(t, tests)
}

func TestDeepCopy(t *testing.T) {
	handler := NewFunctionHandler()

	type testStruct struct {
		A int
	}

	var tests = testCases{
		{"TestDeepCopyInt", `{{$a := 42}}{{$b := deepCopy $a}}{{$b}}`, "42", nil},
		{"TestDeepCopyString", `{{$a := "42"}}{{$b := deepCopy $a}}{{$b}}`, "42", nil},
		{"TestDeepCopySlice", `{{$a := .a}}{{$b := deepCopy $a}}{{$b}}`, "[1 2 3]", map[string]any{"a": []int{1, 2, 3}}},
		{"TestDeepCopyMap", `{{$a := .a}}{{$b := deepCopy $a}}{{$b}}`, `map[a:1 b:2]`, map[string]any{"a": map[string]int{"a": 1, "b": 2}}},
		{"TestDeepCopyStruct", `{{$a := .a}}{{$b := deepCopy $a}}{{$b}}`, "{1}", map[string]any{"a": testStruct{A: 1}}},
		{"TestDeepCopyVariable", `{{$a := 42}}{{$b := deepCopy $a}}{{$b}}`, "42", nil},
		{"TestDeepCopyDifferent", `{{$a := 42}}{{$b := deepCopy "42"}}{{$b}}`, "42", nil},
		{"TestDeepCopyDifferentType", `{{$a := 42}}{{$b := deepCopy 42.0}}{{$b}}`, "42", nil},
		{"TestDeepCopyNil", `{{$b := deepCopy .a}}`, "", map[string]any{"a": nil}},
		{"", `{{- $d := dict "a" 1 "b" 2 | deepCopy }}{{ values $d | sortAlpha | join "," }}`, "1,2", nil},
		{"", `{{- $d := dict "a" 1 "b" 2 | deepCopy }}{{ keys $d | sortAlpha | join "," }}`, "a,b", nil},
		{"", `{{- $one := dict "foo" (dict "bar" "baz") "qux" true -}}{{ deepCopy $one }}`, "map[foo:map[bar:baz] qux:true]", nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tmplResponse, err := runTemplate(t, handler, test.input, test.data)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, tmplResponse)

			if test.data != nil {
				assert.NotEqual(t, test.data["a"], test.expected)
			}
		})
	}
}

func TestFail(t *testing.T) {
	handler := NewFunctionHandler()

	tmplResponse, err := runTemplate(t, handler, `{{fail "This is an error"}}`, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "This is an error")
	assert.Empty(t, tmplResponse)
}
