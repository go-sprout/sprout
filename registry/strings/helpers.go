package strings

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// ellipsis truncates 'value' from both ends, preserving the middle part of
// the string and appending ellipses to both ends if needed.
//
// Parameters:
//
//	offset int - starting position for preserving text.
//	maxWidth int - the maximum width of the string including the ellipsis.
//	value string - the string to truncate.
//
// Returns:
//
//	string - the possibly truncated string with an ellipsis.
func (sr *StringsRegistry) ellipsis(value string, offset int, maxWidth int) string {
	ellipsis := "..."
	// Return the original string if maxWidth is less than 4, or the offset
	// create exclusive dot string,  it's not possible to add an ellipsis.
	if maxWidth < 4 || offset > 0 && maxWidth < 7 {
		return value
	}

	runeCount := utf8.RuneCountInString(value)

	// If the string doesn't need trimming, return it as is.
	if runeCount <= maxWidth || runeCount <= offset {
		return value[offset:]
	}

	// Determine end position for the substring, ensuring room for the ellipsis.
	endPos := offset + maxWidth - 3 // 3 is for the length of the ellipsis
	if offset > 0 {
		endPos -= 3 // remove the left ellipsis
	}

	// Convert the string to a slice of runes to properly handle multi-byte characters.
	runes := []rune(value)

	// Return the substring with an ellipsis, directly constructing the string in the return statement.
	if offset > 0 {
		return ellipsis + string(runes[offset:endPos]) + ellipsis
	}
	return string(runes[offset:endPos]) + ellipsis
}

// initials extracts the initials from 'value', using 'delimiters' to determine
// word boundaries.
//
// Parameters:
//
//	value string - the string from which to extract initials.
//	delimiters string - the string containing delimiter characters.
//
// Returns:
//
//	string - the initials of the words in 'value'.
func (sr *StringsRegistry) initials(value string, delimiters string) string {
	// Define a function to determine if a rune is a delimiter.
	isDelimiter := func(r rune) bool {
		return strings.ContainsRune(delimiters, r)
	}

	words := strings.FieldsFunc(value, isDelimiter)
	runes := make([]rune, len(words))
	for i, word := range strings.FieldsFunc(value, isDelimiter) {
		if i == 0 || unicode.IsLetter(rune(word[0])) {
			runes[i] = rune(word[0])
		}
	}

	return string(runes)
}

// transformString modifies the string 'value' based on various case styling rules
// specified in the 'style' parameter. It can capitalize, lowercase, and insert
// separators according to the rules provided.
//
// Parameters:
//
//	style caseStyle - a struct specifying how to transform the string, including
//	                  capitalization rules, insertion of separators, and whether
//	                  to force lowercase.
//	value string - the string to transform.
//
// Returns:
//
//	string - the transformed string.
func (sr *StringsRegistry) transformString(style caseStyle, value string) string {
	var result strings.Builder
	result.Grow(len(value) + 10) // Allocate a bit more for potential separators

	capitalizeNext := style.CapitalizeNext
	var lastRune, lastLetter, nextRune rune

	if !style.CapitalizeFirst {
		capitalizeNext = false
	}

	for i, r := range value {
		if i+1 < len(value) {
			nextRune = rune(value[i+1])
		} else {
			nextRune = 0
		}

		if r == ' ' || r == '-' || r == '_' {
			if style.Separator != -1 && (lastRune != style.Separator) {
				result.WriteRune(style.Separator)
			}
			if lastLetter != 0 {
				capitalizeNext = true
			}

			lastRune = style.Separator
			continue
		}

		if unicode.IsUpper(r) && result.Len() > 0 && lastRune != style.Separator {
			// Detect word boundaries: transition from lowercase to uppercase (e.g., "firstName")
			// or end of an acronym (e.g., "HTMLParser" where 'P' follows 'L' which follows uppercase).
			var isWordBoundary bool

			if style.Separator != -1 {
				// For styles with separators (snake_case, kebab-case, etc.), use standard detection
				isWordBoundary = unicode.IsLower(lastRune) ||
					(unicode.IsUpper(lastRune) && unicode.IsLower(nextRune))
			} else {
				// For styles without separators (camelCase, PascalCase), detect word boundaries:
				// 1. Lowercase to uppercase transition where the word continues (letter/digit follows)
				//    e.g., "firstName" → "first" + "Name", "NoHTTPS" → "No" + "Https"
				// 2. End of acronym (uppercase to uppercase to lowercase)
				//    e.g., "HTMLParser" → "Html" + "Parser"
				// Exclude stray uppercase at end of word (e.g., "FoO" should not split)
				nextIsContent := unicode.IsLetter(nextRune) || unicode.IsDigit(nextRune)
				isWordBoundary = (unicode.IsLower(lastRune) && nextIsContent) ||
					(unicode.IsUpper(lastRune) && unicode.IsLower(nextRune))
			}

			if isWordBoundary {
				if style.Separator != -1 {
					result.WriteRune(style.Separator)
				}
				// Mark the start of a new word for capitalization
				capitalizeNext = true
			}
		}

		if style.Separator != -1 && result.Len() > 0 && lastRune != style.Separator && (unicode.IsDigit(r) && !unicode.IsDigit(lastRune)) {
			result.WriteRune(style.Separator)
		}

		switch {
		case capitalizeNext && style.CapitalizeNext:
			result.WriteRune(unicode.ToUpper(r))
			capitalizeNext = false
		case style.ForceLowercase:
			result.WriteRune(unicode.ToLower(r))
		case style.ForceUppercase:
			result.WriteRune(unicode.ToUpper(r))
		default:
			result.WriteRune(r)
		}

		lastRune = r // Update lastRune to the current rune
		if unicode.IsLetter(r) {
			lastLetter = r
		}
	}

	return result.String()
}

