package strings_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/strings"
)

func TestNoSpace(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | nospace }}`, ExpectedOutput: ""},
		{Name: "TestSpaceOnly", Input: `{{ " " | nospace }}`, ExpectedOutput: ""},
		{Name: "TestLeadingSpace", Input: `{{ " foo" | nospace }}`, ExpectedOutput: "foo"},
		{Name: "TestTrailingSpace", Input: `{{ "foo " | nospace }}`, ExpectedOutput: "foo"},
		{Name: "TestLeadingAndTrailingSpace", Input: `{{ " foo " | nospace }}`, ExpectedOutput: "foo"},
		{Name: "TestMultipleSpaces", Input: `{{ " foo bar " | nospace }}`, ExpectedOutput: "foobar"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestTrim(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | trim }}`, ExpectedOutput: ""},
		{Name: "TestSpaceOnly", Input: `{{ " " | trim }}`, ExpectedOutput: ""},
		{Name: "TestLeadingSpace", Input: `{{ " foo" | trim }}`, ExpectedOutput: "foo"},
		{Name: "TestTrailingSpace", Input: `{{ "foo " | trim }}`, ExpectedOutput: "foo"},
		{Name: "TestLeadingAndTrailingSpace", Input: `{{ " foo " | trim }}`, ExpectedOutput: "foo"},
		{Name: "TestMultipleSpaces", Input: `{{ " foo bar " | trim }}`, ExpectedOutput: "foo bar"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestTrimAll(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | trimAll "-" }}`, ExpectedOutput: ""},
		{Name: "TestAllDashes", Input: `{{ "---------" | trimAll "-" }}`, ExpectedOutput: ""},
		{Name: "TestNoDashes", Input: `{{ "foo" | trimAll "-" }}`, ExpectedOutput: "foo"},
		{Name: "TestSomeDashes", Input: `{{ "-f--o-o-" | trimAll "-" }}`, ExpectedOutput: "f--o-o"},
		{Name: "TestOtherDashes", Input: `{{ "-f--o-o-" | trimAll "-o" }}`, ExpectedOutput: "f"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestTrimPrefix(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | trimPrefix "-" }}`, ExpectedOutput: ""},
		{Name: "TestDoubleDash", Input: `{{ "--" | trimPrefix "-" }}`, ExpectedOutput: "-"},
		{Name: "TestNoPrefix", Input: `{{ "foo" | trimPrefix "-" }}`, ExpectedOutput: "foo"},
		{Name: "TestSinglePrefix", Input: `{{ "-foo-" | trimPrefix "-" }}`, ExpectedOutput: "foo-"},
		{Name: "TestMultiplePrefix", Input: `{{ "-foo-" | trimPrefix "-f" }}`, ExpectedOutput: "oo-"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestTrimSuffix(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | trimSuffix "-" }}`, ExpectedOutput: ""},
		{Name: "TestDoubleDash", Input: `{{ "--" | trimSuffix "-" }}`, ExpectedOutput: "-"},
		{Name: "TestNoSuffix", Input: `{{ "foo" | trimSuffix "-" }}`, ExpectedOutput: "foo"},
		{Name: "TestSingleSuffix", Input: `{{ "-foo-" | trimSuffix "-" }}`, ExpectedOutput: "-foo"},
		{Name: "TestMultipleSuffix", Input: `{{ "-foo-" | trimSuffix "o-" }}`, ExpectedOutput: "-fo"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestContains(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | contains "-" }}`, ExpectedOutput: "false"},
		{Name: "TestContains", Input: `{{ "foo" | contains "o" }}`, ExpectedOutput: "true"},
		{Name: "TestNotContains", Input: `{{ "foo" | contains "x" }}`, ExpectedOutput: "false"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestHasPrefix(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | hasPrefix "-" }}`, ExpectedOutput: "false"},
		{Name: "TestHasPrefix", Input: `{{ "foo" | hasPrefix "f" }}`, ExpectedOutput: "true"},
		{Name: "TestNotHasPrefix", Input: `{{ "foo" | hasPrefix "o" }}`, ExpectedOutput: "false"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestHasSuffix(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | hasSuffix "-" }}`, ExpectedOutput: "false"},
		{Name: "TestHasSuffix", Input: `{{ "foo" | hasSuffix "o" }}`, ExpectedOutput: "true"},
		{Name: "TestNotHasSuffix", Input: `{{ "foo" | hasSuffix "f" }}`, ExpectedOutput: "false"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestToLower(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | toLower }}`, ExpectedOutput: ""},
		{Name: "TestLower", Input: `{{ "foo" | toLower }}`, ExpectedOutput: "foo"},
		{Name: "TestUpper", Input: `{{ "FOO" | toLower }}`, ExpectedOutput: "foo"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestToUpper(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | toUpper }}`, ExpectedOutput: ""},
		{Name: "TestLower", Input: `{{ "foo" | toUpper }}`, ExpectedOutput: "FOO"},
		{Name: "TestUpper", Input: `{{ "FOO" | toUpper }}`, ExpectedOutput: "FOO"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestReplace(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | replace "-" "x" }}`, ExpectedOutput: ""},
		{Name: "TestReplace", Input: `{{ "foo" | replace "o" "x" }}`, ExpectedOutput: "fxx"},
		{Name: "TestNotReplace", Input: `{{ "foo" | replace "x" "y" }}`, ExpectedOutput: "foo"},
		{Name: "TestMultipleReplace", Input: `{{ "foo" | replace "o" "x" | replace "f" "y" }}`, ExpectedOutput: "yxx"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestRepeat(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | repeat 3 }}`, ExpectedOutput: ""},
		{Name: "TestRepeat", Input: `{{ "foo" | repeat 3 }}`, ExpectedOutput: "foofoofoo"},
		{Name: "TestRepeatZero", Input: `{{ "foo" | repeat 0 }}`, ExpectedOutput: ""},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestJoin(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestNil", Input: `{{ .nil | join "-" }}`, ExpectedOutput: "", Data: map[string]any{"nil": nil}},
		{Name: "TestIntSlice", Input: `{{ .test | join "-" }}`, ExpectedOutput: "1-2-3", Data: map[string]any{"test": []int{1, 2, 3}}},
		{Name: "TestStringSlice", Input: `{{ .test | join "-" }}`, ExpectedOutput: "a-b-c", Data: map[string]any{"test": []string{"a", "b", "c"}}},
		{Name: "TestString", Input: `{{ .test | join "-" }}`, ExpectedOutput: "abc", Data: map[string]any{"test": "abc"}},
		{Name: "TestMixedSlice", Input: `{{ .test | join "-" }}`, ExpectedOutput: "a-1-true", Data: map[string]any{"test": []any{"a", nil, 1, true}}},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestTrunc(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | trunc 3 }}`, ExpectedOutput: ""},
		{Name: "TestTruncate", Input: `{{ "foooooo" | trunc 3 }}`, ExpectedOutput: "foo"},
		{Name: "TestNegativeTruncate", Input: `{{ "foobar" | trunc -3 }}`, ExpectedOutput: "bar"},
		{Name: "TestNegativeLargeTruncate", Input: `{{ "foobar" | trunc -999 }}`, ExpectedOutput: "foobar"},
		{Name: "TestZeroTruncate", Input: `{{ "foobar" | trunc 0 }}`, ExpectedOutput: ""},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestShuffle(t *testing.T) {
	tc := []pesticide.RegexpTestCase{
		{Template: `{{ shuffle "" }}`, Length: 0, Regexp: `^$`},
		{Template: `{{ shuffle "hey" }}`, Length: 3, Regexp: `^[hey]{3}$`},
		{Template: `{{ shuffle "foo bar baz" }}`, Length: 11, Regexp: `^[\sfobazr]{11}$`},
	}

	pesticide.RunRegexpTestCases(t, strings.NewRegistry(), tc)
}

func TestEllipsis(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | ellipsis 3 }}`, ExpectedOutput: ""},
		{Name: "TestShort", Input: `{{ "foo" | ellipsis 5 }}`, ExpectedOutput: "foo"},
		{Name: "TestTruncate", Input: `{{ "foooooo" | ellipsis 6 }}`, ExpectedOutput: "foo..."},
		{Name: "TestZeroTruncate", Input: `{{ "foobar" | ellipsis 0 }}`, ExpectedOutput: "foobar"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestEllipsisBoth(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | ellipsisBoth 3 5 }}`, ExpectedOutput: ""},
		{Name: "TestShort", Input: `{{ "foo" | ellipsisBoth 5 4 }}`, ExpectedOutput: "foo"},
		{Name: "TestTruncate", Input: `{{ "foooboooooo" | ellipsisBoth 4 9 }}`, ExpectedOutput: "...boo..."},
		{Name: "TestZeroTruncate", Input: `{{ "foobar" | ellipsisBoth 0 0 }}`, ExpectedOutput: "foobar"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestInitials(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | initials }}`, ExpectedOutput: ""},
		{Name: "TestSingle", Input: `{{ "f" | initials }}`, ExpectedOutput: "f"},
		{Name: "TestTwo", Input: `{{ "foo" | initials }}`, ExpectedOutput: "f"},
		{Name: "TestThree", Input: `{{ "foo bar" | initials }}`, ExpectedOutput: "fb"},
		{Name: "TestMultipleSpaces", Input: `{{ "foo  bar" | initials }}`, ExpectedOutput: "fb"},
		{Name: "TestWithUppercased", Input: `{{ " Foo bar" | initials }}`, ExpectedOutput: "Fb"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestPlural(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestZero", Input: `{{ 0 | plural "single" "many" }}`, ExpectedOutput: "many"},
		{Name: "TestSingle", Input: `{{ 1 | plural "single" "many" }}`, ExpectedOutput: "single"},
		{Name: "TestMultiple", Input: `{{ 2 | plural "single" "many" }}`, ExpectedOutput: "many"},
		{Name: "TestNegative", Input: `{{ -1 | plural "single" "many" }}`, ExpectedOutput: "many"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestWrap(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | wrap 10 }}`, ExpectedOutput: ""},
		{Name: "TestNegativeWrap", Input: `{{ wrap -1 "With a negative wrap." }}`, ExpectedOutput: "With\na\nnegative\nwrap."},
		{Name: "TestWrap", Input: `{{ "This is a long string that needs to be wrapped." | wrap 10 }}`, ExpectedOutput: "This is a\nlong\nstring\nthat needs\nto be\nwrapped."},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestWrapWith(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | wrapWith 10 "\t" }}`, ExpectedOutput: ""},
		{Name: "TestWrap", Input: `{{ "This is a long string that needs to be wrapped." | wrapWith 10 "\t" }}`, ExpectedOutput: "This is a\tlong\tstring\tthat needs\tto be\twrapped."},
		{Name: "TestWrapWithLongWord", Input: `{{ "This is a long string that needs to be wrapped with a looooooooooooooooooooooooooooooooooooong word." | wrapWith 10 "\t" }}`, ExpectedOutput: "This is a\tlong\tstring\tthat needs\tto be\twrapped\twith a\tlooooooooo\toooooooooo\toooooooooo\toooooooong\tword."},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestQuote(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | quote }}`, ExpectedOutput: `""`},
		{Name: "TestNil", Input: `{{ quote .nil }}`, ExpectedOutput: ``, Data: map[string]any{"nil": nil}},
		{Name: "TestQuote", Input: `{{ "foo" | quote }}`, ExpectedOutput: `"foo"`},
		{Name: "TestSpace", Input: `{{ "foo bar" | quote }}`, ExpectedOutput: `"foo bar"`},
		{Name: "TestQuote", Input: `{{ "foo \"bar\"" | quote }}`, ExpectedOutput: `"foo \"bar\""`},
		{Name: "TestNewline", Input: `{{ "foo\nbar" | quote }}`, ExpectedOutput: `"foo\nbar"`},
		{Name: "TestBackslash", Input: `{{ "foo\\bar" | quote }}`, ExpectedOutput: `"foo\\bar"`},
		{Name: "TestBackslashAndQuote", Input: `{{ "foo\\\"bar" | quote }}`, ExpectedOutput: `"foo\\\"bar"`},
		{Name: "TestUnicode", Input: `{{ quote "foo" "👍" }}`, ExpectedOutput: `"foo" "👍"`},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestSquote(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | squote }}`, ExpectedOutput: "''"},
		{Name: "TestNil", Input: `{{ squote .nil }}`, ExpectedOutput: "", Data: map[string]any{"nil": nil}},
		{Name: "TestQuote", Input: `{{ "foo" | squote }}`, ExpectedOutput: "'foo'"},
		{Name: "TestSpace", Input: `{{ "foo bar" | squote }}`, ExpectedOutput: "'foo bar'"},
		{Name: "TestQuote", Input: `{{ "foo 'bar'" | squote }}`, ExpectedOutput: "'foo 'bar''"},
		{Name: "TestNewline", Input: `{{ "foo\nbar" | squote }}`, ExpectedOutput: "'foo\nbar'"},
		{Name: "TestBackslash", Input: `{{ "foo\\bar" | squote }}`, ExpectedOutput: "'foo\\bar'"},
		{Name: "TestBackslashAndQuote", Input: `{{ "foo\\'bar" | squote }}`, ExpectedOutput: "'foo\\'bar'"},
		{Name: "TestUnicode", Input: `{{ squote "foo" "👍" }}`, ExpectedOutput: "'foo' '👍'"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestToCamelCase(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | toCamelCase }}`, ExpectedOutput: ""},
		{Name: "TestCamelCase", Input: `{{ "foo bar" | toCamelCase }}`, ExpectedOutput: "fooBar"},
		{Name: "TestCamelCaseWithUpperCase", Input: `{{ "FoO  bar" | toCamelCase }}`, ExpectedOutput: "fooBar"},
		{Name: "TestCamelCaseWithSpace", Input: `{{ "foo  bar" | toCamelCase }}`, ExpectedOutput: "fooBar"},
		{Name: "TestCamelCaseWithUnderscore", Input: `{{ "foo_bar" | toCamelCase }}`, ExpectedOutput: "fooBar"},
		{Name: "TestCamelCaseWithHyphen", Input: `{{ "foo-bar" | toCamelCase }}`, ExpectedOutput: "fooBar"},
		{Name: "TestCamelCaseWithMixed", Input: `{{ "foo-bar_baz" | toCamelCase }}`, ExpectedOutput: "fooBarBaz"},
		{Name: "TestComplexCase", Input: `{{ toCamelCase "___complex__case_" }}`, ExpectedOutput: "complexCase"},
		{Name: "TestCamelCaseWithUnderscorePrefix", Input: `{{ toCamelCase "_camel_case" }}`, ExpectedOutput: "camelCase"},
		{Name: "TestCamelCaseWithUnderscoreSuffix", Input: `{{ toCamelCase "http_server" }}`, ExpectedOutput: "httpServer"},
		{Name: "TestCamelCaseWithHyphenSuffix", Input: `{{ toCamelCase "no_https" }}`, ExpectedOutput: "noHttps"},
		{Name: "TestCamelCaseWithAllLowercase", Input: `{{ toCamelCase "all" }}`, ExpectedOutput: "all"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestToKebakCase(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | toKebabCase }}`, ExpectedOutput: ""},
		{Name: "TestKebabCase", Input: `{{ "foo bar" | toKebabCase }}`, ExpectedOutput: "foo-bar"},
		{Name: "TestKebabCaseWithSpace", Input: `{{ "foo  bar" | toKebabCase }}`, ExpectedOutput: "foo-bar"},
		{Name: "TestKebabCaseWithUnderscore", Input: `{{ "foo_bar" | toKebabCase }}`, ExpectedOutput: "foo-bar"},
		{Name: "TestKebabCaseWithHyphen", Input: `{{ "foo-bar" | toKebabCase }}`, ExpectedOutput: "foo-bar"},
		{Name: "TestKebabCaseWithMixed", Input: `{{ "foo-bar_baz" | toKebabCase }}`, ExpectedOutput: "foo-bar-baz"},
		{Input: `{{ toKebabCase "HTTPServer" }}`, ExpectedOutput: "http-server"},
		{Input: `{{ toKebabCase "FirstName" }}`, ExpectedOutput: "first-name"},
		{Input: `{{ toKebabCase "NoHTTPS" }}`, ExpectedOutput: "no-https"},
		{Input: `{{ toKebabCase "GO_PATH" }}`, ExpectedOutput: "go-path"},
		{Input: `{{ toKebabCase "GO PATH" }}`, ExpectedOutput: "go-path"},
		{Input: `{{ toKebabCase "GO-PATH" }}`, ExpectedOutput: "go-path"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestToPascalCase(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | toPascalCase }}`, ExpectedOutput: ""},
		{Name: "TestPascalCase", Input: `{{ "foo bar" | toPascalCase }}`, ExpectedOutput: "FooBar"},
		{Name: "TestPascalCaseWithSpace", Input: `{{ "foo  bar" | toPascalCase }}`, ExpectedOutput: "FooBar"},
		{Name: "TestPascalCaseWithUnderscore", Input: `{{ "foo_bar" | toPascalCase }}`, ExpectedOutput: "FooBar"},
		{Name: "TestPascalCaseWithHyphen", Input: `{{ "foo-bar" | toPascalCase }}`, ExpectedOutput: "FooBar"},
		{Name: "TestPascalCaseWithMixed", Input: `{{ "foo-bar_baz" | toPascalCase }}`, ExpectedOutput: "FooBarBaz"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestToDotCase(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | toDotCase }}`, ExpectedOutput: ""},
		{Name: "TestDotCase", Input: `{{ "foo bar" | toDotCase }}`, ExpectedOutput: "foo.bar"},
		{Name: "TestDotCaseWithSpace", Input: `{{ "foo  bar" | toDotCase }}`, ExpectedOutput: "foo.bar"},
		{Name: "TestDotCaseWithUnderscore", Input: `{{ "foo_bar" | toDotCase }}`, ExpectedOutput: "foo.bar"},
		{Name: "TestDotCaseWithHyphen", Input: `{{ "foo-bar" | toDotCase }}`, ExpectedOutput: "foo.bar"},
		{Name: "TestDotCaseWithMixed", Input: `{{ "foo-bar_baz" | toDotCase }}`, ExpectedOutput: "foo.bar.baz"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestToPathCase(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | toPathCase }}`, ExpectedOutput: ""},
		{Name: "TestPathCase", Input: `{{ "foo bar" | toPathCase }}`, ExpectedOutput: "foo/bar"},
		{Name: "TestPathCaseWithSpace", Input: `{{ "foo  bar" | toPathCase }}`, ExpectedOutput: "foo/bar"},
		{Name: "TestPathCaseWithUnderscore", Input: `{{ "foo_bar" | toPathCase }}`, ExpectedOutput: "foo/bar"},
		{Name: "TestPathCaseWithHyphen", Input: `{{ "foo-bar" | toPathCase }}`, ExpectedOutput: "foo/bar"},
		{Name: "TestPathCaseWithMixed", Input: `{{ "foo-bar_baz" | toPathCase }}`, ExpectedOutput: "foo/bar/baz"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestToConstantCase(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | toConstantCase }}`, ExpectedOutput: ""},
		{Name: "TestConstantCase", Input: `{{ "foo bar" | toConstantCase }}`, ExpectedOutput: "FOO_BAR"},
		{Name: "TestConstantCaseWithSpace", Input: `{{ "foo  bar" | toConstantCase }}`, ExpectedOutput: "FOO_BAR"},
		{Name: "TestConstantCaseWithUnderscore", Input: `{{ "foo_bar" | toConstantCase }}`, ExpectedOutput: "FOO_BAR"},
		{Name: "TestConstantCaseWithHyphen", Input: `{{ "foo-bar" | toConstantCase }}`, ExpectedOutput: "FOO_BAR"},
		{Name: "TestConstantCaseWithMixed", Input: `{{ "foo-bar_baz" | toConstantCase }}`, ExpectedOutput: "FOO_BAR_BAZ"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestToSnakeCase(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | toSnakeCase }}`, ExpectedOutput: ""},
		{Name: "TestSnakeCase", Input: `{{ "foo bar" | toSnakeCase }}`, ExpectedOutput: "foo_bar"},
		{Name: "TestSnakeCaseWithSpace", Input: `{{ "foo  bar" | toSnakeCase }}`, ExpectedOutput: "foo_bar"},
		{Name: "TestSnakeCaseWithUnderscore", Input: `{{ "foo_bar" | toSnakeCase }}`, ExpectedOutput: "foo_bar"},
		{Name: "TestSnakeCaseWithHyphen", Input: `{{ "foo-bar" | toSnakeCase }}`, ExpectedOutput: "foo_bar"},
		{Name: "TestSnakeCaseWithMixed", Input: `{{ "foo-bar_baz" | toSnakeCase }}`, ExpectedOutput: "foo_bar_baz"},
		{Input: `{{ toSnakeCase "http2xx" }}`, ExpectedOutput: "http_2xx"},
		{Input: `{{ toSnakeCase "HTTP20xOK" }}`, ExpectedOutput: "http_20x_ok"},
		{Input: `{{ toSnakeCase "Duration2m3s" }}`, ExpectedOutput: "duration_2m_3s"},
		{Input: `{{ toSnakeCase "Bld4Floor3rd" }}`, ExpectedOutput: "bld_4floor_3rd"},
		{Input: `{{ toSnakeCase "FirstName" }}`, ExpectedOutput: "first_name"},
		{Input: `{{ toSnakeCase "HTTPServer" }}`, ExpectedOutput: "http_server"},
		{Input: `{{ toSnakeCase "NoHTTPS" }}`, ExpectedOutput: "no_https"},
		{Input: `{{ toSnakeCase "GO_PATH" }}`, ExpectedOutput: "go_path"},
		{Input: `{{ toSnakeCase "GO PATH" }}`, ExpectedOutput: "go_path"},
		{Input: `{{ toSnakeCase "GO-PATH" }}`, ExpectedOutput: "go_path"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestToTitleCase(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | toTitleCase }}`, ExpectedOutput: ""},
		{Name: "TestTitleCase", Input: `{{ "foo bar" | toTitleCase }}`, ExpectedOutput: "Foo Bar"},
		{Name: "TestTitleCaseWithSpace", Input: `{{ "foo  bar" | toTitleCase }}`, ExpectedOutput: "Foo  Bar"},
		{Name: "TestTitleCaseWithUnderscore", Input: `{{ "foo_bar" | toTitleCase }}`, ExpectedOutput: "Foo_bar"},
		{Name: "TestTitleCaseWithHyphen", Input: `{{ "foo-bar" | toTitleCase }}`, ExpectedOutput: "Foo-Bar"},
		{Name: "TestTitleCaseWithMixed", Input: `{{ "foo-bar_baz" | toTitleCase }}`, ExpectedOutput: "Foo-Bar_baz"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestUntitle(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | untitle }}`, ExpectedOutput: ""},
		{Name: "TestUnTitle", Input: `{{ "Foo Bar" | untitle }}`, ExpectedOutput: "foo bar"},
		{Name: "TestUnTitleWithSpace", Input: `{{ "Foo  Bar" | untitle }}`, ExpectedOutput: "foo  bar"},
		{Name: "TestUnTitleWithUnderscore", Input: `{{ "Foo_bar" | untitle }}`, ExpectedOutput: "foo_bar"},
		{Name: "TestUnTitleWithHyphen", Input: `{{ "Foo-Bar" | untitle }}`, ExpectedOutput: "foo-Bar"},
		{Name: "TestUnTitleWithMixed", Input: `{{ "Foo-Bar_baz" | untitle }}`, ExpectedOutput: "foo-Bar_baz"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestSwapCase(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | swapCase }}`, ExpectedOutput: ""},
		{Name: "TestSwapCase", Input: `{{ "Foo Bar" | swapCase }}`, ExpectedOutput: "fOO bAR"},
		{Name: "TestSwapCaseWithSpace", Input: `{{ "Foo  Bar" | swapCase }}`, ExpectedOutput: "fOO  bAR"},
		{Name: "TestSwapCaseWithUnderscore", Input: `{{ "Foo_bar" | swapCase }}`, ExpectedOutput: "fOO_BAR"},
		{Name: "TestSwapCaseWithHyphen", Input: `{{ "Foo-Bar" | swapCase }}`, ExpectedOutput: "fOO-bAR"},
		{Name: "TestSwapCaseWithMixed", Input: `{{ "Foo-Bar_baz" | swapCase }}`, ExpectedOutput: "fOO-bAR_BAZ"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestCapitalize(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | capitalize }}`, ExpectedOutput: ""},
		{Name: "CapitalizeAlreadyUpper", Input: `{{ "Foo Bar" | capitalize }}`, ExpectedOutput: "Foo Bar"},
		{Name: "CapitalizeWithSpace", Input: `{{ " fe bar" | capitalize }}`, ExpectedOutput: " Fe bar"},
		{Name: "CapitalizeWithNumber", Input: `{{ "123boo_bar" | capitalize }}`, ExpectedOutput: "123Boo_bar"},
		{Name: "CapitalizeWithUnderscore", Input: `{{ "boo_bar" | capitalize }}`, ExpectedOutput: "Boo_bar"},
		{Name: "CapitalizeWithEmoji", Input: `{{ "👍 good" | capitalize }}`, ExpectedOutput: "👍 Good"},
		{Name: "CapitalizeWithUnicode", Input: `{{ "été" | capitalize }}`, ExpectedOutput: "Été"},
		{Name: "CapitalizeWithArabic", Input: `{{ "مرحبا" | capitalize }}`, ExpectedOutput: "مرحبا"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestUncapitalize(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | uncapitalize }}`, ExpectedOutput: ""},
		{Name: "UncapitalizeAlreadyLower", Input: `{{ "foo bar" | uncapitalize }}`, ExpectedOutput: "foo bar"},
		{Name: "UncapitalizeWithSpace", Input: `{{ " Foo bar" | uncapitalize }}`, ExpectedOutput: " foo bar"},
		{Name: "UncapitalizeWithNumber", Input: `{{ "123Boo_bar" | uncapitalize }}`, ExpectedOutput: "123boo_bar"},
		{Name: "UncapitalizeWithUnderscore", Input: `{{ "Boo_bar" | uncapitalize }}`, ExpectedOutput: "boo_bar"},
		{Name: "UncapitalizeWithEmoji", Input: `{{ "👍 Good" | uncapitalize }}`, ExpectedOutput: "👍 good"},
		{Name: "UncapitalizeWithUnicode", Input: `{{ "Été" | uncapitalize }}`, ExpectedOutput: "été"},
		{Name: "UncapitalizeWithArabic", Input: `{{ "مرحبا" | uncapitalize }}`, ExpectedOutput: "مرحبا"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestSplit(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ $v := ("" | split "-") }}{{$v._0}}`, ExpectedOutput: ""},
		{Name: "TestSplit", Input: `{{ $v := ("foo$bar$baz" | split "$") }}{{$v._0}} {{$v._1}} {{$v._2}}`, ExpectedOutput: "foo bar baz"},
		{Name: "TestSplitWithEmpty", Input: `{{ $v := ("foo$bar$" | split "$") }}{{$v._0}} {{$v._1}} {{$v._2}}`, ExpectedOutput: "foo bar "},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestSplitn(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ $v := ("" | splitn "-" 3) }}{{$v._0}}`, ExpectedOutput: ""},
		{Name: "TestSplit", Input: `{{ $v := ("foo$bar$baz" | splitn "$" 2) }}{{$v._0}} {{$v._1}}`, ExpectedOutput: "foo bar$baz"},
		{Name: "TestSplitWithEmpty", Input: `{{ $v := ("foo$bar$" | splitn "$" 2) }}{{$v._0}} {{$v._1}}`, ExpectedOutput: "foo bar$"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestSubstring(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | substr 0 3 }}`, ExpectedOutput: ""},
		{Name: "TestEmptyWithNegativeValue", Input: `{{ "" | substr -1 -4 }}`, ExpectedOutput: ""},
		{Name: "TestSubstring", Input: `{{ "foobar" | substr 0 3 }}`, ExpectedOutput: "foo"},
		{Name: "TestSubstringNegativeEnd", Input: `{{ "foobar" | substr 0 -3 }}`, ExpectedOutput: "foo"},
		{Name: "TestSubstringNegativeStart", Input: `{{ "foobar" | substr -3 6 }}`, ExpectedOutput: "bar"},
		{Name: "TestSubstringNegativeStartAndEnd", Input: `{{ "foobar" | substr -3 -1 }}`, ExpectedOutput: "ba"},
		{Name: "TestSubstringInvalidRange", Input: `{{ "foobar" | substr -3 -3 }}`, ExpectedOutput: ""},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestIndent(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | indent 3 }}`, ExpectedOutput: "   "},
		{Name: "TestIndent", Input: `{{ "foo\nbar" | indent 3 }}`, ExpectedOutput: "   foo\n   bar"},
		{Name: "TestIndentWithSpace", Input: `{{ "foo\n bar" | indent 3 }}`, ExpectedOutput: "   foo\n    bar"},
		{Name: "TestIndentWithTab", Input: `{{ "foo\n\tbar" | indent 3 }}`, ExpectedOutput: "   foo\n   \tbar"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestNindent(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | nindent 3 }}`, ExpectedOutput: "\n   "},
		{Name: "TestIndent", Input: `{{ "foo\nbar" | nindent 3 }}`, ExpectedOutput: "\n   foo\n   bar"},
		{Name: "TestIndentWithSpace", Input: `{{ "foo\n bar" | nindent 3 }}`, ExpectedOutput: "\n   foo\n    bar"},
		{Name: "TestIndentWithTab", Input: `{{ "foo\n\tbar" | nindent 3 }}`, ExpectedOutput: "\n   foo\n   \tbar"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestSeq(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ seq 0 1 3 }}`, ExpectedOutput: "0 1 2 3"},
		{Input: `{{ seq 0 3 10 }}`, ExpectedOutput: "0 3 6 9"},
		{Input: `{{ seq 3 3 2 }}`, ExpectedOutput: ""},
		{Input: `{{ seq 3 -3 2 }}`, ExpectedOutput: "3"},
		{Input: `{{ seq }}`, ExpectedOutput: ""},
		{Input: `{{ seq 0 4 }}`, ExpectedOutput: "0 1 2 3 4"},
		{Input: `{{ seq 5 }}`, ExpectedOutput: "1 2 3 4 5"},
		{Input: `{{ seq -5 }}`, ExpectedOutput: "1 0 -1 -2 -3 -4 -5"},
		{Input: `{{ seq 0 }}`, ExpectedOutput: "1 0"},
		{Input: `{{ seq 0 1 2 3 }}`, ExpectedOutput: ""},
		{Input: `{{ seq 0 -4 }}`, ExpectedOutput: "0 -1 -2 -3 -4"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}
