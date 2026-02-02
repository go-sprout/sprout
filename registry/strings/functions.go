package strings

import (
	"fmt"
	mathrand "math/rand"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/go-sprout/sprout/internal/helpers"
)

// Nospace removes all whitespace characters from the provided string.
// It uses the unicode package to identify whitespace runes and removes them.
//
// Parameters:
//
//	value string - the string from which to remove whitespace.
//
// Returns:
//
//	string - the modified string with all whitespace characters removed.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: nospace].
//
// [Sprout Documentation: nospace]: https://docs.atom.codes/sprout/registries/strings#nospace
func (sr *StringsRegistry) Nospace(value string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, value)
}

// Trim removes leading and trailing whitespace from the string.
//
// Parameters:
//
//	value string - the string to trim.
//
// Returns:
//
//	string - the trimmed string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: trim].
//
// [Sprout Documentation: trim]: https://docs.atom.codes/sprout/registries/strings#trim
func (sr *StringsRegistry) Trim(value string) string {
	return strings.TrimSpace(value)
}

// TrimAll removes all occurrences of any characters in 'cutset' from both the
// beginning and the end of 'str'.
//
// Parameters:
//
//	cutset string - a string of characters to remove from the string.
//	value string - the string to trim.
//
// Returns:
//
//	string - the string with specified characters removed.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: trimAll].
//
// [Sprout Documentation: trimAll]: https://docs.atom.codes/sprout/registries/strings#trimall
func (sr *StringsRegistry) TrimAll(cutset string, value string) string {
	return strings.Trim(value, cutset)
}

// TrimPrefix removes the 'prefix' from the start of 'str' if present.
//
// Parameters:
//
//	prefix string - the prefix to remove.
//	value string - the string to trim.
//
// Returns:
//
//	string - the string with the prefix removed if it was present.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: trimPrefix].
//
// [Sprout Documentation: trimPrefix]: https://docs.atom.codes/sprout/registries/strings#trimprefix
func (sr *StringsRegistry) TrimPrefix(prefix string, value string) string {
	return strings.TrimPrefix(value, prefix)
}

// TrimSuffix removes the 'suffix' from the end of 'str' if present.
//
// Parameters:
//
//	suffix string - the suffix to remove.
//	value string - the string to trim.
//
// Returns:
//
//	string - the string with the suffix removed if it was present.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: trimSuffix].
//
// [Sprout Documentation: trimSuffix]: https://docs.atom.codes/sprout/registries/strings#trimsuffix
func (sr *StringsRegistry) TrimSuffix(suffix string, value string) string {
	return strings.TrimSuffix(value, suffix)
}

// Contains checks if 'str' contains the 'substring'.
//
// Parameters:
//
//	substring string - the substring to search for.
//	value string - the string to search within.
//
// Returns:
//
//	bool - true if 'str' contains 'substring', false otherwise.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: contains].
//
// [Sprout Documentation: contains]: https://docs.atom.codes/sprout/registries/strings#contains
func (sr *StringsRegistry) Contains(substring string, value string) bool {
	return strings.Contains(value, substring)
}

// HasPrefix checks if 'str' starts with the specified 'prefix'.
//
// Parameters:
//
//	prefix string - the prefix to check.
//	value string - the string to check.
//
// Returns:
//
//	bool - true if 'str' starts with 'prefix', false otherwise.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: hasPrefix].
//
// [Sprout Documentation: hasPrefix]: https://docs.atom.codes/sprout/registries/strings#hasprefix
func (sr *StringsRegistry) HasPrefix(prefix string, value string) bool {
	return strings.HasPrefix(value, prefix)
}

// HasSuffix checks if 'str' ends with the specified 'suffix'.
//
// Parameters:
//
//	suffix string - the suffix to check.
//	value string - the string to check.
//
// Returns:
//
//	bool - true if 'str' ends with 'suffix', false otherwise.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: hasSuffix].
//
// [Sprout Documentation: hasSuffix]: https://docs.atom.codes/sprout/registries/strings#hassuffix
func (sr *StringsRegistry) HasSuffix(suffix string, value string) bool {
	return strings.HasSuffix(value, suffix)
}

