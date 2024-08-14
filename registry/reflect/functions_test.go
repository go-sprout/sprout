package reflect_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/reflect"
	"github.com/stretchr/testify/assert"
)

func TestTypeIs(t *testing.T) {
	type testStruct struct{}

	var tc = []pesticide.TestCase{
		{Name: "TestTypeIsInt", Input: `{{typeIs "int" 42}}`, Expected: "true"},
		{Name: "TestTypeIsString", Input: `{{42 | typeIs "string"}}`, Expected: "false"},
		{Name: "TestTypeIsVariable", Input: `{{$var := 42}}{{typeIs "string" $var}}`, Expected: "false"},
		{Name: "TestTypeIsStruct", Input: `{{.var | typeIs "*reflect_test.testStruct"}}`, Expected: "true", Data: map[string]any{"var": &testStruct{}}},
	}

	pesticide.RunTestCases(t, reflect.NewRegistry(), tc)
}

func TestTypeIsLike(t *testing.T) {
	type testStruct struct{}

	var tc = []pesticide.TestCase{
		{Name: "TestTypeIsLikeInt", Input: `{{typeIsLike "int" 42}}`, Expected: "true"},
		{Name: "TestTypeIsLikeString", Input: `{{42 | typeIsLike "string"}}`, Expected: "false"},
		{Name: "TestTypeIsLikeVariable", Input: `{{$var := 42}}{{typeIsLike "string" $var}}`, Expected: "false"},
		{Name: "TestTypeIsLikeStruct", Input: `{{.var | typeIsLike "*reflect_test.testStruct"}}`, Expected: "true", Data: map[string]any{"var": &testStruct{}}},
		{Name: "TestTypeIsLikeStructWithoutPointerMark", Input: `{{.var | typeIsLike "reflect_test.testStruct"}}`, Expected: "true", Data: map[string]any{"var": &testStruct{}}},
	}

	pesticide.RunTestCases(t, reflect.NewRegistry(), tc)
}

func TestTypeOf(t *testing.T) {
	type testStruct struct{}

	var tc = []pesticide.TestCase{
		{Name: "TestTypeOfInt", Input: `{{typeOf 42}}`, Expected: "int"},
		{Name: "TestTypeOfString", Input: `{{typeOf "42"}}`, Expected: "string"},
		{Name: "TestTypeOfVariable", Input: `{{$var := 42}}{{typeOf $var}}`, Expected: "int"},
		{Name: "TestTypeOfStruct", Input: `{{typeOf .var}}`, Expected: "*reflect_test.testStruct", Data: map[string]any{"var": &testStruct{}}},
	}

	pesticide.RunTestCases(t, reflect.NewRegistry(), tc)
}

func TestKindIs(t *testing.T) {
	type testStruct struct{}

	var tc = []pesticide.TestCase{
		{Name: "TestKindIsInt", Input: `{{kindIs "int" 42}}`, Expected: "true"},
		{Name: "TestKindIsString", Input: `{{42 | kindIs "string"}}`, Expected: "false"},
		{Name: "TestKindIsVariable", Input: `{{$var := 42}}{{kindIs "string" $var}}`, Expected: "false"},
		{Name: "TestKindIsStruct", Input: `{{.var | kindIs "ptr"}}`, Expected: "true", Data: map[string]any{"var": &testStruct{}}},
	}

	pesticide.RunTestCases(t, reflect.NewRegistry(), tc)
}

func TestKindOf(t *testing.T) {
	type testStruct struct{}

	var tc = []pesticide.TestCase{
		{Name: "TestKindOfInt", Input: `{{kindOf 42}}`, Expected: "int"},
		{Name: "TestKindOfString", Input: `{{kindOf "42"}}`, Expected: "string"},
		{Name: "TestKindOfSlice", Input: `{{kindOf .var}}`, Expected: "slice", Data: map[string]any{"var": []int{}}},
		{Name: "TestKindOfVariable", Input: `{{$var := 42}}{{kindOf $var}}`, Expected: "int"},
		{Name: "TestKindOfStruct", Input: `{{kindOf .var}}`, Expected: "ptr", Data: map[string]any{"var": &testStruct{}}},
		{Name: "TestKindOfStructWithoutPointerMark", Input: `{{kindOf .var}}`, Expected: "struct", Data: map[string]any{"var": testStruct{}}},
	}

	pesticide.RunTestCases(t, reflect.NewRegistry(), tc)
}

