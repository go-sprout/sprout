package sprout

import (
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var randSource rand.Source

func init() {
	randSource = rand.NewSource(time.Now().UnixNano())
}

func nospace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func swapCase(s string) string {
	return strings.Map(func(r rune) rune {
		switch {
		case unicode.IsLower(r):
			return unicode.ToUpper(r)
		case unicode.IsUpper(r):
			return unicode.ToLower(r)
		default:
			return r
		}
	}, s)
}

func wordWrap(str string, wrapLength int, newLineCharacter string, wrapLongWords bool) string {
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
		if currentLineLength > 0 && (currentLineLength+1+wordLength > wrapLength && !wrapLongWords || wordLength > wrapLength) {
			resultBuilder.WriteString(newLineCharacter)
			currentLineLength = 0
		}

		if wrapLongWords && wordLength > wrapLength {
			for i, r := range word {
				if currentLineLength == wrapLength {
					resultBuilder.WriteString(newLineCharacter)
					currentLineLength = 0
				}
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
				resultBuilder.WriteString(newLineCharacter)
				currentLineLength++
			}
			resultBuilder.WriteString(word)
			currentLineLength += wordLength
		}
	}

	return resultBuilder.String()
}

func toTitleCase(s string) string {
	return cases.Title(language.English).String(s)
}

// shuffle shuffles a string in a random manner.
func shuffle(str string) string {
	r := []rune(str)
	rand.New(randSource).Shuffle(len(r), func(i, j int) {
		r[i], r[j] = r[j], r[i]
	})
	return string(r)
}

