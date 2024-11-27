package main

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/go-sprout/sprout"
	"github.com/go-sprout/sprout/group/all"
)

var sproutHandler = sprout.New(sprout.WithGroups(all.RegistryGroup()))

// processExample compiles and executes a single template example.
// It checks the execution result against the expected output, and returns an error if they don't match.
// This function is used to verify that code examples in the documentation produce the expected output when executed.
func processExample(example Example) error {
	if example.Skipped {
		return nil
	}

	// Build the template with custom functions
	tmpl, err := template.New("example").Funcs(sproutHandler.Build()).Parse(example.Code)
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
