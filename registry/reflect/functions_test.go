package reflect_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/reflect"
	"github.com/stretchr/testify/assert"
)

var nilPointer *int = nil
var nilInterface interface{}

func TestTypeIs(t *testing.T) {
	type testStruct struct{}

	var tc = []pesticide.TestCase{
		{Name: "TestTypeIsInt", Input: `{{typeIs "int" 42}}`, ExpectedOutput: "true"},
		{Name: "TestTypeIsString", Input: `{{42 | typeIs "string"}}`, ExpectedOutput: "false"},
		{Name: "TestTypeIsVariable", Input: `{{$var := 42}}{{typeIs "string" $var}}`, ExpectedOutput: "false"},
		{Name: "TestTypeIsStruct", Input: `{{.var | typeIs "*reflect_test.testStruct"}}`, ExpectedOutput: "true", Data: map[string]any{"var": &testStruct{}}},
	}

	pesticide.RunTestCases(t, reflect.NewRegistry(), tc)
}

func TestTypeIsLike(t *testing.T) {
	type testStruct struct{}

	var tc = []pesticide.TestCase{
		{Name: "TestTypeIsLikeInt", Input: `{{typeIsLike "int" 42}}`, ExpectedOutput: "true"},
		{Name: "TestTypeIsLikeString", Input: `{{42 | typeIsLike "string"}}`, ExpectedOutput: "false"},
		{Name: "TestTypeIsLikeVariable", Input: `{{$var := 42}}{{typeIsLike "string" $var}}`, ExpectedOutput: "false"},
		{Name: "TestTypeIsLikeStruct", Input: `{{.var | typeIsLike "*reflect_test.testStruct"}}`, ExpectedOutput: "true", Data: map[string]any{"var": &testStruct{}}},
		{Name: "TestTypeIsLikeStructWithoutPointerMark", Input: `{{.var | typeIsLike "reflect_test.testStruct"}}`, ExpectedOutput: "true", Data: map[string]any{"var": &testStruct{}}},
	}

	pesticide.RunTestCases(t, reflect.NewRegistry(), tc)
}

func TestTypeOf(t *testing.T) {
	type testStruct struct{}

	var tc = []pesticide.TestCase{
		{Name: "TestTypeOfInt", Input: `{{typeOf 42}}`, ExpectedOutput: "int"},
		{Name: "TestTypeOfString", Input: `{{typeOf "42"}}`, ExpectedOutput: "string"},
		{Name: "TestTypeOfVariable", Input: `{{$var := 42}}{{typeOf $var}}`, ExpectedOutput: "int"},
		{Name: "TestTypeOfStruct", Input: `{{typeOf .var}}`, ExpectedOutput: "*reflect_test.testStruct", Data: map[string]any{"var": &testStruct{}}},
	}

	pesticide.RunTestCases(t, reflect.NewRegistry(), tc)
}

func TestKindIs(t *testing.T) {
	type testStruct struct{}

	var tc = []pesticide.TestCase{
		{Name: "TestKindIsInt", Input: `{{kindIs "int" 42}}`, ExpectedOutput: "true"},
		{Name: "TestKindIsString", Input: `{{42 | kindIs "string"}}`, ExpectedOutput: "false"},
		{Name: "TestKindIsVariable", Input: `{{$var := 42}}{{kindIs "string" $var}}`, ExpectedOutput: "false"},
		{Name: "TestKindIsStruct", Input: `{{.var | kindIs "ptr"}}`, ExpectedOutput: "true", Data: map[string]any{"var": &testStruct{}}},
		{Name: "TestKindIsInterfaceNil", Input: `{{.var | kindIs "ptr"}}`, ExpectedErr: "src must not be nil", Data: map[string]any{"V": nilInterface}},
	}

	pesticide.RunTestCases(t, reflect.NewRegistry(), tc)
}

