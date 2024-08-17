package std_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/std"
)

// TestHello asserts the Hello method returns the expected greeting.
func TestHello(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestHello", Input: `{{hello}}`, Expected: "Hello!"},
	}
	pesticide.RunTestCases(t, std.NewRegistry(), tc)
}

func TestDefault(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestDefaultEmptyInput", Input: `{{default "default" ""}}`, Expected: "default"},
		{Name: "TestDefaultGivenInput", Input: `{{default "default" "given"}}`, Expected: "given"},
		{Name: "TestDefaultIntInput", Input: `{{default "default" 42}}`, Expected: "42"},
		{Name: "TestDefaultFloatInput", Input: `{{default "default" 2.42}}`, Expected: "2.42"},
		{Name: "TestDefaultTrueInput", Input: `{{default "default" true}}`, Expected: "true"},
		{Name: "TestDefaultFalseInput", Input: `{{default "default" false}}`, Expected: "default"},
		{Name: "TestDefaultNilInput", Input: `{{default "default" nil}}`, Expected: "default"},
		{Name: "TestDefaultNothingInput", Input: `{{default "default" .Nothing}}`, Expected: "default"},
		{Name: "TestDefaultMultipleNothingInput", Input: `{{default "default" .Nothing}}`, Expected: "default"},
		{Name: "TestDefaultMultipleArgument", Input: `{{"first" | default "default" "second"}}`, Expected: "second"},
	}

	pesticide.RunTestCases(t, std.NewRegistry(), tc)
}

func TestEmpty(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyEmptyInput", Input: `{{if empty ""}}1{{else}}0{{end}}`, Expected: "1"},
		{Name: "TestEmptyGivenInput", Input: `{{if empty "given"}}1{{else}}0{{end}}`, Expected: "0"},
		{Name: "TestEmptyIntInput", Input: `{{if empty 42}}1{{else}}0{{end}}`, Expected: "0"},
		{Name: "TestEmptyUintInput", Input: `{{if empty .i}}1{{else}}0{{end}}`, Expected: "0", Data: map[string]any{"i": uint(42)}},
		{Name: "TestEmptyComplexInput", Input: `{{if empty .c}}1{{else}}0{{end}}`, Expected: "0", Data: map[string]any{"c": complex(42, 42)}},
		{Name: "TestEmptyFloatInput", Input: `{{if empty 2.42}}1{{else}}0{{end}}`, Expected: "0"},
		{Name: "TestEmptyTrueInput", Input: `{{if empty true}}1{{else}}0{{end}}`, Expected: "0"},
		{Name: "TestEmptyFalseInput", Input: `{{if empty false}}1{{else}}0{{end}}`, Expected: "1"},
		{Name: "TestEmptyStructInput", Input: `{{if empty .s}}1{{else}}0{{end}}`, Expected: "0", Data: map[string]any{"s": struct{}{}}},
		{Name: "TestEmptyNilInput", Input: `{{if empty nil}}1{{else}}0{{end}}`, Expected: "1"},
		{Name: "TestEmptyNothingInput", Input: `{{if empty .Nothing}}1{{else}}0{{end}}`, Expected: "1"},
		{Name: "TestEmptyNestedInput", Input: `{{if empty .top.NoSuchThing}}1{{else}}0{{end}}`, Expected: "1", Data: map[string]any{"top": map[string]any{}}},
		{Name: "TestEmptyNestedNoDataInput", Input: `{{if empty .bottom.NoSuchThing}}1{{else}}0{{end}}`, Expected: "1"},
		{Name: "TestEmptyNimPointerInput", Input: `{{if empty .nilPtr}}1{{else}}0{{end}}`, Expected: "1", Data: map[string]any{"nilPtr": (*int)(nil)}},
	}

	pesticide.RunTestCases(t, std.NewRegistry(), tc)
}

func TestAll(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestAllEmptyInput", Input: `{{if all ""}}1{{else}}0{{end}}`, Expected: "0"},
		{Name: "TestAllGivenInput", Input: `{{if all "given"}}1{{else}}0{{end}}`, Expected: "1"},
		{Name: "TestAllIntInput", Input: `{{if all 42 0 1}}1{{else}}0{{end}}`, Expected: "0"},
		{Name: "TestAllVariableInput1", Input: `{{ $two := 2 }}{{if all "" 0 nil $two }}1{{else}}0{{end}}`, Expected: "0"},
		{Name: "TestAllVariableInput2", Input: `{{ $two := 2 }}{{if all "" $two 0 0 0 }}1{{else}}0{{end}}`, Expected: "0"},
		{Name: "TestAllVariableInput3", Input: `{{ $two := 2 }}{{if all "" $two 3 4 5 }}1{{else}}0{{end}}`, Expected: "0"},
		{Name: "TestAllNoInput", Input: `{{if all }}1{{else}}0{{end}}`, Expected: "1"},
	}

	pesticide.RunTestCases(t, std.NewRegistry(), tc)
}

