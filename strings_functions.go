package sprout

import (
	"fmt"
	mathrand "math/rand"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// caseStyle defines the rules for transforming strings based on capitalization,
// separator insertion, and case enforcement. This struct is typically used to
// configure functions that modify the case and formatting of strings to match
// specific coding or display conventions.
//
// Fields:
//
//	Separator rune - The character used to separate words in the transformed string.
//	                 For example, underscores (_) or hyphens (-).
//
//	CapitalizeNext bool - Determines if the first character of each word should be
//	                      capitalized in the output. Useful for TitleCase or CamelCase.
//
//	ForceLowercase bool - If set to true, all characters in the output are converted
//	                      to lowercase, overriding any capitalization rules.
//
// Usage:
//
// This struct can be used to configure string transformation functions, allowing for
// flexible adaptation to various text formatting needs, such as generating identifiers
// or user-friendly display text.
//
// Example:
//
//		style := caseStyle{
//		    Separator:       '_',
//		    CapitalizeNext:  false,
//		    ForceLowercase:  true,
//		    InsertSeparator: true,
//		}
//	 Use `style` to transform "ExampleText" to "example_text"
type caseStyle struct {
	Separator      rune // Character that separates words.
	CapitalizeNext bool // Whether to capitalize the first character of each word.
	ForceLowercase bool // Whether to force all characters to lowercase.
	ForceUppercase bool // Whether to force all characters to uppercase.
}

var (
	camelCaseStyle    = caseStyle{Separator: -1, CapitalizeNext: true, ForceLowercase: true}
	kebabCaseStyle    = caseStyle{Separator: '-', ForceLowercase: true}
	pascalCaseStyle   = caseStyle{Separator: -1, CapitalizeNext: true}
	snakeCaseStyle    = caseStyle{Separator: '_', ForceLowercase: true}
	dotCaseStyle      = caseStyle{Separator: '.', ForceLowercase: true}
	pathCaseStyle     = caseStyle{Separator: '/', ForceLowercase: true}
	constantCaseStyle = caseStyle{Separator: '_', ForceUppercase: true}
)

// Nospace removes all whitespace characters from the provided string.
// It uses the unicode package to identify whitespace runes and removes them.
//
// Parameters:
//
//	str string - the string from which to remove whitespace.
//
// Returns:
//
//	string - the modified string with all whitespace characters removed.
//
// Example:
//
//	{{ "Hello World" | nospace }} // Output: "HelloWorld"
func (fh *FunctionHandler) Nospace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

// Trim removes leading and trailing whitespace from the string.
//
// Parameters:
//
//	str string - the string to trim.
//
// Returns:
//
//	string - the trimmed string.
//
// Example:
//
//	{{ " Hello World " | trim }} // Output: "Hello World"
func (fh *FunctionHandler) Trim(str string) string {
	return strings.TrimSpace(str)
}

// TrimAll removes all occurrences of any characters in 'cutset' from both the
// beginning and the end of 'str'.
//
// Parameters:
//
//	cutset string - a string of characters to remove from the string.
//	str string - the string to trim.
//
// Returns:
//
//	string - the string with specified characters removed.
//
// Example:
//
//	{{ "xyzHelloxyz" | trimAll "xyz" }} // Output: "Hello"
func (fh *FunctionHandler) TrimAll(cutset string, str string) string {
	return strings.Trim(str, cutset)
}

// TrimPrefix removes the 'prefix' from the start of 'str' if present.
//
// Parameters:
//
//	prefix string - the prefix to remove.
//	str string - the string to trim.
//
// Returns:
//
//	string - the string with the prefix removed if it was present.
//
// Example:
//
//	{{ "HelloWorld" | trimPrefix "Hello" }} // Output: "World"
func (fh *FunctionHandler) TrimPrefix(prefix string, str string) string {
	return strings.TrimPrefix(str, prefix)
}

// TrimSuffix removes the 'suffix' from the end of 'str' if present.
//
// Parameters:
//
//	suffix string - the suffix to remove.
//	str string - the string to trim.
//
// Returns:
//
//	string - the string with the suffix removed if it was present.
//
// Example:
//
//	{{ "HelloWorld" | trimSuffix "World" }} // Output: "Hello"
func (fh *FunctionHandler) TrimSuffix(suffix string, str string) string {
	return strings.TrimSuffix(str, suffix)
}

// Contains checks if 'str' contains the 'substring'.
//
// Parameters:
//
//	substring string - the substring to search for.
//	str string - the string to search within.
//
// Returns:
//
//	bool - true if 'str' contains 'substring', false otherwise.
//
// Example:
//
//	{{ "Hello" | contains "ell" }} // Output: true
func (fh *FunctionHandler) Contains(substring string, str string) bool {
	return strings.Contains(str, substring)
}

// HasPrefix checks if 'str' starts with the specified 'prefix'.
//
// Parameters:
//
//	prefix string - the prefix to check.
//	str string - the string to check.
//
// Returns:
//
//	bool - true if 'str' starts with 'prefix', false otherwise.
//
// Example:
//
//	{{ "HelloWorld" | hasPrefix "Hello" }} // Output: true
func (fh *FunctionHandler) HasPrefix(prefix string, str string) bool {
	return strings.HasPrefix(str, prefix)
}

// HasSuffix checks if 'str' ends with the specified 'suffix'.
//
// Parameters:
//
//	suffix string - the suffix to check.
//	str string - the string to check.
//
// Returns:
//
//	bool - true if 'str' ends with 'suffix', false otherwise.
//
// Example:
//
//	{{ "HelloWorld" | hasSuffix "World" }} // Output: true
func (fh *FunctionHandler) HasSuffix(suffix string, str string) bool {
	return strings.HasSuffix(str, suffix)
}

// ToLower converts all characters in the provided string to lowercase.
//
// Parameters:
//
//	str string - the string to convert.
//
// Returns:
//
//	string - the lowercase version of the input string.
//
// Example:
//
//	{{ "HELLO WORLD" | toLower }} // Output: "hello world"
func (fh *FunctionHandler) ToLower(str string) string {
	return strings.ToLower(str)
}

// ToUpper converts all characters in the provided string to uppercase.
//
// Parameters:
//
//	str string - the string to convert.
//
// Returns:
//
//	string - the uppercase version of the input string.
//
// Example:
//
//	{{ "hello world" | toUpper }} // Output: "HELLO WORLD"
func (fh *FunctionHandler) ToUpper(str string) string {
	return strings.ToUpper(str)
}

// Replace replaces all occurrences of 'old' in 'src' with 'new'.
//
// Parameters:
//
//	old string - the substring to be replaced.
//	new string - the substring to replace with.
//	src string - the source string where replacements take place.
//
// Returns:
//
//	string - the modified string after all replacements.
//
// Example:
//
//	{{ "banana" | replace "a", "o" }} // Output: "bonono"
func (fh *FunctionHandler) Replace(old, new, src string) string {
	return strings.Replace(src, old, new, -1)
}

// Repeat repeats the string 'str' for 'count' times.
//
// Parameters:
//
//	count int - the number of times to repeat.
//	str string - the string to repeat.
//
// Returns:
//
//	string - the repeated string.
//
// Example:
//
//	{{ "ha" | repeat 3 }} // Output: "hahaha"
func (fh *FunctionHandler) Repeat(count int, str string) string {
	return strings.Repeat(str, count)
}

// Join concatenates the elements of a slice into a single string separated by 'sep'.
// The slice is extracted from 'v', which can be any slice input. The function
// uses 'Strslice' to convert 'v' to a slice of strings if necessary.
//
// Parameters:
//
//	sep string - the separator string.
//	v any - the slice to join, can be of any slice type.
//
// Returns:
//
//	string - the concatenated string.
//
// Example:
//
//	{{ $list := slice "apple" "banana" "cherry" }}
//	{{ $list | join ", " }} // Output: "apple, banana, cherry"
func (fh *FunctionHandler) Join(sep string, v any) string {
	return strings.Join(fh.Strslice(v), sep)
}

// Trunc truncates 's' to a maximum length 'count'. If 'count' is negative, it removes
// '-count' characters from the beginning of the string.
//
// Parameters:
//
//	count int - the number of characters to keep. Negative values indicate truncation
//	            from the beginning.
//	str string - the string to truncate.
//
// Returns:
//
//	string - the truncated string.
//
// Example:
//
//	{{ "Hello World" | trunc 5 }} // Output: "Hello"
//	{{ "Hello World" | trunc -1 }} // Output: "World"
func (fh *FunctionHandler) Trunc(count int, str string) string {
	length := len(str)

	if count < 0 && length+count > 0 {
		return str[length+count:]
	}

	if count >= 0 && length > count {
		return str[:count]
	}

	return str
}

// Shuffle randomly rearranges the characters in 'str'.
//
// Parameters:
//
//	str string - the string to shuffle.
//
// Returns:
//
//	string - the shuffled string.
//
// Example:
//
//	{{ "hello" | shuffle }} // Output: "loleh" (output may vary due to randomness)
func (fh *FunctionHandler) Shuffle(str string) string {
	r := []rune(str)
	mathrand.New(randSource).Shuffle(len(r), func(i, j int) {
		r[i], r[j] = r[j], r[i]
	})
	return string(r)
}

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
func (fh *FunctionHandler) ellipsis(str string, offset int, maxWidth int) string {
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

// Ellipsis truncates 'str' to 'maxWidth' and appends an ellipsis if the string
// is longer than 'maxWidth'.
//
// Parameters:
//
//	maxWidth int - the maximum width of the string including the ellipsis.
//	str string - the string to truncate.
//
// Returns:
//
//	string - the possibly truncated string with an ellipsis.
//
// Example:
//
//	{{ "Hello World" | ellipsis 10 }} // Output: "Hello W..."
func (fh *FunctionHandler) Ellipsis(maxWidth int, str string) string {
	return fh.ellipsis(str, 0, maxWidth)
}

// EllipsisBoth truncates 'str' from both ends, preserving the middle part of
// the string and appending ellipses to both ends if needed.
//
// Parameters:
//
//	offset int - starting position for preserving text.
//	maxWidth int - the total maximum width including ellipses.
//	str string - the string to truncate.
//
// Returns:
//
//	string - the truncated string with ellipses on both ends.
//
// Example:
//
//	{{ "Hello World" | ellipsisBoth 1 10 }} // Output: "...lo Wor..."
func (fh *FunctionHandler) EllipsisBoth(offset int, maxWidth int, str string) string {
	return fh.ellipsis(str, offset, maxWidth)
}

// Initials extracts the initials from 'str', using optional 'delimiters' to
// determine word boundaries.
//
// Parameters:
//
//	str string - the string from which to extract initials.
//	delimiters string - optional string containing delimiter characters.
//
// Returns:
//
//	string - the initials of the words in 'str'.
//
// Example:
//
//	{{ "John Doe" | initials }} // Output: "JD"
func (fh *FunctionHandler) Initials(str string) string {
	return fh.initials(str, " ")
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
func (fh *FunctionHandler) initials(str string, delimiters string) string {
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

// Plural returns 'one' if 'count' is 1, otherwise it returns 'many'.
//
// Parameters:
//
//	one string - the string to return if 'count' is 1.
//	many string - the string to return if 'count' is not 1.
//	count int - the number used to determine which string to return.
//
// Returns:
//
//	string - either 'one' or 'many' based on 'count'.
//
// Example:
//
//	{{ 1 | plural "apple" "apples" }} // Output: "apple"
//	{{ 2 | plural "apple" "apples" }} // Output: "apples"
func (fh *FunctionHandler) Plural(one, many string, count int) string {
	if count == 1 {
		return one
	}
	return many
}

// Wrap breaks 'str' into lines with a maximum length of 'length'.
// It ensures that words are not split across lines unless necessary.
//
// Parameters:
//
//	length int - the maximum length of each line.
//	str string - the string to be wrapped.
//
// Returns:
//
//	string - the wrapped string using newline characters to separate lines.
//
// Example:
//
//	{{ "This is a long string that needs to be wrapped." | wrap 10 }}
//	Output: "This is a\nlong\nstring\nthat needs\nto be\nwrapped."
func (fh *FunctionHandler) Wrap(length int, str string) string {
	return fh.WordWrap(length, "", false, str)
}

// WrapWith breaks 'str' into lines of maximum 'length', using 'newLineCharacter'
// to separate lines. It wraps words only when they exceed the line length.
//
// Parameters:
//
//	length int - the maximum line length.
//	newLineCharacter string - the character(s) used to denote new lines.
//	str string - the string to wrap.
//
// Returns:
//
//	string - the wrapped string.
//
// Example:
//
//	{{ "This is a long string that needs to be wrapped." | wrapWith 10 "<br>" }}
//	Output: "This is a<br>long<br>string<br>that needs<br>to be<br>wrapped."
func (fh *FunctionHandler) WrapWith(length int, newLineCharacter string, str string) string {
	return fh.WordWrap(length, newLineCharacter, true, str)
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
//	{{ "A very longwordindeed that cannot fit on one line." | wordWrap 10 "\n" true }}
//	Output: "A very\nlongwordin\ndeed that\ncannot fit\non one\nline."
func (fh *FunctionHandler) WordWrap(wrapLength int, newLineCharacter string, wrapLongWords bool, str string) string {
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

// Quote wraps each element in 'elements' with double quotes and separates them with spaces.
//
// Parameters:
//
//	elements ...any - the elements to be quoted.
//
// Returns:
//
//	string - a single string with each element double quoted.
//
// Example:
//
//	 {{ $list := slice "hello" "world" 123 }}
//		{{ $list | quote }}
//		Output: "hello" "world" "123"
func (fh *FunctionHandler) Quote(elements ...any) string {
	var build strings.Builder

	for i, elem := range elements {
		if elem == nil {
			continue
		}
		if i > 0 {
			build.WriteRune(' ')
		}
		build.WriteString(fmt.Sprintf("%q", fmt.Sprint(elem)))
	}
	return build.String()
}

// Squote wraps each element in 'elements' with single quotes and separates them with spaces.
//
// Parameters:
//
//	elements ...any - the elements to be single quoted.
//
// Returns:
//
//	string - a single string with each element single quoted.
//
// Example:
//
//	 {{ $list := slice "hello" "world" 123 }}
//		{{ $list | squote }}
//	Output: 'hello' 'world' '123'
func (fh *FunctionHandler) Squote(elements ...any) string {
	var builder strings.Builder
	for i, elem := range elements {
		if elem == nil {
			continue
		}
		if i > 0 {
			builder.WriteRune(' ')
		}
		// Use fmt.Sprint to convert any to string, then quote it.
		builder.WriteRune('\'')
		builder.WriteString(fmt.Sprint(elem))
		builder.WriteRune('\'')
	}
	return builder.String()
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
//	transformed := fh.transformString(style, "hello world")
//	Output: "Hello_World"
//
// Note:
//
//	This example demonstrates how to use the function to capitalize the first letter of
//	each word and insert underscores between words, which is common in identifiers like
//	variable names in programming.
func (fh *FunctionHandler) transformString(style caseStyle, str string) string {
	var result strings.Builder
	result.Grow(len(str) + 10) // Allocate a bit more for potential separators
	capitalizeNext := style.CapitalizeNext
	var lastRune, nextRune rune = 0, 0

	for i, r := range str {
		if i+1 < len(str) {
			nextRune = rune(str[i+1])
		}

		if r == ' ' || r == '-' || r == '_' {
			if style.Separator != -1 && (lastRune != style.Separator) {
				result.WriteRune(style.Separator)
			}
			capitalizeNext = true
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
	}

	return result.String()
}

// ToCamelCase converts a string to camelCase.
//
// Parameters:
//
//	str string - the string to convert.
//
// Returns:
//
//	string - the string converted to camelCase.
//
// Example:
//
//	{{ "hello world" | toCamelCase }} // Output: "helloWorld"
func (fh *FunctionHandler) ToCamelCase(str string) string {
	return fh.transformString(camelCaseStyle, str)
}

// ToKebabCase converts a string to kebab-case.
//
// Parameters:
//
//	str string - the string to convert.
//
// Returns:
//
//	string - the string converted to kebab-case.
//
// Example:
//
//	{{ "hello world" | toKebabCase }} // Output: "hello-world"
func (fh *FunctionHandler) ToKebabCase(str string) string {
	return fh.transformString(kebabCaseStyle, str)
}

// ToPascalCase converts a string to PascalCase.
//
// Parameters:
//
//	str string - the string to convert.
//
// Returns:
//
//	string - the string converted to PascalCase.
//
// Example:
//
//	{{ "hello world" | toPascalCase }} // Output: "HelloWorld"
func (fh *FunctionHandler) ToPascalCase(str string) string {
	return fh.transformString(pascalCaseStyle, str)
}

// ToDotCase converts a string to dot.case.
//
// Parameters:
//
//	str string - the string to convert.
//
// Returns:
//
//	string - the string converted to dot.case.
//
// Example:
//
//	{{ "hello world" | toDotCase }} // Output: "hello.world"
func (fh *FunctionHandler) ToDotCase(str string) string {
	return fh.transformString(dotCaseStyle, str)
}

// ToPathCase converts a string to path/case.
//
// Parameters:
//
//	str string - the string to convert.
//
// Returns:
//
//	string - the string converted to path/case.
//
// Example:
//
//	{{ "hello world" | toPathCase }} // Output: "hello/world"
func (fh *FunctionHandler) ToPathCase(str string) string {
	return fh.transformString(pathCaseStyle, str)
}

// ToConstantCase converts a string to CONSTANT_CASE.
//
// Parameters:
//
//	str string - the string to convert.
//
// Returns:
//
//	string - the string converted to CONSTANT_CASE.
//
// Example:
//
//	{{ "hello world" | toConstantCase }} // Output: "HELLO_WORLD"
func (fh *FunctionHandler) ToConstantCase(str string) string {
	return fh.transformString(constantCaseStyle, str)
}

// ToSnakeCase converts a string to snake_case.
//
// Parameters:
//
//	str string - the string to convert.
//
// Returns:
//
//	string - the string converted to snake_case.
//
// Example:
//
//	{{ "hello world" | toSnakeCase }} // Output: "hello_world"
func (fh *FunctionHandler) ToSnakeCase(str string) string {
	return fh.transformString(snakeCaseStyle, str)
}

// ToTitleCase converts a string to Title Case.
//
// Parameters:
//
//	str string - the string to convert.
//
// Returns:
//
//	string - the string converted to Title Case.
//
// Example:
//
//	{{ "hello world" | toTitleCase }} // Output: "Hello World"
func (fh *FunctionHandler) ToTitleCase(str string) string {
	return cases.Title(language.English).String(str)
}

// Untitle converts the first letter of each word in 'str' to lowercase.
//
// Parameters:
//
//	str string - the string to be converted.
//
// Returns:
//
//	string - the converted string with each word starting in lowercase.
//
// Example:
//
//	{{ "Hello World" | untitle }} // Output: "hello world"
func (fh *FunctionHandler) Untitle(str string) string {
	var result strings.Builder

	// Process each rune in the input string
	startOfWord := true
	for _, r := range str {
		if r == ' ' {
			startOfWord = true
			result.WriteRune(r)
		} else {
			if startOfWord {
				result.WriteRune(unicode.ToLower(r))
				startOfWord = false
			} else {
				result.WriteRune(r)
			}
		}
	}

	return result.String()
}

// SwapCase switches the case of each letter in 'str'. Lowercase letters become
// uppercase and vice versa.
//
// Parameters:
//
//	str string - the string to convert.
//
// Returns:
//
//	string - the string with each character's case switched.
//
// Example:
//
//	{{ "Hello World" | swapCase }} // Output: "hELLO wORLD"
func (fh *FunctionHandler) SwapCase(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLower(r) {
			return unicode.ToUpper(r)
		}
		return unicode.ToLower(r)
	}, str)
}

// Split divides 'orig' into a map of string parts using 'sep' as the separator.
//
// Parameters:
//
//	sep string - the separator string.
//	orig string - the original string to split.
//
// Returns:
//
//	map[string]string - a map of the split parts.
//
// Example:
//
//	{{ "apple,banana,cherry" | split "," }} // Output: { "_0":"apple", "_1":"banana", "_2":"cherry" }
func (fh *FunctionHandler) Split(sep, orig string) map[string]string {
	parts := strings.Split(orig, sep)
	return fh.FillMapWithParts(parts)
}

// Splitn divides 'orig' into a map of string parts using 'sep' as the separator
// up to 'n' parts.
//
// Parameters:
//
//	sep string - the separator string.
//	n int - the maximum number of substrings to return.
//	orig string - the original string to split.
//
// Returns:
//
//	map[string]string - a map of the split parts.
//
// Example:
//
//	{{ "apple,banana,cherry" | split "," 2 }} // Output: { "_0":"apple", "_1":"banana,cherry" }
func (fh *FunctionHandler) Splitn(sep string, n int, orig string) map[string]string {
	parts := strings.SplitN(orig, sep, n)
	return fh.FillMapWithParts(parts)
}

// Substring extracts a substring from 's' starting at 'start' and ending at 'end'.
// Negative values for 'start' or 'end' are interpreted as positions from the end
// of the string.
//
// Parameters:
//
//	start int - the starting index.
//	end int - the ending index, exclusive.
//	str string - the source string.
//
// Returns:
//
//	string - the extracted substring.
//
// Example:
//
//	{{ "Hello World" | substring 0 5 }} // Output: "Hello"
func (fh *FunctionHandler) Substring(start, end int, str string) string {
	if start < 0 {
		start = len(str) + start
	}
	if end < 0 {
		end = len(str) + end
	}
	if start < 0 {
		start = 0
	}
	if end > len(str) {
		end = len(str)
	}
	if start > end {
		return ""
	}
	return str[start:end]
}

// Indent adds spaces to the beginning of each line in 'str'.
//
// Parameters:
//
//	spaces int - the number of spaces to add.
//	str string - the string to indent.
//
// Returns:
//
//	string - the indented string.
//
// Example:
//
//	{{ "Hello\nWorld" | indent 4 }} // Output: "    Hello\n    World"
func (fh *FunctionHandler) Indent(spaces int, str string) string {
	var builder strings.Builder
	pad := strings.Repeat(" ", spaces)
	lines := strings.Split(str, "\n")

	for i, line := range lines {
		if i > 0 {
			builder.WriteString("\n" + pad)
		} else {
			builder.WriteString(pad)
		}
		builder.WriteString(line)
	}

	return builder.String()
}

// Nindent is similar to Indent, but it adds a newline at the start.
//
// Parameters:
//   spaces int - the number of spaces to add after the newline.
//   str string - the string to indent.
//
// Returns:
//   string - the indented string with a newline at the start.
//
// Example:
//   {{ "Hello\nWorld" | nindent 4 }} // Output: "\n    Hello\n    World"

func (fh *FunctionHandler) Nindent(spaces int, str string) string {
	return "\n" + fh.Indent(spaces, str)
}

// Seq generates a sequence of numbers as a string. It can take 0, 1, 2, or 3
// integers as parameters defining the start, end, and step of the sequence.
// NOTE: This function works similarly to the seq command in Unix systems.
//
// Parameters:
//
//	params ...int - sequence parameters (start, step, end).
//
// Returns:
//
//	string - a space-separated string of numbers in the sequence.
//
// Example:
//
//	{{ seq 1, 2, 10 }} // Output: "1 3 5 7 9"
func (fh *FunctionHandler) Seq(params ...int) string {
	increment := 1
	switch len(params) {
	case 0:
		return ""
	case 1:
		start := 1
		end := params[0]
		if end < start {
			increment = -1
		}
		return fh.IntArrayToString(fh.UntilStep(start, end+increment, increment), " ")
	case 3:
		start := params[0]
		end := params[2]
		step := params[1]
		if end < start {
			increment = -1
			if step > 0 {
				return ""
			}
		}
		return fh.IntArrayToString(fh.UntilStep(start, end+increment, step), " ")
	case 2:
		start := params[0]
		end := params[1]
		step := 1
		if end < start {
			step = -1
		}
		return fh.IntArrayToString(fh.UntilStep(start, end+step, step), " ")
	default:
		return ""
	}
}
