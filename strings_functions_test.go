package sprout

import (
	mathrand "math/rand"
	"testing"
)

func TestNoSpace(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | nospace }}`, "", nil},
		{"TestSpaceOnly", `{{ " " | nospace }}`, "", nil},
		{"TestLeadingSpace", `{{ " foo" | nospace }}`, "foo", nil},
		{"TestTrailingSpace", `{{ "foo " | nospace }}`, "foo", nil},
		{"TestLeadingAndTrailingSpace", `{{ " foo " | nospace }}`, "foo", nil},
		{"TestMultipleSpaces", `{{ " foo bar " | nospace }}`, "foobar", nil},
	}

	runTestCases(t, tests)
}

func TestTrim(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | trim }}`, "", nil},
		{"TestSpaceOnly", `{{ " " | trim }}`, "", nil},
		{"TestLeadingSpace", `{{ " foo" | trim }}`, "foo", nil},
		{"TestTrailingSpace", `{{ "foo " | trim }}`, "foo", nil},
		{"TestLeadingAndTrailingSpace", `{{ " foo " | trim }}`, "foo", nil},
		{"TestMultipleSpaces", `{{ " foo bar " | trim }}`, "foo bar", nil},
	}

	runTestCases(t, tests)
}

func TestTrimAll(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | trimAll "-" }}`, "", nil},
		{"TestAllDashes", `{{ "---------" | trimAll "-" }}`, "", nil},
		{"TestNoDashes", `{{ "foo" | trimAll "-" }}`, "foo", nil},
		{"TestSomeDashes", `{{ "-f--o-o-" | trimAll "-" }}`, "f--o-o", nil},
		{"TestOtherDashes", `{{ "-f--o-o-" | trimAll "-o" }}`, "f", nil},
	}

	runTestCases(t, tests)
}

func TestTrimPrefix(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | trimPrefix "-" }}`, "", nil},
		{"TestDoubleDash", `{{ "--" | trimPrefix "-" }}`, "-", nil},
		{"TestNoPrefix", `{{ "foo" | trimPrefix "-" }}`, "foo", nil},
		{"TestSinglePrefix", `{{ "-foo-" | trimPrefix "-" }}`, "foo-", nil},
		{"TestMultiplePrefix", `{{ "-foo-" | trimPrefix "-f" }}`, "oo-", nil},
	}

	runTestCases(t, tests)
}

func TestTrimSuffix(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | trimSuffix "-" }}`, "", nil},
		{"TestDoubleDash", `{{ "--" | trimSuffix "-" }}`, "-", nil},
		{"TestNoSuffix", `{{ "foo" | trimSuffix "-" }}`, "foo", nil},
		{"TestSingleSuffix", `{{ "-foo-" | trimSuffix "-" }}`, "-foo", nil},
		{"TestMultipleSuffix", `{{ "-foo-" | trimSuffix "o-" }}`, "-fo", nil},
	}

	runTestCases(t, tests)
}

func TestContains(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | contains "-" }}`, "false", nil},
		{"TestContains", `{{ "foo" | contains "o" }}`, "true", nil},
		{"TestNotContains", `{{ "foo" | contains "x" }}`, "false", nil},
	}

	runTestCases(t, tests)
}

func TestHasPrefix(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | hasPrefix "-" }}`, "false", nil},
		{"TestHasPrefix", `{{ "foo" | hasPrefix "f" }}`, "true", nil},
		{"TestNotHasPrefix", `{{ "foo" | hasPrefix "o" }}`, "false", nil},
	}

	runTestCases(t, tests)
}

func TestHasSuffix(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | hasSuffix "-" }}`, "false", nil},
		{"TestHasSuffix", `{{ "foo" | hasSuffix "o" }}`, "true", nil},
		{"TestNotHasSuffix", `{{ "foo" | hasSuffix "f" }}`, "false", nil},
	}

	runTestCases(t, tests)
}

func TestToLower(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | toLower }}`, "", nil},
		{"TestLower", `{{ "foo" | toLower }}`, "foo", nil},
		{"TestUpper", `{{ "FOO" | toLower }}`, "foo", nil},
	}

	runTestCases(t, tests)
}

func TestToUpper(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | toUpper }}`, "", nil},
		{"TestLower", `{{ "foo" | toUpper }}`, "FOO", nil},
		{"TestUpper", `{{ "FOO" | toUpper }}`, "FOO", nil},
	}

	runTestCases(t, tests)
}

