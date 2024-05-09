package sprout

import (
	"os"
	"testing"
)

func TestPathBase(t *testing.T) {
	var tests = testCases{
		{"TestEmptyPath", `{{ pathBase "" }}`, ".", nil},
		{"TestRootPath", `{{ pathBase "/" }}`, "/", nil},
		{"TestWithoutExtension", `{{ pathBase "/path/to/file" }}`, "file", nil},
		{"TestWithFileInput", `{{ pathBase "/path/to/file.txt" }}`, "file.txt", nil},
		{"TestPipeSyntax", `{{ "/path/to/file.txt" | pathBase }}`, "file.txt", nil},
		{"TestVariableInput", `{{ .V | pathBase }}`, "file", map[string]any{"V": "/path/to/file"}},
	}

	runTestCases(t, tests)
}

func TestPathDir(t *testing.T) {
	var tests = testCases{
		{"TestEmptyPath", `{{ pathDir "" }}`, ".", nil},
		{"TestRootPath", `{{ pathDir "/" }}`, "/", nil},
		{"TestWithoutExtension", `{{ pathDir "/path/to/file" }}`, "/path/to", nil},
		{"TestWithFileInput", `{{ pathDir "/path/to/file.txt" }}`, "/path/to", nil},
		{"TestPipeSyntax", `{{ "/path/to/file.txt" | pathDir }}`, "/path/to", nil},
		{"TestVariableInput", `{{ .V | pathDir }}`, "/path/to", map[string]any{"V": "/path/to/file"}},
	}

	runTestCases(t, tests)
}

func TestPathExt(t *testing.T) {
	var tests = testCases{
		{"TestEmptyPath", `{{ pathExt "" }}`, "", nil},
		{"TestRootPath", `{{ pathExt "/" }}`, "", nil},
		{"TestWithoutExtension", `{{ pathExt "/path/to/file" }}`, "", nil},
		{"TestWithFileInput", `{{ pathExt "/path/to/file.txt" }}`, ".txt", nil},
		{"TestPipeSyntax", `{{ "/path/to/file.txt" | pathExt }}`, ".txt", nil},
		{"TestVariableInput", `{{ .V | pathExt }}`, ".txt", map[string]any{"V": "/path/to/file.txt"}},
	}

	runTestCases(t, tests)
}

func TestPathClean(t *testing.T) {
	var tests = testCases{
		{"TestEmptyPath", `{{ pathClean "" }}`, ".", nil},
		{"TestRootPath", `{{ pathClean "/" }}`, "/", nil},
		{"TestWithoutExtension", `{{ pathClean "/path/to/file" }}`, "/path/to/file", nil},
		{"TestWithFileInput", `{{ pathClean "/path/to/file.txt" }}`, "/path/to/file.txt", nil},
		{"TestPipeSyntax", `{{ "/path/to/file.txt" | pathClean }}`, "/path/to/file.txt", nil},
		{"TestVariableInput", `{{ .V | pathClean }}`, "/path/to/file", map[string]any{"V": "/path/to/file"}},
		{"TestDoubleSlash", `{{ pathClean "/path//to/file" }}`, "/path/to/file", nil},
		{"TestDotSlash", `{{ pathClean "/path/./to/file" }}`, "/path/to/file", nil},
		{"TestDotDotSlash", `{{ pathClean "/path/../to/file" }}`, "/to/file", nil},
	}

	runTestCases(t, tests)
}

func TestPathIsAbs(t *testing.T) {
	var tests = testCases{
		{"TestEmptyPath", `{{ pathIsAbs "" }}`, "false", nil},
		{"TestRootPath", `{{ pathIsAbs "/" }}`, "true", nil},
		{"TestRelativePath", `{{ pathIsAbs "path/to/file" }}`, "false", nil},
		{"TestAbsolutePath", `{{ pathIsAbs "/path/to/file.txt" }}`, "true", nil},
		{"TestPipeSyntax", `{{ "file.txt" | pathIsAbs }}`, "false", nil},
		{"TestVariableInput", `{{ pathIsAbs .V }}`, "true", map[string]any{"V": "/path/to/file"}},
	}

	runTestCases(t, tests)
}

func TestEnv(t *testing.T) {
	os.Setenv("__SPROUT_TEST_ENV_KEY", "sprout will grow!")
	var tests = testCases{
		{"TestEmpty", `{{ env "" }}`, "", nil},
		{"TestNonExistent", `{{ env "NON_EXISTENT_ENV_VAR" }}`, "", nil},
		{"TestExisting", `{{ env "__SPROUT_TEST_ENV_KEY" }}`, "sprout will grow!", nil},
		{"TestPipeSyntax", `{{ "__SPROUT_TEST_ENV_KEY" | env }}`, "sprout will grow!", nil},
		{"TestVariableInput", `{{ .V | env }}`, "sprout will grow!", map[string]any{"V": "__SPROUT_TEST_ENV_KEY"}},
	}

	runTestCases(t, tests)
}

func TestExpandEnv(t *testing.T) {
	os.Setenv("__SPROUT_TEST_ENV_KEY", "sprout will grow!")
	var tests = testCases{
		{"TestEmpty", `{{ expandEnv "" }}`, "", nil},
		{"TestNonExistent", `{{ expandEnv "Hey" }}`, "Hey", nil},
		{"TestNonExistent", `{{ expandEnv "$NON_EXISTENT_ENV_VAR" }}`, "", nil},
		{"TestExisting", `{{ expandEnv "Hey $__SPROUT_TEST_ENV_KEY" }}`, "Hey sprout will grow!", nil},
		{"TestPipeSyntax", `{{ "Hey $__SPROUT_TEST_ENV_KEY" | expandEnv }}`, "Hey sprout will grow!", nil},
		{"TestVariableInput", `{{ .V | expandEnv }}`, "Hey sprout will grow!", map[string]any{"V": "Hey $__SPROUT_TEST_ENV_KEY"}},
	}

	runTestCases(t, tests)
}