// ToLower converts all characters in the provided string to lowercase.
//
// Parameters:
//
//	value string - the string to convert.
//
// Returns:
//
//	string - the lowercase version of the input string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toLower].
//
// [Sprout Documentation: toLower]: https://docs.atom.codes/sprout/registries/strings#tolower
func (sr *StringsRegistry) ToLower(value string) string {
	return strings.ToLower(value)
}

// ToUpper converts all characters in the provided string to uppercase.
//
// Parameters:
//
//	value string - the string to convert.
//
// Returns:
//
//	string - the uppercase version of the input string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toUpper].
//
// [Sprout Documentation: toUpper]: https://docs.atom.codes/sprout/registries/strings#toupper
func (sr *StringsRegistry) ToUpper(value string) string {
	return strings.ToUpper(value)
}

// Replace replaces all occurrences of 'old' in 'src' with 'new'.
//
// Parameters:
//
//	old string - the substring to be replaced.
//	new string - the substring to replace with.
//	value string - the source string where replacements take place.
//
// Returns:
//
//	string - the modified string after all replacements.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: replace].
//
// [Sprout Documentation: replace]: https://docs.atom.codes/sprout/registries/strings#replace
func (sr *StringsRegistry) Replace(old, new, value string) string {
	return strings.ReplaceAll(value, old, new)
}

// Repeat repeats the string 'str' for 'count' times.
//
// Parameters:
//
//	count int - the number of times to repeat.
//	value string - the string to repeat.
//
// Returns:
//
//	string - the repeated string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: repeat].
//
// [Sprout Documentation: repeat]: https://docs.atom.codes/sprout/registries/strings#repeat
func (sr *StringsRegistry) Repeat(count int, value string) string {
	return strings.Repeat(value, count)
}

// Join concatenates the elements of a slice into a single string separated by 'sep'.
// The slice is extracted from 'v', which can be any slice input. The function
// uses 'Strslice' to convert 'v' to a slice of strings if necessary.
//
// Parameters:
//
//	sep string - the separator string.
//	value any - the slice to join, can be of any slice type.
//
// Returns:
//
//	string - the concatenated string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: join].
//
// [Sprout Documentation: join]: https://docs.atom.codes/sprout/registries/strings#join
func (sr *StringsRegistry) Join(sep string, value any) string {
	return strings.Join(helpers.StrSlice(value), sep)
}

// Trunc truncates 's' to a maximum length 'count'. If 'count' is negative, it removes
// '-count' characters from the beginning of the string.
//
// Parameters:
//
//	count int - the number of characters to keep. Negative values indicate truncation
//	            from the beginning.
//	value string - the string to truncate.
//
// Returns:
//
//	string - the truncated string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: trunc].
//
// [Sprout Documentation: trunc]: https://docs.atom.codes/sprout/registries/strings#trunc
func (sr *StringsRegistry) Trunc(count int, value string) string {
	length := len(value)

	if count < 0 && length+count > 0 {
		return value[length+count:]
	}

	if count >= 0 && length > count {
		return value[:count]
	}

	return value
}

// Shuffle randomly rearranges the characters in 'str'.
//
// Parameters:
//
//	value string - the string to shuffle.
//
// Returns:
//
//	string - the shuffled string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: shuffle].
//
// [Sprout Documentation: shuffle]: https://docs.atom.codes/sprout/registries/strings#shuffle
func (sr *StringsRegistry) Shuffle(value string) string {
	runes := []rune(value)
	mathrand.New(randSource).Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	return string(runes)
}

// Ellipsis truncates 'str' to 'maxWidth' and appends an ellipsis if the string
// is longer than 'maxWidth'.
//
// Parameters:
//
//	maxWidth int - the maximum width of the string including the ellipsis.
//	value string - the string to truncate.
//
// Returns:
//
//	string - the possibly truncated string with an ellipsis.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: ellipsis].
//
// [Sprout Documentation: ellipsis]: https://docs.atom.codes/sprout/registries/strings#ellipsis
func (sr *StringsRegistry) Ellipsis(maxWidth int, value string) string {
	return sr.ellipsis(value, 0, maxWidth)
}

