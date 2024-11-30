package main

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// MockFileSystem is a mock implementation of the FileSystem interface for testing purposes.
type MockFileSystem struct {
	Files map[string]string
}

func (m *MockFileSystem) Glob(pattern string) ([]string, error) {
	var matched []string
	for file := range m.Files {
		matched = append(matched, file)
		if strings.Contains(file, "error_glob") {
			return nil, errors.New("error matching files")
		}
	}
	return matched, nil
}

func (m *MockFileSystem) ReadFile(filename string) ([]byte, error) {
	content, ok := m.Files[filename]
	if !ok {
		return nil, errors.New("file not found")
	}
	if strings.Contains(filename, "error_read") {
		return nil, errors.New("error reading file")
	}

	return []byte(content), nil
}

// TestFilterFiles tests that filterFiles correctly removes specified files from the list.
func TestFilterFiles(t *testing.T) {
	files := []string{
		"docs/registries/example1.md",
		"docs/registries/list-of-all-registries.md",
		"docs/registries/example2.md",
		"docs/registries/backward.md",
	}
	remove := []string{
		"list-of-all-registries.md",
		"backward.md",
	}
	expected := []string{
		"docs/registries/example1.md",
		"docs/registries/example2.md",
	}
	result := filterFiles(files, remove)
	assert.Equal(t, expected, result)
}

func TestRun_ErrorSettingLogLevel(t *testing.T) {
	logLevel = "invalid"
	defer func() { logLevel = "info" }()

	fs := &MockFileSystem{
		Files: map[string]string{},
	}

	err := run(fs)

	assert.ErrorContains(t, err, "invalid log level: invalid")
}

// TestRun_NoFiles tests run function when no markdown files are found.
func TestRun_NoFiles(t *testing.T) {
	fs := &MockFileSystem{
		Files: map[string]string{},
	}

	err := run(fs)
	assert.ErrorContains(t, err, "no Markdown files found in docs directory")
}

func TestRun_ErrorFetchingFiles(t *testing.T) {
	fs := &MockFileSystem{
		Files: map[string]string{
			"docs/registries/none/error_glob.md": "",
		},
	}

	err := run(fs)
	assert.ErrorContains(t, err, "error listing markdown files")
}

func TestRun_ErrorReadingFile(t *testing.T) {
	fs := &MockFileSystem{
		Files: map[string]string{
			"docs/registries/none/error_read.md": "",
		},
	}

	err := run(fs)
	assert.ErrorContains(t, err, "error reading file")
}

func TestRun_ErrorNoExamples(t *testing.T) {
	fs := &MockFileSystem{
		Files: map[string]string{
			"docs/registries/example1.md": "",
		},
	}

	err := run(fs)
	assert.ErrorContains(t, err, "no code examples found in file")
}

// TestRun_AllExamplesProcessed tests run function when all examples are processed successfully.
func TestRun_AllExamplesProcessed(t *testing.T) {
	fs := &MockFileSystem{
		Files: map[string]string{
			"docs/registries/example1.md": `
### <mark>FunctionName</mark>

{% tab title="Template Example" %}
` + "```go" + `
{{ "Hello, World!" }}
// Output: Hello, World!
` + "```" + `
{% endtab %}
`,
		},
	}

	err := run(fs)
	assert.NoError(t, err)
}

// TestRun_FailedExamples tests run function when some examples fail to process.
func TestRun_FailedExamples(t *testing.T) {
	fs := &MockFileSystem{
		Files: map[string]string{
			"docs/registries/example1.md": `
### <mark>FunctionName</mark>

{% tab title="Template Example" %}
` + "```go" + `
{{ "Hello, World!" }} // Output: Hello, Universe!
` + "```" + `
{% endtab %}
`,
		},
	}

	err := run(fs)
	assert.ErrorContains(t, err, "failed to process 1 examples")
}

// TestRun_InvalidLogLevel tests run function with an invalid log level.
func TestSetLogLevel_InvalidLogLevel(t *testing.T) {
	err := setLogLevel("invalid")
	require.Error(t, err)
	assert.ErrorContains(t, err, "invalid log level")
}

// TestSetLogLevel_ValidLogLevel tests run function with a valid log level.
func TestSetLogLevel_ValidLogLevel(t *testing.T) {
	loglevels := []string{"debug", "info", "warn", "error"}
	for _, level := range loglevels {
		err := setLogLevel(level)
		assert.NoError(t, err)
	}
}
