package filesystem_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/filesystem"
)

func TestOsBase(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ osBase "" }}`, Expected: "."},
		{Name: "TestRootPath", Input: `{{ osBase "D:\\" }}`, Expected: "\\"},
		{Name: "TestWithoutExtension", Input: `{{ osBase "D:\\path\\to\\file" }}`, Expected: "file"},
		{Name: "TestWithFileInput", Input: `{{ osBase "D:\\path\\to\\file.txt" }}`, Expected: "file.txt"},
		{Name: "TestPipeSyntax", Input: `{{ "D:\\path\\to\\file.txt" | osBase }}`, Expected: "file.txt"},
		{Name: "TestVariableInput", Input: `{{ .V | osBase }}`, Expected: "file", Data: map[string]any{"V": "\\path\\to\\file"}},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}

func TestOsDir(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ osDir "" }}`, Expected: "."},
		{Name: "TestRootPath", Input: `{{ osDir "D:\\" }}`, Expected: "D:\\"},
		{Name: "TestWithoutExtension", Input: `{{ osDir "D:\\path\\to\\file" }}`, Expected: "D:\\path\\to"},
		{Name: "TestWithFileInput", Input: `{{ osDir "D:\\path\\to\\file.txt" }}`, Expected: "D:\\path\\to"},
		{Name: "TestPipeSyntax", Input: `{{ "D:\\path\\to\\file.txt" | osDir }}`, Expected: "D:\\path\\to"},
		{Name: "TestVariableInput", Input: `{{ .V | osDir }}`, Expected: "\\path\\to", Data: map[string]any{"V": "\\path\\to\\file"}},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}

func TestOsExt(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ osExt "" }}`, Expected: ""},
		{Name: "TestRootPath", Input: `{{ osExt "\\" }}`, Expected: ""},
		{Name: "TestWithoutExtension", Input: `{{ osExt "D:\\path\\to\\file" }}`, Expected: ""},
		{Name: "TestWithFileInput", Input: `{{ osExt "D:\\path\\to\\file.txt" }}`, Expected: ".txt"},
		{Name: "TestPipeSyntax", Input: `{{ "D:\\path\\to\\file.txt" | osExt }}`, Expected: ".txt"},
		{Name: "TestVariableInput", Input: `{{ .V | osExt }}`, Expected: ".txt", Data: map[string]any{"V": "D:\\path\\to\\file.txt"}},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}

func TestOsClean(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ osClean "" }}`, Expected: "."},
		{Name: "TestRootPath", Input: `{{ osClean "D:\\" }}`, Expected: "D:\\"},
		{Name: "TestWithoutExtension", Input: `{{ osClean "D:\\path\\\\to\\file" }}`, Expected: "D:\\path\\to\\file"},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}

func TestOsIsAbs(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmptyPath", Input: `{{ osIsAbs "" }}`, Expected: "false"},
		{Name: "TestRootPath", Input: `{{ osIsAbs "D:\\" }}`, Expected: "true"},
		{Name: "TestRelativePath", Input: `{{ osIsAbs "path\\to\\file" }}`, Expected: "false"},
		{Name: "TestAbsolutePath", Input: `{{ osIsAbs "D:\\path\\to\\file.txt" }}`, Expected: "true"},
		{Name: "TestPipeSyntax", Input: `{{ "file.txt" | osIsAbs }}`, Expected: "false"},
		{Name: "TestVariableInput", Input: `{{ osIsAbs .V }}`, Expected: "true", Data: map[string]any{"V": "D:\\path\\to\\file"}},
	}

	pesticide.RunTestCases(t, filesystem.NewRegistry(), tc)
}
