package sprout

import "testing"

func TestRegexpFind(t *testing.T) {
	var tests = testCases{
		{"TestRegexpFind", `{{ regexFind "a(b+)" "aaabbb" }}`, "abbb", nil},
		{"TestRegexpFindError", `{{ regexFind "a(b+" "aaabbb" }}`, "", nil},
	}

	runTestCases(t, tests)
}

func TestRegexpFindAll(t *testing.T) {
	var tests = testCases{
		{"TestRegexpFindAllWithoutLimit", `{{ regexFindAll "a(b+)" "aaabbb" -1 }}`, "[abbb]", nil},
		{"TestRegexpFindAllWithLimit", `{{ regexFindAll "a{2}" "aaaabbb" -1 }}`, "[aa aa]", nil},
		{"TestRegexpFindAllWithNoMatch", `{{ regexFindAll "a{2}" "none" -1 }}`, "[]", nil},
		{"TestRegexpFindAllWithInvalidPattern", `{{ regexFindAll "a(b+" "aaabbb" -1 }}`, "[]", nil},
	}

	runTestCases(t, tests)
}

func TestRegexMatch(t *testing.T) {
	var tests = testCases{
		{"TestRegexMatchValid", `{{ regexMatch "^[a-zA-Z]+$" "Hello" }}`, "true", nil},
		{"TestRegexMatchInvalidAlphaNumeric", `{{ regexMatch "^[a-zA-Z]+$" "Hello123" }}`, "false", nil},
		{"TestRegexMatchInvalidNumeric", `{{ regexMatch "^[a-zA-Z]+$" "123" }}`, "false", nil},
	}

	runTestCases(t, tests)
}

func TestRegexSplit(t *testing.T) {
	var tests = testCases{
		{"TestRegexpFindAllWithoutLimit", `{{ regexSplit "a" "banana" -1 }}`, "[b n n ]", nil},
		{"TestRegexpSplitZeroLimit", `{{ regexSplit "a" "banana" 0 }}`, "[]", nil},
		{"TestRegexpSplitOneLimit", `{{ regexSplit "a" "banana" 1 }}`, "[banana]", nil},
		{"TestRegexpSplitTwoLimit", `{{ regexSplit "a" "banana" 2 }}`, "[b nana]", nil},
		{"TestRegexpSplitRepetitionLimit", `{{ regexSplit "a+" "banana" 1 }}`, "[banana]", nil},
	}

	runTestCases(t, tests)
}

func TestRegexReplaceAll(t *testing.T) {
	var tests = testCases{
		{"TestRegexReplaceAllValid", `{{ regexReplaceAll "a(x*)b" "-ab-axxb-" "T" }}`, "-T-T-", nil},
		{"TestRegexReplaceAllWithDollarSign", `{{ regexReplaceAll "a(x*)b" "-ab-axxb-" "$1" }}`, "--xx-", nil},
		{"TestRegexReplaceAllWithDollarSignAndLetter", `{{ regexReplaceAll "a(x*)b" "-ab-axxb-" "$1W" }}`, "---", nil},
		{"TestRegexReplaceAllWithDollarSignAndCurlyBraces", `{{ regexReplaceAll "a(x*)b" "-ab-axxb-" "${1}W" }}`, "-W-xxW-", nil},
	}

	runTestCases(t, tests)
}

func TestRegexReplaceAllLiteral(t *testing.T) {
	var tests = testCases{
		{"TestRegexReplaceAllLiteralValid", `{{ regexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "T" }}`, "-T-T-", nil},
		{"TestRegexReplaceAllLiteralWithDollarSign", `{{ regexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "$1" }}`, "-$1-$1-", nil},
		{"TestRegexReplaceAllLiteralWithDollarSignAndLetter", `{{ regexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "$1W" }}`, "-$1W-$1W-", nil},
		{"TestRegexReplaceAllLiteralWithDollarSignAndCurlyBraces", `{{ regexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "${1}W" }}`, "-${1}W-${1}W-", nil},
	}

	runTestCases(t, tests)
}

func TestRegexQuoteMeta(t *testing.T) {
	var tests = testCases{
		{"TestRegexQuoteMetaALongLine", `{{ regexQuoteMeta "Escaping $100? That's a lot." }}`, "Escaping \\$100\\? That's a lot\\.", nil},
		{"TestRegexQuoteMetaASemVer", `{{ regexQuoteMeta "1.2.3" }}`, "1\\.2\\.3", nil},
		{"TestRegexQuoteMetaNothing", `{{ regexQuoteMeta "golang" }}`, "golang", nil},
	}

	runTestCases(t, tests)
}

func TestMustRegexFind(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"TestMustRegexFindValid", `{{ mustRegexFind "a(b+)" "aaabbb" }}`, "abbb", nil}, ""},
		{testCase{"TestMustRegexFindInvalid", `{{ mustRegexFind "a(b+" "aaabbb" }}`, "", nil}, "error parsing regexp: missing closing ): `a(b+`"},
	}

	runMustTestCases(t, tests)
}

