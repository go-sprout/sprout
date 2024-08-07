package regexp_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/regexp"
)

func TestRegexpFind(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestRegexpFind", Input: `{{ regexFind "a(b+)" "aaabbb" }}`, Expected: "abbb"},
		{Name: "TestRegexpFindError", Input: `{{ regexFind "a(b+" "aaabbb" }}`, Expected: ""},
	}

	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestRegexpFindAll(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestRegexpFindAllWithoutLimit", Input: `{{ regexFindAll "a(b+)" "aaabbb" -1 }}`, Expected: "[abbb]"},
		{Name: "TestRegexpFindAllWithLimit", Input: `{{ regexFindAll "a{2}" "aaaabbb" -1 }}`, Expected: "[aa aa]"},
		{Name: "TestRegexpFindAllWithNoMatch", Input: `{{ regexFindAll "a{2}" "none" -1 }}`, Expected: "[]"},
		{Name: "TestRegexpFindAllWithInvalidPattern", Input: `{{ regexFindAll "a(b+" "aaabbb" -1 }}`, Expected: "[]"},
	}

	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestRegexMatch(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestRegexMatchValid", Input: `{{ regexMatch "^[a-zA-Z]+$" "Hello" }}`, Expected: "true"},
		{Name: "TestRegexMatchInvalidAlphaNumeric", Input: `{{ regexMatch "^[a-zA-Z]+$" "Hello123" }}`, Expected: "false"},
		{Name: "TestRegexMatchInvalidNumeric", Input: `{{ regexMatch "^[a-zA-Z]+$" "123" }}`, Expected: "false"},
	}

	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestRegexSplit(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestRegexpFindAllWithoutLimit", Input: `{{ regexSplit "a" "banana" -1 }}`, Expected: "[b n n ]"},
		{Name: "TestRegexpSplitZeroLimit", Input: `{{ regexSplit "a" "banana" 0 }}`, Expected: "[]"},
		{Name: "TestRegexpSplitOneLimit", Input: `{{ regexSplit "a" "banana" 1 }}`, Expected: "[banana]"},
		{Name: "TestRegexpSplitTwoLimit", Input: `{{ regexSplit "a" "banana" 2 }}`, Expected: "[b nana]"},
		{Name: "TestRegexpSplitRepetitionLimit", Input: `{{ regexSplit "a+" "banana" 1 }}`, Expected: "[banana]"},
	}

	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestRegexReplaceAll(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestRegexReplaceAllValid", Input: `{{ regexReplaceAll "a(x*)b" "-ab-axxb-" "T" }}`, Expected: "-T-T-"},
		{Name: "TestRegexReplaceAllWithDollarSign", Input: `{{ regexReplaceAll "a(x*)b" "-ab-axxb-" "$1" }}`, Expected: "--xx-"},
		{Name: "TestRegexReplaceAllWithDollarSignAndLetter", Input: `{{ regexReplaceAll "a(x*)b" "-ab-axxb-" "$1W" }}`, Expected: "---"},
		{Name: "TestRegexReplaceAllWithDollarSignAndCurlyBraces", Input: `{{ regexReplaceAll "a(x*)b" "-ab-axxb-" "${1}W" }}`, Expected: "-W-xxW-"},
	}

	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestRegexReplaceAllLiteral(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestRegexReplaceAllLiteralValid", Input: `{{ regexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "T" }}`, Expected: "-T-T-"},
		{Name: "TestRegexReplaceAllLiteralWithDollarSign", Input: `{{ regexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "$1" }}`, Expected: "-$1-$1-"},
		{Name: "TestRegexReplaceAllLiteralWithDollarSignAndLetter", Input: `{{ regexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "$1W" }}`, Expected: "-$1W-$1W-"},
		{Name: "TestRegexReplaceAllLiteralWithDollarSignAndCurlyBraces", Input: `{{ regexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "${1}W" }}`, Expected: "-${1}W-${1}W-"},
	}

	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestRegexQuoteMeta(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestRegexQuoteMetaALongLine", Input: `{{ regexQuoteMeta "Escaping $100? That's a lot." }}`, Expected: "Escaping \\$100\\? That's a lot\\."},
		{Name: "TestRegexQuoteMetaASemVer", Input: `{{ regexQuoteMeta "1.2.3" }}`, Expected: "1\\.2\\.3"},
		{Name: "TestRegexQuoteMetaNothing", Input: `{{ regexQuoteMeta "golang" }}`, Expected: "golang"},
	}

	pesticide.RunTestCases(t, regexp.NewRegistry(), tc)
}

func TestMustRegexFind(t *testing.T) {
	var tc = []pesticide.MustTestCase{
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexFindValid",
				Input:    `{{ mustRegexFind "a(b+)" "aaabbb" }}`,
				Expected: "abbb",
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexFindInvalid",
				Input:    `{{ mustRegexFind "a(b+" "aaabbb" }}`,
				Expected: "",
			},
			ExpectedErr: "error parsing regexp: missing closing ): `a(b+`",
		},
	}

	pesticide.RunMustTestCases(t, regexp.NewRegistry(), tc)
}