// EllipsisBoth truncates 'str' from both ends, preserving the middle part of
// the string and appending ellipses to both ends if needed.
//
// Parameters:
//
//	offset int - starting position for preserving text.
//	maxWidth int - the total maximum width including ellipses.
//	value string - the string to truncate.
//
// Returns:
//
//	string - the truncated string with ellipses on both ends.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: ellipsisBoth].
//
// [Sprout Documentation: ellipsisBoth]: https://docs.atom.codes/sprout/registries/strings#ellipsisboth
func (sr *StringsRegistry) EllipsisBoth(offset int, maxWidth int, value string) string {
	return sr.ellipsis(value, offset, maxWidth)
}

// Initials extracts the initials from 'str', using optional 'delimiters' to
// determine word boundaries.
//
// Parameters:
//
//	value string - the string from which to extract initials.
//
// Returns:
//
//	string - the initials of the words in 'str'.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: initials].
//
// [Sprout Documentation: initials]: https://docs.atom.codes/sprout/registries/strings#initials
func (sr *StringsRegistry) Initials(value string) string {
	return sr.initials(value, " ")
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
// For an example of this function in a Go template, refer to [Sprout Documentation: plural].
//
// [Sprout Documentation: plural]: https://docs.atom.codes/sprout/registries/strings#plural
func (sr *StringsRegistry) Plural(one, many string, count int) string {
	if count == 1 {
		return one
	}
	return many
}

// Wrap breaks 'value' into lines with a maximum length of 'length'.
// It ensures that words are not split across lines unless necessary.
//
// Parameters:
//
//	length int - the maximum length of each line.
//	value string - the string to be wrapped.
//
// Returns:
//
//	string - the wrapped string using newline characters to separate lines.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: wrap].
//
// [Sprout Documentation: wrap]: https://docs.atom.codes/sprout/registries/strings#wrap
func (sr *StringsRegistry) Wrap(length int, value string) string {
	return sr.wordWrap(length, "", false, value)
}

// WrapWith breaks 'value' into lines of maximum 'length', using 'newLineCharacter'
// to separate lines. It wraps words only when they exceed the line length.
//
// Parameters:
//
//	length int - the maximum line length.
//	newLineCharacter string - the character(s) used to denote new lines.
//	value string - the string to wrap.
//
// Returns:
//
//	string - the wrapped string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: wrapWith].
//
// [Sprout Documentation: wrapWith]: https://docs.atom.codes/sprout/registries/strings#wrapwith
func (sr *StringsRegistry) WrapWith(length int, newLineCharacter string, value string) string {
	return sr.wordWrap(length, newLineCharacter, true, value)
}

// Quote wraps each element in 'values' with double quotes and separates them with spaces.
//
// Parameters:
//
//	values ...any - the elements to be quoted.
//
// Returns:
//
//	string - a single string with each element double quoted.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: quote].
//
// [Sprout Documentation: quote]: https://docs.atom.codes/sprout/registries/strings#quote
func (sr *StringsRegistry) Quote(values ...any) string {
	var build strings.Builder

	for i, elem := range values {
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

// Squote wraps each element in 'values' with single quotes and separates them with spaces.
//
// Parameters:
//
//	values ...any - the elements to be single quoted.
//
// Returns:
//
//	string - a single string with each element single quoted.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: squote].
//
// [Sprout Documentation: squote]: https://docs.atom.codes/sprout/registries/strings#squote
func (sr *StringsRegistry) Squote(values ...any) string {
	var builder strings.Builder
	for i, elem := range values {
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

// ToCamelCase converts a string to camelCase.
//
// Parameters:
//
//	value string - the string to convert.
//
// Returns:
//
//	string - the string converted to camelCase.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toCamelCase].
//
// [Sprout Documentation: toCamelCase]: https://docs.atom.codes/sprout/registries/strings#tocamelcase
func (sr *StringsRegistry) ToCamelCase(value string) string {
	return sr.transformString(camelCaseStyle, value)
}

// ToKebabCase converts a string to kebab-case.
//
// Parameters:
//
//	value string - the string to convert.
//
// Returns:
//
//	string - the string converted to kebab-case.
//
// Example:
//
//	{{ "hello world" | toKebabCase }} // Output: "hello-world"
func (sr *StringsRegistry) ToKebabCase(value string) string {
	return sr.transformString(kebabCaseStyle, value)
}

// ToPascalCase converts a string to PascalCase.
//
// Parameters:
//
//	value string - the string to convert.
//
// Returns:
//
//	string - the string converted to PascalCase.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toPascalCase].
//
// [Sprout Documentation: toPascalCase]: https://docs.atom.codes/sprout/registries/strings#topascalcase
func (sr *StringsRegistry) ToPascalCase(value string) string {
	return sr.transformString(pascalCaseStyle, value)
}

// ToDotCase converts a string to dot.case.
//
// Parameters:
//
//	value string - the string to convert.
//
// Returns:
//
//	string - the string converted to dot.case.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toDotCase].
//
// [Sprout Documentation: toDotCase]: https://docs.atom.codes/sprout/registries/strings#todotcase
func (sr *StringsRegistry) ToDotCase(value string) string {
	return sr.transformString(dotCaseStyle, value)
}

// ToPathCase converts a string to path/case.
//
// Parameters:
//
//	value string - the string to convert.
//
// Returns:
//
//	string - the string converted to path/case.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toPathCase].
//
// [Sprout Documentation: toPathCase]: https://docs.atom.codes/sprout/registries/strings#topathcase
func (sr *StringsRegistry) ToPathCase(value string) string {
	return sr.transformString(pathCaseStyle, value)
}

// ToConstantCase converts a string to CONSTANT_CASE.
//
// Parameters:
//
//	value string - the string to convert.
//
// Returns:
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toConstantCase].
//
// [Sprout Documentation: toConstantCase]: https://docs.atom.codes/sprout/registries/strings#toconstantcase
func (sr *StringsRegistry) ToConstantCase(value string) string {
	return sr.transformString(constantCaseStyle, value)
}

// ToSnakeCase converts a string to snake_case.
//
// Parameters:
//
//	value string - the string to convert.
//
// Returns:
//
//	string - the string converted to snake_case.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toSnakeCase].
//
// [Sprout Documentation: toSnakeCase]: https://docs.atom.codes/sprout/registries/strings#tosnakecase
func (sr *StringsRegistry) ToSnakeCase(value string) string {
	return sr.transformString(snakeCaseStyle, value)
}

// ToTitleCase converts a string to Title Case.
//
// Parameters:
//
//	value string - the string to convert.
//
// Returns:
//
//	string - the string converted to Title Case.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toTitleCase].
//
// [Sprout Documentation: toTitleCase]: https://docs.atom.codes/sprout/registries/strings#totitlecase
func (sr *StringsRegistry) ToTitleCase(value string) string {
	return cases.Title(language.English).String(value)
}

// Untitle converts the first letter of each word in 'str' to lowercase.
//
// Parameters:
//
//	value string - the string to be converted.
//
// Returns:
//
//	string - the converted string with each word starting in lowercase.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: untitle].
//
// [Sprout Documentation: untitle]: https://docs.atom.codes/sprout/registries/strings#untitle
func (sr *StringsRegistry) Untitle(value string) string {
	var result strings.Builder

	// Process each rune in the input string
	startOfWord := true
	for _, r := range value {
		if unicode.IsSpace(r) {
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

// SwapCase switches the case of each letter in 'value'. Lowercase letters become
// uppercase and vice versa.
//
// Parameters:
//
//	value string - the string to convert.
//
// Returns:
//
//	string - the string with each character's case switched.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: swapCase].
//
// [Sprout Documentation: swapCase]: https://docs.atom.codes/sprout/registries/strings#swapcase
func (sr *StringsRegistry) SwapCase(value string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLower(r) {
			return unicode.ToUpper(r)
		}
		return unicode.ToLower(r)
	}, value)
}

// Capitalize capitalizes the first letter of 'value'.
//
// Parameters:
//
//	value string - the string to capitalize.
//
// Returns:
//
//	string - the string with the first letter capitalized.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: capitalize].
//
// [Sprout Documentation: capitalize]: https://docs.atom.codes/sprout/registries/strings#capitalize
func (sr *StringsRegistry) Capitalize(value string) string {
	return swapFirstLetter(value, true)
}

// Uncapitalize converts the first letter of 'value' to lowercase.
//
// Parameters:
//
//	value string - the string to uncapitalize.
//
// Returns:
//
//	string - the string with the first letter in lowercase.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: uncapitalize].
//
// [Sprout Documentation: uncapitalize]: https://docs.atom.codes/sprout/registries/strings#uncapitalize
func (sr *StringsRegistry) Uncapitalize(value string) string {
	return swapFirstLetter(value, false)
}

// Split divides 'value' into a map of string parts using 'sep' as the separator.
//
// Parameters:
//
//	sep string - the separator string.
//	value string - the original string to split.
//
// Returns:
//
//	map[string]string - a map of the split parts.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: split].
//
// [Sprout Documentation: split]: https://docs.atom.codes/sprout/registries/strings#split
func (sr *StringsRegistry) Split(sep, value string) map[string]string {
	parts := strings.Split(value, sep)
	return sr.populateMapWithParts(parts)
}

// Splitn divides 'value' into a map of string parts using 'sep' as the separator
// up to 'n' parts.
//
// Parameters:
//
//	sep string - the separator string.
//	n int - the maximum number of substrings to return.
//	value string - the original string to split.
//
// Returns:
//
//	map[string]string - a map of the split parts.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: splitn].
//
// [Sprout Documentation: splitn]: https://docs.atom.codes/sprout/registries/strings#splitn
func (sr *StringsRegistry) Splitn(sep string, n int, value string) map[string]string {
	parts := strings.SplitN(value, sep, n)
	return sr.populateMapWithParts(parts)
}

// Substring extracts a substring from 's' starting at 'start' and ending at 'end'.
// Negative values for 'start' or 'end' are interpreted as positions from the end
// of the string.
//
// Parameters:
//
//	start int - the starting index.
//	end int - the ending index, exclusive.
//	value string - the source string.
//
// Returns:
//
//	string - the extracted substring.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: substr].
//
// [Sprout Documentation: substr]: https://docs.atom.codes/sprout/registries/strings#substr
func (sr *StringsRegistry) Substring(start, end int, value string) string {
	length := len(value)

	if start < 0 {
		start = length + start
	}
	if end < 0 {
		end = length + end
	}
	if start < 0 {
		start = 0
	}
	if end > length || end == 0 {
		end = length
	}
	if start > end {
		return ""
	}
	return value[start:end]
}

// Indent adds spaces to the beginning of each line in 'value'.
//
// Parameters:
//
//	spaces int - the number of spaces to add.
//	value string - the string to indent.
//
// Returns:
//
//	string - the indented string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: indent].
//
// [Sprout Documentation: indent]: https://docs.atom.codes/sprout/registries/strings#indent
func (sr *StringsRegistry) Indent(spaces int, value string) string {
	var builder strings.Builder
	pad := strings.Repeat(" ", spaces)
	lines := strings.Split(value, "\n")

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
//	spaces int - the number of spaces to add after the newline.
//	value string - the string to indent.
//
// Returns:
//	string - the indented string with a newline at the start.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: nindent].
//
// [Sprout Documentation: nindent]: https://docs.atom.codes/sprout/registries/strings#nindent

func (sr *StringsRegistry) Nindent(spaces int, value string) string {
	return "\n" + sr.Indent(spaces, value)
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
// For an example of this function in a Go template, refer to [Sprout Documentation: seq].
//
// [Sprout Documentation: seq]: https://docs.atom.codes/sprout/registries/strings#seq
func (sr *StringsRegistry) Seq(params ...int) string {
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
		return sr.convertIntArrayToString(helpers.UntilStep(start, end+increment, increment), " ")
	case 2:
		start := params[0]
		end := params[1]
		step := 1
		if end < start {
			step = -1
		}
		return sr.convertIntArrayToString(helpers.UntilStep(start, end+step, step), " ")
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
		return sr.convertIntArrayToString(helpers.UntilStep(start, end+increment, step), " ")
	default:
		return ""
	}
}