// populateMapWithParts converts an array of strings into a map with keys based
// on the index of each string.
//
// Parameters:
//
//	parts []string - the array of strings to be converted into a map.
//
// Returns:
//
//	map[string]string - a map where each key corresponds to an index (with an underscore prefix) of the string in the input array.
func (sr *StringsRegistry) populateMapWithParts(parts []string) map[string]string {
	res := make(map[string]string, len(parts))
	for i, v := range parts {
		res[fmt.Sprintf("_%d", i)] = v
	}
	return res
}

// convertIntArrayToString converts an array of integers into a single string
// with elements separated by a given delimiter.
//
// Parameters:
//
//	slice []int - the array of integers to convert.
//	delimiter string - the string to use as a delimiter between the integers in the output string.
//
// Returns:
//
//	string - the resulting string that concatenates all the integers in the array separated by the specified delimiter.
func (sr *StringsRegistry) convertIntArrayToString(slice []int, delimiter string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), delimiter), "[]")
}

// WordWrap formats 'value' into lines of maximum 'wrapLength', optionally wrapping
// long words and using 'newLineCharacter' for line breaks.
//
// Parameters:
//
//	value string - the string to wrap.
//	wrapLength int - the maximum length of each line.
//	newLineCharacter string - the string used to denote new lines.
//	wrapLongWords bool - true to wrap long words that exceed the line length.
//
// Returns:
//
//	string - the wrapped string.
func (sr *StringsRegistry) wordWrap(wrapLength int, newLineCharacter string, wrapLongWords bool, value string) string {
	if wrapLength < 1 {
		wrapLength = 1
	}
	if newLineCharacter == "" {
		newLineCharacter = "\n"
	}

	var resultBuilder strings.Builder
	var currentLineLength int

	for _, word := range strings.Fields(value) {
		wordLength := utf8.RuneCountInString(word)

		// If the word is too long and should be wrapped, or it fits in the remaining line length
		if currentLineLength > 0 && (currentLineLength+1+wordLength > wrapLength && !wrapLongWords || currentLineLength+1+wordLength > wrapLength) {
			resultBuilder.WriteString(newLineCharacter)
			currentLineLength = 0
		}

		if wrapLongWords && wordLength > wrapLength {
			for i, r := range word {
				resultBuilder.WriteRune(r)
				currentLineLength++
				// Avoid adding a new line immediately after wrapping a long word
				if i < len(word)-1 && currentLineLength == wrapLength {
					resultBuilder.WriteString(newLineCharacter)
					currentLineLength = 0
				}
			}
		} else {
			if currentLineLength > 0 {
				resultBuilder.WriteRune(' ')
				currentLineLength++
			}
			resultBuilder.WriteString(word)
			currentLineLength += wordLength
		}
	}

	return resultBuilder.String()
}

// swapFirstLetter swaps the first letter of the string 'value' to uppercase or
// lowercase. The casing is determined by the 'casing' parameter.
//
// Parameters:
//
//	value string - the string to modify.
//	shouldUppercaseFirst bool - the casing to apply to the first letter.
//
// Returns:
//
//	string - the modified string with the first letter in the desired casing.
func swapFirstLetter(value string, shouldUppercase bool) string {
	var conditionFunc func(r rune) bool
	var updateFunc func(r rune) rune

	if shouldUppercase {
		conditionFunc = unicode.IsUpper
		updateFunc = unicode.ToUpper
	} else {
		conditionFunc = unicode.IsLower
		updateFunc = unicode.ToLower
	}

	buf := []byte(value)
	for i := 0; i < len(buf); {
		r, size := utf8.DecodeRune(buf[i:])

		if unicode.IsLetter(r) {
			if conditionFunc(r) {
				return value
			}

			upperRune := updateFunc(r)
			utf8.EncodeRune(buf[i:i+size], upperRune)

			return string(buf)
		}

		i += size
	}
	return value
}
