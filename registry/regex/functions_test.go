package regex_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/regex"
)

func TestRegexFind(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestRegexFind", Input: `{{ "aaabbb" | regexFind "a(b+)" }}`, ExpectedOutput: "abbb"},
		{Name: "TestRegexFindError", Input: `{{ "aaabbb" | regexFind "a(b+" }}`, ExpectedErr: "error parsing regexp"},
	}

	pesticide.RunTestCases(t, regex.NewRegistry(), tc)
}

func TestRegexFindAll(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestRegexFindAllWithoutLimit", Input: `{{ "aaabbb" | regexFindAll "a(b+)" -1 }}`, ExpectedOutput: "[abbb]"},
		{Name: "TestRegexFindAllWithLimit", Input: `{{ "aaaabbb" | regexFindAll "a{2}" -1 }}`, ExpectedOutput: "[aa aa]"},
		{Name: "TestRegexFindAllWithNoMatch", Input: `{{ "none" | regexFindAll "a{2}" -1 }}`, ExpectedOutput: "[]"},
		{Name: "TestRegexFindAllWithInvalidPattern", Input: `{{ "aaabbb" | regexFindAll "a(b+" -1 }}`, ExpectedErr: "error parsing regexp"},
	}

	pesticide.RunTestCases(t, regex.NewRegistry(), tc)
}

func TestRegexMatch(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestRegexMatchValid", Input: `{{ "Hello" | regexMatch "^[a-zA-Z]+$" }}`, ExpectedOutput: "true"},
		{Name: "TestRegexMatchInvalidAlphaNumeric", Input: `{{ "Hello123" | regexMatch "^[a-zA-Z]+$" }}`, ExpectedOutput: "false"},
		{Name: "TestRegexMatchInvalidNumeric", Input: `{{ "123" | regexMatch "^[a-zA-Z]+$" }}`, ExpectedOutput: "false"},
		{Name: "TestRegexMatchInvalidPattern", Input: `{{ "Hello" | regexMatch "^[a-zA+$" }}`, ExpectedErr: "error parsing regexp"},
	}

	pesticide.RunTestCases(t, regex.NewRegistry(), tc)
}

func TestRegexSplit(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestRegexSplitWithoutLimit", Input: `{{ "banana" | regexSplit "a" -1 }}`, ExpectedOutput: "[b n n ]"},
		{Name: "TestRegexSplitZeroLimit", Input: `{{ "banana" | regexSplit "a" 0 }}`, ExpectedOutput: "[]"},
		{Name: "TestRegexSplitOneLimit", Input: `{{ "banana" | regexSplit "a" 1 }}`, ExpectedOutput: "[banana]"},
		{Name: "TestRegexSplitTwoLimit", Input: `{{ "banana" | regexSplit "a" 2 }}`, ExpectedOutput: "[b nana]"},
		{Name: "TestRegexSplitRepetitionLimit", Input: `{{ "banana" | regexSplit "a+" 1 }}`, ExpectedOutput: "[banana]"},
		{Name: "TestRegexSplitInvalidPattern", Input: `{{ "banana" | regexSplit "+" 0 }}`, ExpectedErr: "error parsing regexp"},
	}

	pesticide.RunTestCases(t, regex.NewRegistry(), tc)
}

func TestRegexReplaceAll(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestRegexReplaceAllValid", Input: `{{ "-ab-axxb-" | regexReplaceAll "a(x*)b" "T" }}`, ExpectedOutput: "-T-T-"},
		{Name: "TestRegexReplaceAllWithDollarSign", Input: `{{ "-ab-axxb-" | regexReplaceAll "a(x*)b" "$1" }}`, ExpectedOutput: "--xx-"},
		{Name: "TestRegexReplaceAllWithDollarSignAndLetter", Input: `{{ "-ab-axxb-" | regexReplaceAll "a(x*)b" "$1W" }}`, ExpectedOutput: "---"},
		{Name: "TestRegexReplaceAllWithDollarSignAndCurlyBraces", Input: `{{ "-ab-axxb-" | regexReplaceAll "a(x*)b" "${1}W" }}`, ExpectedOutput: "-W-xxW-"},
		{Name: "TestRegexReplaceAllWithInvalidPattern", Input: `{{ "-ab-axxb-" | regexReplaceAll "a(x*}" "T" }}`, ExpectedErr: "error parsing regexp"},
	}

	pesticide.RunTestCases(t, regex.NewRegistry(), tc)
}

func TestRegexReplaceAllLiteral(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestRegexReplaceAllLiteralValid", Input: `{{ "-ab-axxb-" | regexReplaceAllLiteral "a(x*)b" "T" }}`, ExpectedOutput: "-T-T-"},
		{Name: "TestRegexReplaceAllLiteralWithDollarSign", Input: `{{ "-ab-axxb-" | regexReplaceAllLiteral "a(x*)b" "$1" }}`, ExpectedOutput: "-$1-$1-"},
		{Name: "TestRegexReplaceAllLiteralWithDollarSignAndLetter", Input: `{{ "-ab-axxb-" | regexReplaceAllLiteral "a(x*)b" "$1W" }}`, ExpectedOutput: "-$1W-$1W-"},
		{Name: "TestRegexReplaceAllLiteralWithDollarSignAndCurlyBraces", Input: `{{ "-ab-axxb-" | regexReplaceAllLiteral "a(x*)b" "${1}W" }}`, ExpectedOutput: "-${1}W-${1}W-"},
		{Name: "TestRegexReplaceAllLiteralWithInvalidPattern", Input: `{{ "-ab-axxb-" | regexReplaceAllLiteral "a(x*}" "T" }}`, ExpectedErr: "error parsing regexp"},
	}

	pesticide.RunTestCases(t, regex.NewRegistry(), tc)
}

