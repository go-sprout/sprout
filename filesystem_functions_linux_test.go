package sprout

import "testing"

func TestOsBase(t *testing.T) {
	var tests = testCases{
		{"TestEmptyPath", `{{ osBase "" }}`, ".", nil},
		{"TestRootPath", `{{ osBase "/" }}`, "/", nil},
		{"TestWithoutExtension", `{{ osBase "/path/to/file" }}`, "file", nil},
		{"TestWithFileInput", `{{ osBase "/path/to/file.txt" }}`, "file.txt", nil},
		{"TestPipeSyntax", `{{ "/path/to/file.txt" | osBase }}`, "file.txt", nil},
		{"TestVariableInput", `{{ .V | osBase }}`, "file", map[string]any{"V": "/path/to/file"}},
	}

	runTestCases(t, tests)
}

func TestOsDir(t *testing.T) {
	var tests = testCases{
		{"TestEmptyPath", `{{ osDir "" }}`, ".", nil},
		{"TestRootPath", `{{ osDir "/" }}`, "/", nil},
		{"TestWithoutExtension", `{{ osDir "/path/to/file" }}`, "/path/to", nil},
		{"TestWithFileInput", `{{ osDir "/path/to/file.txt" }}`, "/path/to", nil},
		{"TestPipeSyntax", `{{ "/path/to/file.txt" | osDir }}`, "/path/to", nil},
		{"TestVariableInput", `{{ .V | osDir }}`, "/path/to", map[string]any{"V": "/path/to/file"}},
	}

	runTestCases(t, tests)
}

func TestOsExt(t *testing.T) {
	var tests = testCases{
		{"TestEmptyPath", `{{ osExt "" }}`, "", nil},
		{"TestRootPath", `{{ osExt "/" }}`, "", nil},
		{"TestWithoutExtension", `{{ osExt "/path/to/file" }}`, "", nil},
		{"TestWithFileInput", `{{ osExt "/path/to/file.txt" }}`, ".txt", nil},
		{"TestPipeSyntax", `{{ "/path/to/file.txt" | osExt }}`, ".txt", nil},
		{"TestVariableInput", `{{ .V | osExt }}`, ".txt", map[string]any{"V": "/path/to/file.txt"}},
	}

	runTestCases(t, tests)
}

func TestOsClean(t *testing.T) {
	var tests = testCases{
		{"TestEmptyPath", `{{ osClean "" }}`, ".", nil},
		{"TestRootPath", `{{ osClean "/" }}`, "/", nil},
		{"TestWithoutExtension", `{{ osClean "/path///to/file" }}`, "/path/to/file", nil},
	}

	runTestCases(t, tests)
}

func TestOsIsAbs(t *testing.T) {
	var tests = testCases{
		{"TestEmptyPath", `{{ osIsAbs "" }}`, "false", nil},
		{"TestRootPath", `{{ osIsAbs "/" }}`, "true", nil},
		{"TestRelativePath", `{{ osIsAbs "path/to/file" }}`, "false", nil},
		{"TestAbsolutePath", `{{ osIsAbs "/path/to/file.txt" }}`, "true", nil},
		{"TestPipeSyntax", `{{ "file.txt" | osIsAbs }}`, "false", nil},
		{"TestVariableInput", `{{ osIsAbs .V }}`, "true", map[string]any{"V": "/path/to/file"}},
	}

	runTestCases(t, tests)
}