func TestReplace(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | replace "-" "x" }}`, "", nil},
		{"TestReplace", `{{ "foo" | replace "o" "x" }}`, "fxx", nil},
		{"TestNotReplace", `{{ "foo" | replace "x" "y" }}`, "foo", nil},
		{"TestMultipleReplace", `{{ "foo" | replace "o" "x" | replace "f" "y" }}`, "yxx", nil},
	}

	runTestCases(t, tests)
}

func TestRepeat(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | repeat 3 }}`, "", nil},
		{"TestRepeat", `{{ "foo" | repeat 3 }}`, "foofoofoo", nil},
		{"TestRepeatZero", `{{ "foo" | repeat 0 }}`, "", nil},
	}

	runTestCases(t, tests)
}

func TestJoin(t *testing.T) {
	var tests = testCases{
		{"TestNil", `{{ .nil | join "-" }}`, "", map[string]any{"nil": nil}},
		{"TestIntSlice", `{{ .test | join "-" }}`, "1-2-3", map[string]any{"test": []int{1, 2, 3}}},
		{"TestStringSlice", `{{ .test | join "-" }}`, "a-b-c", map[string]any{"test": []string{"a", "b", "c"}}},
		{"TestString", `{{ .test | join "-" }}`, "abc", map[string]any{"test": "abc"}},
		{"TestMixedSlice", `{{ .test | join "-" }}`, "a-1-true", map[string]any{"test": []any{"a", nil, 1, true}}},
	}

	runTestCases(t, tests)
}

func TestTrunc(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | trunc 3 }}`, "", nil},
		{"TestTruncate", `{{ "foooooo" | trunc 3 }}`, "foo", nil},
		{"TestNegativeTruncate", `{{ "foobar" | trunc -3 }}`, "bar", nil},
		{"TestNegativeLargeTruncate", `{{ "foobar" | trunc -999 }}`, "foobar", nil},
		{"TestZeroTruncate", `{{ "foobar" | trunc 0 }}`, "", nil},
	}

	runTestCases(t, tests)
}

func TestShuffle(t *testing.T) {
	originalRandSource := randSource
	defer func() {
		randSource = originalRandSource
	}()

	randSource = mathrand.NewSource(0)

	var tests = testCases{
		{"TestEmpty", `{{ "" | shuffle }}`, "", nil},
		{"TestShuffle", `{{ "foobar" | shuffle }}`, "abfoor", nil},
	}

	runTestCases(t, tests)
}

func TestEllipsis(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | ellipsis 3 }}`, "", nil},
		{"TestShort", `{{ "foo" | ellipsis 5 }}`, "foo", nil},
		{"TestTruncate", `{{ "foooooo" | ellipsis 6 }}`, "foo...", nil},
		{"TestZeroTruncate", `{{ "foobar" | ellipsis 0 }}`, "foobar", nil},
	}

	runTestCases(t, tests)
}

func TestEllipsisBoth(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | ellipsisBoth 3 5 }}`, "", nil},
		{"TestShort", `{{ "foo" | ellipsisBoth 5 4 }}`, "foo", nil},
		{"TestTruncate", `{{ "foooboooooo" | ellipsisBoth 4 9 }}`, "...boo...", nil},
		{"TestZeroTruncate", `{{ "foobar" | ellipsisBoth 0 0 }}`, "foobar", nil},
	}

	runTestCases(t, tests)
}

func TestInitials(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | initials }}`, "", nil},
		{"TestSingle", `{{ "f" | initials }}`, "f", nil},
		{"TestTwo", `{{ "foo" | initials }}`, "f", nil},
		{"TestThree", `{{ "foo bar" | initials }}`, "fb", nil},
		{"TestMultipleSpaces", `{{ "foo  bar" | initials }}`, "fb", nil},
		{"TestWithUppercased", `{{ " Foo bar" | initials }}`, "Fb", nil},
	}

	runTestCases(t, tests)
}

