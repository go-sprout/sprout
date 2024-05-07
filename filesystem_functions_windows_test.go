package sprout

import (
	"testing"
)

func TestOsBase(t *testing.T) {
	var tests = testCases{
		{"TestEmptyPath", `{{ osBase "" }}`, ".", nil},
		{"TestRootPath", `{{ osBase "D:\\" }}`, "\\", nil},
		{"TestWithoutExtension", `{{ osBase "D:\\path\\to\\file" }}`, "file", nil},
		{"TestWithFileInput", `{{ osBase "D:\\path\\to\\file.txt" }}`, "file.txt", nil},
		{"TestPipeSyntax", `{{ "D:\\path\\to\\file.txt" | osBase }}`, "file.txt", nil},
		{"TestVariableInput", `{{ .V | osBase }}`, "file", map[string]any{"V": "\\path\\to\\file"}},
	}

	runTestCases(t, tests)
}

func TestOsDir(t *testing.T) {
	var tests = testCases{
		{"TestEmptyPath", `{{ osDir "" }}`, ".", nil},
		{"TestRootPath", `{{ osDir "D:\\" }}`, "D:\\", nil},
		{"TestWithoutExtension", `{{ osDir "D:\\path\\to\\file" }}`, "D:\\path\\to", nil},
		{"TestWithFileInput", `{{ osDir "D:\\path\\to\\file.txt" }}`, "D:\\path\\to", nil},
		{"TestPipeSyntax", `{{ "D:\\path\\to\\file.txt" | osDir }}`, "D:\\path\\to", nil},
		{"TestVariableInput", `{{ .V | osDir }}`, "\\path\\to", map[string]any{"V": "\\path\\to\\file"}},
	}

	runTestCases(t, tests)
}

func TestOsExt(t *testing.T) {
	var tests = testCases{
		{"TestEmptyPath", `{{ osExt "" }}`, "", nil},
		{"TestRootPath", `{{ osExt "\\" }}`, "", nil},
		{"TestWithoutExtension", `{{ osExt "D:\\path\\to\\file" }}`, "", nil},
		{"TestWithFileInput", `{{ osExt "D:\\path\\to\\file.txt" }}`, ".txt", nil},
		{"TestPipeSyntax", `{{ "D:\\path\\to\\file.txt" | osExt }}`, ".txt", nil},
		{"TestVariableInput", `{{ .V | osExt }}`, ".txt", map[string]any{"V": "D:\\path\\to\\file.txt"}},
	}

	runTestCases(t, tests)
}

func TestOsIsAbs(t *testing.T) {
	var tests = testCases{
		{"TestEmptyPath", `{{ osIsAbs "" }}`, "false", nil},
		{"TestRootPath", `{{ osIsAbs "D:\\" }}`, "true", nil},
		{"TestRelativePath", `{{ osIsAbs "path\\to\\file" }}`, "false", nil},
		{"TestAbsolutePath", `{{ osIsAbs "D:\\path\\to\\file.txt" }}`, "true", nil},
		{"TestPipeSyntax", `{{ "file.txt" | osIsAbs }}`, "false", nil},
		{"TestVariableInput", `{{ osIsAbs .V }}`, "true", map[string]any{"V": "D:\\path\\to\\file"}},
	}

	runTestCases(t, tests)
}
