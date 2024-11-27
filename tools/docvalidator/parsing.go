package main

import (
	"log/slog"
	"regexp"
	"strconv"
	"strings"
)

// extractExamples scans the content of a markdown file and extracts code examples.
// It returns a slice of Example structs, each representing a code example found in the content.
// This function is used to parse markdown files and find code examples under specific headings.
func extractExamples(content, filePath string) []Example {
	var examples []Example

	// Regex to find function names (assuming they are in headings)
	headingRe := regexp.MustCompile(`(?m)^###\s+<mark[^>]*>([^<]+)</mark>`)

	// Find all headings with function names
	headings := headingRe.FindAllStringSubmatchIndex(content, -1)

	if len(headings) == 0 {
		return examples
	}

	// Corrected regex pattern to match code blocks under 'Template Example' tabs
	codePattern := `(?s){% tab title="Template Example" %}.*?` + "```go(.*?)```" + `.*?{% endtab %}`
	codeRe := regexp.MustCompile(codePattern)

	for i, heading := range headings {
		funcName := content[heading[2]:heading[3]]
		startIdx := heading[1] // end of the heading

		var endIdx int
		if i+1 < len(headings) {
			endIdx = headings[i+1][0] // start of next heading
		} else {
			endIdx = len(content)
		}

		section := content[startIdx:endIdx]

		// Find code examples in this section
		codeMatches := codeRe.FindAllStringSubmatch(section, -1)

		for _, codeMatch := range codeMatches {
			codeBlock := strings.TrimSpace(codeMatch[1])

			// Parse the code block to extract multiple examples
			examplesInBlock := parseCodeBlock(codeBlock, funcName, filePath)
			examples = append(examples, examplesInBlock...)
		}
	}

	return examples
}

// parseCodeBlock processes a code block and extracts individual code examples.
// It handles code examples that may have multiple code snippets and their expected outputs.
// This function is used to split code blocks into individual examples, especially when multiple examples are in the same code block.
func parseCodeBlock(codeBlock, funcName, filePath string) []Example {
	var examples []Example
	lines := strings.Split(codeBlock, "\n")

	i := 0
	for i < len(lines) {
		codeBuffer := []string{}
		expectedOutput := ""

		// Collect code lines
		for i < len(lines) {
			line := lines[i]
			trimmedLine := strings.TrimSpace(line)

			switch {
			case strings.Contains(trimmedLine, "// Output:"):
				parts := strings.SplitN(line, "// Output:", 2)
				codeLine := strings.TrimSpace(parts[0])
				outputPart := strings.TrimSpace(parts[1])
				expectedOutput = outputPart
				if codeLine != "" {
					codeBuffer = append(codeBuffer, codeLine)
				}
			case strings.Contains(trimmedLine, "// Error"):
				parts := strings.SplitN(line, "// Error", 2)
				codeLine := strings.TrimSpace(parts[0])
				expectedOutput = "Error"
				if codeLine != "" {
					codeBuffer = append(codeBuffer, codeLine)
				}
			case strings.Contains(trimmedLine, "(will be different)") || strings.Contains(trimmedLine, "human readable"):
				parts := strings.SplitN(line, "// Output(will be different):", 2)
				codeLine := strings.TrimSpace(parts[0])
				expectedOutput = "(will be different)"
				if codeLine != "" {
					codeBuffer = append(codeBuffer, codeLine)
				}
			default:
				codeBuffer = append(codeBuffer, line)
				i++
				continue
			}

			// If expected output is empty or continues on the next lines
			if expectedOutput == "" && i+1 < len(lines) {
				i++
				outputLines := []string{}
				for i < len(lines) {
					nextLine := strings.TrimSpace(lines[i])
					if nextLine == "" || strings.HasPrefix(nextLine, "{{") || strings.Contains(nextLine, "// Output:") || strings.Contains(nextLine, "// Error") {
						i--
						break
					}
					outputLines = append(outputLines, nextLine)
					i++
				}
				expectedOutput = strings.Join(outputLines, "\n")
			}

			if expectedOutput != "\"\"" {
				unescaped, err := strconv.Unquote(`"` + strings.Trim(expectedOutput, `"'`) + `"`)
				if err != nil {
					slog.Warn("Error unquoting string", "error", err, "string", expectedOutput)
				} else {
					expectedOutput = unescaped
				}
			}

			i++
			break
		}

		code := strings.Join(codeBuffer, "\n")
		code = strings.TrimSpace(code)

		// Skip examples with no code or expected output
		skip := false
		if expectedOutput == "(will be different)" {
			skip = true
		}

		if !skip && (code == "" || expectedOutput == "") {
			slog.Warn("Incomplete example", "func", funcName, "file", filePath)
		} else {
			if expectedOutput == "\"\"" {
				expectedOutput = ""
			}

			example := Example{
				FuncName: funcName,
				File:     filePath,
				Code:     code,
				Expected: expectedOutput,
				Skipped:  skip,
			}
			examples = append(examples, example)
		}
	}

	return examples
}
