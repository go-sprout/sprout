package filesystem_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/filesystem"
)

func TestOsBase(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ osBase "" }}`, ExpectedOutput: "."},
		{Name: "TestRootPath", Input: `{{ osBase "D:\\" }}`, ExpectedOutput: "\\"},
		{Name: "TestWithoutExtension", Input: `{{ osBase "D:\\path\\to\\file" }}`, ExpectedOutput: "file"},
		{Name: "TestWithFileInput", Input: `{{ osBase "D:\\path\\to\\file.txt" }}`, ExpectedOutput: "file.txt"},
		{Name: "TestPipeSyntax", Input: `{{ "D:\\path\\to\\file.txt" | osBase }}`, ExpectedOutput: "file.txt"},
		{Name: "TestVariableInput", Input: `{{ .V | osBase }}`, ExpectedOutput: "file", Data: map[string]any{"V": "\\path\\to\\file"}},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}

func TestOsDir(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ osDir "" }}`, ExpectedOutput: "."},
		{Name: "TestRootPath", Input: `{{ osDir "D:\\" }}`, ExpectedOutput: "D:\\"},
		{Name: "TestWithoutExtension", Input: `{{ osDir "D:\\path\\to\\file" }}`, ExpectedOutput: "D:\\path\\to"},
		{Name: "TestWithFileInput", Input: `{{ osDir "D:\\path\\to\\file.txt" }}`, ExpectedOutput: "D:\\path\\to"},
		{Name: "TestPipeSyntax", Input: `{{ "D:\\path\\to\\file.txt" | osDir }}`, ExpectedOutput: "D:\\path\\to"},
		{Name: "TestVariableInput", Input: `{{ .V | osDir }}`, ExpectedOutput: "\\path\\to", Data: map[string]any{"V": "\\path\\to\\file"}},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}

func TestOsExt(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ osExt "" }}`, ExpectedOutput: ""},
		{Name: "TestRootPath", Input: `{{ osExt "\\" }}`, ExpectedOutput: ""},
		{Name: "TestWithoutExtension", Input: `{{ osExt "D:\\path\\to\\file" }}`, ExpectedOutput: ""},
		{Name: "TestWithFileInput", Input: `{{ osExt "D:\\path\\to\\file.txt" }}`, ExpectedOutput: ".txt"},
		{Name: "TestPipeSyntax", Input: `{{ "D:\\path\\to\\file.txt" | osExt }}`, ExpectedOutput: ".txt"},
		{Name: "TestVariableInput", Input: `{{ .V | osExt }}`, ExpectedOutput: ".txt", Data: map[string]any{"V": "D:\\path\\to\\file.txt"}},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}

func TestOsClean(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ osClean "" }}`, ExpectedOutput: "."},
		{Name: "TestRootPath", Input: `{{ osClean "D:\\" }}`, ExpectedOutput: "D:\\"},
		{Name: "TestWithoutExtension", Input: `{{ osClean "D:\\path\\\\to\\file" }}`, ExpectedOutput: "D:\\path\\to\\file"},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}

func TestOsIsAbs(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ osIsAbs "" }}`, ExpectedOutput: "false"},
		{Name: "TestRootPath", Input: `{{ osIsAbs "D:\\" }}`, ExpectedOutput: "true"},
		{Name: "TestRelativePath", Input: `{{ osIsAbs "path\\to\\file" }}`, ExpectedOutput: "false"},
		{Name: "TestAbsolutePath", Input: `{{ osIsAbs "D:\\path\\to\\file.txt" }}`, ExpectedOutput: "true"},
		{Name: "TestPipeSyntax", Input: `{{ "file.txt" | osIsAbs }}`, ExpectedOutput: "false"},
		{Name: "TestVariableInput", Input: `{{ osIsAbs .V }}`, ExpectedOutput: "true", Data: map[string]any{"V": "D:\\path\\to\\file"}},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}