func TestMustRegexFindAll(t *testing.T) {
	var tc = []pesticide.MustTestCase{
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexFindAllValid",
				Input:    `{{ mustRegexFindAll "a(b+)" "aaabbb" -1 }}`,
				Expected: "[abbb]",
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexFindAllWithLimit",
				Input:    `{{ mustRegexFindAll "a{2}" "aaaabbb" -1 }}`,
				Expected: "[aa aa]",
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexFindAllWithNoMatch",
				Input:    `{{ mustRegexFindAll "a{2}" "none" -1 }}`,
				Expected: "[]",
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexFindAllWithInvalidPattern",
				Input:    `{{ mustRegexFindAll "a(b+" "aaabbb" -1 }}`,
				Expected: "",
			},
			ExpectedErr: "error parsing regexp: missing closing ): `a(b+`",
		},
	}

	pesticide.RunMustTestCases(t, regexp.NewRegistry(), tc)
}

func TestMustRegexMatch(t *testing.T) {
	var tc = []pesticide.MustTestCase{
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexMatchValid",
				Input:    `{{ mustRegexMatch "^[a-zA-Z]+$" "Hello" }}`,
				Expected: "true",
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexMatchInvalidAlphaNumeric",
				Input:    `{{ mustRegexMatch "^[a-zA-Z]+$" "Hello123" }}`,
				Expected: "false",
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexMatchInvalidNumeric",
				Input:    `{{ mustRegexMatch "^[a-zA-Z]+$" "123" }}`,
				Expected: "false",
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexMatchInvalidPattern",
				Input:    `{{ mustRegexMatch "^[a-zA+$" "Hello" }}`,
				Expected: "",
			},
			ExpectedErr: "error parsing regexp: missing closing ]: `[a-zA+$`",
		},
	}

	pesticide.RunMustTestCases(t, regexp.NewRegistry(), tc)
}

func TestMustRegexSplit(t *testing.T) {
	var tc = []pesticide.MustTestCase{
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexSplitWithoutLimit",
				Input:    `{{ mustRegexSplit "a" "banana" -1 }}`,
				Expected: "[b n n ]",
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexSplitZeroLimit",
				Input:    `{{ mustRegexSplit "a" "banana" 0 }}`,
				Expected: "[]",
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexSplitOneLimit",
				Input:    `{{ mustRegexSplit "a" "banana" 1 }}`,
				Expected: "[banana]",
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexSplitTwoLimit",
				Input:    `{{ mustRegexSplit "a" "banana" 2 }}`,
				Expected: "[b nana]",
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexSplitRepetitionLimit",
				Input:    `{{ mustRegexSplit "a+" "banana" 1 }}`,
				Expected: "[banana]",
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexSplitInvalidPattern",
				Input:    `{{ mustRegexSplit "+" "banana" 0 }}`,
				Expected: "",
			},
			ExpectedErr: "error parsing regexp: missing argument to repetition operator: `+`",
		},
	}

	pesticide.RunMustTestCases(t, regexp.NewRegistry(), tc)
}

func TestMustRegexReplaceAll(t *testing.T) {
	var tc = []pesticide.MustTestCase{
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexReplaceAllValid",
				Input:    `{{ mustRegexReplaceAll "a(x*)b" "-ab-axxb-" "T" }}`,
				Expected: "-T-T-",
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexReplaceAllWithDollarSign",
				Input:    `{{ mustRegexReplaceAll "a(x*)b" "-ab-axxb-" "$1" }}`,
				Expected: "--xx-",
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexReplaceAllWithDollarSignAndLetter",
				Input:    `{{ mustRegexReplaceAll "a(x*)b" "-ab-axxb-" "$1W" }}`,
				Expected: "---",
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexReplaceAllWithDollarSignAndCurlyBraces",
				Input:    `{{ mustRegexReplaceAll "a(x*)b" "-ab-axxb-" "${1}W" }}`,
				Expected: "-W-xxW-",
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexReplaceAllWithInvalidPattern",
				Input:    `{{ mustRegexReplaceAll "a(x*}" "-ab-axxb-" "T" }}`,
				Expected: "",
			},
			ExpectedErr: "error parsing regexp: missing closing ): `a(x*}`",
		},
	}

	pesticide.RunMustTestCases(t, regexp.NewRegistry(), tc)
}

func TestMustRegexReplaceAllLiteral(t *testing.T) {
	var tc = []pesticide.MustTestCase{
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexReplaceAllLiteralValid",
				Input:    `{{ mustRegexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "T" }}`,
				Expected: "-T-T-",
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexReplaceAllLiteralWithDollarSign",
				Input:    `{{ mustRegexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "$1" }}`,
				Expected: "-$1-$1-",
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexReplaceAllLiteralWithDollarSignAndLetter",
				Input:    `{{ mustRegexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "$1W" }}`,
				Expected: "-$1W-$1W-",
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexReplaceAllLiteralWithDollarSignAndCurlyBraces",
				Input:    `{{ mustRegexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "${1}W" }}`,
				Expected: "-${1}W-${1}W-",
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestMustRegexReplaceAllLiteralWithInvalidPattern",
				Input:    `{{ mustRegexReplaceAllLiteral "a(x*}" "-ab-axxb-" "T" }}`,
				Expected: "",
			},
			ExpectedErr: "error parsing regexp: missing closing ): `a(x*}`",
		},
	}

	pesticide.RunMustTestCases(t, regexp.NewRegistry(), tc)
}
