package std_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/std"
)

// TestHello asserts the Hello method returns the expected greeting.
func TestHello(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestHello", Input: `{{hello}}`, ExpectedOutput: "Hello!"},
	}
	pesticide.RunTestCases(t, std.NewRegistry(), tc)
}

func TestDefault(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestDefaultEmptyInput", Input: `{{default "default" ""}}`, ExpectedOutput: "default"},
		{Name: "TestDefaultGivenInput", Input: `{{default "default" "given"}}`, ExpectedOutput: "given"},
		{Name: "TestDefaultIntInput", Input: `{{default "default" 42}}`, ExpectedOutput: "42"},
		{Name: "TestDefaultFloatInput", Input: `{{default "default" 2.42}}`, ExpectedOutput: "2.42"},
		{Name: "TestDefaultTrueInput", Input: `{{default "default" true}}`, ExpectedOutput: "true"},
		{Name: "TestDefaultFalseInput", Input: `{{default "default" false}}`, ExpectedOutput: "default"},
		{Name: "TestDefaultNilInput", Input: `{{default "default" nil}}`, ExpectedOutput: "default"},
		{Name: "TestDefaultNothingInput", Input: `{{default "default" .Nothing}}`, ExpectedOutput: "default"},
		{Name: "TestDefaultMultipleNothingInput", Input: `{{default "default" .Nothing}}`, ExpectedOutput: "default"},
		{Name: "TestDefaultMultipleArgument", Input: `{{"first" | default "default" "second"}}`, ExpectedOutput: "second"},
	}

	pesticide.RunTestCases(t, std.NewRegistry(), tc)
}

func TestEmpty(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyEmptyInput", Input: `{{if empty ""}}1{{else}}0{{end}}`, ExpectedOutput: "1"},
		{Name: "TestEmptyGivenInput", Input: `{{if empty "given"}}1{{else}}0{{end}}`, ExpectedOutput: "0"},
		{Name: "TestEmptyIntInput", Input: `{{if empty 42}}1{{else}}0{{end}}`, ExpectedOutput: "0"},
		{Name: "TestEmptyUintInput", Input: `{{if empty .i}}1{{else}}0{{end}}`, ExpectedOutput: "0", Data: map[string]any{"i": uint(42)}},
		{Name: "TestEmptyComplexInput", Input: `{{if empty .c}}1{{else}}0{{end}}`, ExpectedOutput: "0", Data: map[string]any{"c": complex(42, 42)}},
		{Name: "TestEmptyFloatInput", Input: `{{if empty 2.42}}1{{else}}0{{end}}`, ExpectedOutput: "0"},
		{Name: "TestEmptyTrueInput", Input: `{{if empty true}}1{{else}}0{{end}}`, ExpectedOutput: "0"},
		{Name: "TestEmptyFalseInput", Input: `{{if empty false}}1{{else}}0{{end}}`, ExpectedOutput: "1"},
		{Name: "TestEmptyStructInput", Input: `{{if empty .s}}1{{else}}0{{end}}`, ExpectedOutput: "0", Data: map[string]any{"s": struct{}{}}},
		{Name: "TestEmptyNilInput", Input: `{{if empty nil}}1{{else}}0{{end}}`, ExpectedOutput: "1"},
		{Name: "TestEmptyNothingInput", Input: `{{if empty .Nothing}}1{{else}}0{{end}}`, ExpectedOutput: "1"},
		{Name: "TestEmptyNestedInput", Input: `{{if empty .top.NoSuchThing}}1{{else}}0{{end}}`, ExpectedOutput: "1", Data: map[string]any{"top": map[string]any{}}},
		{Name: "TestEmptyNestedNoDataInput", Input: `{{if empty .bottom.NoSuchThing}}1{{else}}0{{end}}`, ExpectedOutput: "1"},
		{Name: "TestEmptyNimPointerInput", Input: `{{if empty .nilPtr}}1{{else}}0{{end}}`, ExpectedOutput: "1", Data: map[string]any{"nilPtr": (*int)(nil)}},
	}

	pesticide.RunTestCases(t, std.NewRegistry(), tc)
}

func TestAll(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestAllEmptyInput", Input: `{{if all ""}}1{{else}}0{{end}}`, ExpectedOutput: "0"},
		{Name: "TestAllGivenInput", Input: `{{if all "given"}}1{{else}}0{{end}}`, ExpectedOutput: "1"},
		{Name: "TestAllIntInput", Input: `{{if all 42 0 1}}1{{else}}0{{end}}`, ExpectedOutput: "0"},
		{Name: "TestAllVariableInput1", Input: `{{ $two := 2 }}{{if all "" 0 nil $two }}1{{else}}0{{end}}`, ExpectedOutput: "0"},
		{Name: "TestAllVariableInput2", Input: `{{ $two := 2 }}{{if all "" $two 0 0 0 }}1{{else}}0{{end}}`, ExpectedOutput: "0"},
		{Name: "TestAllVariableInput3", Input: `{{ $two := 2 }}{{if all "" $two 3 4 5 }}1{{else}}0{{end}}`, ExpectedOutput: "0"},
		{Name: "TestAllNoInput", Input: `{{if all }}1{{else}}0{{end}}`, ExpectedOutput: "1"},
	}

	pesticide.RunTestCases(t, std.NewRegistry(), tc)
}