// ellipsis adds an ellipsis to the string `str` starting at `offset` if the length exceeds `maxWidth`.
// `maxWidth` must be at least 4, to accommodate the ellipsis and at least one character.
func ellipsis(str string, offset, maxWidth int) string {
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

// initials extracts the initials from the given string using the specified delimiters.
// If delimiters are empty, it defaults to using whitespace.
func initials(str string, delimiters string) string {
	// Define a function to determine if a rune is a delimiter.
	isDelimiter := func(r rune) bool {
		if delimiters == "" {
			return unicode.IsSpace(r)
		}
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

// uncapitalize transforms the first letter of each word in the string to lowercase.
// It uses specified delimiters or whitespace to determine word boundaries.
func uncapitalize(str string, delimiters string) string {
	var result strings.Builder
	// Convert delimiters to a map for efficient checking
	delimMap := make(map[rune]bool)
	for _, d := range delimiters {
		delimMap[d] = true
	}

	// Helper function to check if a rune is a delimiter
	isDelim := func(r rune) bool {
		if delimiters == "" {
			return unicode.IsSpace(r)
		}
		return delimMap[r]
	}

	// Process each rune in the input string
	startOfWord := true
	for _, r := range str {
		if isDelim(r) {
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

// base64encode encodes a string to Base64.
func base64encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// base64decode decodes a Base64 encoded string.
func base64decode(s string) string {
	bytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}

// base32encode encodes a string to Base32.
func base32encode(s string) string {
	return base32.StdEncoding.EncodeToString([]byte(s))
}

// base32decode decodes a Base32 encoded string.
func base32decode(s string) string {
	bytes, err := base32.StdEncoding.DecodeString(s)
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}

func randAlphaNumeric(count int) string {
	return cryptoRandomString(count, cryptoRandomStringOpts{letters: true, numbers: true})
}

func randAlpha(count int) string {
	return cryptoRandomString(count, cryptoRandomStringOpts{letters: true})
}

func randAscii(count int) string {
	return cryptoRandomString(count, cryptoRandomStringOpts{ascii: true})
}

func randNumeric(count int) string {
	return cryptoRandomString(count, cryptoRandomStringOpts{numbers: true})
}

func untitle(str string) string {
	return uncapitalize(str, "")
}

func quote(str ...interface{}) string {
	var build strings.Builder
	for i, s := range str {
		if s == nil {
			continue
		}
		if i > 0 {
			build.WriteRune(' ')
		}
		build.WriteString(fmt.Sprintf("%q", fmt.Sprint(s)))
	}
	return build.String()
}

func squote(str ...interface{}) string {
	var builder strings.Builder
	for i, s := range str {
		if s == nil {
			continue
		}
		if i > 0 {
			builder.WriteRune(' ')
		}
		// Use fmt.Sprint to convert interface{} to string, then quote it.
		builder.WriteRune('\'')
		builder.WriteString(fmt.Sprint(s))
		builder.WriteRune('\'')
	}
	return builder.String()
}

// Efficiently concatenates non-nil elements of v, separated by spaces.
func cat(v ...interface{}) string {
	var builder strings.Builder
	for i, item := range v {
		if item == nil {
			continue // Skip nil elements
		}
		if i > 0 {
			builder.WriteRune(' ') // Add space between elements
		}
		// Append the string representation of the item
		builder.WriteString(fmt.Sprint(item))
	}
	// Return the concatenated string without trailing spaces
	return builder.String()
}

// Efficiently indents each line of the input string `v` with `spaces` number of spaces.
func indent(spaces int, v string) string {
	var builder strings.Builder
	pad := strings.Repeat(" ", spaces)
	lines := strings.Split(v, "\n")

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

// Adds a newline at the start and then indents each line of `v` with `spaces` number of spaces.
func nindent(spaces int, v string) string {
	return "\n" + indent(spaces, v)
}

func replace(old, new, src string) string {
	return strings.Replace(src, old, new, -1)
}

func plural(one, many string, count int) string {
	if count == 1 {
		return one
	}
	return many
}

func join(sep string, v interface{}) string {
	return strings.Join(strslice(v), sep)
}

func strval(v interface{}) string {
	switch v := v.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case error:
		return v.Error()
	case fmt.Stringer:
		return v.String()
	default:
		// Handles any other types by leveraging fmt.Sprintf for a string representation.
		return fmt.Sprintf("%v", v)
	}
}

// strslice attempts to convert a variety of slice types to a slice of strings, optimizing performance and minimizing assignments.
func strslice(v interface{}) []string {
	if v == nil {
		return []string{}
	}

	// Handle []string type efficiently without reflection.
	if strs, ok := v.([]string); ok {
		return strs
	}

	// For slices of interface{}, convert each element to a string, skipping nil values.
	if interfaces, ok := v.([]interface{}); ok {
		var result []string
		for _, s := range interfaces {
			if s != nil {
				result = append(result, strval(s))
			}
		}
		return result
	}

	// Use reflection for other slice types to convert them to []string.
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		var result []string
		for i := 0; i < val.Len(); i++ {
			value := val.Index(i).Interface()
			if value != nil {
				result = append(result, strval(value))
			}
		}
		return result
	}

	// If it's not a slice, array, or nil, return a slice with the string representation of v.
	return []string{strval(v)}
}

func trunc(c int, s string) string {
	length := len(s)

	if c < 0 && length+c > 0 {
		return s[length+c:]
	}

	if c >= 0 && length > c {
		return s[:c]
	}

	return s
}

// Splits `orig` string by `sep` and returns a map of the resulting parts.
// Each key is prefixed with an underscore followed by the part's index.
func split(sep, orig string) map[string]string {
	parts := strings.Split(orig, sep)
	return fillMapWithParts(parts)
}

// Splits `orig` string by `sep` into at most `n` parts and returns a map of the parts.
// Each key is prefixed with an underscore followed by the part's index.
func splitn(sep string, n int, orig string) map[string]string {
	parts := strings.SplitN(orig, sep, n)
	return fillMapWithParts(parts)
}

// fillMapWithParts fills a map with the provided parts, using a key format.
func fillMapWithParts(parts []string) map[string]string {
	res := make(map[string]string, len(parts))
	for i, v := range parts {
		res[fmt.Sprintf("_%d", i)] = v
	}
	return res
}

func substring(start, end int, s string) string {
	if start < 0 {
		start = len(s) + start
	}
	if end < 0 {
		end = len(s) + end
	}
	if start < 0 {
		start = 0
	}
	if end > len(s) {
		end = len(s)
	}
	if start > end {
		return ""
	}
	return s[start:end]
}