func TestMustRegexFindAll(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"TestMustRegexFindAllValid", `{{ mustRegexFindAll "a(b+)" "aaabbb" -1 }}`, "[abbb]", nil}, ""},
		{testCase{"TestMustRegexFindAllWithLimit", `{{ mustRegexFindAll "a{2}" "aaaabbb" -1 }}`, "[aa aa]", nil}, ""},
		{testCase{"TestMustRegexFindAllWithNoMatch", `{{ mustRegexFindAll "a{2}" "none" -1 }}`, "[]", nil}, ""},
		{testCase{"TestMustRegexFindAllWithInvalidPattern", `{{ mustRegexFindAll "a(b+" "aaabbb" -1 }}`, "", nil}, "error parsing regexp: missing closing ): `a(b+`"},
	}

	runMustTestCases(t, tests)
}

func TestMustRegexMatch(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"TestMustRegexMatchValid", `{{ mustRegexMatch "^[a-zA-Z]+$" "Hello" }}`, "true", nil}, ""},
		{testCase{"TestMustRegexMatchInvalidAlphaNumeric", `{{ mustRegexMatch "^[a-zA-Z]+$" "Hello123" }}`, "false", nil}, ""},
		{testCase{"TestMustRegexMatchInvalidNumeric", `{{ mustRegexMatch "^[a-zA-Z]+$" "123" }}`, "false", nil}, ""},
		{testCase{"TestMustRegexMatchInvalidPattern", `{{ mustRegexMatch "^[a-zA+$" "Hello" }}`, "", nil}, "error parsing regexp: missing closing ]: `[a-zA+$`"},
	}

	runMustTestCases(t, tests)
}

func TestMustRegexSplit(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"TestMustRegexSplitWithoutLimit", `{{ mustRegexSplit "a" "banana" -1 }}`, "[b n n ]", nil}, ""},
		{testCase{"TestMustRegexSplitZeroLimit", `{{ mustRegexSplit "a" "banana" 0 }}`, "[]", nil}, ""},
		{testCase{"TestMustRegexSplitOneLimit", `{{ mustRegexSplit "a" "banana" 1 }}`, "[banana]", nil}, ""},
		{testCase{"TestMustRegexSplitTwoLimit", `{{ mustRegexSplit "a" "banana" 2 }}`, "[b nana]", nil}, ""},
		{testCase{"TestMustRegexSplitRepetitionLimit", `{{ mustRegexSplit "a+" "banana" 1 }}`, "[banana]", nil}, ""},
		{testCase{"TestMustRegexSplitInvalidPattern", `{{ mustRegexSplit "+" "banana" 0 }}`, "", nil}, "error parsing regexp: missing argument to repetition operator: `+`"},
	}

	runMustTestCases(t, tests)
}

func TestMustRegexReplaceAll(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"TestMustRegexReplaceAllValid", `{{ mustRegexReplaceAll "a(x*)b" "-ab-axxb-" "T" }}`, "-T-T-", nil}, ""},
		{testCase{"TestMustRegexReplaceAllWithDollarSign", `{{ mustRegexReplaceAll "a(x*)b" "-ab-axxb-" "$1" }}`, "--xx-", nil}, ""},
		{testCase{"TestMustRegexReplaceAllWithDollarSignAndLetter", `{{ mustRegexReplaceAll "a(x*)b" "-ab-axxb-" "$1W" }}`, "---", nil}, ""},
		{testCase{"TestMustRegexReplaceAllWithDollarSignAndCurlyBraces", `{{ mustRegexReplaceAll "a(x*)b" "-ab-axxb-" "${1}W" }}`, "-W-xxW-", nil}, ""},
		{testCase{"TestMustRegexReplaceAllWithInvalidPattern", `{{ mustRegexReplaceAll "a(x*}" "-ab-axxb-" "T" }}`, "", nil}, "error parsing regexp: missing closing ): `a(x*}`"},
	}

	runMustTestCases(t, tests)
}

func TestMustRegexReplaceAllLiteral(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"TestMustRegexReplaceAllLiteralValid", `{{ mustRegexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "T" }}`, "-T-T-", nil}, ""},
		{testCase{"TestMustRegexReplaceAllLiteralWithDollarSign", `{{ mustRegexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "$1" }}`, "-$1-$1-", nil}, ""},
		{testCase{"TestMustRegexReplaceAllLiteralWithDollarSignAndLetter", `{{ mustRegexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "$1W" }}`, "-$1W-$1W-", nil}, ""},
		{testCase{"TestMustRegexReplaceAllLiteralWithDollarSignAndCurlyBraces", `{{ mustRegexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "${1}W" }}`, "-${1}W-${1}W-", nil}, ""},
		{testCase{"TestMustRegexReplaceAllLiteralWithInvalidPattern", `{{ mustRegexReplaceAllLiteral "a(x*}" "-ab-axxb-" "T" }}`, "", nil}, "error parsing regexp: missing closing ): `a(x*}`"},
	}

	runMustTestCases(t, tests)
}
