package env_test

import (
	"os"
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/env"
)

func TestEnv(t *testing.T) {
	os.Setenv("__SPROUT_TEST_ENV_KEY", "sprout will grow!")
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ env "" }}`, ExpectedOutput: ""},
		{Name: "TestNonExistent", Input: `{{ env "NON_EXISTENT_ENV_VAR" }}`, ExpectedOutput: ""},
		{Name: "TestExisting", Input: `{{ env "__SPROUT_TEST_ENV_KEY" }}`, ExpectedOutput: "sprout will grow!"},
		{Name: "TestPipeSyntax", Input: `{{ "__SPROUT_TEST_ENV_KEY" | env }}`, ExpectedOutput: "sprout will grow!"},
		{Name: "TestVariableInput", Input: `{{ .V | env }}`, ExpectedOutput: "sprout will grow!", Data: map[string]any{"V": "__SPROUT_TEST_ENV_KEY"}},
	}

	pesticide.RunTestCases(t, env.NewRegistry(), tc)
}

func TestExpandEnv(t *testing.T) {
	os.Setenv("__SPROUT_TEST_ENV_KEY", "sprout will grow!")
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ expandEnv "" }}`, ExpectedOutput: ""},
		{Name: "TestNonExistent", Input: `{{ expandEnv "Hey" }}`, ExpectedOutput: "Hey"},
		{Name: "TestNonExistent", Input: `{{ expandEnv "$NON_EXISTENT_ENV_VAR" }}`, ExpectedOutput: ""},
		{Name: "TestExisting", Input: `{{ expandEnv "Hey $__SPROUT_TEST_ENV_KEY" }}`, ExpectedOutput: "Hey sprout will grow!"},
		{Name: "TestPipeSyntax", Input: `{{ "Hey $__SPROUT_TEST_ENV_KEY" | expandEnv }}`, ExpectedOutput: "Hey sprout will grow!"},
		{Name: "TestVariableInput", Input: `{{ .V | expandEnv }}`, ExpectedOutput: "Hey sprout will grow!", Data: map[string]any{"V": "Hey $__SPROUT_TEST_ENV_KEY"}},
	}

	pesticide.RunTestCases(t, env.NewRegistry(), tc)
}
