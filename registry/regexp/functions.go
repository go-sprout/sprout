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
// For an example of this function in a Go template, refer to [Sprout Documentation: regexQuoteMeta].
//
// [Sprout Documentation: regexQuoteMeta]: https://docs.atom.codes/sprout/registries/regexp#regexquotemeta
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
// For an example of this function in a Go template, refer to [Sprout Documentation: regexFind].
//
// [Sprout Documentation: regexFind]: https://docs.atom.codes/sprout/registries/regexp#regexfind
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
// For an example of this function in a Go template, refer to [Sprout Documentation: regexFindAll].
//
// [Sprout Documentation: regexFindAll]: https://docs.atom.codes/sprout/registries/regexp#regexfindall
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
// For an example of this function in a Go template, refer to [Sprout Documentation: regexMatch].
//
// [Sprout Documentation: regexMatch]: https://docs.atom.codes/sprout/registries/regexp#regexmatch
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
// For an example of this function in a Go template, refer to [Sprout Documentation: regexSplit].
//
// [Sprout Documentation: regexSplit]: https://docs.atom.codes/sprout/registries/regexp#regexsplit
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
// For an example of this function in a Go template, refer to [Sprout Documentation: regexReplaceAll].
//
// [Sprout Documentation: regexReplaceAll]: https://docs.atom.codes/sprout/registries/regexp#regexreplaceall
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
// For an example of this function in a Go template, refer to [Sprout Documentation: regexReplaceAllLiteral].
//
// [Sprout Documentation: regexReplaceAllLiteral]: https://docs.atom.codes/sprout/registries/regexp#regexreplaceallliteral
func (rr *RegexpRegistry) RegexReplaceAllLiteral(regex string, s string, replacedBy string) (string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}
	return r.ReplaceAllLiteralString(s, replacedBy), nil
}

// RegexFindGroups finds the first match of a regex pattern in a string and
// returns the matched groups, with error handling.
//
// Parameters:
//
//	regex string - the regular expression to search with.
//	str string - the string to search within.
//
// Returns:
//
//	[]string - the matched groups from the first regex match found.
//	error - error if the regex fails to compile.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: regexFindGroups].
//
// [Sprout Documentation: regexFindGroups]: https://docs.atom.codes/sprout/registries/regexp#regexfindgroups
func (rr *RegexpRegistry) RegexFindGroups(regex string, str string) ([]string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return []string{}, err
	}
	matches := r.FindStringSubmatch(str)
	if len(matches) == 0 {
		return []string{}, nil
	}
	return matches, nil
}

// RegexFindAllGroups finds all matches of a regex pattern in a string up
// to a specified limit and returns the matched groups, with error handling.
//
// Parameters:
//
//	regex string - the regular expression to search with.
//	str string - the string to search within.
//	n int - the maximum number of matches to return; use -1 for no limit.
//
// Returns:
//
//	[][]string - a slice containing the matched groups for each match found.
//	error - error if the regex fails to compile.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: regexFindAllGroups].
//
// [Sprout Documentation: regexFindAllGroups]: https://docs.atom.codes/sprout/registries/regexp#regexfindallgroups
func (rr *RegexpRegistry) RegexFindAllGroups(regex string, n int, str string) ([][]string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return [][]string{}, err
	}
	matches := r.FindAllStringSubmatch(str, n)
	if len(matches) == 0 {
		return [][]string{}, nil
	}
	return matches, nil
}

// RegexFindNamed finds the first match of a regex pattern with named
// capturing groups in a string and returns a map of group names to matched
// strings, with error handling.
//
// Parameters:
//
//	regex string - the regular expression to search with, containing named capturing groups.
//	str string - the string to search within.
//
// Returns:
//
//	map[string]string - a map of group names to their corresponding matched strings.
//	error - error if the regex fails to compile.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: regexFindNamed].
//
// [Sprout Documentation: regexFindNamed]: https://docs.atom.codes/sprout/registries/regexp#regexfindnamed
func (rr *RegexpRegistry) RegexFindNamed(regex string, str string) (map[string]string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return map[string]string{}, err
	}
	matches := r.FindStringSubmatch(str)
	if len(matches) == 0 {
		return map[string]string{}, nil
	}

	result := make(map[string]string, len(r.SubexpNames())-1)
	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = matches[i]
		}
	}
	return result, nil
}

// RegexFindAllNamed finds all matches of a regex pattern with named capturing
// groups in a string up to a specified limit and returns a slice of maps of
// group names to matched strings, with error handling.
//
// Parameters:
//
//	regex string - the regular expression to search with, containing named capturing groups.
//	str string - the string to search within.
//	n int - the maximum number of matches to return; use -1 for no limit.
//
// Returns:
//
//	[]map[string]string - a slice containing a map of group names to their corresponding matched strings for each match found.
//	error - error if the regex fails to compile.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: regexFindAllNamed].
//
// [Sprout Documentation: regexFindAllNamed]: https://docs.atom.codes/sprout/registries/regexp#regexfindallnamed
func (rr *RegexpRegistry) RegexFindAllNamed(regex string, n int, str string) ([]map[string]string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return []map[string]string{}, err
	}
	matches := r.FindAllStringSubmatch(str, n)
	if len(matches) == 0 {
		return []map[string]string{}, nil
	}

	subexpNames := r.SubexpNames()
	results := make([]map[string]string, 0, len(matches))

	for _, match := range matches {
		m := make(map[string]string, len(subexpNames))
		for i, name := range subexpNames {
			if i != 0 && name != "" {
				m[name] = match[i]
			}
		}
		results = append(results, m)
	}
	return results, nil
}