func TestPlural(t *testing.T) {
	var tests = testCases{
		{"TestZero", `{{ 0 | plural "single" "many" }}`, "many", nil},
		{"TestSingle", `{{ 1 | plural "single" "many" }}`, "single", nil},
		{"TestMultiple", `{{ 2 | plural "single" "many" }}`, "many", nil},
		{"TestNegative", `{{ -1 | plural "single" "many" }}`, "many", nil},
	}

	runTestCases(t, tests)
}

func TestWrap(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | wrap 10 }}`, "", nil},
		{"TestNegativeWrap", `{{ wrap -1 "With a negative wrap." }}`, "With\na\nnegative\nwrap.", nil},
		{"TestWrap", `{{ "This is a long string that needs to be wrapped." | wrap 10 }}`, "This is a\nlong\nstring\nthat needs\nto be\nwrapped.", nil},
	}

	runTestCases(t, tests)
}

func TestWrapWith(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | wrapWith 10 "\t" }}`, "", nil},
		{"TestWrap", `{{ "This is a long string that needs to be wrapped." | wrapWith 10 "\t" }}`, "This is a\tlong\tstring\tthat needs\tto be\twrapped.", nil},
		{"TestWrapWithLongWord", `{{ "This is a long string that needs to be wrapped with a looooooooooooooooooooooooooooooooooooong word." | wrapWith 10 "\t" }}`, "This is a\tlong\tstring\tthat needs\tto be\twrapped\twith a\tlooooooooo\toooooooooo\toooooooooo\toooooooong\tword.", nil},
	}

	runTestCases(t, tests)
}

func TestQuote(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | quote }}`, `""`, nil},
		{"TestNil", `{{ quote .nil }}`, ``, map[string]any{"nil": nil}},
		{"TestQuote", `{{ "foo" | quote }}`, `"foo"`, nil},
		{"TestSpace", `{{ "foo bar" | quote }}`, `"foo bar"`, nil},
		{"TestQuote", `{{ "foo \"bar\"" | quote }}`, `"foo \"bar\""`, nil},
		{"TestNewline", `{{ "foo\nbar" | quote }}`, `"foo\nbar"`, nil},
		{"TestBackslash", `{{ "foo\\bar" | quote }}`, `"foo\\bar"`, nil},
		{"TestBackslashAndQuote", `{{ "foo\\\"bar" | quote }}`, `"foo\\\"bar"`, nil},
		{"TestUnicode", `{{ quote "foo" "👍" }}`, `"foo" "👍"`, nil},
	}

	runTestCases(t, tests)
}

func TestSquote(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | squote }}`, "''", nil},
		{"TestNil", `{{ squote .nil }}`, "", map[string]any{"nil": nil}},
		{"TestQuote", `{{ "foo" | squote }}`, "'foo'", nil},
		{"TestSpace", `{{ "foo bar" | squote }}`, "'foo bar'", nil},
		{"TestQuote", `{{ "foo 'bar'" | squote }}`, "'foo 'bar''", nil},
		{"TestNewline", `{{ "foo\nbar" | squote }}`, "'foo\nbar'", nil},
		{"TestBackslash", `{{ "foo\\bar" | squote }}`, "'foo\\bar'", nil},
		{"TestBackslashAndQuote", `{{ "foo\\'bar" | squote }}`, "'foo\\'bar'", nil},
		{"TestUnicode", `{{ squote "foo" "👍" }}`, "'foo' '👍'", nil},
	}

	runTestCases(t, tests)
}

