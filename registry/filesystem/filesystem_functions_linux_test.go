package filesystem_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/filesystem"
)

func TestOsBase(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ osBase "" }}`, ExpectedOutput: "."},
		{Name: "TestRootPath", Input: `{{ osBase "/" }}`, ExpectedOutput: "/"},
		{Name: "TestWithoutExtension", Input: `{{ osBase "/path/to/file" }}`, ExpectedOutput: "file"},
		{Name: "TestWithFileInput", Input: `{{ osBase "/path/to/file.txt" }}`, ExpectedOutput: "file.txt"},
		{Name: "TestPipeSyntax", Input: `{{ "/path/to/file.txt" | osBase }}`, ExpectedOutput: "file.txt"},
		{Name: "TestVariableInput", Input: `{{ .V | osBase }}`, ExpectedOutput: "file", Data: map[string]any{"V": "/path/to/file"}},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}

func TestOsDir(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ osDir "" }}`, ExpectedOutput: "."},
		{Name: "TestRootPath", Input: `{{ osDir "/" }}`, ExpectedOutput: "/"},
		{Name: "TestWithoutExtension", Input: `{{ osDir "/path/to/file" }}`, ExpectedOutput: "/path/to"},
		{Name: "TestWithFileInput", Input: `{{ osDir "/path/to/file.txt" }}`, ExpectedOutput: "/path/to"},
		{Name: "TestPipeSyntax", Input: `{{ "/path/to/file.txt" | osDir }}`, ExpectedOutput: "/path/to"},
		{Name: "TestVariableInput", Input: `{{ .V | osDir }}`, ExpectedOutput: "/path/to", Data: map[string]any{"V": "/path/to/file"}},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}

func TestOsExt(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ osExt "" }}`, ExpectedOutput: ""},
		{Name: "TestRootPath", Input: `{{ osExt "/" }}`, ExpectedOutput: ""},
		{Name: "TestWithoutExtension", Input: `{{ osExt "/path/to/file" }}`, ExpectedOutput: ""},
		{Name: "TestWithFileInput", Input: `{{ osExt "/path/to/file.txt" }}`, ExpectedOutput: ".txt"},
		{Name: "TestPipeSyntax", Input: `{{ "/path/to/file.txt" | osExt }}`, ExpectedOutput: ".txt"},
		{Name: "TestVariableInput", Input: `{{ .V | osExt }}`, ExpectedOutput: ".txt", Data: map[string]any{"V": "/path/to/file.txt"}},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}

func TestOsClean(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ osClean "" }}`, ExpectedOutput: "."},
		{Name: "TestRootPath", Input: `{{ osClean "/" }}`, ExpectedOutput: "/"},
		{Name: "TestWithoutExtension", Input: `{{ osClean "/path///to/file" }}`, ExpectedOutput: "/path/to/file"},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}

func TestOsIsAbs(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ osIsAbs "" }}`, ExpectedOutput: "false"},
		{Name: "TestRootPath", Input: `{{ osIsAbs "/" }}`, ExpectedOutput: "true"},
		{Name: "TestRelativePath", Input: `{{ osIsAbs "path/to/file" }}`, ExpectedOutput: "false"},
		{Name: "TestAbsolutePath", Input: `{{ osIsAbs "/path/to/file.txt" }}`, ExpectedOutput: "true"},
		{Name: "TestPipeSyntax", Input: `{{ "file.txt" | osIsAbs }}`, ExpectedOutput: "false"},
		{Name: "TestVariableInput", Input: `{{ osIsAbs .V }}`, ExpectedOutput: "true", Data: map[string]any{"V": "/path/to/file"}},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}