func TestAny(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestAnyEmptyInput", Input: `{{if any ""}}1{{else}}0{{end}}`, Expected: "0"},
		{Name: "TestAnyGivenInput", Input: `{{if any "given"}}1{{else}}0{{end}}`, Expected: "1"},
		{Name: "TestAnyIntInput", Input: `{{if any 42 0 1}}1{{else}}0{{end}}`, Expected: "1"},
		{Name: "TestAnyVariableInput1", Input: `{{ $two := 2 }}{{if any "" 0 nil $two }}1{{else}}0{{end}}`, Expected: "1"},
		{Name: "TestAnyVariableInput2", Input: `{{ $two := 2 }}{{if any "" $two 3 4 5 }}1{{else}}0{{end}}`, Expected: "1"},
		{Name: "TestAnyVariableInput3", Input: `{{ $zero := 0 }}{{if any "" $zero 0 0 0 }}1{{else}}0{{end}}`, Expected: "0"},
		{Name: "TestAnyNoInput", Input: `{{if any }}1{{else}}0{{end}}`, Expected: "0"},
	}

	pesticide.RunTestCases(t, std.NewRegistry(), tc)
}

func TestCoalesce(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestCoalesceEmptyInput", Input: `{{coalesce ""}}`, Expected: "<no value>"},
		{Name: "TestCoalesceGivenInput", Input: `{{coalesce "given"}}`, Expected: "given"},
		{Name: "TestCoalesceIntInput", Input: `{{ coalesce "" 0 nil 42 }}`, Expected: "42"},
		{Name: "TestCoalesceVariableInput1", Input: `{{ $two := 2 }}{{ coalesce "" 0 nil $two }}`, Expected: "2"},
		{Name: "TestCoalesceVariableInput2", Input: `{{ $two := 2 }}{{ coalesce "" $two 0 0 0 }}`, Expected: "2"},
		{Name: "TestCoalesceVariableInput3", Input: `{{ $two := 2 }}{{ coalesce "" $two 3 4 5 }}`, Expected: "2"},
		{Name: "TestCoalesceNoInput", Input: `{{ coalesce }}`, Expected: "<no value>"},
	}

	pesticide.RunTestCases(t, std.NewRegistry(), tc)
}

func TestTernary(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{true | ternary "foo" "bar"}}`, Expected: "foo"},
		{Input: `{{ternary "foo" "bar" true}}`, Expected: "foo"},
		{Input: `{{false | ternary "foo" "bar"}}`, Expected: "bar"},
		{Input: `{{ternary "foo" "bar" false}}`, Expected: "bar"},
	}

	pesticide.RunTestCases(t, std.NewRegistry(), tc)
}

func TestCat(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestCatEmptyInput", Input: `{{cat ""}}`, Expected: ""},
		{Name: "TestCatGivenInput", Input: `{{cat "given"}}`, Expected: "given"},
		{Name: "TestCatIntInput", Input: `{{cat 42}}`, Expected: "42"},
		{Name: "TestCatFloatInput", Input: `{{cat 2.42}}`, Expected: "2.42"},
		{Name: "TestCatTrueInput", Input: `{{cat true}}`, Expected: "true"},
		{Name: "TestCatFalseInput", Input: `{{cat false}}`, Expected: "false"},
		{Name: "TestCatNilInput", Input: `{{cat nil}}`, Expected: ""},
		{Name: "TestCatNothingInput", Input: `{{cat .Nothing}}`, Expected: ""},
		{Name: "TestCatMultipleInput", Input: `{{cat "first" "second"}}`, Expected: "first second"},
		{Name: "TestCatMultipleArgument", Input: `{{"first" | cat "second"}}`, Expected: "second first"},
		{Name: "TestCatVariableInput", Input: `{{$b := "b"}}{{"c" | cat "a" $b}}`, Expected: "a b c"},
		{Name: "TestCatDataInput", Input: `{{.text | cat "a" "b"}}`, Expected: "a b cd", Data: map[string]any{"text": "cd"}},
	}

	pesticide.RunTestCases(t, std.NewRegistry(), tc)
}