func TestKindOf(t *testing.T) {
	type testStruct struct{}

	var tc = []pesticide.TestCase{
		{Name: "TestKindOfInt", Input: `{{kindOf 42}}`, ExpectedOutput: "int"},
		{Name: "TestKindOfString", Input: `{{kindOf "42"}}`, ExpectedOutput: "string"},
		{Name: "TestKindOfSlice", Input: `{{kindOf .var}}`, ExpectedOutput: "slice", Data: map[string]any{"var": []int{}}},
		{Name: "TestKindOfVariable", Input: `{{$var := 42}}{{kindOf $var}}`, ExpectedOutput: "int"},
		{Name: "TestKindOfStruct", Input: `{{kindOf .var}}`, ExpectedOutput: "ptr", Data: map[string]any{"var": &testStruct{}}},
		{Name: "TestKindOfStructWithoutPointerMark", Input: `{{kindOf .var}}`, ExpectedOutput: "struct", Data: map[string]any{"var": testStruct{}}},
		{Name: "TestKindOfIntNil", Input: `{{kindOf .V }}`, ExpectedOutput: "ptr", Data: map[string]any{"V": nilPointer}},
		{Name: "TestKindOfInterfaceNil", Input: `{{kindOf .V }}`, ExpectedErr: "src must not be nil", Data: map[string]any{"V": nilInterface}},
		{Name: "TestKindOfAnyNil", Input: `{{kindOf nil}}`, ExpectedErr: "src must not be nil"},
	}

	pesticide.RunTestCases(t, reflect.NewRegistry(), tc)
}

func TestHasField(t *testing.T) {
	type A struct {
		Foo string
	}
	type B struct {
		Bar string
	}

	var tc = []pesticide.TestCase{
		{Name: "TestHasFieldStructPointerTrue", Input: `{{ .V | hasField "Foo" }}`, ExpectedOutput: "true", Data: map[string]any{"V": &A{Foo: "bar"}}},
		{Name: "TestHasFieldStructPointerFalse", Input: `{{ .V | hasField "Foo" }}`, ExpectedOutput: "false", Data: map[string]any{"V": &B{Bar: "boo"}}},
		{Name: "TestHasFieldStructTrue", Input: `{{ .V | hasField "Foo" }}`, ExpectedOutput: "true", Data: map[string]any{"V": A{Foo: "bar"}}},
		{Name: "TestHasFieldStructFalse", Input: `{{ .V | hasField "Foo" }}`, ExpectedOutput: "false", Data: map[string]any{"V": B{Bar: "boo"}}},
		{Name: "TestHasFieldMap", Input: `{{ .V | hasField "Foo" }}`, ExpectedErr: "last argument must be a struct", Data: map[string]any{"V": map[string]string{"Foo": "bar"}}},
		{Name: "TestHasFieldInt", Input: `{{ .V | hasField "Foo" }}`, ExpectedErr: "last argument must be a struct", Data: map[string]any{"V": 123}},
		{Name: "TestHasFieldSlice", Input: `{{ .V | hasField "Foo" }}`, ExpectedErr: "last argument must be a struct", Data: map[string]any{"V": []int{1, 2, 3}}},
		{Name: "TestHasFieldString", Input: `{{ .V | hasField "Foo" }}`, ExpectedErr: "last argument must be a struct", Data: map[string]any{"V": "foobar"}},
		{Name: "TestHasFieldNil", Input: `{{ .V | hasField "Foo" }}`, ExpectedErr: "last argument must be a struct", Data: map[string]any{"V": nil}},
	}

	pesticide.RunTestCases(t, reflect.NewRegistry(), tc)
}

