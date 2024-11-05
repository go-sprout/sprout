package maps_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/maps"
)

func TestDeprecatedGet(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{get  . "a" }}`, ExpectedOutput: ""},
		{Name: "TestWithKey", Input: `{{get . "a" }}`, ExpectedOutput: "1", Data: map[string]any{"a": 1}},
		{Name: "TestWithNestedKeyNotFound", Input: `{{get . "b" }}`, ExpectedOutput: "", Data: map[string]any{"a": 1}},
		{Name: "TestInvalidArguments", Input: `{{get . "a" "b" }}`, ExpectedErr: "expected 2 arguments, got 3", Data: map[string]any{"a": 1}},
		{Name: "TestInvalidArgumentsType", Input: `{{get 1 2 }}`, ExpectedErr: "expected map or string, got int", Data: map[string]any{"a": 1}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestDeprecatedSet(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestWithKey", Input: `{{$d := set . "a" 2}}{{$d}}`, ExpectedOutput: "map[a:2]", Data: map[string]any{"a": 1}},
		{Name: "TestWithNewKey", Input: `{{$d := set . "b" 3}}{{$d}}`, ExpectedOutput: "map[a:1 b:3]", Data: map[string]any{"a": 1}},
		{Name: "TestWithNilValue", Input: `{{$d := set .V "a" .Nil}}{{$d}}`, ExpectedOutput: "map[a:<nil>]", Data: map[string]any{"V": map[string]any{"a": 1}, "Nil": nil}},
		{Name: "TestInvalidArguments", Input: `{{$d := set . "a" 2 "b" 3}}{{$d}}`, ExpectedErr: "expected 3 arguments, got 5", Data: map[string]any{"a": 1}},
		{Name: "TestInvalidArgumentsType", Input: `{{$d := set "a" "a" 2}}{{$d}}`, ExpectedErr: "last argument must be a map[string]any", Data: map[string]any{"a": 1}},
		{Name: "TestInvalidArgumentsType", Input: `{{$d := set 1 "a" 2}}{{$d}}`, ExpectedErr: "expected map or string, got int", Data: map[string]any{"a": 1}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestDeprecatedUnset(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestWithKey", Input: `{{$d := unset . "a"}}{{$d}}`, ExpectedOutput: "map[]", Data: map[string]any{"a": 1}},
		{Name: "TestWithNestedKeyNotFound", Input: `{{$d := unset . "b"}}{{$d}}`, ExpectedOutput: "map[a:1]", Data: map[string]any{"a": 1}},
		{Name: "TestInvalidArguments", Input: `{{$d := unset . "a" "b"}}{{$d}}`, ExpectedErr: "expected 2 arguments, got 3", Data: map[string]any{"a": 1}},
		{Name: "TestInvalidArgumentsType", Input: `{{$d := unset "a" "a"}}{{$d}}`, ExpectedErr: "last argument must be a map[string]any", Data: map[string]any{"a": 1}},
		{Name: "TestInvalidArgumentsType", Input: `{{$d := unset 1 "a"}}{{$d}}`, ExpectedErr: "expected map or string, got int", Data: map[string]any{"a": 1}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestDeprecatedHasKey(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{hasKey . "a"}}`, ExpectedOutput: "false"},
		{Name: "TestWithKey", Input: `{{hasKey . "a"}}`, ExpectedOutput: "true", Data: map[string]any{"a": 1}},
		{Name: "TestWithNestedKeyNotFound", Input: `{{hasKey . "b"}}`, ExpectedOutput: "false", Data: map[string]any{"a": 1}},
		{Name: "TestInvalidArguments", Input: `{{hasKey . "a" "b"}}`, ExpectedErr: "expected 2 arguments, got 3", Data: map[string]any{"a": 1}},
		{Name: "TestInvalidArgumentsType", Input: `{{hasKey 1 "a"}}`, ExpectedErr: "expected map or string, got int", Data: map[string]any{"a": 1}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestDeprecatedPick(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{pick . "a" "b"}}`, ExpectedOutput: "map[]"},
		{Name: "TestWithKeys", Input: `{{pick . "a" "b"}}`, ExpectedOutput: "map[a:1 b:2]", Data: map[string]any{"a": 1, "b": 2}},
		{Name: "TestWithNestedKeyNotFound", Input: `{{pick . "a" "b"}}`, ExpectedOutput: "map[a:1]", Data: map[string]any{"a": 1}},
		{Name: "TestWithInvalidKeys", Input: `{{pick . "a" 1}}`, ExpectedErr: "all keys must be strings", Data: map[string]any{"a": 1}},
		{Name: "TestInvalidArguments", Input: `{{pick . }}`, ExpectedErr: "expected 2 arguments, got 1", Data: map[string]any{"a": 1}},
		{Name: "TestInvalidArgumentsType", Input: `{{pick 1 "a"}}`, ExpectedErr: "expected map or string, got int", Data: map[string]any{"a": 1}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}

func TestDeprecatedOmit(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{omit . "a" "b"}}`, ExpectedOutput: "map[]"},
		{Name: "TestWithKeys", Input: `{{omit . "a" "b"}}`, ExpectedOutput: "map[]", Data: map[string]any{"a": 1, "b": 2}},
		{Name: "TestWithNestedKeyNotFound", Input: `{{omit . "b"}}`, ExpectedOutput: "map[a:1]", Data: map[string]any{"a": 1}},
		{Name: "TestWithInvalidKeys", Input: `{{omit . "a" 1}}`, ExpectedErr: "all keys must be strings", Data: map[string]any{"a": 1}},
		{Name: "TestInvalidArguments", Input: `{{omit . }}`, ExpectedErr: "expected 2 arguments, got 1", Data: map[string]any{"a": 1}},
		{Name: "TestInvalidArgumentsType", Input: `{{omit 1 "a"}}`, ExpectedErr: "expected map or string, got int", Data: map[string]any{"a": 1}},
	}

	pesticide.RunTestCases(t, maps.NewRegistry(), tc)
}
