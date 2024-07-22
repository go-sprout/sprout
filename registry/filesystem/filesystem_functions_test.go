package filesystem_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/filesystem"
)

func TestPathBase(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ pathBase "" }}`, Expected: "."},
		{Name: "TestRootPath", Input: `{{ pathBase "/" }}`, Expected: "/"},
		{Name: "TestWithoutExtension", Input: `{{ pathBase "/path/to/file" }}`, Expected: "file"},
		{Name: "TestWithFileInput", Input: `{{ pathBase "/path/to/file.txt" }}`, Expected: "file.txt"},
		{Name: "TestPipeSyntax", Input: `{{ "/path/to/file.txt" | pathBase }}`, Expected: "file.txt"},
		{Name: "TestVariableInput", Input: `{{ .V | pathBase }}`, Expected: "file", Data: map[string]any{"V": "/path/to/file"}},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}

func TestPathDir(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ pathDir "" }}`, Expected: "."},
		{Name: "TestRootPath", Input: `{{ pathDir "/" }}`, Expected: "/"},
		{Name: "TestWithoutExtension", Input: `{{ pathDir "/path/to/file" }}`, Expected: "/path/to"},
		{Name: "TestWithFileInput", Input: `{{ pathDir "/path/to/file.txt" }}`, Expected: "/path/to"},
		{Name: "TestPipeSyntax", Input: `{{ "/path/to/file.txt" | pathDir }}`, Expected: "/path/to"},
		{Name: "TestVariableInput", Input: `{{ .V | pathDir }}`, Expected: "/path/to", Data: map[string]any{"V": "/path/to/file"}},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}

func TestPathExt(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ pathExt "" }}`, Expected: ""},
		{Name: "TestRootPath", Input: `{{ pathExt "/" }}`, Expected: ""},
		{Name: "TestWithoutExtension", Input: `{{ pathExt "/path/to/file" }}`, Expected: ""},
		{Name: "TestWithFileInput", Input: `{{ pathExt "/path/to/file.txt" }}`, Expected: ".txt"},
		{Name: "TestPipeSyntax", Input: `{{ "/path/to/file.txt" | pathExt }}`, Expected: ".txt"},
		{Name: "TestVariableInput", Input: `{{ .V | pathExt }}`, Expected: ".txt", Data: map[string]any{"V": "/path/to/file.txt"}},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}

func TestPathClean(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ pathClean "" }}`, Expected: "."},
		{Name: "TestRootPath", Input: `{{ pathClean "/" }}`, Expected: "/"},
		{Name: "TestWithoutExtension", Input: `{{ pathClean "/path/to/file" }}`, Expected: "/path/to/file"},
		{Name: "TestWithFileInput", Input: `{{ pathClean "/path/to/file.txt" }}`, Expected: "/path/to/file.txt"},
		{Name: "TestPipeSyntax", Input: `{{ "/path/to/file.txt" | pathClean }}`, Expected: "/path/to/file.txt"},
		{Name: "TestVariableInput", Input: `{{ .V | pathClean }}`, Expected: "/path/to/file", Data: map[string]any{"V": "/path/to/file"}},
		{Name: "TestDoubleSlash", Input: `{{ pathClean "/path//to/file" }}`, Expected: "/path/to/file"},
		{Name: "TestDotSlash", Input: `{{ pathClean "/path/./to/file" }}`, Expected: "/path/to/file"},
		{Name: "TestDotDotSlash", Input: `{{ pathClean "/path/../to/file" }}`, Expected: "/to/file"},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}

func TestPathIsAbs(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ pathIsAbs "" }}`, Expected: "false"},
		{Name: "TestRootPath", Input: `{{ pathIsAbs "/" }}`, Expected: "true"},
		{Name: "TestRelativePath", Input: `{{ pathIsAbs "path/to/file" }}`, Expected: "false"},
		{Name: "TestAbsolutePath", Input: `{{ pathIsAbs "/path/to/file.txt" }}`, Expected: "true"},
		{Name: "TestPipeSyntax", Input: `{{ "file.txt" | pathIsAbs }}`, Expected: "false"},
		{Name: "TestVariableInput", Input: `{{ pathIsAbs .V }}`, Expected: "true", Data: map[string]any{"V": "/path/to/file"}},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}
