package regexp

import (
	"regexp"
)

// RegexFind returns the first match of the regex pattern in the string.
//
// Parameters:
//
//	regex string - the regular expression pattern to search for.
//	s string - the string to search.
//
// Returns:
//
//	string - the first matching string.
//
// Example:
//
//	{{ regexFind "a(b+)" "aaabbb" }} // Output: "abbb"
func (rr *RegexpRegistry) RegexFind(regex string, s string) string {
	result, _ := rr.MustRegexFind(regex, s)
	return result
}

// RegexFindAll returns all matches of the regex pattern in the string up to n
// matches.
//
// Parameters:
//
//	regex string - the regular expression pattern to search for.
//	s string - the string to search.
//	n int - the maximum number of matches to return.
//
// Returns:
//
//	[]string - a slice of all matches.
//
// Example:
//
//	{{ regexFindAll "a(b+)" "aaabbb" 2 }} // Output: ["abbb"]
func (rr *RegexpRegistry) RegexFindAll(regex string, s string, n int) []string {
	result, _ := rr.MustRegexFindAll(regex, s, n)
	return result
}

// RegexMatch checks if the string matches the regex pattern.
//
// Parameters:
//
//	regex string - the regular expression pattern to match against.
//	s string - the string to check.
//
// Returns:
//
//	bool - true if the string matches the regex pattern, otherwise false.
//
// Example:
//
//	{{ regexMatch "^[a-zA-Z]+$" "Hello" }} // Output: true
func (rr *RegexpRegistry) RegexMatch(regex string, s string) bool {
	result, _ := rr.MustRegexMatch(regex, s)
	return result
}

// RegexSplit splits the string by the regex pattern up to n times.
//
// Parameters:
//
//	regex string - the regular expression pattern to split by.
//	s string - the string to split.
//	n int - the number of times to split.
//
// Returns:
//
//	[]string - a slice of the substrings split by the regex.
//
// Example:
//
//	{{regexSplit "\\s+" "hello world" -1 }} // Output: ["hello", "world"]
func (rr *RegexpRegistry) RegexSplit(regex string, s string, n int) []string {
	result, _ := rr.MustRegexSplit(regex, s, n)
	return result
}

// RegexReplaceAll replaces all occurrences of the regex pattern in the string
// with the replacement string.
//
// Parameters:
//
//	regex string - the regular expression pattern to replace.
//	s string - the string to perform replacements on.
//	repl string - the replacement string.
//
// Returns:
//
//	string - the string with all replacements made.
//
// Example:
//
//	{{ regexReplaceAll "[aeiou]" "hello" "i" }} // Output: "hillo"
func (rr *RegexpRegistry) RegexReplaceAll(regex string, s string, repl string) string {
	result, _ := rr.MustRegexReplaceAll(regex, s, repl)
	return result
}

// RegexReplaceAllLiteral replaces all occurrences of the regex pattern in the
// string with the literal replacement string.
//
// Parameters:
//
//	regex string - the regular expression pattern to replace.
//	s string - the string to perform replacements on.
//	repl string - the replacement string, inserted literally.
//
// Returns:
//
//	string - the string with all replacements made, without treating the replacement string as a regex replacement pattern.
//
// Example:
//
//	{{ regexReplaceAllLiteral "[aeiou]" "hello" "$&" }} // Output: "h$&ll$&"
func (rr *RegexpRegistry) RegexReplaceAllLiteral(regex string, s string, repl string) string {
	result, _ := rr.MustRegexReplaceAllLiteral(regex, s, repl)
	return result
}

// RegexQuoteMeta returns a literal pattern string for the provided string.
//
// Parameters:
//
//	s string - the string to be escaped.
//
// Returns:
//
//	string - the escaped regex pattern.
//
// Example:
//
//	{{ regexQuoteMeta ".+*?^$()[]{}|" }} // Output: "\.\+\*\?\^\$\(\)\[\]\{\}\|"
func (rr *RegexpRegistry) RegexQuoteMeta(s string) string {
	return regexp.QuoteMeta(s)
}