func TestDeepEqual(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestDeepEqualInt", Input: `{{deepEqual 42 42}}`, ExpectedOutput: "true"},
		{Name: "TestDeepEqualString", Input: `{{deepEqual "42" "42"}}`, ExpectedOutput: "true"},
		{Name: "TestDeepEqualSlice", Input: `{{deepEqual .a .b}}`, ExpectedOutput: "true", Data: map[string]any{"a": []int{1, 2, 3}, "b": []int{1, 2, 3}}},
		{Name: "TestDeepEqualMap", Input: `{{deepEqual .a .b}}`, ExpectedOutput: "true", Data: map[string]any{"a": map[string]int{"a": 1, "b": 2}, "b": map[string]int{"a": 1, "b": 2}}},
		{Name: "TestDeepEqualStruct", Input: `{{deepEqual .a .b}}`, ExpectedOutput: "true", Data: map[string]any{"a": struct{ A int }{A: 1}, "b": struct{ A int }{A: 1}}},
		{Name: "TestDeepEqualVariable", Input: `{{$a := 42}}{{$b := 42}}{{deepEqual $a $b}}`, ExpectedOutput: "true"},
		{Name: "TestDeepEqualDifferent", Input: `{{deepEqual 42 32}}`, ExpectedOutput: "false"},
		{Name: "TestDeepEqualDifferentType", Input: `{{deepEqual 42 "42"}}`, ExpectedOutput: "false"},
	}

	pesticide.RunTestCases(t, reflect.NewRegistry(), tc)
}

func TestDeepCopy(t *testing.T) {
	type testStruct struct {
		A int
	}

	var tc = []pesticide.TestCase{
		{Name: "TestDeepCopyInt", Input: `{{$a := 42}}{{$b := deepCopy $a}}{{$b}}`, ExpectedOutput: "42"},
		{Name: "TestDeepCopyString", Input: `{{$a := "42"}}{{$b := deepCopy $a}}{{$b}}`, ExpectedOutput: "42"},
		{Name: "TestDeepCopySlice", Input: `{{$a := .a}}{{$b := deepCopy $a}}{{$b}}`, ExpectedOutput: "[1 2 3]", Data: map[string]any{"a": []int{1, 2, 3}}},
		{Name: "TestDeepCopyMap", Input: `{{$a := .a}}{{$b := deepCopy $a}}{{$b}}`, ExpectedOutput: `map[a:1 b:2]`, Data: map[string]any{"a": map[string]int{"a": 1, "b": 2}}},
		{Name: "TestDeepCopyStruct", Input: `{{$a := .a}}{{$b := deepCopy $a}}{{$b}}`, ExpectedOutput: "{1}", Data: map[string]any{"a": testStruct{A: 1}}},
		{Name: "TestDeepCopyVariable", Input: `{{$a := 42}}{{$b := deepCopy $a}}{{$b}}`, ExpectedOutput: "42"},
		{Name: "TestDeepCopyDifferent", Input: `{{$a := 42}}{{$b := deepCopy "42"}}{{$b}}`, ExpectedOutput: "42"},
		{Name: "TestDeepCopyDifferentType", Input: `{{$a := 42}}{{$b := deepCopy 42.0}}{{$b}}`, ExpectedOutput: "42"},
		{Name: "TestDeepCopyNil", Input: `{{$b := deepCopy .a}}`, ExpectedErr: "element cannot be nil", Data: map[string]any{"a": nil}},
		{Input: `{{- $d := dict "a" 1 "b" 2 | deepCopy }}{{ values $d | sortAlpha | join "," }}`, ExpectedOutput: "1,2"},
		{Input: `{{- $d := dict "a" 1 "b" 2 | deepCopy }}{{ keys $d | sortAlpha | join "," }}`, ExpectedOutput: "a,b"},
		{Input: `{{- $one := dict "foo" (dict "bar" "baz") "qux" true -}}{{ deepCopy $one }}`, ExpectedOutput: "map[foo:map[bar:baz] qux:true]"},
	}

	for _, test := range tc {
		t.Run(test.Name, func(t *testing.T) {
			tmplResponse, err := pesticide.TestTemplate(t, reflect.NewRegistry(), test.Input, test.Data)
			if test.ExpectedErr != "" {
				assert.ErrorContains(t, err, test.ExpectedErr)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.ExpectedOutput, tmplResponse)

			if test.Data != nil {
				assert.NotEqual(t, test.Data["a"], test.ExpectedOutput)
			}
		})
	}
}