func TestAny(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestAnyEmptyInput", Input: `{{if any ""}}1{{else}}0{{end}}`, ExpectedOutput: "0"},
		{Name: "TestAnyGivenInput", Input: `{{if any "given"}}1{{else}}0{{end}}`, ExpectedOutput: "1"},
		{Name: "TestAnyIntInput", Input: `{{if any 42 0 1}}1{{else}}0{{end}}`, ExpectedOutput: "1"},
		{Name: "TestAnyVariableInput1", Input: `{{ $two := 2 }}{{if any "" 0 nil $two }}1{{else}}0{{end}}`, ExpectedOutput: "1"},
		{Name: "TestAnyVariableInput2", Input: `{{ $two := 2 }}{{if any "" $two 3 4 5 }}1{{else}}0{{end}}`, ExpectedOutput: "1"},
		{Name: "TestAnyVariableInput3", Input: `{{ $zero := 0 }}{{if any "" $zero 0 0 0 }}1{{else}}0{{end}}`, ExpectedOutput: "0"},
		{Name: "TestAnyNoInput", Input: `{{if any }}1{{else}}0{{end}}`, ExpectedOutput: "0"},
	}

	pesticide.RunTestCases(t, std.NewRegistry(), tc)
}

func TestCoalesce(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestCoalesceEmptyInput", Input: `{{coalesce ""}}`, ExpectedOutput: "<no value>"},
		{Name: "TestCoalesceGivenInput", Input: `{{coalesce "given"}}`, ExpectedOutput: "given"},
		{Name: "TestCoalesceIntInput", Input: `{{ coalesce "" 0 nil 42 }}`, ExpectedOutput: "42"},
		{Name: "TestCoalesceVariableInput1", Input: `{{ $two := 2 }}{{ coalesce "" 0 nil $two }}`, ExpectedOutput: "2"},
		{Name: "TestCoalesceVariableInput2", Input: `{{ $two := 2 }}{{ coalesce "" $two 0 0 0 }}`, ExpectedOutput: "2"},
		{Name: "TestCoalesceVariableInput3", Input: `{{ $two := 2 }}{{ coalesce "" $two 3 4 5 }}`, ExpectedOutput: "2"},
		{Name: "TestCoalesceNoInput", Input: `{{ coalesce }}`, ExpectedOutput: "<no value>"},
	}

	pesticide.RunTestCases(t, std.NewRegistry(), tc)
}

func TestTernary(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{true | ternary "foo" "bar"}}`, ExpectedOutput: "foo"},
		{Input: `{{ternary "foo" "bar" true}}`, ExpectedOutput: "foo"},
		{Input: `{{false | ternary "foo" "bar"}}`, ExpectedOutput: "bar"},
		{Input: `{{ternary "foo" "bar" false}}`, ExpectedOutput: "bar"},
	}

	pesticide.RunTestCases(t, std.NewRegistry(), tc)
}

func TestCat(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestCatEmptyInput", Input: `{{cat ""}}`, ExpectedOutput: ""},
		{Name: "TestCatGivenInput", Input: `{{cat "given"}}`, ExpectedOutput: "given"},
		{Name: "TestCatIntInput", Input: `{{cat 42}}`, ExpectedOutput: "42"},
		{Name: "TestCatFloatInput", Input: `{{cat 2.42}}`, ExpectedOutput: "2.42"},
		{Name: "TestCatTrueInput", Input: `{{cat true}}`, ExpectedOutput: "true"},
		{Name: "TestCatFalseInput", Input: `{{cat false}}`, ExpectedOutput: "false"},
		{Name: "TestCatNilInput", Input: `{{cat nil}}`, ExpectedOutput: ""},
		{Name: "TestCatNothingInput", Input: `{{cat .Nothing}}`, ExpectedOutput: ""},
		{Name: "TestCatMultipleInput", Input: `{{cat "first" "second"}}`, ExpectedOutput: "first second"},
		{Name: "TestCatMultipleArgument", Input: `{{"first" | cat "second"}}`, ExpectedOutput: "second first"},
		{Name: "TestCatVariableInput", Input: `{{$b := "b"}}{{"c" | cat "a" $b}}`, ExpectedOutput: "a b c"},
		{Name: "TestCatDataInput", Input: `{{.text | cat "a" "b"}}`, ExpectedOutput: "a b cd", Data: map[string]any{"text": "cd"}},
	}

	pesticide.RunTestCases(t, std.NewRegistry(), tc)
}