func TestToCamelCase(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | toCamelCase }}`, "", nil},
		{"TestCamelCase", `{{ "foo bar" | toCamelCase }}`, "fooBar", nil},
		{"TestCamelCaseWithUpperCase", `{{ "FoO  bar" | toCamelCase }}`, "fooBar", nil},
		{"TestCamelCaseWithSpace", `{{ "foo  bar" | toCamelCase }}`, "fooBar", nil},
		{"TestCamelCaseWithUnderscore", `{{ "foo_bar" | toCamelCase }}`, "fooBar", nil},
		{"TestCamelCaseWithHyphen", `{{ "foo-bar" | toCamelCase }}`, "fooBar", nil},
		{"TestCamelCaseWithMixed", `{{ "foo-bar_baz" | toCamelCase }}`, "fooBarBaz", nil},
		{"", `{{ toCamelCase "___complex__case_" }}`, "complexCase", nil},
		{"", `{{ toCamelCase "_camel_case" }}`, "camelCase", nil},
		{"", `{{ toCamelCase "http_server" }}`, "httpServer", nil},
		{"", `{{ toCamelCase "no_https" }}`, "noHttps", nil},
		{"", `{{ toCamelCase "all" }}`, "all", nil},
	}

	runTestCases(t, tests)
}

func TestToKebakCase(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | toKebabCase }}`, "", nil},
		{"TestKebabCase", `{{ "foo bar" | toKebabCase }}`, "foo-bar", nil},
		{"TestKebabCaseWithSpace", `{{ "foo  bar" | toKebabCase }}`, "foo-bar", nil},
		{"TestKebabCaseWithUnderscore", `{{ "foo_bar" | toKebabCase }}`, "foo-bar", nil},
		{"TestKebabCaseWithHyphen", `{{ "foo-bar" | toKebabCase }}`, "foo-bar", nil},
		{"TestKebabCaseWithMixed", `{{ "foo-bar_baz" | toKebabCase }}`, "foo-bar-baz", nil},
		{"", `{{ toKebabCase "HTTPServer" }}`, "http-server", nil},
		{"", `{{ toKebabCase "FirstName" }}`, "first-name", nil},
		{"", `{{ toKebabCase "NoHTTPS" }}`, "no-https", nil},
		{"", `{{ toKebabCase "GO_PATH" }}`, "go-path", nil},
		{"", `{{ toKebabCase "GO PATH" }}`, "go-path", nil},
		{"", `{{ toKebabCase "GO-PATH" }}`, "go-path", nil},
	}

	runTestCases(t, tests)
}

func TestToPascalCase(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | toPascalCase }}`, "", nil},
		{"TestPascalCase", `{{ "foo bar" | toPascalCase }}`, "FooBar", nil},
		{"TestPascalCaseWithSpace", `{{ "foo  bar" | toPascalCase }}`, "FooBar", nil},
		{"TestPascalCaseWithUnderscore", `{{ "foo_bar" | toPascalCase }}`, "FooBar", nil},
		{"TestPascalCaseWithHyphen", `{{ "foo-bar" | toPascalCase }}`, "FooBar", nil},
		{"TestPascalCaseWithMixed", `{{ "foo-bar_baz" | toPascalCase }}`, "FooBarBaz", nil},
	}

	runTestCases(t, tests)
}

func TestToDotCase(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | toDotCase }}`, "", nil},
		{"TestDotCase", `{{ "foo bar" | toDotCase }}`, "foo.bar", nil},
		{"TestDotCaseWithSpace", `{{ "foo  bar" | toDotCase }}`, "foo.bar", nil},
		{"TestDotCaseWithUnderscore", `{{ "foo_bar" | toDotCase }}`, "foo.bar", nil},
		{"TestDotCaseWithHyphen", `{{ "foo-bar" | toDotCase }}`, "foo.bar", nil},
		{"TestDotCaseWithMixed", `{{ "foo-bar_baz" | toDotCase }}`, "foo.bar.baz", nil},
	}

	runTestCases(t, tests)
}

func TestToPathCase(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | toPathCase }}`, "", nil},
		{"TestPathCase", `{{ "foo bar" | toPathCase }}`, "foo/bar", nil},
		{"TestPathCaseWithSpace", `{{ "foo  bar" | toPathCase }}`, "foo/bar", nil},
		{"TestPathCaseWithUnderscore", `{{ "foo_bar" | toPathCase }}`, "foo/bar", nil},
		{"TestPathCaseWithHyphen", `{{ "foo-bar" | toPathCase }}`, "foo/bar", nil},
		{"TestPathCaseWithMixed", `{{ "foo-bar_baz" | toPathCase }}`, "foo/bar/baz", nil},
	}

	runTestCases(t, tests)
}

func TestToConstantCase(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | toConstantCase }}`, "", nil},
		{"TestConstantCase", `{{ "foo bar" | toConstantCase }}`, "FOO_BAR", nil},
		{"TestConstantCaseWithSpace", `{{ "foo  bar" | toConstantCase }}`, "FOO_BAR", nil},
		{"TestConstantCaseWithUnderscore", `{{ "foo_bar" | toConstantCase }}`, "FOO_BAR", nil},
		{"TestConstantCaseWithHyphen", `{{ "foo-bar" | toConstantCase }}`, "FOO_BAR", nil},
		{"TestConstantCaseWithMixed", `{{ "foo-bar_baz" | toConstantCase }}`, "FOO_BAR_BAZ", nil},
	}

	runTestCases(t, tests)
}

