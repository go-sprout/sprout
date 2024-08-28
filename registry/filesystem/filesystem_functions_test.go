package filesystem_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/filesystem"
)

func TestPathBase(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ pathBase "" }}`, ExpectedOutput: "."},
		{Name: "TestRootPath", Input: `{{ pathBase "/" }}`, ExpectedOutput: "/"},
		{Name: "TestWithoutExtension", Input: `{{ pathBase "/path/to/file" }}`, ExpectedOutput: "file"},
		{Name: "TestWithFileInput", Input: `{{ pathBase "/path/to/file.txt" }}`, ExpectedOutput: "file.txt"},
		{Name: "TestPipeSyntax", Input: `{{ "/path/to/file.txt" | pathBase }}`, ExpectedOutput: "file.txt"},
		{Name: "TestVariableInput", Input: `{{ .V | pathBase }}`, ExpectedOutput: "file", Data: map[string]any{"V": "/path/to/file"}},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}

func TestPathDir(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ pathDir "" }}`, ExpectedOutput: "."},
		{Name: "TestRootPath", Input: `{{ pathDir "/" }}`, ExpectedOutput: "/"},
		{Name: "TestWithoutExtension", Input: `{{ pathDir "/path/to/file" }}`, ExpectedOutput: "/path/to"},
		{Name: "TestWithFileInput", Input: `{{ pathDir "/path/to/file.txt" }}`, ExpectedOutput: "/path/to"},
		{Name: "TestPipeSyntax", Input: `{{ "/path/to/file.txt" | pathDir }}`, ExpectedOutput: "/path/to"},
		{Name: "TestVariableInput", Input: `{{ .V | pathDir }}`, ExpectedOutput: "/path/to", Data: map[string]any{"V": "/path/to/file"}},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}

func TestPathExt(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ pathExt "" }}`, ExpectedOutput: ""},
		{Name: "TestRootPath", Input: `{{ pathExt "/" }}`, ExpectedOutput: ""},
		{Name: "TestWithoutExtension", Input: `{{ pathExt "/path/to/file" }}`, ExpectedOutput: ""},
		{Name: "TestWithFileInput", Input: `{{ pathExt "/path/to/file.txt" }}`, ExpectedOutput: ".txt"},
		{Name: "TestPipeSyntax", Input: `{{ "/path/to/file.txt" | pathExt }}`, ExpectedOutput: ".txt"},
		{Name: "TestVariableInput", Input: `{{ .V | pathExt }}`, ExpectedOutput: ".txt", Data: map[string]any{"V": "/path/to/file.txt"}},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}

func TestPathClean(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ pathClean "" }}`, ExpectedOutput: "."},
		{Name: "TestRootPath", Input: `{{ pathClean "/" }}`, ExpectedOutput: "/"},
		{Name: "TestWithoutExtension", Input: `{{ pathClean "/path/to/file" }}`, ExpectedOutput: "/path/to/file"},
		{Name: "TestWithFileInput", Input: `{{ pathClean "/path/to/file.txt" }}`, ExpectedOutput: "/path/to/file.txt"},
		{Name: "TestPipeSyntax", Input: `{{ "/path/to/file.txt" | pathClean }}`, ExpectedOutput: "/path/to/file.txt"},
		{Name: "TestVariableInput", Input: `{{ .V | pathClean }}`, ExpectedOutput: "/path/to/file", Data: map[string]any{"V": "/path/to/file"}},
		{Name: "TestDoubleSlash", Input: `{{ pathClean "/path//to/file" }}`, ExpectedOutput: "/path/to/file"},
		{Name: "TestDotSlash", Input: `{{ pathClean "/path/./to/file" }}`, ExpectedOutput: "/path/to/file"},
		{Name: "TestDotDotSlash", Input: `{{ pathClean "/path/../to/file" }}`, ExpectedOutput: "/to/file"},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}

func TestPathIsAbs(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ pathIsAbs "" }}`, ExpectedOutput: "false"},
		{Name: "TestRootPath", Input: `{{ pathIsAbs "/" }}`, ExpectedOutput: "true"},
		{Name: "TestRelativePath", Input: `{{ pathIsAbs "path/to/file" }}`, ExpectedOutput: "false"},
		{Name: "TestAbsolutePath", Input: `{{ pathIsAbs "/path/to/file.txt" }}`, ExpectedOutput: "true"},
		{Name: "TestPipeSyntax", Input: `{{ "file.txt" | pathIsAbs }}`, ExpectedOutput: "false"},
		{Name: "TestVariableInput", Input: `{{ pathIsAbs .V }}`, ExpectedOutput: "true", Data: map[string]any{"V": "/path/to/file"}},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}
