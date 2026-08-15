package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/go-sprout/sprout"
	"github.com/go-sprout/sprout/group/all"
	"github.com/go-sprout/sprout/registry/regex"
)

var sproutHandler = sprout.New(sprout.WithGroups(all.RegistryGroup()))

// dedicatedHandlers associates a documentation file with its own handler. It is
// needed for registries sharing their function names with a registry of the
// `all` group, like `regex` and `regexp`, which cannot be registered together.
// The dedicated registry is registered first so its functions take precedence.
var dedicatedHandlers = map[string]*sprout.DefaultHandler{
	filepath.Join("docs", "registries", "regex.md"): sprout.New(
		sprout.WithRegistries(regex.NewRegistry()),
		sprout.WithGroups(all.RegistryGroup()),
	),
}

// handlerFor returns the handler to use to validate the examples of the given
// documentation file, falling back to the default one built from all registries.
func handlerFor(file string) *sprout.DefaultHandler {
	if handler, ok := dedicatedHandlers[filepath.Clean(file)]; ok {
		return handler
	}
	return sproutHandler
}

// processExample compiles and executes a single template example.
// It checks the execution result against the expected output, and returns an error if they don't match.
// This function is used to verify that code examples in the documentation produce the expected output when executed.
func processExample(example Example) error {
	if example.Skipped {
		return nil
	}

	// Build the template with custom functions
	tmpl, err := template.New("example").Funcs(handlerFor(example.File).Build()).Parse(example.Code)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	// Check if the expected output is 'Error'
	isExpectedError := strings.EqualFold(strings.TrimSpace(example.Expected), "Error")

	// Execute the template
	var builder strings.Builder
	err = tmpl.Execute(&builder, map[string]any{
		"Nil":    nil,
		"Struct": struct{ V string }{"value"},
		"SecondStruct": struct {
			A int
			B string
		}{0, "second"},
	})

	if isExpectedError {
		if err != nil {
			// We expected an error and got an error
			return nil
		} else {
			// We expected an error but got none
			return fmt.Errorf("expected an error but the template executed successfully")
		}
	} else {
		if err != nil {
			// We did not expect an error but got one
			return fmt.Errorf("unexpected error during template execution: %w", err)
		}
		output := builder.String()
		// Compare output with expected output
		if output != example.Expected {
			return fmt.Errorf("output mismatch. Expected: %s Got: %s", example.Expected, output)
		}
	}

	return nil
}
