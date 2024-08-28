package regexp

import (
	"regexp"
)

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
func (rr *RegexpRegistry) RegexQuoteMeta(str string) string {
	return regexp.QuoteMeta(str)
}

// RegexFind searches for the first match of a regex pattern in a string
// and returns it, with error handling.
//
// Parameters:
//
//	regex string - the regular expression to search with.
//	str string - the string to search within.
//
// Returns:
//
//	string - the first regex match found.
//	error - error if the regex fails to compile.
//
// Example:
//
//	{{ "hello world" | RegexFind "hello" }} // Output: "hello", nil
func (rr *RegexpRegistry) RegexFind(regex string, str string) (string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}
	return r.FindString(str), nil
}

// RegexFindAll finds all matches of a regex pattern in a string up to a
// specified limit, with error handling.
//
// Parameters:
//
//	regex string - the regular expression to search with.
//	str string - the string to search within.
//	n int - the maximum number of matches to return; use -1 for no limit.
//
// Returns:
//
//	[]string - all regex matches found.
//	error - error if the regex fails to compile.
//
// Example:
//
//	{{ RegexFindAll "a.", "aba acada afa", 3 }} // Output: ["ab", "ac", "af"], nil
func (rr *RegexpRegistry) RegexFindAll(regex string, str string, n int) ([]string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return []string{}, err
	}
	return r.FindAllString(str, n), nil
}

// RegexMatch checks if a string matches a regex pattern, with error handling.
//
// Parameters:
//
//	regex string - the regular expression to match against.
//	str string - the string to check.
//
// Returns:
//
//	bool - true if the string matches the regex pattern, otherwise false.
//	error - error if the regex fails to compile.
//
// Example:
//
//	{{ RegexMatch "^[a-zA-Z]+$", "Hello" }} // Output: true, nil
func (rr *RegexpRegistry) RegexMatch(regex string, str string) (bool, error) {
	return regexp.MatchString(regex, str)
}

// RegexSplit splits a string by a regex pattern up to a specified number of
// substrings, with error handling.
//
// Parameters:
//
//	regex string - the regular expression to split by.
//	str string - the string to split.
//	n int - the maximum number of substrings to return; use -1 for no limit.
//
// Returns:
//
//	[]string - the substrings resulting from the split.
//	error - error if the regex fails to compile.
//
// Example:
//
//	{{ RegexSplit "\\s+", "hello world from Go", 2 }} // Output: ["hello", "world from Go"], nil
func (rr *RegexpRegistry) RegexSplit(regex string, str string, n int) ([]string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return []string{}, err
	}
	return r.Split(str, n), nil
}

// RegexReplaceAll replaces all occurrences of a regex pattern in a string
// with a replacement string, with error handling.
//
// Parameters:
//
//	regex string - the regular expression to replace.
//	str string - the string containing the original text.
//	replacedBy string - the replacement text.
//
// Returns:
//
//	string - the modified string after all replacements.
//	error - error if the regex fails to compile.
//
// Example:
//
//	{{ RegexReplaceAll "\\d", "R2D2 C3PO", "X" }} // Output: "RXDX CXPO", nil
func (rr *RegexpRegistry) RegexReplaceAll(regex string, str string, replacedBy string) (string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}
	return r.ReplaceAllString(str, replacedBy), nil
}

// RegexReplaceAllLiteral replaces all occurrences of a regex pattern in a
// string with a literal replacement string, with error handling.
//
// Parameters:
//
//	regex string - the regular expression to replace.
//	s string - the string containing the original text.
//	replacedBy string - the literal replacement text.
//
// Returns:
//
//	string - the modified string after all replacements, treating the replacement text as literal text.
//	error - error if the regex fails to compile.
//
// Example:
//
//	{{ RegexReplaceAllLiteral "world", "hello world", "$1" }} // Output: "hello $1", nil
func (rr *RegexpRegistry) RegexReplaceAllLiteral(regex string, s string, replacedBy string) (string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}
	return r.ReplaceAllLiteralString(s, replacedBy), nil
}
