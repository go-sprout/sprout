package strings

import (
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// allCaseStyles contains all available case styles for comprehensive testing
var allCaseStyles = map[string]caseStyle{
	"camelCase":    camelCaseStyle,
	"kebabCase":    kebabCaseStyle,
	"pascalCase":   pascalCaseStyle,
	"snakeCase":    snakeCaseStyle,
	"dotCase":      dotCaseStyle,
	"pathCase":     pathCaseStyle,
	"constantCase": constantCaseStyle,
}

func TestTransformString(t *testing.T) {
	sr := NewRegistry()

	tests := []struct {
		name     string
		style    caseStyle
		input    string
		expected string
	}{
		// ==================== camelCase ====================
		{name: "camelCase/empty", style: camelCaseStyle, input: "", expected: ""},
		{name: "camelCase/single_lower", style: camelCaseStyle, input: "a", expected: "a"},
		{name: "camelCase/single_upper", style: camelCaseStyle, input: "A", expected: "a"},
		{name: "camelCase/simple_words", style: camelCaseStyle, input: "foo bar", expected: "fooBar"},
		{name: "camelCase/underscore_sep", style: camelCaseStyle, input: "foo_bar", expected: "fooBar"},
		{name: "camelCase/hyphen_sep", style: camelCaseStyle, input: "foo-bar", expected: "fooBar"},
		{name: "camelCase/mixed_sep", style: camelCaseStyle, input: "foo-bar_baz", expected: "fooBarBaz"},
		{name: "camelCase/upper_case_input", style: camelCaseStyle, input: "UPPER_CASE", expected: "upperCase"},
		{name: "camelCase/preserve_word_boundary", style: camelCaseStyle, input: "firstName", expected: "firstName"},
		{name: "camelCase/stray_uppercase", style: camelCaseStyle, input: "FoO bar", expected: "fooBar"},
		{name: "camelCase/acronym", style: camelCaseStyle, input: "HTMLParser", expected: "htmlParser"},
		{name: "camelCase/acronym_end", style: camelCaseStyle, input: "NoHTTPS", expected: "noHttps"},
		{name: "camelCase/multiple_spaces", style: camelCaseStyle, input: "foo  bar", expected: "fooBar"},
		{name: "camelCase/leading_separator", style: camelCaseStyle, input: "_foo_bar", expected: "fooBar"},
		{name: "camelCase/trailing_separator", style: camelCaseStyle, input: "foo_bar_", expected: "fooBar"},
		{name: "camelCase/with_numbers", style: camelCaseStyle, input: "http2xx", expected: "http2xx"},
		{name: "camelCase/complex", style: camelCaseStyle, input: "___complex__case_", expected: "complexCase"},

		// ==================== PascalCase ====================
		{name: "pascalCase/empty", style: pascalCaseStyle, input: "", expected: ""},
		{name: "pascalCase/single_lower", style: pascalCaseStyle, input: "a", expected: "A"},
		{name: "pascalCase/single_upper", style: pascalCaseStyle, input: "A", expected: "A"},
		{name: "pascalCase/simple_words", style: pascalCaseStyle, input: "foo bar", expected: "FooBar"},
		{name: "pascalCase/underscore_sep", style: pascalCaseStyle, input: "foo_bar", expected: "FooBar"},
		{name: "pascalCase/upper_case_input", style: pascalCaseStyle, input: "UPPER_CASE", expected: "UpperCase"},
		{name: "pascalCase/preserve_word_boundary", style: pascalCaseStyle, input: "FirstName", expected: "FirstName"},
		{name: "pascalCase/from_camel", style: pascalCaseStyle, input: "firstName", expected: "FirstName"},
		{name: "pascalCase/acronym", style: pascalCaseStyle, input: "HTMLParser", expected: "HtmlParser"},
		{name: "pascalCase/stray_uppercase", style: pascalCaseStyle, input: "FoO", expected: "Foo"},

		// ==================== snake_case ====================
		{name: "snakeCase/empty", style: snakeCaseStyle, input: "", expected: ""},
		{name: "snakeCase/simple_words", style: snakeCaseStyle, input: "foo bar", expected: "foo_bar"},
		{name: "snakeCase/from_camel", style: snakeCaseStyle, input: "fooBar", expected: "foo_bar"},
		{name: "snakeCase/from_pascal", style: snakeCaseStyle, input: "FooBar", expected: "foo_bar"},
		{name: "snakeCase/acronym", style: snakeCaseStyle, input: "HTTPServer", expected: "http_server"},
		{name: "snakeCase/acronym_middle", style: snakeCaseStyle, input: "NoHTTPS", expected: "no_https"},
		{name: "snakeCase/with_numbers", style: snakeCaseStyle, input: "http2xx", expected: "http_2xx"},
		{name: "snakeCase/complex_numbers", style: snakeCaseStyle, input: "HTTP20xOK", expected: "http_20x_ok"},
		{name: "snakeCase/number_word", style: snakeCaseStyle, input: "Duration2m3s", expected: "duration_2m_3s"},
		{name: "snakeCase/already_snake", style: snakeCaseStyle, input: "foo_bar", expected: "foo_bar"},

		// ==================== kebab-case ====================
		{name: "kebabCase/empty", style: kebabCaseStyle, input: "", expected: ""},
		{name: "kebabCase/simple_words", style: kebabCaseStyle, input: "foo bar", expected: "foo-bar"},
		{name: "kebabCase/from_camel", style: kebabCaseStyle, input: "fooBar", expected: "foo-bar"},
		{name: "kebabCase/from_pascal", style: kebabCaseStyle, input: "FooBar", expected: "foo-bar"},
		{name: "kebabCase/acronym", style: kebabCaseStyle, input: "HTTPServer", expected: "http-server"},
		{name: "kebabCase/already_kebab", style: kebabCaseStyle, input: "foo-bar", expected: "foo-bar"},

		// ==================== dot.case ====================
		{name: "dotCase/empty", style: dotCaseStyle, input: "", expected: ""},
		{name: "dotCase/simple_words", style: dotCaseStyle, input: "foo bar", expected: "foo.bar"},
		{name: "dotCase/from_camel", style: dotCaseStyle, input: "fooBar", expected: "foo.bar"},

		// ==================== path/case ====================
		{name: "pathCase/empty", style: pathCaseStyle, input: "", expected: ""},
		{name: "pathCase/simple_words", style: pathCaseStyle, input: "foo bar", expected: "foo/bar"},
		{name: "pathCase/from_camel", style: pathCaseStyle, input: "fooBar", expected: "foo/bar"},

		// ==================== CONSTANT_CASE ====================
		{name: "constantCase/empty", style: constantCaseStyle, input: "", expected: ""},
		{name: "constantCase/simple_words", style: constantCaseStyle, input: "foo bar", expected: "FOO_BAR"},
		{name: "constantCase/from_camel", style: constantCaseStyle, input: "fooBar", expected: "FOO_BAR"},
		{name: "constantCase/from_snake", style: constantCaseStyle, input: "foo_bar", expected: "FOO_BAR"},
		{name: "constantCase/already_constant", style: constantCaseStyle, input: "FOO_BAR", expected: "FOO_BAR"},

		// ==================== Edge cases ====================
		{name: "edge/only_separators", style: camelCaseStyle, input: "___", expected: ""},
		{name: "edge/only_spaces", style: camelCaseStyle, input: "   ", expected: ""},
		{name: "edge/mixed_separators", style: snakeCaseStyle, input: "a-b_c d", expected: "a_b_c_d"},
		{name: "edge/unicode_letters", style: camelCaseStyle, input: "café latté", expected: "caféLatté"},
		{name: "edge/numbers_only", style: snakeCaseStyle, input: "123", expected: "123"},
		{name: "edge/number_start", style: camelCaseStyle, input: "2fast2furious", expected: "2fast2furious"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := sr.transformString(tc.style, tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestTransformStringInvariants(t *testing.T) {
	sr := NewRegistry()

	// Test inputs covering various patterns
	inputs := []string{
		"",
		"a",
		"A",
		"abc",
		"ABC",
		"aBC",
		"Abc",
		"abcDef",
		"AbcDef",
		"ABCDef",
		"abcDEF",
		"abc_def",
		"abc-def",
		"abc def",
		"ABC_DEF",
		"__abc__",
		"HTTPServer",
		"noHTTPS",
		"http2xx",
		"firstName",
		"FirstName",
		"UPPER_CASE",
		"café",
		"日本語",
		"mixed123Numbers",
		"123start",
		"end123",
		"  spaces  ",
	}

	for styleName, style := range allCaseStyles {
		for _, input := range inputs {
			t.Run(styleName+"/"+input, func(t *testing.T) {
				result := sr.transformString(style, input)

				// Invariant 1: Result should be valid UTF-8
				assert.True(t, utf8.ValidString(result),
					"Result should be valid UTF-8 for input %q", input)

				// Invariant 2: For styles without separators, no separator chars should appear
				if style.Separator == -1 {
					assert.NotContains(t, result, "_",
						"camelCase/PascalCase should not contain underscores")
					assert.NotContains(t, result, "-",
						"camelCase/PascalCase should not contain hyphens")
				}

				// Invariant 3: For ForceLowercase styles, check casing rules
				if style.ForceLowercase && !style.CapitalizeNext {
					for _, r := range result {
						if unicode.IsLetter(r) {
							assert.True(t, unicode.IsLower(r) || !unicode.IsUpper(r),
								"ForceLowercase without CapitalizeNext should have all lowercase letters")
						}
					}
				}

				// Invariant 4: For ForceUppercase styles, all letters should be uppercase
				if style.ForceUppercase {
					for _, r := range result {
						if unicode.IsLetter(r) {
							assert.True(t, unicode.IsUpper(r) || !unicode.IsLower(r),
								"ForceUppercase should have all uppercase letters")
						}
					}
				}

				// Invariant 5: No consecutive separators should appear
				if style.Separator != -1 {
					sep := string(style.Separator)
					assert.NotContains(t, result, sep+sep,
						"Result should not contain consecutive separators")
				}

				// Invariant 6: Result should not start with separator when input starts with content
				if style.Separator != -1 && len(result) > 0 && len(input) > 0 {
					firstRune, _ := utf8.DecodeRuneInString(result)
					firstInputRune, _ := utf8.DecodeRuneInString(input)
					// Only check if input starts with content (not a separator)
					inputStartsWithContent := unicode.IsLetter(firstInputRune) || unicode.IsDigit(firstInputRune)
					if inputStartsWithContent {
						assert.NotEqual(t, style.Separator, firstRune,
							"Result should not start with separator when input starts with content")
					}
				}
			})
		}
	}
}

// FuzzTransformString tests transformString with random inputs
func FuzzTransformString(f *testing.F) {
	// Add seed corpus with interesting patterns
	seeds := []string{
		"",
		"a",
		"A",
		"abc",
		"ABC",
		"abcDef",
		"AbcDef",
		"abc_def",
		"abc-def",
		"abc def",
		"HTTPServer",
		"noHTTPS",
		"firstName",
		"FirstName",
		"UPPER_CASE",
		"FoO",
		"café latté",
		"http2xx",
		"HTTP20xOK",
		"___test___",
		"  spaces  ",
		"MixedCASE_with-separators",
		"123Numbers456",
		"a1b2c3",
		"XMLHttpRequest",
		"getHTTPResponseCode",
		"already_snake_case",
		"already-kebab-case",
		"ALREADY_CONSTANT_CASE",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	sr := NewRegistry()
	styles := []caseStyle{
		camelCaseStyle,
		kebabCaseStyle,
		pascalCaseStyle,
		snakeCaseStyle,
		dotCaseStyle,
		pathCaseStyle,
		constantCaseStyle,
	}

	f.Fuzz(func(t *testing.T, input string) {
		// Skip invalid UTF-8 inputs
		if !utf8.ValidString(input) {
			t.Skip("Invalid UTF-8 input")
		}

		for _, style := range styles {
			// The function should never panic
			result := sr.transformString(style, input)

			// Invariant 1: Result must be valid UTF-8
			require.True(t, utf8.ValidString(result),
				"Result must be valid UTF-8 for input %q", input)

			// Invariant 2: Result length should be bounded
			// (at most 2x input length + some buffer for separators)
			maxExpectedLen := len(input)*2 + 10
			require.LessOrEqual(t, len(result), maxExpectedLen,
				"Result length should be bounded")

			// Invariant 3: For styles without separators, verify no separators
			if style.Separator == -1 {
				for _, r := range result {
					require.False(t, r == '_' || r == '-' || r == ' ',
						"Styles without separators should not introduce separators")
				}
			}

			// Invariant 4: For ForceUppercase, letters with uppercase forms must be uppercase
			// (some Unicode letters like 'ɿ' have no uppercase form and are kept as-is)
			if style.ForceUppercase {
				for _, r := range result {
					if unicode.IsLetter(r) && unicode.IsLower(r) && unicode.ToUpper(r) != r {
						t.Errorf("ForceUppercase style produced lowercase letter in result %q from input %q", result, input)
					}
				}
			}

			// Invariant 5: No consecutive separators (unless input contains separator char)
			if style.Separator != -1 {
				// Skip this check if input contains the separator character itself
				// (e.g., input " ." with dotCase will naturally produce "..")
				inputContainsSeparator := strings.ContainsRune(input, style.Separator)
				if !inputContainsSeparator {
					prev := rune(0)
					for _, r := range result {
						if r == style.Separator && prev == style.Separator {
							t.Errorf("Consecutive separators found in result %q from input %q", result, input)
						}
						prev = r
					}
				}
			}

			// Invariant 6: Letter count should be preserved (letters are transformed, not removed)
			inputLetters := countLetters(input)
			resultLetters := countLetters(result)
			require.Equal(t, inputLetters, resultLetters,
				"Letter count should be preserved: input %q (%d letters) -> result %q (%d letters)",
				input, inputLetters, result, resultLetters)

			// Invariant 7: Digit count should be preserved
			inputDigits := countDigits(input)
			resultDigits := countDigits(result)
			require.Equal(t, inputDigits, resultDigits,
				"Digit count should be preserved: input %q (%d digits) -> result %q (%d digits)",
				input, inputDigits, result, resultDigits)
		}
	})
}

// FuzzTransformStringRoundTrip tests idempotency for ASCII-only inputs
// Note: Idempotency doesn't hold universally for all Unicode inputs due to
// edge cases with characters that have unusual case properties (e.g., 'ϒ').
func FuzzTransformStringRoundTrip(f *testing.F) {
	seeds := []string{
		"fooBar",
		"FooBar",
		"foo_bar",
		"foo-bar",
		"FOO_BAR",
		"httpServer",
		"HTTPServer",
		"simple",
		"ALLCAPS",
		"mixedCase123",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	sr := NewRegistry()

	// Styles with separators should be idempotent for ASCII inputs
	idempotentStyles := map[string]caseStyle{
		"snakeCase":    snakeCaseStyle,
		"kebabCase":    kebabCaseStyle,
		"dotCase":      dotCaseStyle,
		"pathCase":     pathCaseStyle,
		"constantCase": constantCaseStyle,
	}

	f.Fuzz(func(t *testing.T, input string) {
		if !utf8.ValidString(input) {
			t.Skip("Invalid UTF-8 input")
		}

		// Skip non-ASCII inputs as they may have unusual case properties
		for _, r := range input {
			if r > 127 {
				t.Skip("Non-ASCII input may have unusual case properties")
			}
		}

		// Test idempotency: applying the same transformation twice should give the same result
		for styleName, style := range idempotentStyles {
			first := sr.transformString(style, input)
			second := sr.transformString(style, first)

			require.Equal(t, first, second,
				"%s should be idempotent: transforming %q gave %q, but transforming again gave %q",
				styleName, input, first, second)
		}
	})
}

// FuzzTransformStringAllStyles tests that all style transformations produce valid output
// and don't lose content (letters and digits are preserved).
func FuzzTransformStringAllStyles(f *testing.F) {
	seeds := []string{
		"helloWorld",
		"HelloWorld",
		"hello_world",
		"hello-world",
		"HELLO_WORLD",
		"test123",
		"123test",
		"mixedCase123Test",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	sr := NewRegistry()

	f.Fuzz(func(t *testing.T, input string) {
		if !utf8.ValidString(input) {
			t.Skip("Invalid UTF-8 input")
		}

		// Test that all transformations preserve letter and digit counts
		inputLetters := countLetters(input)
		inputDigits := countDigits(input)

		for styleName, style := range allCaseStyles {
			result := sr.transformString(style, input)

			// Letters should be preserved
			resultLetters := countLetters(result)
			require.Equal(t, inputLetters, resultLetters,
				"%s should preserve letter count: input %q (%d) -> result %q (%d)",
				styleName, input, inputLetters, result, resultLetters)

			// Digits should be preserved
			resultDigits := countDigits(result)
			require.Equal(t, inputDigits, resultDigits,
				"%s should preserve digit count: input %q (%d) -> result %q (%d)",
				styleName, input, inputDigits, result, resultDigits)

			// Result should be valid UTF-8
			require.True(t, utf8.ValidString(result),
				"%s should produce valid UTF-8 for input %q", styleName, input)
		}
	})
}

// countLetters counts the number of Unicode letters in a string
func countLetters(s string) int {
	count := 0
	for _, r := range s {
		if unicode.IsLetter(r) {
			count++
		}
	}
	return count
}

// countDigits counts the number of Unicode digits in a string
func countDigits(s string) int {
	count := 0
	for _, r := range s {
		if unicode.IsDigit(r) {
			count++
		}
	}
	return count
}

func TestSwapFirstLetter(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		shouldUppercase bool
		expected        string
	}{
		{name: "empty", input: "", shouldUppercase: true, expected: ""},
		{name: "uppercase_first_lower", input: "hello", shouldUppercase: true, expected: "Hello"},
		{name: "uppercase_first_already_upper", input: "Hello", shouldUppercase: true, expected: "Hello"},
		{name: "lowercase_first_upper", input: "Hello", shouldUppercase: false, expected: "hello"},
		{name: "lowercase_first_already_lower", input: "hello", shouldUppercase: false, expected: "hello"},
		{name: "single_char_upper", input: "a", shouldUppercase: true, expected: "A"},
		{name: "single_char_lower", input: "A", shouldUppercase: false, expected: "a"},
		{name: "number_prefix", input: "123abc", shouldUppercase: true, expected: "123Abc"},
		{name: "unicode", input: "über", shouldUppercase: true, expected: "Über"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := swapFirstLetter(tc.input, tc.shouldUppercase)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestEllipsis(t *testing.T) {
	sr := NewRegistry()

	tests := []struct {
		name     string
		value    string
		offset   int
		maxWidth int
		expected string
	}{
		{name: "no_truncation_needed", value: "hello", offset: 0, maxWidth: 10, expected: "hello"},
		{name: "truncate_end", value: "hello world", offset: 0, maxWidth: 8, expected: "hello..."},
		{name: "maxWidth_too_small", value: "hello", offset: 0, maxWidth: 3, expected: "hello"},
		{name: "with_offset", value: "hello world", offset: 2, maxWidth: 10, expected: "...llo ..."},
		{name: "empty_string", value: "", offset: 0, maxWidth: 10, expected: ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := sr.ellipsis(tc.value, tc.offset, tc.maxWidth)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestInitials(t *testing.T) {
	sr := NewRegistry()

	tests := []struct {
		name       string
		value      string
		delimiters string
		expected   string
	}{
		{name: "simple", value: "John Doe", delimiters: " ", expected: "JD"},
		{name: "multiple_words", value: "John Michael Doe", delimiters: " ", expected: "JMD"},
		{name: "custom_delimiter", value: "John-Michael-Doe", delimiters: "-", expected: "JMD"},
		{name: "empty", value: "", delimiters: " ", expected: ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := sr.initials(tc.value, tc.delimiters)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestPopulateMapWithParts(t *testing.T) {
	sr := NewRegistry()

	tests := []struct {
		name     string
		parts    []string
		expected map[string]string
	}{
		{
			name:     "simple",
			parts:    []string{"a", "b", "c"},
			expected: map[string]string{"_0": "a", "_1": "b", "_2": "c"},
		},
		{
			name:     "empty",
			parts:    []string{},
			expected: map[string]string{},
		},
		{
			name:     "single",
			parts:    []string{"only"},
			expected: map[string]string{"_0": "only"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := sr.populateMapWithParts(tc.parts)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestConvertIntArrayToString(t *testing.T) {
	sr := NewRegistry()

	tests := []struct {
		name      string
		slice     []int
		delimiter string
		expected  string
	}{
		{name: "simple", slice: []int{1, 2, 3}, delimiter: ",", expected: "1,2,3"},
		{name: "single", slice: []int{42}, delimiter: ",", expected: "42"},
		{name: "empty", slice: []int{}, delimiter: ",", expected: ""},
		{name: "different_delimiter", slice: []int{1, 2, 3}, delimiter: "-", expected: "1-2-3"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := sr.convertIntArrayToString(tc.slice, tc.delimiter)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestWordWrap(t *testing.T) {
	sr := NewRegistry()

	tests := []struct {
		name             string
		value            string
		wrapLength       int
		newLineCharacter string
		wrapLongWords    bool
		expected         string
	}{
		{
			name:             "simple_wrap",
			value:            "hello world foo bar",
			wrapLength:       10,
			newLineCharacter: "\n",
			wrapLongWords:    false,
			expected:         "hello\nworld foo\nbar",
		},
		{
			name:             "no_wrap_needed",
			value:            "hello",
			wrapLength:       10,
			newLineCharacter: "\n",
			wrapLongWords:    false,
			expected:         "hello",
		},
		{
			name:             "wrap_long_word",
			value:            "superlongword",
			wrapLength:       5,
			newLineCharacter: "\n",
			wrapLongWords:    true,
			expected:         "super\nlongw\nord",
		},
		{
			name:             "custom_newline",
			value:            "hello world",
			wrapLength:       5,
			newLineCharacter: "<br>",
			wrapLongWords:    false,
			expected:         "hello<br>world",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := sr.wordWrap(tc.wrapLength, tc.newLineCharacter, tc.wrapLongWords, tc.value)
			assert.Equal(t, tc.expected, result)
		})
	}
}

// BenchmarkTransformString benchmarks the transformString function
func BenchmarkTransformString(b *testing.B) {
	sr := NewRegistry()
	inputs := []string{
		"simple",
		"camelCaseInput",
		"PascalCaseInput",
		"snake_case_input",
		"CONSTANT_CASE_INPUT",
		"MixedHTTPServerInput",
		"very_long_snake_case_string_with_many_words",
	}

	for styleName, style := range allCaseStyles {
		for _, input := range inputs {
			b.Run(styleName+"/"+strings.ReplaceAll(input, "_", ""), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					sr.transformString(style, input)
				}
			})
		}
	}
}
