// Package main provides a utility to process markdown documentation files,
// extracts code examples, executes them, and verifies their outputs.
package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strings"
	"time"
)

// Example represents an individual code example extracted from the markdown files.
// It contains the function name, the file it was extracted from, the code snippet, and the expected output.
type Example struct {
	FuncName string
	File     string
	Code     string
	Expected string
	Skipped  bool
}

var logLevel string

// init initializes the logging level from command-line flags.
// It sets the default logging level to "info", but can be overridden by
// specifying a different level (debug, info, warn, error) using the -log flag.
func init() {
	flag.StringVar(&logLevel, "log", "info", "Set the logging level (debug, info, warn, error)")
}

// main is the entry point of the application that processes markdown documentation files
// to extract, execute, and verify code examples. It performs the following steps:
// 1. Parses command-line flags to set the logging level.
// 2. Sets the log level based on the flag value and forces the timezone to UTC.
// 3. Discovers markdown files in the 'docs/registries' directory, filtering out specific files.
// 4. Iterates over each file, extracting code examples and processing them.
// 5. For each example, logs the processing status and counts errors.
// 6. Exits with an error code if any examples fail processing, or with success if all pass.
func main() {
	// Parse command-line flags
	flag.Parse()

	// Create an instance of the real file system
	fs := &OSFileSystem{}

	// Call the run function
	if err := run(fs); err != nil {
		// If an error occurs, log it and exit with code 1
		slog.Error("An error occurred during validation", "error", err)
		os.Exit(1)
	}
	// Exit successfully
	os.Exit(0)
}

// run contains the main logic of the program.
// It processes markdown files, extracts code examples, executes them, and verifies their outputs.
// This function returns an error if any step fails, making it easier to test.
func run(fs FileSystem) error {
	// Set the log level based on the flag
	if err := setLogLevel(logLevel); err != nil {
		return err
	}

	// Force timezone to UTC (for date parsing)
	time.Local = time.UTC

	// Discover markdown files in the docs directory
	files, err := fs.Glob("docs/registries/*.md")
	if err != nil {
		return fmt.Errorf("error listing markdown files: %w", err)
	}

	// Remove specified files from the list of files
	files = filterFiles(files, []string{
		"list-of-all-registries.md",
		"registryname.md",
		"backward.md",
		"crypto.md",
	})

	if len(files) == 0 {
		slog.Error("No Markdown files found in docs directory.")
		return fmt.Errorf("no Markdown files found in docs directory")
	}

	var examples []Example

	// Process each markdown file
	for _, file := range files {
		slog.Info("Processing file", "path", file)
		content, err := fs.ReadFile(file)
		if err != nil {
			return fmt.Errorf("error reading file %s: %w", file, err)
		}

		fileExamples := extractExamples(string(content), file)
		if len(fileExamples) == 0 {
			return fmt.Errorf("no code examples found in file %s", file)
		}

		examples = append(examples, fileExamples...)
	}

	// Process each example
	var errCount int
	for _, example := range examples {
		slog.Info("Processing example", "function", example.FuncName, "file", example.File, "code", example.Code)
		err := processExample(example)
		if err != nil {
			errCount++
			slog.Error("Error processing example", "function", example.FuncName, "file", example.File, "code", example.Code, "error", err)
		}
	}

	if errCount > 0 {
		slog.Error("Failed to process some examples", "count", errCount)
		return fmt.Errorf("failed to process %d examples", errCount)
	}
	slog.Info("All examples processed successfully")
	return nil
}

// filterFiles filters out unwanted files from the provided list of files.
// It returns a new slice containing only the files that are not in the 'remove' list.
// This function is used to exclude certain markdown files that should not be processed.
func filterFiles(files []string, remove []string) []string {
	return slices.DeleteFunc(files, func(file string) bool {
		for _, rem := range remove {
			if strings.Contains(file, rem) {
				return true
			}
		}
		return false
	})
}

// setLogLevel sets the log level to the specified level. It returns an error if the level is
// invalid. This function is used to set the log level from the command-line flag.
func setLogLevel(levelString string) error {
	var level slog.Level
	switch strings.ToLower(levelString) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		return fmt.Errorf("invalid log level: %s", levelString)
	}
	slog.SetLogLoggerLevel(level)
	return nil
}