func TestDeepEqual(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestDeepEqualInt", Input: `{{deepEqual 42 42}}`, Expected: "true"},
		{Name: "TestDeepEqualString", Input: `{{deepEqual "42" "42"}}`, Expected: "true"},
		{Name: "TestDeepEqualSlice", Input: `{{deepEqual .a .b}}`, Expected: "true", Data: map[string]any{"a": []int{1, 2, 3}, "b": []int{1, 2, 3}}},
		{Name: "TestDeepEqualMap", Input: `{{deepEqual .a .b}}`, Expected: "true", Data: map[string]any{"a": map[string]int{"a": 1, "b": 2}, "b": map[string]int{"a": 1, "b": 2}}},
		{Name: "TestDeepEqualStruct", Input: `{{deepEqual .a .b}}`, Expected: "true", Data: map[string]any{"a": struct{ A int }{A: 1}, "b": struct{ A int }{A: 1}}},
		{Name: "TestDeepEqualVariable", Input: `{{$a := 42}}{{$b := 42}}{{deepEqual $a $b}}`, Expected: "true"},
		{Name: "TestDeepEqualDifferent", Input: `{{deepEqual 42 32}}`, Expected: "false"},
		{Name: "TestDeepEqualDifferentType", Input: `{{deepEqual 42 "42"}}`, Expected: "false"},
	}

	pesticide.RunTestCases(t, reflect.NewRegistry(), tc)
}

func TestDeepCopy(t *testing.T) {
	type testStruct struct {
		A int
	}

	var tc = []pesticide.TestCase{
		{Name: "TestDeepCopyInt", Input: `{{$a := 42}}{{$b := deepCopy $a}}{{$b}}`, Expected: "42"},
		{Name: "TestDeepCopyString", Input: `{{$a := "42"}}{{$b := deepCopy $a}}{{$b}}`, Expected: "42"},
		{Name: "TestDeepCopySlice", Input: `{{$a := .a}}{{$b := deepCopy $a}}{{$b}}`, Expected: "[1 2 3]", Data: map[string]any{"a": []int{1, 2, 3}}},
		{Name: "TestDeepCopyMap", Input: `{{$a := .a}}{{$b := deepCopy $a}}{{$b}}`, Expected: `map[a:1 b:2]`, Data: map[string]any{"a": map[string]int{"a": 1, "b": 2}}},
		{Name: "TestDeepCopyStruct", Input: `{{$a := .a}}{{$b := deepCopy $a}}{{$b}}`, Expected: "{1}", Data: map[string]any{"a": testStruct{A: 1}}},
		{Name: "TestDeepCopyVariable", Input: `{{$a := 42}}{{$b := deepCopy $a}}{{$b}}`, Expected: "42"},
		{Name: "TestDeepCopyDifferent", Input: `{{$a := 42}}{{$b := deepCopy "42"}}{{$b}}`, Expected: "42"},
		{Name: "TestDeepCopyDifferentType", Input: `{{$a := 42}}{{$b := deepCopy 42.0}}{{$b}}`, Expected: "42"},
		{Name: "TestDeepCopyNil", Input: `{{$b := deepCopy .a}}`, Expected: "", Data: map[string]any{"a": nil}},
		{Input: `{{- $d := dict "a" 1 "b" 2 | deepCopy }}{{ values $d | sortAlpha | join "," }}`, Expected: "1,2"},
		{Input: `{{- $d := dict "a" 1 "b" 2 | deepCopy }}{{ keys $d | sortAlpha | join "," }}`, Expected: "a,b"},
		{Input: `{{- $one := dict "foo" (dict "bar" "baz") "qux" true -}}{{ deepCopy $one }}`, Expected: "map[foo:map[bar:baz] qux:true]"},
	}

	for _, test := range tc {
		t.Run(test.Name, func(t *testing.T) {
			tmplResponse, err := pesticide.TestTemplate(t, reflect.NewRegistry(), test.Input, test.Data)
			assert.NoError(t, err)
			assert.Equal(t, test.Expected, tmplResponse)

			if test.Data != nil {
				assert.NotEqual(t, test.Data["a"], test.Expected)
			}
		})
	}
}
