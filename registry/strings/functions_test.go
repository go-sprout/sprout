package strings_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/strings"
)

func TestNoSpace(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | nospace }}`, Expected: ""},
		{Name: "TestSpaceOnly", Input: `{{ " " | nospace }}`, Expected: ""},
		{Name: "TestLeadingSpace", Input: `{{ " foo" | nospace }}`, Expected: "foo"},
		{Name: "TestTrailingSpace", Input: `{{ "foo " | nospace }}`, Expected: "foo"},
		{Name: "TestLeadingAndTrailingSpace", Input: `{{ " foo " | nospace }}`, Expected: "foo"},
		{Name: "TestMultipleSpaces", Input: `{{ " foo bar " | nospace }}`, Expected: "foobar"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestTrim(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | trim }}`, Expected: ""},
		{Name: "TestSpaceOnly", Input: `{{ " " | trim }}`, Expected: ""},
		{Name: "TestLeadingSpace", Input: `{{ " foo" | trim }}`, Expected: "foo"},
		{Name: "TestTrailingSpace", Input: `{{ "foo " | trim }}`, Expected: "foo"},
		{Name: "TestLeadingAndTrailingSpace", Input: `{{ " foo " | trim }}`, Expected: "foo"},
		{Name: "TestMultipleSpaces", Input: `{{ " foo bar " | trim }}`, Expected: "foo bar"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestTrimAll(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | trimAll "-" }}`, Expected: ""},
		{Name: "TestAllDashes", Input: `{{ "---------" | trimAll "-" }}`, Expected: ""},
		{Name: "TestNoDashes", Input: `{{ "foo" | trimAll "-" }}`, Expected: "foo"},
		{Name: "TestSomeDashes", Input: `{{ "-f--o-o-" | trimAll "-" }}`, Expected: "f--o-o"},
		{Name: "TestOtherDashes", Input: `{{ "-f--o-o-" | trimAll "-o" }}`, Expected: "f"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestTrimPrefix(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | trimPrefix "-" }}`, Expected: ""},
		{Name: "TestDoubleDash", Input: `{{ "--" | trimPrefix "-" }}`, Expected: "-"},
		{Name: "TestNoPrefix", Input: `{{ "foo" | trimPrefix "-" }}`, Expected: "foo"},
		{Name: "TestSinglePrefix", Input: `{{ "-foo-" | trimPrefix "-" }}`, Expected: "foo-"},
		{Name: "TestMultiplePrefix", Input: `{{ "-foo-" | trimPrefix "-f" }}`, Expected: "oo-"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestTrimSuffix(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | trimSuffix "-" }}`, Expected: ""},
		{Name: "TestDoubleDash", Input: `{{ "--" | trimSuffix "-" }}`, Expected: "-"},
		{Name: "TestNoSuffix", Input: `{{ "foo" | trimSuffix "-" }}`, Expected: "foo"},
		{Name: "TestSingleSuffix", Input: `{{ "-foo-" | trimSuffix "-" }}`, Expected: "-foo"},
		{Name: "TestMultipleSuffix", Input: `{{ "-foo-" | trimSuffix "o-" }}`, Expected: "-fo"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestContains(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | contains "-" }}`, Expected: "false"},
		{Name: "TestContains", Input: `{{ "foo" | contains "o" }}`, Expected: "true"},
		{Name: "TestNotContains", Input: `{{ "foo" | contains "x" }}`, Expected: "false"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestHasPrefix(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | hasPrefix "-" }}`, Expected: "false"},
		{Name: "TestHasPrefix", Input: `{{ "foo" | hasPrefix "f" }}`, Expected: "true"},
		{Name: "TestNotHasPrefix", Input: `{{ "foo" | hasPrefix "o" }}`, Expected: "false"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestHasSuffix(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | hasSuffix "-" }}`, Expected: "false"},
		{Name: "TestHasSuffix", Input: `{{ "foo" | hasSuffix "o" }}`, Expected: "true"},
		{Name: "TestNotHasSuffix", Input: `{{ "foo" | hasSuffix "f" }}`, Expected: "false"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestToLower(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | toLower }}`, Expected: ""},
		{Name: "TestLower", Input: `{{ "foo" | toLower }}`, Expected: "foo"},
		{Name: "TestUpper", Input: `{{ "FOO" | toLower }}`, Expected: "foo"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestToUpper(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | toUpper }}`, Expected: ""},
		{Name: "TestLower", Input: `{{ "foo" | toUpper }}`, Expected: "FOO"},
		{Name: "TestUpper", Input: `{{ "FOO" | toUpper }}`, Expected: "FOO"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestReplace(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | replace "-" "x" }}`, Expected: ""},
		{Name: "TestReplace", Input: `{{ "foo" | replace "o" "x" }}`, Expected: "fxx"},
		{Name: "TestNotReplace", Input: `{{ "foo" | replace "x" "y" }}`, Expected: "foo"},
		{Name: "TestMultipleReplace", Input: `{{ "foo" | replace "o" "x" | replace "f" "y" }}`, Expected: "yxx"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestRepeat(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | repeat 3 }}`, Expected: ""},
		{Name: "TestRepeat", Input: `{{ "foo" | repeat 3 }}`, Expected: "foofoofoo"},
		{Name: "TestRepeatZero", Input: `{{ "foo" | repeat 0 }}`, Expected: ""},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestJoin(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestNil", Input: `{{ .nil | join "-" }}`, Expected: "", Data: map[string]any{"nil": nil}},
		{Name: "TestIntSlice", Input: `{{ .test | join "-" }}`, Expected: "1-2-3", Data: map[string]any{"test": []int{1, 2, 3}}},
		{Name: "TestStringSlice", Input: `{{ .test | join "-" }}`, Expected: "a-b-c", Data: map[string]any{"test": []string{"a", "b", "c"}}},
		{Name: "TestString", Input: `{{ .test | join "-" }}`, Expected: "abc", Data: map[string]any{"test": "abc"}},
		{Name: "TestMixedSlice", Input: `{{ .test | join "-" }}`, Expected: "a-1-true", Data: map[string]any{"test": []any{"a", nil, 1, true}}},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestTrunc(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | trunc 3 }}`, Expected: ""},
		{Name: "TestTruncate", Input: `{{ "foooooo" | trunc 3 }}`, Expected: "foo"},
		{Name: "TestNegativeTruncate", Input: `{{ "foobar" | trunc -3 }}`, Expected: "bar"},
		{Name: "TestNegativeLargeTruncate", Input: `{{ "foobar" | trunc -999 }}`, Expected: "foobar"},
		{Name: "TestZeroTruncate", Input: `{{ "foobar" | trunc 0 }}`, Expected: ""},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestShuffle(t *testing.T) {
	var tc = []pesticide.RegexpTestCase{
		{Template: `{{ shuffle "" }}`, Length: 0, Regexp: `^$`},
		{Template: `{{ shuffle "hey" }}`, Length: 3, Regexp: `^[hey]{3}$`},
		{Template: `{{ shuffle "foo bar baz" }}`, Length: 11, Regexp: `^[\sfobazr]{11}$`},
	}

	pesticide.RunRegexpTestCases(t, strings.NewRegistry(), tc)
}

func TestEllipsis(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | ellipsis 3 }}`, Expected: ""},
		{Name: "TestShort", Input: `{{ "foo" | ellipsis 5 }}`, Expected: "foo"},
		{Name: "TestTruncate", Input: `{{ "foooooo" | ellipsis 6 }}`, Expected: "foo..."},
		{Name: "TestZeroTruncate", Input: `{{ "foobar" | ellipsis 0 }}`, Expected: "foobar"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestEllipsisBoth(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | ellipsisBoth 3 5 }}`, Expected: ""},
		{Name: "TestShort", Input: `{{ "foo" | ellipsisBoth 5 4 }}`, Expected: "foo"},
		{Name: "TestTruncate", Input: `{{ "foooboooooo" | ellipsisBoth 4 9 }}`, Expected: "...boo..."},
		{Name: "TestZeroTruncate", Input: `{{ "foobar" | ellipsisBoth 0 0 }}`, Expected: "foobar"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestInitials(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | initials }}`, Expected: ""},
		{Name: "TestSingle", Input: `{{ "f" | initials }}`, Expected: "f"},
		{Name: "TestTwo", Input: `{{ "foo" | initials }}`, Expected: "f"},
		{Name: "TestThree", Input: `{{ "foo bar" | initials }}`, Expected: "fb"},
		{Name: "TestMultipleSpaces", Input: `{{ "foo  bar" | initials }}`, Expected: "fb"},
		{Name: "TestWithUppercased", Input: `{{ " Foo bar" | initials }}`, Expected: "Fb"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestPlural(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestZero", Input: `{{ 0 | plural "single" "many" }}`, Expected: "many"},
		{Name: "TestSingle", Input: `{{ 1 | plural "single" "many" }}`, Expected: "single"},
		{Name: "TestMultiple", Input: `{{ 2 | plural "single" "many" }}`, Expected: "many"},
		{Name: "TestNegative", Input: `{{ -1 | plural "single" "many" }}`, Expected: "many"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestWrap(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | wrap 10 }}`, Expected: ""},
		{Name: "TestNegativeWrap", Input: `{{ wrap -1 "With a negative wrap." }}`, Expected: "With\na\nnegative\nwrap."},
		{Name: "TestWrap", Input: `{{ "This is a long string that needs to be wrapped." | wrap 10 }}`, Expected: "This is a\nlong\nstring\nthat needs\nto be\nwrapped."},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestWrapWith(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | wrapWith 10 "\t" }}`, Expected: ""},
		{Name: "TestWrap", Input: `{{ "This is a long string that needs to be wrapped." | wrapWith 10 "\t" }}`, Expected: "This is a\tlong\tstring\tthat needs\tto be\twrapped."},
		{Name: "TestWrapWithLongWord", Input: `{{ "This is a long string that needs to be wrapped with a looooooooooooooooooooooooooooooooooooong word." | wrapWith 10 "\t" }}`, Expected: "This is a\tlong\tstring\tthat needs\tto be\twrapped\twith a\tlooooooooo\toooooooooo\toooooooooo\toooooooong\tword."},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestQuote(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | quote }}`, Expected: `""`},
		{Name: "TestNil", Input: `{{ quote .nil }}`, Expected: ``, Data: map[string]any{"nil": nil}},
		{Name: "TestQuote", Input: `{{ "foo" | quote }}`, Expected: `"foo"`},
		{Name: "TestSpace", Input: `{{ "foo bar" | quote }}`, Expected: `"foo bar"`},
		{Name: "TestQuote", Input: `{{ "foo \"bar\"" | quote }}`, Expected: `"foo \"bar\""`},
		{Name: "TestNewline", Input: `{{ "foo\nbar" | quote }}`, Expected: `"foo\nbar"`},
		{Name: "TestBackslash", Input: `{{ "foo\\bar" | quote }}`, Expected: `"foo\\bar"`},
		{Name: "TestBackslashAndQuote", Input: `{{ "foo\\\"bar" | quote }}`, Expected: `"foo\\\"bar"`},
		{Name: "TestUnicode", Input: `{{ quote "foo" "👍" }}`, Expected: `"foo" "👍"`},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestSquote(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | squote }}`, Expected: "''"},
		{Name: "TestNil", Input: `{{ squote .nil }}`, Expected: "", Data: map[string]any{"nil": nil}},
		{Name: "TestQuote", Input: `{{ "foo" | squote }}`, Expected: "'foo'"},
		{Name: "TestSpace", Input: `{{ "foo bar" | squote }}`, Expected: "'foo bar'"},
		{Name: "TestQuote", Input: `{{ "foo 'bar'" | squote }}`, Expected: "'foo 'bar''"},
		{Name: "TestNewline", Input: `{{ "foo\nbar" | squote }}`, Expected: "'foo\nbar'"},
		{Name: "TestBackslash", Input: `{{ "foo\\bar" | squote }}`, Expected: "'foo\\bar'"},
		{Name: "TestBackslashAndQuote", Input: `{{ "foo\\'bar" | squote }}`, Expected: "'foo\\'bar'"},
		{Name: "TestUnicode", Input: `{{ squote "foo" "👍" }}`, Expected: "'foo' '👍'"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestToCamelCase(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | toCamelCase }}`, Expected: ""},
		{Name: "TestCamelCase", Input: `{{ "foo bar" | toCamelCase }}`, Expected: "fooBar"},
		{Name: "TestCamelCaseWithUpperCase", Input: `{{ "FoO  bar" | toCamelCase }}`, Expected: "fooBar"},
		{Name: "TestCamelCaseWithSpace", Input: `{{ "foo  bar" | toCamelCase }}`, Expected: "fooBar"},
		{Name: "TestCamelCaseWithUnderscore", Input: `{{ "foo_bar" | toCamelCase }}`, Expected: "fooBar"},
		{Name: "TestCamelCaseWithHyphen", Input: `{{ "foo-bar" | toCamelCase }}`, Expected: "fooBar"},
		{Name: "TestCamelCaseWithMixed", Input: `{{ "foo-bar_baz" | toCamelCase }}`, Expected: "fooBarBaz"},
		{Name: "TestComplexCase", Input: `{{ toCamelCase "___complex__case_" }}`, Expected: "complexCase"},
		{Name: "TestCamelCaseWithUnderscorePrefix", Input: `{{ toCamelCase "_camel_case" }}`, Expected: "camelCase"},
		{Name: "TestCamelCaseWithUnderscoreSuffix", Input: `{{ toCamelCase "http_server" }}`, Expected: "httpServer"},
		{Name: "TestCamelCaseWithHyphenSuffix", Input: `{{ toCamelCase "no_https" }}`, Expected: "noHttps"},
		{Name: "TestCamelCaseWithAllLowercase", Input: `{{ toCamelCase "all" }}`, Expected: "all"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestToKebakCase(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | toKebabCase }}`, Expected: ""},
		{Name: "TestKebabCase", Input: `{{ "foo bar" | toKebabCase }}`, Expected: "foo-bar"},
		{Name: "TestKebabCaseWithSpace", Input: `{{ "foo  bar" | toKebabCase }}`, Expected: "foo-bar"},
		{Name: "TestKebabCaseWithUnderscore", Input: `{{ "foo_bar" | toKebabCase }}`, Expected: "foo-bar"},
		{Name: "TestKebabCaseWithHyphen", Input: `{{ "foo-bar" | toKebabCase }}`, Expected: "foo-bar"},
		{Name: "TestKebabCaseWithMixed", Input: `{{ "foo-bar_baz" | toKebabCase }}`, Expected: "foo-bar-baz"},
		{Input: `{{ toKebabCase "HTTPServer" }}`, Expected: "http-server"},
		{Input: `{{ toKebabCase "FirstName" }}`, Expected: "first-name"},
		{Input: `{{ toKebabCase "NoHTTPS" }}`, Expected: "no-https"},
		{Input: `{{ toKebabCase "GO_PATH" }}`, Expected: "go-path"},
		{Input: `{{ toKebabCase "GO PATH" }}`, Expected: "go-path"},
		{Input: `{{ toKebabCase "GO-PATH" }}`, Expected: "go-path"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestToPascalCase(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | toPascalCase }}`, Expected: ""},
		{Name: "TestPascalCase", Input: `{{ "foo bar" | toPascalCase }}`, Expected: "FooBar"},
		{Name: "TestPascalCaseWithSpace", Input: `{{ "foo  bar" | toPascalCase }}`, Expected: "FooBar"},
		{Name: "TestPascalCaseWithUnderscore", Input: `{{ "foo_bar" | toPascalCase }}`, Expected: "FooBar"},
		{Name: "TestPascalCaseWithHyphen", Input: `{{ "foo-bar" | toPascalCase }}`, Expected: "FooBar"},
		{Name: "TestPascalCaseWithMixed", Input: `{{ "foo-bar_baz" | toPascalCase }}`, Expected: "FooBarBaz"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestToDotCase(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | toDotCase }}`, Expected: ""},
		{Name: "TestDotCase", Input: `{{ "foo bar" | toDotCase }}`, Expected: "foo.bar"},
		{Name: "TestDotCaseWithSpace", Input: `{{ "foo  bar" | toDotCase }}`, Expected: "foo.bar"},
		{Name: "TestDotCaseWithUnderscore", Input: `{{ "foo_bar" | toDotCase }}`, Expected: "foo.bar"},
		{Name: "TestDotCaseWithHyphen", Input: `{{ "foo-bar" | toDotCase }}`, Expected: "foo.bar"},
		{Name: "TestDotCaseWithMixed", Input: `{{ "foo-bar_baz" | toDotCase }}`, Expected: "foo.bar.baz"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestToPathCase(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | toPathCase }}`, Expected: ""},
		{Name: "TestPathCase", Input: `{{ "foo bar" | toPathCase }}`, Expected: "foo/bar"},
		{Name: "TestPathCaseWithSpace", Input: `{{ "foo  bar" | toPathCase }}`, Expected: "foo/bar"},
		{Name: "TestPathCaseWithUnderscore", Input: `{{ "foo_bar" | toPathCase }}`, Expected: "foo/bar"},
		{Name: "TestPathCaseWithHyphen", Input: `{{ "foo-bar" | toPathCase }}`, Expected: "foo/bar"},
		{Name: "TestPathCaseWithMixed", Input: `{{ "foo-bar_baz" | toPathCase }}`, Expected: "foo/bar/baz"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestToConstantCase(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | toConstantCase }}`, Expected: ""},
		{Name: "TestConstantCase", Input: `{{ "foo bar" | toConstantCase }}`, Expected: "FOO_BAR"},
		{Name: "TestConstantCaseWithSpace", Input: `{{ "foo  bar" | toConstantCase }}`, Expected: "FOO_BAR"},
		{Name: "TestConstantCaseWithUnderscore", Input: `{{ "foo_bar" | toConstantCase }}`, Expected: "FOO_BAR"},
		{Name: "TestConstantCaseWithHyphen", Input: `{{ "foo-bar" | toConstantCase }}`, Expected: "FOO_BAR"},
		{Name: "TestConstantCaseWithMixed", Input: `{{ "foo-bar_baz" | toConstantCase }}`, Expected: "FOO_BAR_BAZ"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestToSnakeCase(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | toSnakeCase }}`, Expected: ""},
		{Name: "TestSnakeCase", Input: `{{ "foo bar" | toSnakeCase }}`, Expected: "foo_bar"},
		{Name: "TestSnakeCaseWithSpace", Input: `{{ "foo  bar" | toSnakeCase }}`, Expected: "foo_bar"},
		{Name: "TestSnakeCaseWithUnderscore", Input: `{{ "foo_bar" | toSnakeCase }}`, Expected: "foo_bar"},
		{Name: "TestSnakeCaseWithHyphen", Input: `{{ "foo-bar" | toSnakeCase }}`, Expected: "foo_bar"},
		{Name: "TestSnakeCaseWithMixed", Input: `{{ "foo-bar_baz" | toSnakeCase }}`, Expected: "foo_bar_baz"},
		{Input: `{{ toSnakeCase "http2xx" }}`, Expected: "http_2xx"},
		{Input: `{{ toSnakeCase "HTTP20xOK" }}`, Expected: "http_20x_ok"},
		{Input: `{{ toSnakeCase "Duration2m3s" }}`, Expected: "duration_2m_3s"},
		{Input: `{{ toSnakeCase "Bld4Floor3rd" }}`, Expected: "bld_4floor_3rd"},
		{Input: `{{ toSnakeCase "FirstName" }}`, Expected: "first_name"},
		{Input: `{{ toSnakeCase "HTTPServer" }}`, Expected: "http_server"},
		{Input: `{{ toSnakeCase "NoHTTPS" }}`, Expected: "no_https"},
		{Input: `{{ toSnakeCase "GO_PATH" }}`, Expected: "go_path"},
		{Input: `{{ toSnakeCase "GO PATH" }}`, Expected: "go_path"},
		{Input: `{{ toSnakeCase "GO-PATH" }}`, Expected: "go_path"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestToTitleCase(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | toTitleCase }}`, Expected: ""},
		{Name: "TestTitleCase", Input: `{{ "foo bar" | toTitleCase }}`, Expected: "Foo Bar"},
		{Name: "TestTitleCaseWithSpace", Input: `{{ "foo  bar" | toTitleCase }}`, Expected: "Foo  Bar"},
		{Name: "TestTitleCaseWithUnderscore", Input: `{{ "foo_bar" | toTitleCase }}`, Expected: "Foo_bar"},
		{Name: "TestTitleCaseWithHyphen", Input: `{{ "foo-bar" | toTitleCase }}`, Expected: "Foo-Bar"},
		{Name: "TestTitleCaseWithMixed", Input: `{{ "foo-bar_baz" | toTitleCase }}`, Expected: "Foo-Bar_baz"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestUntitle(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | untitle }}`, Expected: ""},
		{Name: "TestUnTitle", Input: `{{ "Foo Bar" | untitle }}`, Expected: "foo bar"},
		{Name: "TestUnTitleWithSpace", Input: `{{ "Foo  Bar" | untitle }}`, Expected: "foo  bar"},
		{Name: "TestUnTitleWithUnderscore", Input: `{{ "Foo_bar" | untitle }}`, Expected: "foo_bar"},
		{Name: "TestUnTitleWithHyphen", Input: `{{ "Foo-Bar" | untitle }}`, Expected: "foo-Bar"},
		{Name: "TestUnTitleWithMixed", Input: `{{ "Foo-Bar_baz" | untitle }}`, Expected: "foo-Bar_baz"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestSwapCase(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | swapCase }}`, Expected: ""},
		{Name: "TestSwapCase", Input: `{{ "Foo Bar" | swapCase }}`, Expected: "fOO bAR"},
		{Name: "TestSwapCaseWithSpace", Input: `{{ "Foo  Bar" | swapCase }}`, Expected: "fOO  bAR"},
		{Name: "TestSwapCaseWithUnderscore", Input: `{{ "Foo_bar" | swapCase }}`, Expected: "fOO_BAR"},
		{Name: "TestSwapCaseWithHyphen", Input: `{{ "Foo-Bar" | swapCase }}`, Expected: "fOO-bAR"},
		{Name: "TestSwapCaseWithMixed", Input: `{{ "Foo-Bar_baz" | swapCase }}`, Expected: "fOO-bAR_BAZ"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestSplit(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ $v := ("" | split "-") }}{{$v._0}}`, Expected: ""},
		{Name: "TestSplit", Input: `{{ $v := ("foo$bar$baz" | split "$") }}{{$v._0}} {{$v._1}} {{$v._2}}`, Expected: "foo bar baz"},
		{Name: "TestSplitWithEmpty", Input: `{{ $v := ("foo$bar$" | split "$") }}{{$v._0}} {{$v._1}} {{$v._2}}`, Expected: "foo bar "},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestSplitn(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ $v := ("" | splitn "-" 3) }}{{$v._0}}`, Expected: ""},
		{Name: "TestSplit", Input: `{{ $v := ("foo$bar$baz" | splitn "$" 2) }}{{$v._0}} {{$v._1}}`, Expected: "foo bar$baz"},
		{Name: "TestSplitWithEmpty", Input: `{{ $v := ("foo$bar$" | splitn "$" 2) }}{{$v._0}} {{$v._1}}`, Expected: "foo bar$"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestSubstring(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | substr 0 3 }}`, Expected: ""},
		{Name: "TestEmptyWithNegativeValue", Input: `{{ "" | substr -1 -4 }}`, Expected: ""},
		{Name: "TestSubstring", Input: `{{ "foobar" | substr 0 3 }}`, Expected: "foo"},
		{Name: "TestSubstringNegativeEnd", Input: `{{ "foobar" | substr 0 -3 }}`, Expected: "foo"},
		{Name: "TestSubstringNegativeStart", Input: `{{ "foobar" | substr -3 6 }}`, Expected: "bar"},
		{Name: "TestSubstringNegativeStartAndEnd", Input: `{{ "foobar" | substr -3 -1 }}`, Expected: "ba"},
		{Name: "TestSubstringInvalidRange", Input: `{{ "foobar" | substr -3 -3 }}`, Expected: ""},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestIndent(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | indent 3 }}`, Expected: "   "},
		{Name: "TestIndent", Input: `{{ "foo\nbar" | indent 3 }}`, Expected: "   foo\n   bar"},
		{Name: "TestIndentWithSpace", Input: `{{ "foo\n bar" | indent 3 }}`, Expected: "   foo\n    bar"},
		{Name: "TestIndentWithTab", Input: `{{ "foo\n\tbar" | indent 3 }}`, Expected: "   foo\n   \tbar"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestNindent(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Name: "TestEmpty", Input: `{{ "" | nindent 3 }}`, Expected: "\n   "},
		{Name: "TestIndent", Input: `{{ "foo\nbar" | nindent 3 }}`, Expected: "\n   foo\n   bar"},
		{Name: "TestIndentWithSpace", Input: `{{ "foo\n bar" | nindent 3 }}`, Expected: "\n   foo\n    bar"},
		{Name: "TestIndentWithTab", Input: `{{ "foo\n\tbar" | nindent 3 }}`, Expected: "\n   foo\n   \tbar"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}

func TestSeq(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ seq 0 1 3 }}`, Expected: "0 1 2 3"},
		{Input: `{{ seq 0 3 10 }}`, Expected: "0 3 6 9"},
		{Input: `{{ seq 3 3 2 }}`, Expected: ""},
		{Input: `{{ seq 3 -3 2 }}`, Expected: "3"},
		{Input: `{{ seq }}`, Expected: ""},
		{Input: `{{ seq 0 4 }}`, Expected: "0 1 2 3 4"},
		{Input: `{{ seq 5 }}`, Expected: "1 2 3 4 5"},
		{Input: `{{ seq -5 }}`, Expected: "1 0 -1 -2 -3 -4 -5"},
		{Input: `{{ seq 0 }}`, Expected: "1 0"},
		{Input: `{{ seq 0 1 2 3 }}`, Expected: ""},
		{Input: `{{ seq 0 -4 }}`, Expected: "0 -1 -2 -3 -4"},
	}

	pesticide.RunTestCases(t, strings.NewRegistry(), tc)
}
