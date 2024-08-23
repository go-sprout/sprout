package strings

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// ellipsis truncates 'str' from both ends, preserving the middle part of
// the string and appending ellipses to both ends if needed.
//
// Parameters:
//
//	offset int - starting position for preserving text.
//	maxWidth int - the maximum width of the string including the ellipsis.
//	str string - the string to truncate.
//
// Returns:
//
//	string - the possibly truncated string with an ellipsis.
func (sr *StringsRegistry) ellipsis(str string, offset int, maxWidth int) string {
	ellipsis := "..."
	// Return the original string if maxWidth is less than 4, or the offset
	// create exclusive dot string,  it's not possible to add an ellipsis.
	if maxWidth < 4 || offset > 0 && maxWidth < 7 {
		return str
	}

	runeCount := utf8.RuneCountInString(str)

	// If the string doesn't need trimming, return it as is.
	if runeCount <= maxWidth || runeCount <= offset {
		return str[offset:]
	}

	// Determine end position for the substring, ensuring room for the ellipsis.
	endPos := offset + maxWidth - 3 // 3 is for the length of the ellipsis
	if offset > 0 {
		endPos -= 3 // remove the left ellipsis
	}

	// Convert the string to a slice of runes to properly handle multi-byte characters.
	runes := []rune(str)

	// Return the substring with an ellipsis, directly constructing the string in the return statement.
	if offset > 0 {
		return ellipsis + string(runes[offset:endPos]) + ellipsis
	}
	return string(runes[offset:endPos]) + ellipsis
}

// initials extracts the initials from 'str', using 'delimiters' to determine
// word boundaries.
//
// Parameters:
//
//	str string - the string from which to extract initials.
//	delimiters string - the string containing delimiter characters.
//
// Returns:
//
//	string - the initials of the words in 'str'.
func (sr *StringsRegistry) initials(str string, delimiters string) string {
	// Define a function to determine if a rune is a delimiter.
	isDelimiter := func(r rune) bool {
		return strings.ContainsRune(delimiters, r)
	}

	words := strings.FieldsFunc(str, isDelimiter)
	var runes = make([]rune, len(words))
	for i, word := range strings.FieldsFunc(str, isDelimiter) {
		if i == 0 || unicode.IsLetter(rune(word[0])) {
			runes[i] = rune(word[0])
		}
	}

	return string(runes)
}

// transformString modifies the string 'str' based on various case styling rules
// specified in the 'style' parameter. It can capitalize, lowercase, and insert
// separators according to the rules provided.
//
// Parameters:
//
//	style caseStyle - a struct specifying how to transform the string, including
//	                  capitalization rules, insertion of separators, and whether
//	                  to force lowercase.
//	str string - the string to transform.
//
// Returns:
//
//	string - the transformed string.
//
// Example:
//
//	style := caseStyle{
//	    Separator:       '_',
//	    CapitalizeNext:  true,
//	    ForceLowercase:  false,
//	    InsertSeparator: true,
//	}
//	transformed := sr.transformString(style, "hello world")
//	Output: "Hello_World"
//
// Note:
//
//	This example demonstrates how to use the function to capitalize the first letter of
//	each word and insert underscores between words, which is common in identifiers like
//	variable names in programming.
func (sr *StringsRegistry) transformString(style caseStyle, str string) string {
	var result strings.Builder
	result.Grow(len(str) + 10) // Allocate a bit more for potential separators

	var capitalizeNext = style.CapitalizeNext
	var lastRune, lastLetter, nextRune rune = 0, 0, 0

	if !style.CapitalizeFirst {
		capitalizeNext = false
	}

	for i, r := range str {
		if i+1 < len(str) {
			nextRune = rune(str[i+1])
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

		if unicode.IsUpper(r) && style.Separator != -1 && result.Len() > 0 && lastRune != style.Separator {
			if (unicode.IsUpper(lastRune) && unicode.IsUpper(r) && unicode.IsLower(nextRune)) || (unicode.IsUpper(r) && unicode.IsLower(lastRune)) {
				result.WriteRune(style.Separator)
			}
		}

		if style.Separator != -1 && lastRune != style.Separator && (unicode.IsDigit(r) && !unicode.IsDigit(lastRune)) {
			result.WriteRune(style.Separator)
		}

		if capitalizeNext && style.CapitalizeNext {
			result.WriteRune(unicode.ToUpper(r))
			capitalizeNext = false
		} else if style.ForceLowercase {
			result.WriteRune(unicode.ToLower(r))
		} else if style.ForceUppercase {
			result.WriteRune(unicode.ToUpper(r))
		} else {
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
//
// Example:
//
//	parts := []string{"apple", "banana", "cherry"}
//	result := sr.populateMapWithParts(parts)
//	fmt.Println(result) // Output: {"_0": "apple", "_1": "banana", "_2": "cherry"}
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
//
// Example:
//
//	slice := []int{1, 2, 3, 4, 5}
//	result := sr.convertIntArrayToString(slice, ", ")
//	fmt.Println(result) // Output: "1, 2, 3, 4, 5"
func (sr *StringsRegistry) convertIntArrayToString(slice []int, delimeter string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), delimeter), "[]")
}

// WordWrap formats 'str' into lines of maximum 'wrapLength', optionally wrapping
// long words and using 'newLineCharacter' for line breaks.
//
// Parameters:
//
//	str string - the string to wrap.
//	wrapLength int - the maximum length of each line.
//	newLineCharacter string - the string used to denote new lines.
//	wrapLongWords bool - true to wrap long words that exceed the line length.
//
// Returns:
//
//	string - the wrapped string.
//
// Example:
//
//	wrapped := sr.wordWrap(10, "\n", true, "This is a long wordwrap example")
//	fmt.Println(wrapped)
func (sr *StringsRegistry) wordWrap(wrapLength int, newLineCharacter string, wrapLongWords bool, str string) string {
	if wrapLength < 1 {
		wrapLength = 1
	}
	if newLineCharacter == "" {
		newLineCharacter = "\n"
	}

	var resultBuilder strings.Builder
	var currentLineLength int

	for _, word := range strings.Fields(str) {
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

// swapFirstLetter swaps the first letter of the string 'str' to uppercase or
// lowercase. The casing is determined by the 'casing' parameter.
//
// Parameters:
//
//	str string - the string to modify.
//	shouldUppercaseFirst bool - the casing to apply to the first letter.
//
// Returns:
//
//	string - the modified string with the first letter in the desired casing.
//
// Example:
//
//	result := sr.swapFirstLetter("123hello", cassingUpper)
//	fmt.Println(result) // Output: "123Hello"
func swapFirstLetter(str string, shouldUppercase bool) string {
	var conditionFunc func(r rune) bool
	var updateFunc func(r rune) rune

	if shouldUppercase {
		conditionFunc = unicode.IsUpper
		updateFunc = unicode.ToUpper
	} else {
		conditionFunc = unicode.IsLower
		updateFunc = unicode.ToLower
	}

	buf := []byte(str)
	for i := 0; i < len(buf); {
		r, size := utf8.DecodeRune(buf[i:])

		if unicode.IsLetter(r) {
			if conditionFunc(r) {
				return str
			}

			upperRune := updateFunc(r)
			utf8.EncodeRune(buf[i:i+size], upperRune)

			return string(buf)
		}

		i += size
	}
	return str
}
