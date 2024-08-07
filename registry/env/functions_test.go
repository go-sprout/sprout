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
		{Name: "TestEmpty", Input: `{{ env "" }}`, Expected: ""},
		{Name: "TestNonExistent", Input: `{{ env "NON_EXISTENT_ENV_VAR" }}`, Expected: ""},
		{Name: "TestExisting", Input: `{{ env "__SPROUT_TEST_ENV_KEY" }}`, Expected: "sprout will grow!"},
		{Name: "TestPipeSyntax", Input: `{{ "__SPROUT_TEST_ENV_KEY" | env }}`, Expected: "sprout will grow!"},
		{Name: "TestVariableInput", Input: `{{ .V | env }}`, Expected: "sprout will grow!", Data: map[string]any{"V": "__SPROUT_TEST_ENV_KEY"}},
	}

	pesticide.RunTestCases(t, env.NewRegistry(), tc)
}

func TestExpandEnv(t *testing.T) {
	os.Setenv("__SPROUT_TEST_ENV_KEY", "sprout will grow!")
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ expandEnv "" }}`, Expected: ""},
		{Name: "TestNonExistent", Input: `{{ expandEnv "Hey" }}`, Expected: "Hey"},
		{Name: "TestNonExistent", Input: `{{ expandEnv "$NON_EXISTENT_ENV_VAR" }}`, Expected: ""},
		{Name: "TestExisting", Input: `{{ expandEnv "Hey $__SPROUT_TEST_ENV_KEY" }}`, Expected: "Hey sprout will grow!"},
		{Name: "TestPipeSyntax", Input: `{{ "Hey $__SPROUT_TEST_ENV_KEY" | expandEnv }}`, Expected: "Hey sprout will grow!"},
		{Name: "TestVariableInput", Input: `{{ .V | expandEnv }}`, Expected: "Hey sprout will grow!", Data: map[string]any{"V": "Hey $__SPROUT_TEST_ENV_KEY"}},
	}

	pesticide.RunTestCases(t, env.NewRegistry(), tc)
}
