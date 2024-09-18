package regexp_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/regexp"
)

func TestRegexpFind(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestRegexpFind", Input: `{{ regexFind "a(b+)" "aaabbb" }}`, ExpectedOutput: "abbb"},
		{Name: "TestRegexpFindError", Input: `{{ regexFind "a(b+" "aaabbb" }}`, ExpectedErr: "error parsing regexp"},
	}

	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestRegexpFindAll(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestRegexpFindAllWithoutLimit", Input: `{{ regexFindAll "a(b+)" "aaabbb" -1 }}`, ExpectedOutput: "[abbb]"},
		{Name: "TestRegexpFindAllWithLimit", Input: `{{ regexFindAll "a{2}" "aaaabbb" -1 }}`, ExpectedOutput: "[aa aa]"},
		{Name: "TestRegexpFindAllWithNoMatch", Input: `{{ regexFindAll "a{2}" "none" -1 }}`, ExpectedOutput: "[]"},
		{Name: "TestRegexpFindAllWithInvalidPattern", Input: `{{ regexFindAll "a(b+" "aaabbb" -1 }}`, ExpectedErr: "error parsing regexp"},
	}

	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestRegexMatch(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestRegexMatchValid", Input: `{{ regexMatch "^[a-zA-Z]+$" "Hello" }}`, ExpectedOutput: "true"},
		{Name: "TestRegexMatchInvalidAlphaNumeric", Input: `{{ regexMatch "^[a-zA-Z]+$" "Hello123" }}`, ExpectedOutput: "false"},
		{Name: "TestRegexMatchInvalidNumeric", Input: `{{ regexMatch "^[a-zA-Z]+$" "123" }}`, ExpectedOutput: "false"},
	}

	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestRegexSplit(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestRegexpFindAllWithoutLimit", Input: `{{ regexSplit "a" "banana" -1 }}`, ExpectedOutput: "[b n n ]"},
		{Name: "TestRegexpSplitZeroLimit", Input: `{{ regexSplit "a" "banana" 0 }}`, ExpectedOutput: "[]"},
		{Name: "TestRegexpSplitOneLimit", Input: `{{ regexSplit "a" "banana" 1 }}`, ExpectedOutput: "[banana]"},
		{Name: "TestRegexpSplitTwoLimit", Input: `{{ regexSplit "a" "banana" 2 }}`, ExpectedOutput: "[b nana]"},
		{Name: "TestRegexpSplitRepetitionLimit", Input: `{{ regexSplit "a+" "banana" 1 }}`, ExpectedOutput: "[banana]"},
	}

	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestRegexReplaceAll(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestRegexReplaceAllValid", Input: `{{ regexReplaceAll "a(x*)b" "-ab-axxb-" "T" }}`, ExpectedOutput: "-T-T-"},
		{Name: "TestRegexReplaceAllWithDollarSign", Input: `{{ regexReplaceAll "a(x*)b" "-ab-axxb-" "$1" }}`, ExpectedOutput: "--xx-"},
		{Name: "TestRegexReplaceAllWithDollarSignAndLetter", Input: `{{ regexReplaceAll "a(x*)b" "-ab-axxb-" "$1W" }}`, ExpectedOutput: "---"},
		{Name: "TestRegexReplaceAllWithDollarSignAndCurlyBraces", Input: `{{ regexReplaceAll "a(x*)b" "-ab-axxb-" "${1}W" }}`, ExpectedOutput: "-W-xxW-"},
	}

	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestRegexReplaceAllLiteral(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestRegexReplaceAllLiteralValid", Input: `{{ regexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "T" }}`, ExpectedOutput: "-T-T-"},
		{Name: "TestRegexReplaceAllLiteralWithDollarSign", Input: `{{ regexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "$1" }}`, ExpectedOutput: "-$1-$1-"},
		{Name: "TestRegexReplaceAllLiteralWithDollarSignAndLetter", Input: `{{ regexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "$1W" }}`, ExpectedOutput: "-$1W-$1W-"},
		{Name: "TestRegexReplaceAllLiteralWithDollarSignAndCurlyBraces", Input: `{{ regexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "${1}W" }}`, ExpectedOutput: "-${1}W-${1}W-"},
	}

	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestRegexQuoteMeta(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestRegexQuoteMetaALongLine", Input: `{{ regexQuoteMeta "Escaping $100? That's a lot." }}`, ExpectedOutput: "Escaping \\$100\\? That's a lot\\."},
		{Name: "TestRegexQuoteMetaASemVer", Input: `{{ regexQuoteMeta "1.2.3" }}`, ExpectedOutput: "1\\.2\\.3"},
		{Name: "TestRegexQuoteMetaNothing", Input: `{{ regexQuoteMeta "golang" }}`, ExpectedOutput: "golang"},
	}

	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestMustRegexFind(t *testing.T) {
	tc := []pesticide.TestCase{
		{
			Name:           "TestMustRegexFindValid",
			Input:          `{{ mustRegexFind "a(b+)" "aaabbb" }}`,
			ExpectedOutput: "abbb",
			ExpectedErr:    "",
		},
		{
			Name:           "TestMustRegexFindInvalid",
			Input:          `{{ mustRegexFind "a(b+" "aaabbb" }}`,
			ExpectedOutput: "",
			ExpectedErr:    "error parsing regexp: missing closing ): `a(b+`",
		},
	}

	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestMustRegexFindAll(t *testing.T) {
	tc := []pesticide.TestCase{
		{
			Name:           "TestMustRegexFindAllValid",
			Input:          `{{ mustRegexFindAll "a(b+)" "aaabbb" -1 }}`,
			ExpectedOutput: "[abbb]",
			ExpectedErr:    "",
		},
		{
			Name:           "TestMustRegexFindAllWithLimit",
			Input:          `{{ mustRegexFindAll "a{2}" "aaaabbb" -1 }}`,
			ExpectedOutput: "[aa aa]",
			ExpectedErr:    "",
		},
		{
			Name:           "TestMustRegexFindAllWithNoMatch",
			Input:          `{{ mustRegexFindAll "a{2}" "none" -1 }}`,
			ExpectedOutput: "[]",
			ExpectedErr:    "",
		},
		{
			Name:           "TestMustRegexFindAllWithInvalidPattern",
			Input:          `{{ mustRegexFindAll "a(b+" "aaabbb" -1 }}`,
			ExpectedOutput: "",
			ExpectedErr:    "error parsing regexp: missing closing ): `a(b+`",
		},
	}

	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestMustRegexMatch(t *testing.T) {
	tc := []pesticide.TestCase{
		{
			Name:           "TestMustRegexMatchValid",
			Input:          `{{ mustRegexMatch "^[a-zA-Z]+$" "Hello" }}`,
			ExpectedOutput: "true",
			ExpectedErr:    "",
		},
		{
			Name:           "TestMustRegexMatchInvalidAlphaNumeric",
			Input:          `{{ mustRegexMatch "^[a-zA-Z]+$" "Hello123" }}`,
			ExpectedOutput: "false",
			ExpectedErr:    "",
		},
		{
			Name:           "TestMustRegexMatchInvalidNumeric",
			Input:          `{{ mustRegexMatch "^[a-zA-Z]+$" "123" }}`,
			ExpectedOutput: "false",
			ExpectedErr:    "",
		},
		{
			Name:           "TestMustRegexMatchInvalidPattern",
			Input:          `{{ mustRegexMatch "^[a-zA+$" "Hello" }}`,
			ExpectedOutput: "",
			ExpectedErr:    "error parsing regexp: missing closing ]: `[a-zA+$`",
		},
	}

	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestMustRegexSplit(t *testing.T) {
	tc := []pesticide.TestCase{
		{
			Name:           "TestMustRegexSplitWithoutLimit",
			Input:          `{{ mustRegexSplit "a" "banana" -1 }}`,
			ExpectedOutput: "[b n n ]",
			ExpectedErr:    "",
		},
		{
			Name:           "TestMustRegexSplitZeroLimit",
			Input:          `{{ mustRegexSplit "a" "banana" 0 }}`,
			ExpectedOutput: "[]",
			ExpectedErr:    "",
		},
		{
			Name:           "TestMustRegexSplitOneLimit",
			Input:          `{{ mustRegexSplit "a" "banana" 1 }}`,
			ExpectedOutput: "[banana]",
			ExpectedErr:    "",
		},
		{
			Name:           "TestMustRegexSplitTwoLimit",
			Input:          `{{ mustRegexSplit "a" "banana" 2 }}`,
			ExpectedOutput: "[b nana]",
			ExpectedErr:    "",
		},
		{
			Name:           "TestMustRegexSplitRepetitionLimit",
			Input:          `{{ mustRegexSplit "a+" "banana" 1 }}`,
			ExpectedOutput: "[banana]",
			ExpectedErr:    "",
		},
		{
			Name:           "TestMustRegexSplitInvalidPattern",
			Input:          `{{ mustRegexSplit "+" "banana" 0 }}`,
			ExpectedOutput: "",
			ExpectedErr:    "error parsing regexp: missing argument to repetition operator: `+`",
		},
	}

	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestMustRegexReplaceAll(t *testing.T) {
	tc := []pesticide.TestCase{
		{
			Name:           "TestMustRegexReplaceAllValid",
			Input:          `{{ mustRegexReplaceAll "a(x*)b" "-ab-axxb-" "T" }}`,
			ExpectedOutput: "-T-T-",
			ExpectedErr:    "",
		},
		{
			Name:           "TestMustRegexReplaceAllWithDollarSign",
			Input:          `{{ mustRegexReplaceAll "a(x*)b" "-ab-axxb-" "$1" }}`,
			ExpectedOutput: "--xx-",
			ExpectedErr:    "",
		},
		{
			Name:           "TestMustRegexReplaceAllWithDollarSignAndLetter",
			Input:          `{{ mustRegexReplaceAll "a(x*)b" "-ab-axxb-" "$1W" }}`,
			ExpectedOutput: "---",
			ExpectedErr:    "",
		},
		{
			Name:           "TestMustRegexReplaceAllWithDollarSignAndCurlyBraces",
			Input:          `{{ mustRegexReplaceAll "a(x*)b" "-ab-axxb-" "${1}W" }}`,
			ExpectedOutput: "-W-xxW-",
			ExpectedErr:    "",
		},
		{
			Name:           "TestMustRegexReplaceAllWithInvalidPattern",
			Input:          `{{ mustRegexReplaceAll "a(x*}" "-ab-axxb-" "T" }}`,
			ExpectedOutput: "",
			ExpectedErr:    "error parsing regexp: missing closing ): `a(x*}`",
		},
	}

	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestMustRegexReplaceAllLiteral(t *testing.T) {
	tc := []pesticide.TestCase{
		{
			Name:           "TestMustRegexReplaceAllLiteralValid",
			Input:          `{{ mustRegexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "T" }}`,
			ExpectedOutput: "-T-T-",
			ExpectedErr:    "",
		},
		{
			Name:           "TestMustRegexReplaceAllLiteralWithDollarSign",
			Input:          `{{ mustRegexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "$1" }}`,
			ExpectedOutput: "-$1-$1-",
			ExpectedErr:    "",
		},
		{
			Name:           "TestMustRegexReplaceAllLiteralWithDollarSignAndLetter",
			Input:          `{{ mustRegexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "$1W" }}`,
			ExpectedOutput: "-$1W-$1W-",
			ExpectedErr:    "",
		},
		{
			Name:           "TestMustRegexReplaceAllLiteralWithDollarSignAndCurlyBraces",
			Input:          `{{ mustRegexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "${1}W" }}`,
			ExpectedOutput: "-${1}W-${1}W-",
			ExpectedErr:    "",
		},
		{
			Name:           "TestMustRegexReplaceAllLiteralWithInvalidPattern",
			Input:          `{{ mustRegexReplaceAllLiteral "a(x*}" "-ab-axxb-" "T" }}`,
			ExpectedOutput: "",
			ExpectedErr:    "error parsing regexp: missing closing ): `a(x*}`",
		},
	}

	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestRegexFindNamed(t *testing.T) {
	tc := []pesticide.TestCase{
		{
			Name:           "TestRegexFindNamedValid",
			Input:          `{{ .V | regexFindNamed "(?P<username>[A-Za-z]+)@(?P<domain>example\\.com)" }}`,
			Data:           map[string]any{"V": "Contact us at noreply@example.com"},
			ExpectedOutput: "map[domain:example.com username:noreply]",
		},
		{
			Name:           "TestRegexpFindNamedWithUnnamedGroup",
			Input:          `{{ .V | regexFindNamed "(?P<username>[A-Za-z]+)@(example\\.com)" }}`,
			Data:           map[string]any{"V": "Contact us at noreply@example.com"},
			ExpectedOutput: "map[username:noreply]",
		},
		{
			Name:           "TestRegexFindNamedNoMatch",
			Input:          `{{ .V | regexFindNamed "(?P<username>[A-Za-z]+)@(?P<domain>example\\.org)" }}`,
			Data:           map[string]any{"V": "Contact us at noreply@example.com"},
			ExpectedOutput: "map[]",
		},
		{
			Name:        "TestRegexFindNamedInvalidPattern",
			Input:       `{{ .V | regexFindNamed "(?P<username>[A-Za-z]+)@(?P<domain>example\\.com" }}`,
			Data:        map[string]any{"V": "Contact us at noreply@example.com"},
			ExpectedErr: "error parsing regexp",
		},
	}
	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestRegexFindAllNamed(t *testing.T) {
	tc := []pesticide.TestCase{
		{
			Name:           "TestRegexFindAllNamedValid",
			Input:          `{{ regexFindAllNamed "(?P<param>\\w+)=(?P<value>\\w+)" -1 "var1=value1&var2=value2&var3=value3" }}`,
			ExpectedOutput: "[map[param:var1 value:value1] map[param:var2 value:value2] map[param:var3 value:value3]]",
		},
		{
			Name:           "TestRegexFindAllNamedWithLimit",
			Input:          `{{ regexFindAllNamed "(?P<param>\\w+)=(?P<value>\\w+)" 2 "var1=value1&var2=value2&var3=value3" }}`,
			ExpectedOutput: "[map[param:var1 value:value1] map[param:var2 value:value2]]",
		},
		{
			Name:           "TestRegexFindAllNamedNoMatch",
			Input:          `{{ regexFindAllNamed "(?P<param>\\d+)=(?P<value>\\d+)" -1 "var1=value1&var2=value2&var3=value3" }}`,
			ExpectedOutput: "[]",
		},
		{
			Name:        "TestRegexFindAllNamedInvalidPattern",
			Input:       `{{ regexFindAllNamed "(?P<param>\\w+)=(?P<value>\\w+" -1 "var1=value1&var2=value2&var3=value3" }}`,
			ExpectedErr: "error parsing regexp",
		},
	}
	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestRegexFindGroups(t *testing.T) {
	tc := []pesticide.TestCase{
		{
			Name:           "TestRegexFindGroupsValid",
			Input:          `{{ regexFindGroups "([A-Za-z]+)@(example\\.com)" "Contact us at support@example.com" }}`,
			ExpectedOutput: "[support@example.com support example.com]",
		},
		{
			Name:           "TestRegexFindGroupsNoMatch",
			Input:          `{{ regexFindGroups "([A-Za-z]+)@(example\\.org)" "Contact us at support@example.com" }}`,
			ExpectedOutput: "[]",
		},
		{
			Name:        "TestRegexFindGroupsInvalidPattern",
			Input:       `{{ regexFindGroups "([A-Za-z]+)@(example\\.com" "Contact us at support@example.com" }}`,
			ExpectedErr: "error parsing regexp",
		},
	}
	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestRegexFindAllGroups(t *testing.T) {
	tc := []pesticide.TestCase{
		{
			Name:           "TestRegexFindAllGroupsValid",
			Input:          `{{ regexFindAllGroups "(\\w+)=(\\w+)" -1 "var1=value1&var2=value2&var3=value3" }}`,
			ExpectedOutput: "[[var1=value1 var1 value1] [var2=value2 var2 value2] [var3=value3 var3 value3]]",
		},
		{
			Name:           "TestRegexFindAllGroupsWithLimit",
			Input:          `{{ regexFindAllGroups "(\\w+)=(\\w+)" 2 "var1=value1&var2=value2&var3=value3" }}`,
			ExpectedOutput: "[[var1=value1 var1 value1] [var2=value2 var2 value2]]",
		},
		{
			Name:           "TestRegexFindAllGroupsNoMatch",
			Input:          `{{ regexFindAllGroups "(\\d+)=(\\d+)" -1 "var1=value1&var2=value2&var3=value3" }}`,
			ExpectedOutput: "[]",
		},
		{
			Name:        "TestRegexFindAllGroupsInvalidPattern",
			Input:       `{{ regexFindAllGroups "(\\w+)=(\\w+" -1 "var1=value1&var2=value2&var3=value3" }}`,
			ExpectedErr: "error parsing regexp",
		},
	}
	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}