func TestRegexQuoteMeta(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestRegexQuoteMetaALongLine", Input: `{{ regexQuoteMeta "Escaping $100? That's a lot." }}`, ExpectedOutput: "Escaping \\$100\\? That's a lot\\."},
		{Name: "TestRegexQuoteMetaASemVer", Input: `{{ regexQuoteMeta "1.2.3" }}`, ExpectedOutput: "1\\.2\\.3"},
		{Name: "TestRegexQuoteMetaNothing", Input: `{{ regexQuoteMeta "golang" }}`, ExpectedOutput: "golang"},
	}

	pesticide.RunTestCases(t, regex.NewRegistry(), tc)
}

func TestRegexFindGroups(t *testing.T) {
	tc := []pesticide.TestCase{
		{
			Name:           "TestRegexFindGroupsValid",
			Input:          `{{ .V | regexFindGroups "([A-Za-z]+)@(example\\.com)" }}`,
			Data:           map[string]any{"V": "Contact us at support@example.com"},
			ExpectedOutput: "[support@example.com support example.com]",
		},
		{
			Name:           "TestRegexFindGroupsNoMatch",
			Input:          `{{ .V | regexFindGroups "([A-Za-z]+)@(example\\.org)" }}`,
			Data:           map[string]any{"V": "Contact us at support@example.com"},
			ExpectedOutput: "[]",
		},
		{
			Name:        "TestRegexFindGroupsInvalidPattern",
			Input:       `{{ .V | regexFindGroups "([A-Za-z]+)@(example\\.com" }}`,
			Data:        map[string]any{"V": "Contact us at support@example.com"},
			ExpectedErr: "error parsing regexp",
		},
	}

	pesticide.RunTestCases(t, regex.NewRegistry(), tc)
}

func TestRegexFindAllGroups(t *testing.T) {
	tc := []pesticide.TestCase{
		{
			Name:           "TestRegexFindAllGroupsValid",
			Input:          `{{ .V | regexFindAllGroups "(\\w+)=(\\w+)" -1 }}`,
			Data:           map[string]any{"V": "var1=value1&var2=value2&var3=value3"},
			ExpectedOutput: "[[var1=value1 var1 value1] [var2=value2 var2 value2] [var3=value3 var3 value3]]",
		},
		{
			Name:           "TestRegexFindAllGroupsWithLimit",
			Input:          `{{ .V | regexFindAllGroups "(\\w+)=(\\w+)" 2 }}`,
			Data:           map[string]any{"V": "var1=value1&var2=value2&var3=value3"},
			ExpectedOutput: "[[var1=value1 var1 value1] [var2=value2 var2 value2]]",
		},
		{
			Name:           "TestRegexFindAllGroupsNoMatch",
			Input:          `{{ .V | regexFindAllGroups "(\\d+)=(\\d+)" -1 }}`,
			Data:           map[string]any{"V": "var1=value1&var2=value2&var3=value3"},
			ExpectedOutput: "[]",
		},
		{
			Name:        "TestRegexFindAllGroupsInvalidPattern",
			Input:       `{{ .V | regexFindAllGroups "(\\w+)=(\\w+" -1 }}`,
			Data:        map[string]any{"V": "var1=value1&var2=value2&var3=value3"},
			ExpectedErr: "error parsing regexp",
		},
	}

	pesticide.RunTestCases(t, regex.NewRegistry(), tc)
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
			Name:           "TestRegexFindNamedWithUnnamedGroup",
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

	pesticide.RunTestCases(t, regex.NewRegistry(), tc)
}

func TestRegexFindAllNamed(t *testing.T) {
	tc := []pesticide.TestCase{
		{
			Name:           "TestRegexFindAllNamedValid",
			Input:          `{{ .V | regexFindAllNamed "(?P<param>\\w+)=(?P<value>\\w+)" -1 }}`,
			Data:           map[string]any{"V": "var1=value1&var2=value2&var3=value3"},
			ExpectedOutput: "[map[param:var1 value:value1] map[param:var2 value:value2] map[param:var3 value:value3]]",
		},
		{
			Name:           "TestRegexFindAllNamedWithLimit",
			Input:          `{{ .V | regexFindAllNamed "(?P<param>\\w+)=(?P<value>\\w+)" 2 }}`,
			Data:           map[string]any{"V": "var1=value1&var2=value2&var3=value3"},
			ExpectedOutput: "[map[param:var1 value:value1] map[param:var2 value:value2]]",
		},
		{
			Name:           "TestRegexFindAllNamedNoMatch",
			Input:          `{{ .V | regexFindAllNamed "(?P<param>\\d+)=(?P<value>\\d+)" -1 }}`,
			Data:           map[string]any{"V": "var1=value1&var2=value2&var3=value3"},
			ExpectedOutput: "[]",
		},
		{
			Name:        "TestRegexFindAllNamedInvalidPattern",
			Input:       `{{ .V | regexFindAllNamed "(?P<param>\\w+)=(?P<value>\\w+" -1 }}`,
			Data:        map[string]any{"V": "var1=value1&var2=value2&var3=value3"},
			ExpectedErr: "error parsing regexp",
		},
	}

	pesticide.RunTestCases(t, regex.NewRegistry(), tc)
}