func TestToSnakeCase(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | toSnakeCase }}`, "", nil},
		{"TestSnakeCase", `{{ "foo bar" | toSnakeCase }}`, "foo_bar", nil},
		{"TestSnakeCaseWithSpace", `{{ "foo  bar" | toSnakeCase }}`, "foo_bar", nil},
		{"TestSnakeCaseWithUnderscore", `{{ "foo_bar" | toSnakeCase }}`, "foo_bar", nil},
		{"TestSnakeCaseWithHyphen", `{{ "foo-bar" | toSnakeCase }}`, "foo_bar", nil},
		{"TestSnakeCaseWithMixed", `{{ "foo-bar_baz" | toSnakeCase }}`, "foo_bar_baz", nil},
		{"", `{{ toSnakeCase "http2xx" }}`, "http_2xx", nil},
		{"", `{{ toSnakeCase "HTTP20xOK" }}`, "http_20x_ok", nil},
		{"", `{{ toSnakeCase "Duration2m3s" }}`, "duration_2m_3s", nil},
		{"", `{{ toSnakeCase "Bld4Floor3rd" }}`, "bld_4floor_3rd", nil},
		{"", `{{ toSnakeCase "FirstName" }}`, "first_name", nil},
		{"", `{{ toSnakeCase "HTTPServer" }}`, "http_server", nil},
		{"", `{{ toSnakeCase "NoHTTPS" }}`, "no_https", nil},
		{"", `{{ toSnakeCase "GO_PATH" }}`, "go_path", nil},
		{"", `{{ toSnakeCase "GO PATH" }}`, "go_path", nil},
		{"", `{{ toSnakeCase "GO-PATH" }}`, "go_path", nil},
	}

	runTestCases(t, tests)
}

func TestToTitleCase(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | toTitleCase }}`, "", nil},
		{"TestTitleCase", `{{ "foo bar" | toTitleCase }}`, "Foo Bar", nil},
		{"TestTitleCaseWithSpace", `{{ "foo  bar" | toTitleCase }}`, "Foo  Bar", nil},
		{"TestTitleCaseWithUnderscore", `{{ "foo_bar" | toTitleCase }}`, "Foo_bar", nil},
		{"TestTitleCaseWithHyphen", `{{ "foo-bar" | toTitleCase }}`, "Foo-Bar", nil},
		{"TestTitleCaseWithMixed", `{{ "foo-bar_baz" | toTitleCase }}`, "Foo-Bar_baz", nil},
	}

	runTestCases(t, tests)
}

func TestUntitle(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | untitle }}`, "", nil},
		{"TestUnTitle", `{{ "Foo Bar" | untitle }}`, "foo bar", nil},
		{"TestUnTitleWithSpace", `{{ "Foo  Bar" | untitle }}`, "foo  bar", nil},
		{"TestUnTitleWithUnderscore", `{{ "Foo_bar" | untitle }}`, "foo_bar", nil},
		{"TestUnTitleWithHyphen", `{{ "Foo-Bar" | untitle }}`, "foo-Bar", nil},
		{"TestUnTitleWithMixed", `{{ "Foo-Bar_baz" | untitle }}`, "foo-Bar_baz", nil},
	}

	runTestCases(t, tests)
}

func TestSwapCase(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | swapCase }}`, "", nil},
		{"TestSwapCase", `{{ "Foo Bar" | swapCase }}`, "fOO bAR", nil},
		{"TestSwapCaseWithSpace", `{{ "Foo  Bar" | swapCase }}`, "fOO  bAR", nil},
		{"TestSwapCaseWithUnderscore", `{{ "Foo_bar" | swapCase }}`, "fOO_BAR", nil},
		{"TestSwapCaseWithHyphen", `{{ "Foo-Bar" | swapCase }}`, "fOO-bAR", nil},
		{"TestSwapCaseWithMixed", `{{ "Foo-Bar_baz" | swapCase }}`, "fOO-bAR_BAZ", nil},
	}

	runTestCases(t, tests)
}