// MustRegexFind searches for the first match of a regex pattern in a string
// and returns it, with error handling.
//
// Parameters:
//
//	regex string - the regular expression to search with.
//	s string - the string to search within.
//
// Returns:
//
//	string - the first regex match found.
//	error - error if the regex fails to compile.
//
// Example:
//
//	{{ "hello world" | mustRegexFind "hello" }} // Output: "hello", nil
func (rr *RegexpRegistry) MustRegexFind(regex string, s string) (string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}
	return r.FindString(s), nil
}

// MustRegexFindAll finds all matches of a regex pattern in a string up to a
// specified limit, with error handling.
//
// Parameters:
//
//	regex string - the regular expression to search with.
//	s string - the string to search within.
//	n int - the maximum number of matches to return; use -1 for no limit.
//
// Returns:
//
//	[]string - all regex matches found.
//	error - error if the regex fails to compile.
//
// Example:
//
//	{{ mustRegexFindAll "a.", "aba acada afa", 3 }} // Output: ["ab", "ac", "af"], nil
func (rr *RegexpRegistry) MustRegexFindAll(regex string, s string, n int) ([]string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return []string{}, err
	}
	return r.FindAllString(s, n), nil
}

// MustRegexMatch checks if a string matches a regex pattern, with error handling.
//
// Parameters:
//
//	regex string - the regular expression to match against.
//	s string - the string to check.
//
// Returns:
//
//	bool - true if the string matches the regex pattern, otherwise false.
//	error - error if the regex fails to compile.
//
// Example:
//
//	{{ mustRegexMatch "^[a-zA-Z]+$", "Hello" }} // Output: true, nil
func (rr *RegexpRegistry) MustRegexMatch(regex string, s string) (bool, error) {
	return regexp.MatchString(regex, s)
}

// MustRegexSplit splits a string by a regex pattern up to a specified number of
// substrings, with error handling.
//
// Parameters:
//
//	regex string - the regular expression to split by.
//	s string - the string to split.
//	n int - the maximum number of substrings to return; use -1 for no limit.
//
// Returns:
//
//	[]string - the substrings resulting from the split.
//	error - error if the regex fails to compile.
//
// Example:
//
//	{{ mustRegexSplit "\\s+", "hello world from Go", 2 }} // Output: ["hello", "world from Go"], nil
func (rr *RegexpRegistry) MustRegexSplit(regex string, s string, n int) ([]string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return []string{}, err
	}
	return r.Split(s, n), nil
}

// MustRegexReplaceAll replaces all occurrences of a regex pattern in a string
// with a replacement string, with error handling.
//
// Parameters:
//
//	regex string - the regular expression to replace.
//	s string - the string containing the original text.
//	repl string - the replacement text.
//
// Returns:
//
//	string - the modified string after all replacements.
//	error - error if the regex fails to compile.
//
// Example:
//
//	{{ mustRegexReplaceAll "\\d", "R2D2 C3PO", "X" }} // Output: "RXDX CXPO", nil
func (rr *RegexpRegistry) MustRegexReplaceAll(regex string, s string, repl string) (string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}
	return r.ReplaceAllString(s, repl), nil
}

// MustRegexReplaceAllLiteral replaces all occurrences of a regex pattern in a
// string with a literal replacement string, with error handling.
//
// Parameters:
//
//	regex string - the regular expression to replace.
//	s string - the string containing the original text.
//	repl string - the literal replacement text.
//
// Returns:
//
//	string - the modified string after all replacements, treating the replacement text as literal text.
//	error - error if the regex fails to compile.
//
// Example:
//
//	{{ mustRegexReplaceAllLiteral "world", "hello world", "$1" }} // Output: "hello $1", nil
func (rr *RegexpRegistry) MustRegexReplaceAllLiteral(regex string, s string, repl string) (string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}
	return r.ReplaceAllLiteralString(s, repl), nil
}