func TestSplit(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ $v := ("" | split "-") }}{{$v._0}}`, "", nil},
		{"TestSplit", `{{ $v := ("foo$bar$baz" | split "$") }}{{$v._0}} {{$v._1}} {{$v._2}}`, "foo bar baz", nil},
		{"TestSplitWithEmpty", `{{ $v := ("foo$bar$" | split "$") }}{{$v._0}} {{$v._1}} {{$v._2}}`, "foo bar ", nil},
	}

	runTestCases(t, tests)
}

func TestSplitn(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ $v := ("" | splitn "-" 3) }}{{$v._0}}`, "", nil},
		{"TestSplit", `{{ $v := ("foo$bar$baz" | splitn "$" 2) }}{{$v._0}} {{$v._1}}`, "foo bar$baz", nil},
		{"TestSplitWithEmpty", `{{ $v := ("foo$bar$" | splitn "$" 2) }}{{$v._0}} {{$v._1}}`, "foo bar$", nil},
	}

	runTestCases(t, tests)
}

func TestSubstring(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | substr 0 3 }}`, "", nil},
		{"TestEmptyWithNegativeValue", `{{ "" | substr -1 -4 }}`, "", nil},
		{"TestSubstring", `{{ "foobar" | substr 0 3 }}`, "foo", nil},
		{"TestSubstringNegativeEnd", `{{ "foobar" | substr 0 -3 }}`, "foo", nil},
		{"TestSubstringNegativeStart", `{{ "foobar" | substr -3 6 }}`, "bar", nil},
		{"TestSubstringNegativeStartAndEnd", `{{ "foobar" | substr -3 -1 }}`, "ba", nil},
		{"TestSubstringInvalidRange", `{{ "foobar" | substr -3 -3 }}`, "", nil},
	}

	runTestCases(t, tests)
}

func TestIndent(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | indent 3 }}`, "   ", nil},
		{"TestIndent", `{{ "foo\nbar" | indent 3 }}`, "   foo\n   bar", nil},
		{"TestIndentWithSpace", `{{ "foo\n bar" | indent 3 }}`, "   foo\n    bar", nil},
		{"TestIndentWithTab", `{{ "foo\n\tbar" | indent 3 }}`, "   foo\n   \tbar", nil},
	}

	runTestCases(t, tests)
}

func TestNindent(t *testing.T) {
	var tests = testCases{
		{"TestEmpty", `{{ "" | nindent 3 }}`, "\n   ", nil},
		{"TestIndent", `{{ "foo\nbar" | nindent 3 }}`, "\n   foo\n   bar", nil},
		{"TestIndentWithSpace", `{{ "foo\n bar" | nindent 3 }}`, "\n   foo\n    bar", nil},
		{"TestIndentWithTab", `{{ "foo\n\tbar" | nindent 3 }}`, "\n   foo\n   \tbar", nil},
	}

	runTestCases(t, tests)
}

func TestSeq(t *testing.T) {
	var tests = testCases{
		{"", `{{ seq 0 1 3 }}`, "0 1 2 3", nil},
		{"", `{{ seq 0 3 10 }}`, "0 3 6 9", nil},
		{"", `{{ seq 3 3 2 }}`, "", nil},
		{"", `{{ seq 3 -3 2 }}`, "3", nil},
		{"", `{{ seq }}`, "", nil},
		{"", `{{ seq 0 4 }}`, "0 1 2 3 4", nil},
		{"", `{{ seq 5 }}`, "1 2 3 4 5", nil},
		{"", `{{ seq -5 }}`, "1 0 -1 -2 -3 -4 -5", nil},
		{"", `{{ seq 0 }}`, "1 0", nil},
		{"", `{{ seq 0 1 2 3 }}`, "", nil},
		{"", `{{ seq 0 -4 }}`, "0 -1 -2 -3 -4", nil},
	}

	runTestCases(t, tests)
}
