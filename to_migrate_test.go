package sprout

import (
	"bytes"
	"crypto/x509"
	"encoding/base32"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"testing"
	"text/template"
	"time"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
	bcrypt_lib "golang.org/x/crypto/bcrypt"
)

var urlTests = map[string]map[string]interface{}{
	"proto://auth@host:80/path?query#fragment": {
		"fragment": "fragment",
		"host":     "host:80",
		"hostname": "host",
		"opaque":   "",
		"path":     "/path",
		"query":    "query",
		"scheme":   "proto",
		"userinfo": "auth",
	},
	"proto://host:80/path": {
		"fragment": "",
		"host":     "host:80",
		"hostname": "host",
		"opaque":   "",
		"path":     "/path",
		"query":    "",
		"scheme":   "proto",
		"userinfo": "",
	},
	"something": {
		"fragment": "",
		"host":     "",
		"hostname": "",
		"opaque":   "",
		"path":     "something",
		"query":    "",
		"scheme":   "",
		"userinfo": "",
	},
	"proto://user:passwor%20d@host:80/path": {
		"fragment": "",
		"host":     "host:80",
		"hostname": "host",
		"opaque":   "",
		"path":     "/path",
		"query":    "",
		"scheme":   "proto",
		"userinfo": "user:passwor%20d",
	},
	"proto://host:80/pa%20th?key=val%20ue": {
		"fragment": "",
		"host":     "host:80",
		"hostname": "host",
		"opaque":   "",
		"path":     "/pa th",
		"query":    "key=val%20ue",
		"scheme":   "proto",
		"userinfo": "",
	},
}

func TestUrlParse(t *testing.T) {
	// testing that function is exported and working properly
	assert.NoError(t, runt(
		`{{ index ( urlParse "proto://auth@host:80/path?query#fragment" ) "host" }}`,
		"host:80"))

	// testing scenarios
	fh := NewFunctionHandler()
	for url, expected := range urlTests {
		assert.EqualValues(t, expected, fh.UrlParse(url))
	}
}

func TestUrlJoin(t *testing.T) {
	tests := map[string]string{
		`{{ urlJoin (dict "fragment" "fragment" "host" "host:80" "path" "/path" "query" "query" "scheme" "proto") }}`:       "proto://host:80/path?query#fragment",
		`{{ urlJoin (dict "fragment" "fragment" "host" "host:80" "path" "/path" "scheme" "proto" "userinfo" "ASDJKJSD") }}`: "proto://ASDJKJSD@host:80/path#fragment",
	}
	for tpl, expected := range tests {
		assert.NoError(t, runt(tpl, expected))
	}

	fh := NewFunctionHandler()
	for expected, urlMap := range urlTests {
		assert.EqualValues(t, expected, fh.UrlJoin(urlMap))
	}

}

func TestToString(t *testing.T) {
	tpl := `{{ toString 1 | kindOf }}`
	assert.NoError(t, runt(tpl, "string"))
}

func TestToStrings(t *testing.T) {
	tpl := `{{ $s := list 1 2 3 | toStrings }}{{ index $s 1 | kindOf }}`
	assert.NoError(t, runt(tpl, "string"))
	tpl = `{{ list 1 .value 2 | toStrings }}`
	values := map[string]interface{}{"value": nil}
	if err := runtv(tpl, `[1 2]`, values); err != nil {
		t.Error(err)
	}
}

func TestSortAlpha(t *testing.T) {
	// Named `append` in the function map
	tests := map[string]string{
		`{{ list "c" "a" "b" | sortAlpha | join "" }}`: "abc",
		`{{ list 2 1 4 3 | sortAlpha | join "" }}`:     "1234",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}
func TestBase64EncodeDecode(t *testing.T) {
	magicWord := "coffee"
	expect := base64.StdEncoding.EncodeToString([]byte(magicWord))

	if expect == magicWord {
		t.Fatal("Encoder doesn't work.")
	}

	tpl := `{{b64enc "coffee"}}`
	if err := runt(tpl, expect); err != nil {
		t.Error(err)
	}
	tpl = fmt.Sprintf("{{b64dec %q}}", expect)
	if err := runt(tpl, magicWord); err != nil {
		t.Error(err)
	}
}
func TestBase32EncodeDecode(t *testing.T) {
	magicWord := "coffee"
	expect := base32.StdEncoding.EncodeToString([]byte(magicWord))

	if expect == magicWord {
		t.Fatal("Encoder doesn't work.")
	}

	tpl := `{{b32enc "coffee"}}`
	if err := runt(tpl, expect); err != nil {
		t.Error(err)
	}
	tpl = fmt.Sprintf("{{b32dec %q}}", expect)
	if err := runt(tpl, magicWord); err != nil {
		t.Error(err)
	}
}

func TestRandomString(t *testing.T) {
	// Random strings are now using Masterminds/goutils's cryptographically secure random string functions
	// by default. Consequently, these tests now have no predictable character sequence. No checks for exact
	// string output are necessary.

	// {{randAlphaNum 5}} should yield five random characters
	if x, _ := runRaw(`{{randAlphaNum 5}}`, nil); utf8.RuneCountInString(x) != 5 {
		t.Errorf("String should be 5 characters; string was %v characters", utf8.RuneCountInString(x))
	}

	// {{randAlpha 5}} should yield five random characters
	if x, _ := runRaw(`{{randAlpha 5}}`, nil); utf8.RuneCountInString(x) != 5 {
		t.Errorf("String should be 5 characters; string was %v characters", utf8.RuneCountInString(x))
	}

	// {{randAscii 5}} should yield five random characters
	if x, _ := runRaw(`{{randAscii 5}}`, nil); utf8.RuneCountInString(x) != 5 {
		t.Errorf("String should be 5 characters; string was %v characters", utf8.RuneCountInString(x))
	}

	// {{randNumeric 5}} should yield five random characters
	if x, _ := runRaw(`{{randNumeric 5}}`, nil); utf8.RuneCountInString(x) != 5 {
		t.Errorf("String should be 5 characters; string was %v characters", utf8.RuneCountInString(x))
	}
}

func TestSemverCompare(t *testing.T) {
	tests := map[string]string{
		`{{ semverCompare "1.2.3" "1.2.3" }}`:  `true`,
		`{{ semverCompare "^1.2.0" "1.2.3" }}`: `true`,
		`{{ semverCompare "^1.2.0" "2.2.3" }}`: `false`,
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestSemver(t *testing.T) {
	tests := map[string]string{
		`{{ $s := semver "1.2.3-beta.1+c0ff33" }}{{ $s.Prerelease }}`: "beta.1",
		`{{ $s := semver "1.2.3-beta.1+c0ff33" }}{{ $s.Major}}`:       "1",
		`{{ semver "1.2.3" | (semver "1.2.3").Compare }}`:             `0`,
		`{{ semver "1.2.3" | (semver "1.3.3").Compare }}`:             `1`,
		`{{ semver "1.4.3" | (semver "1.2.3").Compare }}`:             `-1`,
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestRegexMatch(t *testing.T) {
	fh := NewFunctionHandler()
	regex := "[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}"

	assert.True(t, fh.RegexMatch(regex, "test@acme.com"))
	assert.True(t, fh.RegexMatch(regex, "Test@Acme.Com"))
	assert.False(t, fh.RegexMatch(regex, "test"))
	assert.False(t, fh.RegexMatch(regex, "test.com"))
	assert.False(t, fh.RegexMatch(regex, "test@acme"))
}

func TestMustRegexMatch(t *testing.T) {
	fh := NewFunctionHandler()
	regex := "[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}"

	o, err := fh.MustRegexMatch(regex, "test@acme.com")
	assert.True(t, o)
	assert.Nil(t, err)

	o, err = fh.MustRegexMatch(regex, "Test@Acme.Com")
	assert.True(t, o)
	assert.Nil(t, err)

	o, err = fh.MustRegexMatch(regex, "test")
	assert.False(t, o)
	assert.Nil(t, err)

	o, err = fh.MustRegexMatch(regex, "test.com")
	assert.False(t, o)
	assert.Nil(t, err)

	o, err = fh.MustRegexMatch(regex, "test@acme")
	assert.False(t, o)
	assert.Nil(t, err)
}

func TestRegexFindAll(t *testing.T) {
	fh := NewFunctionHandler()
	regex := "a{2}"
	assert.Equal(t, 1, len(fh.RegexFindAll(regex, "aa", -1)))
	assert.Equal(t, 1, len(fh.RegexFindAll(regex, "aaaaaaaa", 1)))
	assert.Equal(t, 2, len(fh.RegexFindAll(regex, "aaaa", -1)))
	assert.Equal(t, 0, len(fh.RegexFindAll(regex, "none", -1)))
}

func TestMustRegexFindAll(t *testing.T) {
	type args struct {
		regex, s string
		n        int
	}
	cases := []struct {
		expected int
		args     args
	}{
		{1, args{"a{2}", "aa", -1}},
		{1, args{"a{2}", "aaaaaaaa", 1}},
		{2, args{"a{2}", "aaaa", -1}},
		{0, args{"a{2}", "none", -1}},
	}

	fh := NewFunctionHandler()
	for _, c := range cases {
		res, err := fh.MustRegexFindAll(c.args.regex, c.args.s, c.args.n)
		if err != nil {
			t.Errorf("regexFindAll test case %v failed with err %s", c, err)
		}
		assert.Equal(t, c.expected, len(res), "case %#v", c.args)
	}
}

func TestRegexFindl(t *testing.T) {
	fh := NewFunctionHandler()
	regex := "fo.?"
	assert.Equal(t, "foo", fh.RegexFind(regex, "foorbar"))
	assert.Equal(t, "foo", fh.RegexFind(regex, "foo foe fome"))
	assert.Equal(t, "", fh.RegexFind(regex, "none"))
}

func TestMustRegexFindl(t *testing.T) {
	type args struct{ regex, s string }
	cases := []struct {
		expected string
		args     args
	}{
		{"foo", args{"fo.?", "foorbar"}},
		{"foo", args{"fo.?", "foo foe fome"}},
		{"", args{"fo.?", "none"}},
	}

	fh := NewFunctionHandler()
	for _, c := range cases {
		res, err := fh.MustRegexFind(c.args.regex, c.args.s)
		if err != nil {
			t.Errorf("regexFind test case %v failed with err %s", c, err)
		}
		assert.Equal(t, c.expected, res, "case %#v", c.args)
	}
}

func TestRegexReplaceAll(t *testing.T) {
	fh := NewFunctionHandler()
	regex := "a(x*)b"
	assert.Equal(t, "-T-T-", fh.RegexReplaceAll(regex, "-ab-axxb-", "T"))
	assert.Equal(t, "--xx-", fh.RegexReplaceAll(regex, "-ab-axxb-", "$1"))
	assert.Equal(t, "---", fh.RegexReplaceAll(regex, "-ab-axxb-", "$1W"))
	assert.Equal(t, "-W-xxW-", fh.RegexReplaceAll(regex, "-ab-axxb-", "${1}W"))
}

func TestMustRegexReplaceAll(t *testing.T) {
	type args struct{ regex, s, repl string }
	cases := []struct {
		expected string
		args     args
	}{
		{"-T-T-", args{"a(x*)b", "-ab-axxb-", "T"}},
		{"--xx-", args{"a(x*)b", "-ab-axxb-", "$1"}},
		{"---", args{"a(x*)b", "-ab-axxb-", "$1W"}},
		{"-W-xxW-", args{"a(x*)b", "-ab-axxb-", "${1}W"}},
	}

	fh := NewFunctionHandler()
	for _, c := range cases {
		res, err := fh.MustRegexReplaceAll(c.args.regex, c.args.s, c.args.repl)
		if err != nil {
			t.Errorf("regexReplaceAll test case %v failed with err %s", c, err)
		}
		assert.Equal(t, c.expected, res, "case %#v", c.args)
	}
}

func TestRegexReplaceAllLiteral(t *testing.T) {
	fh := NewFunctionHandler()
	regex := "a(x*)b"
	assert.Equal(t, "-T-T-", fh.RegexReplaceAllLiteral(regex, "-ab-axxb-", "T"))
	assert.Equal(t, "-$1-$1-", fh.RegexReplaceAllLiteral(regex, "-ab-axxb-", "$1"))
	assert.Equal(t, "-${1}-${1}-", fh.RegexReplaceAllLiteral(regex, "-ab-axxb-", "${1}"))
}

func TestMustRegexReplaceAllLiteral(t *testing.T) {
	type args struct{ regex, s, repl string }
	cases := []struct {
		expected string
		args     args
	}{
		{"-T-T-", args{"a(x*)b", "-ab-axxb-", "T"}},
		{"-$1-$1-", args{"a(x*)b", "-ab-axxb-", "$1"}},
		{"-${1}-${1}-", args{"a(x*)b", "-ab-axxb-", "${1}"}},
	}

	fh := NewFunctionHandler()
	for _, c := range cases {
		res, err := fh.MustRegexReplaceAllLiteral(c.args.regex, c.args.s, c.args.repl)
		if err != nil {
			t.Errorf("regexReplaceAllLiteral test case %v failed with err %s", c, err)
		}
		assert.Equal(t, c.expected, res, "case %#v", c.args)
	}
}

func TestRegexSplit(t *testing.T) {
	fh := NewFunctionHandler()
	regex := "a"
	assert.Equal(t, 4, len(fh.RegexSplit(regex, "banana", -1)))
	assert.Equal(t, 0, len(fh.RegexSplit(regex, "banana", 0)))
	assert.Equal(t, 1, len(fh.RegexSplit(regex, "banana", 1)))
	assert.Equal(t, 2, len(fh.RegexSplit(regex, "banana", 2)))

	regex = "z+"
	assert.Equal(t, 2, len(fh.RegexSplit(regex, "pizza", -1)))
	assert.Equal(t, 0, len(fh.RegexSplit(regex, "pizza", 0)))
	assert.Equal(t, 1, len(fh.RegexSplit(regex, "pizza", 1)))
	assert.Equal(t, 2, len(fh.RegexSplit(regex, "pizza", 2)))
}

func TestMustRegexSplit(t *testing.T) {
	type args struct {
		regex, s string
		n        int
	}
	cases := []struct {
		expected int
		args     args
	}{
		{4, args{"a", "banana", -1}},
		{0, args{"a", "banana", 0}},
		{1, args{"a", "banana", 1}},
		{2, args{"a", "banana", 2}},
		{2, args{"z+", "pizza", -1}},
		{0, args{"z+", "pizza", 0}},
		{1, args{"z+", "pizza", 1}},
		{2, args{"z+", "pizza", 2}},
	}

	for _, c := range cases {
		res, err := NewFunctionHandler().MustRegexSplit(c.args.regex, c.args.s, c.args.n)
		if err != nil {
			t.Errorf("regexSplit test case %v failed with err %s", c, err)
		}
		assert.Equal(t, c.expected, len(res), "case %#v", c.args)
	}
}

func TestRegexQuoteMeta(t *testing.T) {
	fh := NewFunctionHandler()
	assert.Equal(t, "1\\.2\\.3", fh.RegexQuoteMeta("1.2.3"))
	assert.Equal(t, "pretzel", fh.RegexQuoteMeta("pretzel"))
}

func TestBiggest(t *testing.T) {
	tpl := `{{ biggest 1 2 3 345 5 6 7}}`
	if err := runt(tpl, `345`); err != nil {
		t.Error(err)
	}

	tpl = `{{ max 345}}`
	if err := runt(tpl, `345`); err != nil {
		t.Error(err)
	}
}
func TestMaxf(t *testing.T) {
	tpl := `{{ maxf 1 2 3 345.7 5 6 7}}`
	if err := runt(tpl, `345.7`); err != nil {
		t.Error(err)
	}

	tpl = `{{ max 345 }}`
	if err := runt(tpl, `345`); err != nil {
		t.Error(err)
	}
}
func TestMin(t *testing.T) {
	tpl := `{{ min 1 2 3 345 5 6 7}}`
	if err := runt(tpl, `1`); err != nil {
		t.Error(err)
	}

	tpl = `{{ min 345}}`
	if err := runt(tpl, `345`); err != nil {
		t.Error(err)
	}
}

func TestMinf(t *testing.T) {
	tpl := `{{ minf 1.4 2 3 345.6 5 6 7}}`
	if err := runt(tpl, `1.4`); err != nil {
		t.Error(err)
	}

	tpl = `{{ minf 345 }}`
	if err := runt(tpl, `345`); err != nil {
		t.Error(err)
	}
}

func TestToFloat64(t *testing.T) {
	fh := NewFunctionHandler()
	target := float64(102)
	if target != fh.ToFloat64(int8(102)) {
		t.Errorf("Expected 102")
	}
	if target != fh.ToFloat64(int(102)) {
		t.Errorf("Expected 102")
	}
	if target != fh.ToFloat64(int32(102)) {
		t.Errorf("Expected 102")
	}
	if target != fh.ToFloat64(int16(102)) {
		t.Errorf("Expected 102")
	}
	if target != fh.ToFloat64(int64(102)) {
		t.Errorf("Expected 102")
	}
	if target != fh.ToFloat64("102") {
		t.Errorf("Expected 102")
	}
	if 0 != fh.ToFloat64("frankie") {
		t.Errorf("Expected 0")
	}
	if target != fh.ToFloat64(uint16(102)) {
		t.Errorf("Expected 102")
	}
	if target != fh.ToFloat64(uint64(102)) {
		t.Errorf("Expected 102")
	}
	if 102.1234 != fh.ToFloat64(float64(102.1234)) {
		t.Errorf("Expected 102.1234")
	}
	if 1 != fh.ToFloat64(true) {
		t.Errorf("Expected 102")
	}
}
func TestToInt64(t *testing.T) {
	fh := NewFunctionHandler()
	target := int64(102)
	if target != fh.ToInt64(int8(102)) {
		t.Errorf("Expected 102")
	}
	if target != fh.ToInt64(int(102)) {
		t.Errorf("Expected 102")
	}
	if target != fh.ToInt64(int32(102)) {
		t.Errorf("Expected 102")
	}
	if target != fh.ToInt64(int16(102)) {
		t.Errorf("Expected 102")
	}
	if target != fh.ToInt64(int64(102)) {
		t.Errorf("Expected 102")
	}
	if target != fh.ToInt64("102") {
		t.Errorf("Expected 102")
	}
	if 0 != fh.ToInt64("frankie") {
		t.Errorf("Expected 0")
	}
	if target != fh.ToInt64(uint16(102)) {
		t.Errorf("Expected 102")
	}
	if target != fh.ToInt64(uint64(102)) {
		t.Errorf("Expected 102")
	}
	if target != fh.ToInt64(float64(102.1234)) {
		t.Errorf("Expected 102")
	}
	if 1 != fh.ToInt64(true) {
		t.Errorf("Expected 102")
	}
}

func TestToInt(t *testing.T) {
	fh := NewFunctionHandler()
	target := int(102)
	if target != fh.ToInt(int8(102)) {
		t.Errorf("Expected 102")
	}
	if target != fh.ToInt(int(102)) {
		t.Errorf("Expected 102")
	}
	if target != fh.ToInt(int32(102)) {
		t.Errorf("Expected 102")
	}
	if target != fh.ToInt(int16(102)) {
		t.Errorf("Expected 102")
	}
	if target != fh.ToInt(int64(102)) {
		t.Errorf("Expected 102")
	}
	if target != fh.ToInt("102") {
		t.Errorf("Expected 102")
	}
	if 0 != fh.ToInt("frankie") {
		t.Errorf("Expected 0")
	}
	if target != fh.ToInt(uint16(102)) {
		t.Errorf("Expected 102")
	}
	if target != fh.ToInt(uint64(102)) {
		t.Errorf("Expected 102")
	}
	if target != fh.ToInt(float64(102.1234)) {
		t.Errorf("Expected 102")
	}
	if 1 != fh.ToInt(true) {
		t.Errorf("Expected 102")
	}
}

func TestToDecimal(t *testing.T) {
	tests := map[interface{}]int64{
		"777": 511,
		777:   511,
		770:   504,
		755:   493,
	}

	for input, expectedResult := range tests {
		result := NewFunctionHandler().ToDecimal(input)
		if result != expectedResult {
			t.Errorf("Expected %v but got %v", expectedResult, result)
		}
	}
}

func TestAdd1(t *testing.T) {
	tpl := `{{ 3 | add1 }}`
	if err := runt(tpl, `4`); err != nil {
		t.Error(err)
	}
}

func TestAdd1f(t *testing.T) {
	tpl := `{{ 3.4 | add1f }}`
	if err := runt(tpl, `4.4`); err != nil {
		t.Error(err)
	}
}

func TestAdd(t *testing.T) {
	tpl := `{{ 3 | add 1 2}}`
	if err := runt(tpl, `6`); err != nil {
		t.Error(err)
	}
}

func TestAddf(t *testing.T) {
	tpl := `{{ 3 | addf 1.5 2.2}}`
	if err := runt(tpl, `6.7`); err != nil {
		t.Error(err)
	}
}

func TestDiv(t *testing.T) {
	tpl := `{{ 4 | div 5 }}`
	if err := runt(tpl, `1`); err != nil {
		t.Error(err)
	}
}

func TestDivf(t *testing.T) {
	tpl := `{{ 2 | divf 5 4 }}`
	if err := runt(tpl, `0.625`); err != nil {
		t.Error(err)
	}
}

func TestMul(t *testing.T) {
	tpl := `{{ 1 | mul "2" 3 "4"}}`
	if err := runt(tpl, `24`); err != nil {
		t.Error(err)
	}
}

func TestMulf(t *testing.T) {
	tpl := `{{ 1.2 | mulf "2.4" 10 "4"}}`
	if err := runt(tpl, `115.19999999999999`); err != nil {
		t.Error(err)
	}
}

func TestSub(t *testing.T) {
	tpl := `{{ 3 | sub 14 }}`
	if err := runt(tpl, `11`); err != nil {
		t.Error(err)
	}
}

func TestSubf(t *testing.T) {
	tpl := `{{ 3 | subf 4.5 1 }}`
	if err := runt(tpl, `0.5`); err != nil {
		t.Error(err)
	}
}

func TestCeil(t *testing.T) {
	fh := NewFunctionHandler()
	assert.Equal(t, 123.0, fh.Ceil(123))
	assert.Equal(t, 123.0, fh.Ceil("123"))
	assert.Equal(t, 124.0, fh.Ceil(123.01))
	assert.Equal(t, 124.0, fh.Ceil("123.01"))
}

func TestFloor(t *testing.T) {
	fh := NewFunctionHandler()
	assert.Equal(t, 123.0, fh.Floor(123))
	assert.Equal(t, 123.0, fh.Floor("123"))
	assert.Equal(t, 123.0, fh.Floor(123.9999))
	assert.Equal(t, 123.0, fh.Floor("123.9999"))
}

func TestRound(t *testing.T) {
	fh := NewFunctionHandler()
	assert.Equal(t, 123.556, fh.Round(123.5555, 3))
	assert.Equal(t, 123.556, fh.Round("123.55555", 3))
	assert.Equal(t, 124.0, fh.Round(123.500001, 0))
	assert.Equal(t, 123.0, fh.Round(123.49999999, 0))
	assert.Equal(t, 123.23, fh.Round(123.2329999, 2, .3))
	assert.Equal(t, 123.24, fh.Round(123.233, 2, .3))
}

func TestRandomInt(t *testing.T) {
	var tests = []struct {
		min int
		max int
	}{
		{10, 11},
		{10, 13},
		{0, 1},
		{5, 50},
	}
	for _, v := range tests {
		x, _ := runRaw(fmt.Sprintf(`{{ randInt %d %d }}`, v.min, v.max), nil)
		r, err := strconv.Atoi(x)
		assert.NoError(t, err)
		assert.True(t, func(min, max, r int) bool {
			return r >= v.min && r < v.max
		}(v.min, v.max, r))
	}
}

func TestGetHostByName(t *testing.T) {
	tpl := `{{"www.google.com" | getHostByName}}`

	resolvedIP, _ := runRaw(tpl, nil)

	ip := net.ParseIP(resolvedIP)
	assert.NotNil(t, ip)
	assert.NotEmpty(t, ip)
}

func TestTuple(t *testing.T) {
	tpl := `{{$t := tuple 1 "a" "foo"}}{{index $t 2}}{{index $t 0 }}{{index $t 1}}`
	if err := runt(tpl, "foo1a"); err != nil {
		t.Error(err)
	}
}

func TestList(t *testing.T) {
	tpl := `{{$t := list 1 "a" "foo"}}{{index $t 2}}{{index $t 0 }}{{index $t 1}}`
	if err := runt(tpl, "foo1a"); err != nil {
		t.Error(err)
	}
}

func TestPush(t *testing.T) {
	// Named `append` in the function map
	tests := map[string]string{
		`{{ $t := tuple 1 2 3  }}{{ append $t 4 | len }}`:                             "4",
		`{{ $t := tuple 1 2 3 4  }}{{ append $t 5 | join "-" }}`:                      "1-2-3-4-5",
		`{{ $t := regexSplit "/" "foo/bar/baz" -1 }}{{ append $t "qux" | join "-" }}`: "foo-bar-baz-qux",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestMustPush(t *testing.T) {
	// Named `append` in the function map
	tests := map[string]string{
		`{{ $t := tuple 1 2 3  }}{{ mustAppend $t 4 | len }}`:                           "4",
		`{{ $t := tuple 1 2 3 4  }}{{ mustAppend $t 5 | join "-" }}`:                    "1-2-3-4-5",
		`{{ $t := regexSplit "/" "foo/bar/baz" -1 }}{{ mustPush $t "qux" | join "-" }}`: "foo-bar-baz-qux",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestChunk(t *testing.T) {
	tests := map[string]string{
		`{{ tuple 1 2 3 4 5 6 7 | chunk 3 | len }}`:                                 "3",
		`{{ tuple | chunk 3 | len }}`:                                               "0",
		`{{ range ( tuple 1 2 3 4 5 6 7 8 9 | chunk 3 ) }}{{. | join "-"}}|{{end}}`: "1-2-3|4-5-6|7-8-9|",
		`{{ range ( tuple 1 2 3 4 5 6 7 8 | chunk 3 ) }}{{. | join "-"}}|{{end}}`:   "1-2-3|4-5-6|7-8|",
		`{{ range ( tuple 1 2 | chunk 3 ) }}{{. | join "-"}}|{{end}}`:               "1-2|",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestMustChunk(t *testing.T) {
	tests := map[string]string{
		`{{ tuple 1 2 3 4 5 6 7 | mustChunk 3 | len }}`:                                 "3",
		`{{ tuple | mustChunk 3 | len }}`:                                               "0",
		`{{ range ( tuple 1 2 3 4 5 6 7 8 9 | mustChunk 3 ) }}{{. | join "-"}}|{{end}}`: "1-2-3|4-5-6|7-8-9|",
		`{{ range ( tuple 1 2 3 4 5 6 7 8 | mustChunk 3 ) }}{{. | join "-"}}|{{end}}`:   "1-2-3|4-5-6|7-8|",
		`{{ range ( tuple 1 2 | mustChunk 3 ) }}{{. | join "-"}}|{{end}}`:               "1-2|",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestPrepend(t *testing.T) {
	tests := map[string]string{
		`{{ $t := tuple 1 2 3  }}{{ prepend $t 0 | len }}`:                             "4",
		`{{ $t := tuple 1 2 3 4  }}{{ prepend $t 0 | join "-" }}`:                      "0-1-2-3-4",
		`{{ $t := regexSplit "/" "foo/bar/baz" -1 }}{{ prepend $t "qux" | join "-" }}`: "qux-foo-bar-baz",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestMustPrepend(t *testing.T) {
	tests := map[string]string{
		`{{ $t := tuple 1 2 3  }}{{ mustPrepend $t 0 | len }}`:                             "4",
		`{{ $t := tuple 1 2 3 4  }}{{ mustPrepend $t 0 | join "-" }}`:                      "0-1-2-3-4",
		`{{ $t := regexSplit "/" "foo/bar/baz" -1 }}{{ mustPrepend $t "qux" | join "-" }}`: "qux-foo-bar-baz",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestFirst(t *testing.T) {
	tests := map[string]string{
		`{{ list 1 2 3 | first }}`:                          "1",
		`{{ list | first }}`:                                "<no value>",
		`{{ regexSplit "/src/" "foo/src/bar" -1 | first }}`: "foo",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestMustFirst(t *testing.T) {
	tests := map[string]string{
		`{{ list 1 2 3 | mustFirst }}`:                          "1",
		`{{ list | mustFirst }}`:                                "<no value>",
		`{{ regexSplit "/src/" "foo/src/bar" -1 | mustFirst }}`: "foo",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestLast(t *testing.T) {
	tests := map[string]string{
		`{{ list 1 2 3 | last }}`:                          "3",
		`{{ list | last }}`:                                "<no value>",
		`{{ regexSplit "/src/" "foo/src/bar" -1 | last }}`: "bar",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestMustLast(t *testing.T) {
	tests := map[string]string{
		`{{ list 1 2 3 | mustLast }}`:                          "3",
		`{{ list | mustLast }}`:                                "<no value>",
		`{{ regexSplit "/src/" "foo/src/bar" -1 | mustLast }}`: "bar",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestInitial(t *testing.T) {
	tests := map[string]string{
		`{{ list 1 2 3 | initial | len }}`:                "2",
		`{{ list 1 2 3 | initial | last }}`:               "2",
		`{{ list 1 2 3 | initial | first }}`:              "1",
		`{{ list | initial }}`:                            "[]",
		`{{ regexSplit "/" "foo/bar/baz" -1 | initial }}`: "[foo bar]",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestMustInitial(t *testing.T) {
	tests := map[string]string{
		`{{ list 1 2 3 | mustInitial | len }}`:                "2",
		`{{ list 1 2 3 | mustInitial | last }}`:               "2",
		`{{ list 1 2 3 | mustInitial | first }}`:              "1",
		`{{ list | mustInitial }}`:                            "[]",
		`{{ regexSplit "/" "foo/bar/baz" -1 | mustInitial }}`: "[foo bar]",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestRest(t *testing.T) {
	tests := map[string]string{
		`{{ list 1 2 3 | rest | len }}`:                "2",
		`{{ list 1 2 3 | rest | last }}`:               "3",
		`{{ list 1 2 3 | rest | first }}`:              "2",
		`{{ list | rest }}`:                            "[]",
		`{{ regexSplit "/" "foo/bar/baz" -1 | rest }}`: "[bar baz]",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestMustRest(t *testing.T) {
	tests := map[string]string{
		`{{ list 1 2 3 | mustRest | len }}`:                "2",
		`{{ list 1 2 3 | mustRest | last }}`:               "3",
		`{{ list 1 2 3 | mustRest | first }}`:              "2",
		`{{ list | mustRest }}`:                            "[]",
		`{{ regexSplit "/" "foo/bar/baz" -1 | mustRest }}`: "[bar baz]",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestReverse(t *testing.T) {
	tests := map[string]string{
		`{{ list 1 2 3 | reverse | first }}`:              "3",
		`{{ list 1 2 3 | reverse | rest | first }}`:       "2",
		`{{ list 1 2 3 | reverse | last }}`:               "1",
		`{{ list 1 2 3 4 | reverse }}`:                    "[4 3 2 1]",
		`{{ list 1 | reverse }}`:                          "[1]",
		`{{ list | reverse }}`:                            "[]",
		`{{ regexSplit "/" "foo/bar/baz" -1 | reverse }}`: "[baz bar foo]",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestMustReverse(t *testing.T) {
	tests := map[string]string{
		`{{ list 1 2 3 | mustReverse | first }}`:              "3",
		`{{ list 1 2 3 | mustReverse | rest | first }}`:       "2",
		`{{ list 1 2 3 | mustReverse | last }}`:               "1",
		`{{ list 1 2 3 4 | mustReverse }}`:                    "[4 3 2 1]",
		`{{ list 1 | mustReverse }}`:                          "[1]",
		`{{ list | mustReverse }}`:                            "[]",
		`{{ regexSplit "/" "foo/bar/baz" -1 | mustReverse }}`: "[baz bar foo]",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestCompact(t *testing.T) {
	tests := map[string]string{
		`{{ list 1 0 "" "hello" | compact }}`:          `[1 hello]`,
		`{{ list "" "" | compact }}`:                   `[]`,
		`{{ list | compact }}`:                         `[]`,
		`{{ regexSplit "/" "foo//bar" -1 | compact }}`: "[foo bar]",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestMustCompact(t *testing.T) {
	tests := map[string]string{
		`{{ list 1 0 "" "hello" | mustCompact }}`:          `[1 hello]`,
		`{{ list "" "" | mustCompact }}`:                   `[]`,
		`{{ list | mustCompact }}`:                         `[]`,
		`{{ regexSplit "/" "foo//bar" -1 | mustCompact }}`: "[foo bar]",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestUniq(t *testing.T) {
	tests := map[string]string{
		`{{ list 1 2 3 4 | uniq }}`:                    `[1 2 3 4]`,
		`{{ list "a" "b" "c" "d" | uniq }}`:            `[a b c d]`,
		`{{ list 1 1 1 1 2 2 2 2 | uniq }}`:            `[1 2]`,
		`{{ list "foo" 1 1 1 1 "foo" "foo" | uniq }}`:  `[foo 1]`,
		`{{ list | uniq }}`:                            `[]`,
		`{{ regexSplit "/" "foo/foo/bar" -1 | uniq }}`: "[foo bar]",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestMustUniq(t *testing.T) {
	tests := map[string]string{
		`{{ list 1 2 3 4 | mustUniq }}`:                    `[1 2 3 4]`,
		`{{ list "a" "b" "c" "d" | mustUniq }}`:            `[a b c d]`,
		`{{ list 1 1 1 1 2 2 2 2 | mustUniq }}`:            `[1 2]`,
		`{{ list "foo" 1 1 1 1 "foo" "foo" | mustUniq }}`:  `[foo 1]`,
		`{{ list | mustUniq }}`:                            `[]`,
		`{{ regexSplit "/" "foo/foo/bar" -1 | mustUniq }}`: "[foo bar]",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestWithout(t *testing.T) {
	tests := map[string]string{
		`{{ without (list 1 2 3 4) 1 }}`:                         `[2 3 4]`,
		`{{ without (list "a" "b" "c" "d") "a" }}`:               `[b c d]`,
		`{{ without (list 1 1 1 1 2) 1 }}`:                       `[2]`,
		`{{ without (list) 1 }}`:                                 `[]`,
		`{{ without (list 1 2 3) }}`:                             `[1 2 3]`,
		`{{ without list }}`:                                     `[]`,
		`{{ without (regexSplit "/" "foo/bar/baz" -1 ) "foo" }}`: "[bar baz]",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestMustWithout(t *testing.T) {
	tests := map[string]string{
		`{{ mustWithout (list 1 2 3 4) 1 }}`:                         `[2 3 4]`,
		`{{ mustWithout (list "a" "b" "c" "d") "a" }}`:               `[b c d]`,
		`{{ mustWithout (list 1 1 1 1 2) 1 }}`:                       `[2]`,
		`{{ mustWithout (list) 1 }}`:                                 `[]`,
		`{{ mustWithout (list 1 2 3) }}`:                             `[1 2 3]`,
		`{{ mustWithout list }}`:                                     `[]`,
		`{{ mustWithout (regexSplit "/" "foo/bar/baz" -1 ) "foo" }}`: "[bar baz]",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestHas(t *testing.T) {
	tests := map[string]string{
		`{{ list 1 2 3 | has 1 }}`:                          `true`,
		`{{ list 1 2 3 | has 4 }}`:                          `false`,
		`{{ regexSplit "/" "foo/bar/baz" -1 | has "bar" }}`: `true`,
		`{{ has "bar" nil }}`:                               `false`,
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestMustHas(t *testing.T) {
	tests := map[string]string{
		`{{ list 1 2 3 | mustHas 1 }}`:                          `true`,
		`{{ list 1 2 3 | mustHas 4 }}`:                          `false`,
		`{{ regexSplit "/" "foo/bar/baz" -1 | mustHas "bar" }}`: `true`,
		`{{ mustHas "bar" nil }}`:                               `false`,
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestSlice(t *testing.T) {
	tests := map[string]string{
		`{{ slice (list 1 2 3) }}`:                          "[1 2 3]",
		`{{ slice (list 1 2 3) 0 1 }}`:                      "[1]",
		`{{ slice (list 1 2 3) 1 3 }}`:                      "[2 3]",
		`{{ slice (list 1 2 3) 1 }}`:                        "[2 3]",
		`{{ slice (regexSplit "/" "foo/bar/baz" -1) 1 2 }}`: "[bar]",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestMustSlice(t *testing.T) {
	tests := map[string]string{
		`{{ mustSlice (list 1 2 3) }}`:                          "[1 2 3]",
		`{{ mustSlice (list 1 2 3) 0 1 }}`:                      "[1]",
		`{{ mustSlice (list 1 2 3) 1 3 }}`:                      "[2 3]",
		`{{ mustSlice (list 1 2 3) 1 }}`:                        "[2 3]",
		`{{ mustSlice (regexSplit "/" "foo/bar/baz" -1) 1 2 }}`: "[bar]",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestConcat(t *testing.T) {
	tests := map[string]string{
		`{{ concat (list 1 2 3) }}`:                                   "[1 2 3]",
		`{{ concat (list 1 2 3) (list 4 5) }}`:                        "[1 2 3 4 5]",
		`{{ concat (list 1 2 3) (list 4 5) (list) }}`:                 "[1 2 3 4 5]",
		`{{ concat (list 1 2 3) (list 4 5) (list nil) }}`:             "[1 2 3 4 5 <nil>]",
		`{{ concat (list 1 2 3) (list 4 5) (list ( list "foo" ) ) }}`: "[1 2 3 4 5 [foo]]",
	}
	for tpl, expect := range tests {
		assert.NoError(t, runt(tpl, expect))
	}
}

func TestIssue188(t *testing.T) {
	tests := map[string]string{

		// This first test shows two merges and the merge is NOT A DEEP COPY MERGE.
		// The first merge puts $one on to $target. When the second merge of $two
		// on to $target the nested dict brought over from $one is changed on
		// $one as well as $target.
		`{{- $target := dict -}}
			{{- $one := dict "foo" (dict "bar" "baz") "qux" true -}}
			{{- $two := dict "foo" (dict "bar" "baz2") "qux" false -}}
			{{- mergeOverwrite $target $one | toString | trunc 0 }}{{ $__ := mergeOverwrite $target $two }}{{ $one }}`: "map[foo:map[bar:baz2] qux:true]",

		// This test uses deepCopy on $one to create a deep copy and then merge
		// that. In this case the merge of $two on to $target does not affect
		// $one because a deep copy was used for that merge.
		`{{- $target := dict -}}
			{{- $one := dict "foo" (dict "bar" "baz") "qux" true -}}
			{{- $two := dict "foo" (dict "bar" "baz2") "qux" false -}}
			{{- deepCopy $one | mergeOverwrite $target | toString | trunc 0 }}{{ $__ := mergeOverwrite $target $two }}{{ $one }}`: "map[foo:map[bar:baz] qux:true]",
	}

	for tpl, expect := range tests {
		if err := runt(tpl, expect); err != nil {
			t.Error(err)
		}
	}
}

func TestEnv(t *testing.T) {
	os.Setenv("FOO", "bar")
	tpl := `{{env "FOO"}}`
	if err := runt(tpl, "bar"); err != nil {
		t.Error(err)
	}
}

func TestExpandEnv(t *testing.T) {
	os.Setenv("FOO", "bar")
	tpl := `{{expandenv "Hello $FOO"}}`
	if err := runt(tpl, "Hello bar"); err != nil {
		t.Error(err)
	}
}

func TestBase(t *testing.T) {
	assert.NoError(t, runt(`{{ base "foo/bar" }}`, "bar"))
}

func TestDir(t *testing.T) {
	assert.NoError(t, runt(`{{ dir "foo/bar/baz" }}`, "foo/bar"))
}

func TestIsAbs(t *testing.T) {
	assert.NoError(t, runt(`{{ isAbs "/foo" }}`, "true"))
	assert.NoError(t, runt(`{{ isAbs "foo" }}`, "false"))
}

func TestClean(t *testing.T) {
	assert.NoError(t, runt(`{{ clean "/foo/../foo/../bar" }}`, "/bar"))
}

func TestExt(t *testing.T) {
	assert.NoError(t, runt(`{{ ext "/foo/bar/baz.txt" }}`, ".txt"))
}

func TestRegex(t *testing.T) {
	assert.NoError(t, runt(`{{ regexQuoteMeta "1.2.3" }}`, "1\\.2\\.3"))
	assert.NoError(t, runt(`{{ regexQuoteMeta "pretzel" }}`, "pretzel"))
}

// runt runs a template and checks that the output exactly matches the expected string.
func runt(tpl, expect string) error {
	return runtv(tpl, expect, map[string]string{})
}

// runtv takes a template, and expected return, and values for substitution.
//
// It runs the template and verifies that the output is an exact match.
func runtv(tpl, expect string, vars interface{}) error {
	t := template.Must(template.New("test").Funcs(FuncMap()).Parse(tpl))
	var b bytes.Buffer
	err := t.Execute(&b, vars)
	if err != nil {
		return err
	}
	if expect != b.String() {
		return fmt.Errorf("Expected '%v', got '%v'", expect, b.String())
	}
	return nil
}

// runRaw runs a template with the given variables and returns the result.
func runRaw(tpl string, vars interface{}) (string, error) {
	t := template.Must(template.New("test").Funcs(FuncMap()).Parse(tpl))
	var b bytes.Buffer
	err := t.Execute(&b, vars)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func Example() {
	// Set up variables and template.
	vars := map[string]interface{}{"Name": "  John Jacob Jingleheimer Schmidt "}
	tpl := `Hello {{.Name | trim | lower}}`

	// Get the sprout function map.
	t := template.Must(template.New("test").Funcs(FuncMap()).Parse(tpl))

	err := t.Execute(os.Stdout, vars)
	if err != nil {
		fmt.Printf("Error during template execution: %s", err)
		return
	}
	// Output:
	// Hello john jacob jingleheimer schmidt
}

func TestDict(t *testing.T) {
	tpl := `{{$d := dict 1 2 "three" "four" 5}}{{range $k, $v := $d}}{{$k}}{{$v}}{{end}}`
	out, err := runRaw(tpl, nil)
	if err != nil {
		t.Error(err)
	}
	if len(out) != 12 {
		t.Errorf("Expected length 12, got %d", len(out))
	}
	// dict does not guarantee ordering because it is backed by a map.
	if !strings.Contains(out, "12") {
		t.Error("Expected grouping 12")
	}
	if !strings.Contains(out, "threefour") {
		t.Error("Expected grouping threefour")
	}
	if !strings.Contains(out, "5") {
		t.Error("Expected 5")
	}
	tpl = `{{$t := dict "I" "shot" "the" "albatross"}}{{$t.the}} {{$t.I}}`
	if err := runt(tpl, "albatross shot"); err != nil {
		t.Error(err)
	}
}

func TestUnset(t *testing.T) {
	tpl := `{{- $d := dict "one" 1 "two" 222222 -}}
	{{- $_ := unset $d "two" -}}
	{{- range $k, $v := $d}}{{$k}}{{$v}}{{- end -}}
	`

	expect := "one1"
	if err := runt(tpl, expect); err != nil {
		t.Error(err)
	}
}
func TestHasKey(t *testing.T) {
	tpl := `{{- $d := dict "one" 1 "two" 222222 -}}
	{{- if hasKey $d "one" -}}1{{- end -}}
	`

	expect := "1"
	if err := runt(tpl, expect); err != nil {
		t.Error(err)
	}
}

func TestPluck(t *testing.T) {
	tpl := `
	{{- $d := dict "one" 1 "two" 222222 -}}
	{{- $d2 := dict "one" 1 "two" 33333 -}}
	{{- $d3 := dict "one" 1 -}}
	{{- $d4 := dict "one" 1 "two" 4444 -}}
	{{- pluck "two" $d $d2 $d3 $d4 -}}
	`

	expect := "[222222 33333 4444]"
	if err := runt(tpl, expect); err != nil {
		t.Error(err)
	}
}

func TestKeys(t *testing.T) {
	tests := map[string]string{
		`{{ dict "foo" 1 "bar" 2 | keys | sortAlpha }}`: "[bar foo]",
		`{{ dict | keys }}`:                             "[]",
		`{{ keys (dict "foo" 1) (dict "bar" 2) (dict "bar" 3) | uniq | sortAlpha }}`: "[bar foo]",
	}
	for tpl, expect := range tests {
		if err := runt(tpl, expect); err != nil {
			t.Error(err)
		}
	}
}

func TestPick(t *testing.T) {
	tests := map[string]string{
		`{{- $d := dict "one" 1 "two" 222222 }}{{ pick $d "two" | len -}}`:               "1",
		`{{- $d := dict "one" 1 "two" 222222 }}{{ pick $d "two" -}}`:                     "map[two:222222]",
		`{{- $d := dict "one" 1 "two" 222222 }}{{ pick $d "one" "two" | len -}}`:         "2",
		`{{- $d := dict "one" 1 "two" 222222 }}{{ pick $d "one" "two" "three" | len -}}`: "2",
		`{{- $d := dict }}{{ pick $d "two" | len -}}`:                                    "0",
	}
	for tpl, expect := range tests {
		if err := runt(tpl, expect); err != nil {
			t.Error(err)
		}
	}
}
func TestOmit(t *testing.T) {
	tests := map[string]string{
		`{{- $d := dict "one" 1 "two" 222222 }}{{ omit $d "one" | len -}}`:         "1",
		`{{- $d := dict "one" 1 "two" 222222 }}{{ omit $d "one" -}}`:               "map[two:222222]",
		`{{- $d := dict "one" 1 "two" 222222 }}{{ omit $d "one" "two" | len -}}`:   "0",
		`{{- $d := dict "one" 1 "two" 222222 }}{{ omit $d "two" "three" | len -}}`: "1",
		`{{- $d := dict }}{{ omit $d "two" | len -}}`:                              "0",
	}
	for tpl, expect := range tests {
		if err := runt(tpl, expect); err != nil {
			t.Error(err)
		}
	}
}

func TestGet(t *testing.T) {
	tests := map[string]string{
		`{{- $d := dict "one" 1 }}{{ get $d "one" -}}`:           "1",
		`{{- $d := dict "one" 1 "two" "2" }}{{ get $d "two" -}}`: "2",
		`{{- $d := dict }}{{ get $d "two" -}}`:                   "",
	}
	for tpl, expect := range tests {
		if err := runt(tpl, expect); err != nil {
			t.Error(err)
		}
	}
}

func TestSet(t *testing.T) {
	tpl := `{{- $d := dict "one" 1 "two" 222222 -}}
	{{- $_ := set $d "two" 2 -}}
	{{- $_ := set $d "three" 3 -}}
	{{- if hasKey $d "one" -}}{{$d.one}}{{- end -}}
	{{- if hasKey $d "two" -}}{{$d.two}}{{- end -}}
	{{- if hasKey $d "three" -}}{{$d.three}}{{- end -}}
	`

	expect := "123"
	if err := runt(tpl, expect); err != nil {
		t.Error(err)
	}
}

func TestMerge(t *testing.T) {
	dict := map[string]interface{}{
		"src2": map[string]interface{}{
			"h": 10,
			"i": "i",
			"j": "j",
		},
		"src1": map[string]interface{}{
			"a": 1,
			"b": 2,
			"d": map[string]interface{}{
				"e": "four",
			},
			"g": []int{6, 7},
			"i": "aye",
			"j": "jay",
			"k": map[string]interface{}{
				"l": false,
			},
		},
		"dst": map[string]interface{}{
			"a": "one",
			"c": 3,
			"d": map[string]interface{}{
				"f": 5,
			},
			"g": []int{8, 9},
			"i": "eye",
			"k": map[string]interface{}{
				"l": true,
			},
		},
	}
	tpl := `{{merge .dst .src1 .src2}}`
	_, err := runRaw(tpl, dict)
	if err != nil {
		t.Error(err)
	}
	expected := map[string]interface{}{
		"a": "one", // key overridden
		"b": 2,     // merged from src1
		"c": 3,     // merged from dst
		"d": map[string]interface{}{ // deep merge
			"e": "four",
			"f": 5,
		},
		"g": []int{8, 9}, // overridden - arrays are not merged
		"h": 10,          // merged from src2
		"i": "eye",       // overridden twice
		"j": "jay",       // overridden and merged
		"k": map[string]interface{}{
			"l": true, // overridden
		},
	}
	assert.Equal(t, expected, dict["dst"])
}

func TestMergeOverwrite(t *testing.T) {
	dict := map[string]interface{}{
		"src2": map[string]interface{}{
			"h": 10,
			"i": "i",
			"j": "j",
		},
		"src1": map[string]interface{}{
			"a": 1,
			"b": 2,
			"d": map[string]interface{}{
				"e": "four",
			},
			"g": []int{6, 7},
			"i": "aye",
			"j": "jay",
			"k": map[string]interface{}{
				"l": false,
			},
		},
		"dst": map[string]interface{}{
			"a": "one",
			"c": 3,
			"d": map[string]interface{}{
				"f": 5,
			},
			"g": []int{8, 9},
			"i": "eye",
			"k": map[string]interface{}{
				"l": true,
			},
		},
	}
	tpl := `{{mergeOverwrite .dst .src1 .src2}}`
	_, err := runRaw(tpl, dict)
	if err != nil {
		t.Error(err)
	}
	expected := map[string]interface{}{
		"a": 1, // key overwritten from src1
		"b": 2, // merged from src1
		"c": 3, // merged from dst
		"d": map[string]interface{}{ // deep merge
			"e": "four",
			"f": 5,
		},
		"g": []int{6, 7}, // overwritten src1 wins
		"h": 10,          // merged from src2
		"i": "i",         // overwritten twice src2 wins
		"j": "j",         // overwritten twice src2 wins
		"k": map[string]interface{}{ // deep merge
			"l": false, // overwritten src1 wins
		},
	}
	assert.Equal(t, expected, dict["dst"])
}

func TestValues(t *testing.T) {
	tests := map[string]string{
		`{{- $d := dict "a" 1 "b" 2 }}{{ values $d | sortAlpha | join "," }}`:       "1,2",
		`{{- $d := dict "a" "first" "b" 2 }}{{ values $d | sortAlpha | join "," }}`: "2,first",
	}

	for tpl, expect := range tests {
		if err := runt(tpl, expect); err != nil {
			t.Error(err)
		}
	}
}
func TestDig(t *testing.T) {
	tests := map[string]string{
		`{{- $d := dict "a" (dict "b" (dict "c" 1)) }}{{ dig "a" "b" "c" "" $d }}`:  "1",
		`{{- $d := dict "a" (dict "b" (dict "c" 1)) }}{{ dig "a" "b" "z" "2" $d }}`: "2",
		`{{ dict "a" 1 | dig "a" "" }}`:                                             "1",
		`{{ dict "a" 1 | dig "z" "2" }}`:                                            "2",
	}

	for tpl, expect := range tests {
		if err := runt(tpl, expect); err != nil {
			t.Error(err)
		}
	}
}

func TestFromJson(t *testing.T) {
	dict := map[string]interface{}{"Input": `{"foo": 55}`}

	tpl := `{{.Input | fromJson}}`
	expected := `map[foo:55]`
	if err := runtv(tpl, expected, dict); err != nil {
		t.Error(err)
	}

	tpl = `{{(.Input | fromJson).foo}}`
	expected = `55`
	if err := runtv(tpl, expected, dict); err != nil {
		t.Error(err)
	}
}

func TestToJson(t *testing.T) {
	dict := map[string]interface{}{"Top": map[string]interface{}{"bool": true, "string": "test", "number": 42}}

	tpl := `{{.Top | toJson}}`
	expected := `{"bool":true,"number":42,"string":"test"}`
	if err := runtv(tpl, expected, dict); err != nil {
		t.Error(err)
	}
}

func TestToPrettyJson(t *testing.T) {
	dict := map[string]interface{}{"Top": map[string]interface{}{"bool": true, "string": "test", "number": 42}}
	tpl := `{{.Top | toPrettyJson}}`
	expected := `{
  "bool": true,
  "number": 42,
  "string": "test"
}`
	if err := runtv(tpl, expected, dict); err != nil {
		t.Error(err)
	}
}

func TestToRawJson(t *testing.T) {
	dict := map[string]interface{}{"Top": map[string]interface{}{"bool": true, "string": "test", "number": 42, "html": "<HEAD>"}}
	tpl := `{{.Top | toRawJson}}`
	expected := `{"bool":true,"html":"<HEAD>","number":42,"string":"test"}`

	if err := runtv(tpl, expected, dict); err != nil {
		t.Error(err)
	}
}

func TestHtmlDate(t *testing.T) {
	tpl := `{{ htmlDate 0}}`
	if err := runt(tpl, "1970-01-01"); err != nil {
		t.Error(err)
	}
}

func TestAgo(t *testing.T) {
	tpl := "{{ ago .Time }}"
	if err := runtv(tpl, "2m5s", map[string]interface{}{"Time": time.Now().Add(-125 * time.Second)}); err != nil {
		t.Error(err)
	}

	if err := runtv(tpl, "2h34m17s", map[string]interface{}{"Time": time.Now().Add(-(2*3600 + 34*60 + 17) * time.Second)}); err != nil {
		t.Error(err)
	}

	if err := runtv(tpl, "-5s", map[string]interface{}{"Time": time.Now().Add(5 * time.Second)}); err != nil {
		t.Error(err)
	}
}

func TestToDate(t *testing.T) {
	tpl := `{{toDate "2006-01-02" "2017-12-31" | date "02/01/2006"}}`
	if err := runt(tpl, "31/12/2017"); err != nil {
		t.Error(err)
	}
}

func TestUnixEpoch(t *testing.T) {
	tm, err := time.Parse("02 Jan 06 15:04:05 MST", "13 Jun 19 20:39:39 GMT")
	if err != nil {
		t.Error(err)
	}
	tpl := `{{unixEpoch .Time}}`

	if err = runtv(tpl, "1560458379", map[string]interface{}{"Time": tm}); err != nil {
		t.Error(err)
	}
}

func TestDateInZone(t *testing.T) {
	tm, err := time.Parse("02 Jan 06 15:04:05 MST", "13 Jun 19 20:39:39 GMT")
	if err != nil {
		t.Error(err)
	}
	tpl := `{{ date_in_zone "02 Jan 06 15:04 -0700" .Time "UTC" }}`

	// Test time.Time input
	if err = runtv(tpl, "13 Jun 19 20:39 +0000", map[string]interface{}{"Time": tm}); err != nil {
		t.Error(err)
	}

	// Test pointer to time.Time input
	if err = runtv(tpl, "13 Jun 19 20:39 +0000", map[string]interface{}{"Time": &tm}); err != nil {
		t.Error(err)
	}

	// Test no time input. This should be close enough to time.Now() we can test
	loc, _ := time.LoadLocation("UTC")
	if err = runtv(tpl, time.Now().In(loc).Format("02 Jan 06 15:04 -0700"), map[string]interface{}{"Time": ""}); err != nil {
		t.Error(err)
	}

	// Test unix timestamp as int64
	if err = runtv(tpl, "13 Jun 19 20:39 +0000", map[string]interface{}{"Time": int64(1560458379)}); err != nil {
		t.Error(err)
	}

	// Test unix timestamp as int32
	if err = runtv(tpl, "13 Jun 19 20:39 +0000", map[string]interface{}{"Time": int32(1560458379)}); err != nil {
		t.Error(err)
	}

	// Test unix timestamp as int
	if err = runtv(tpl, "13 Jun 19 20:39 +0000", map[string]interface{}{"Time": int(1560458379)}); err != nil {
		t.Error(err)
	}

	// Test case of invalid timezone
	tpl = `{{ date_in_zone "02 Jan 06 15:04 -0700" .Time "foobar" }}`
	if err = runtv(tpl, "13 Jun 19 20:39 +0000", map[string]interface{}{"Time": tm}); err != nil {
		t.Error(err)
	}
}

func TestDuration(t *testing.T) {
	tpl := "{{ duration .Secs }}"
	if err := runtv(tpl, "1m1s", map[string]interface{}{"Secs": "61"}); err != nil {
		t.Error(err)
	}
	if err := runtv(tpl, "1h0m0s", map[string]interface{}{"Secs": "3600"}); err != nil {
		t.Error(err)
	}
	// 1d2h3m4s but go is opinionated
	if err := runtv(tpl, "26h3m4s", map[string]interface{}{"Secs": "93784"}); err != nil {
		t.Error(err)
	}
}

func TestDurationRound(t *testing.T) {
	tpl := "{{ durationRound .Time }}"
	if err := runtv(tpl, "2h", map[string]interface{}{"Time": "2h5s"}); err != nil {
		t.Error(err)
	}
	if err := runtv(tpl, "1d", map[string]interface{}{"Time": "24h5s"}); err != nil {
		t.Error(err)
	}
	if err := runtv(tpl, "3mo", map[string]interface{}{"Time": "2400h5s"}); err != nil {
		t.Error(err)
	}
}

const (
	beginCertificate = "-----BEGIN CERTIFICATE-----"
	endCertificate   = "-----END CERTIFICATE-----"
)

var (
	// fastCertKeyAlgos is the list of private key algorithms that are supported for certificate use, and
	// are fast to generate.
	fastCertKeyAlgos = []string{
		"ecdsa",
		"ed25519",
	}
)

func TestSha256Sum(t *testing.T) {
	tpl := `{{"abc" | sha256sum}}`
	if err := runt(tpl, "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"); err != nil {
		t.Error(err)
	}
}
func TestSha1Sum(t *testing.T) {
	tpl := `{{"abc" | sha1sum}}`
	if err := runt(tpl, "a9993e364706816aba3e25717850c26c9cd0d89d"); err != nil {
		t.Error(err)
	}
}

func TestAdler32Sum(t *testing.T) {
	tpl := `{{"abc" | adler32sum}}`
	if err := runt(tpl, "38600999"); err != nil {
		t.Error(err)
	}
}

func TestBcrypt(t *testing.T) {
	out, err := runRaw(`{{"abc" | bcrypt}}`, nil)
	if err != nil {
		t.Error(err)
	}
	if bcrypt_lib.CompareHashAndPassword([]byte(out), []byte("abc")) != nil {
		t.Error("Generated hash is not the equivalent for password:", "abc")
	}
}

type HtpasswdCred struct {
	Username string
	Password string
	Valid    bool
}

func TestHtpasswd(t *testing.T) {
	expectations := []HtpasswdCred{
		{Username: "myUser", Password: "myPassword", Valid: true},
		{Username: "special'o79Cv_*qFe,)<user", Password: "special<j7+3p#6-.Jx2U:m8G;kGypassword", Valid: true},
		{Username: "wrongus:er", Password: "doesn'tmatter", Valid: false}, // ':' isn't allowed in the username - https://tools.ietf.org/html/rfc2617#page-6
	}

	for _, credential := range expectations {
		out, err := runRaw(`{{htpasswd .Username .Password}}`, credential)
		if err != nil {
			t.Error(err)
		}
		result := strings.Split(out, ":")
		if 0 != strings.Compare(credential.Username, result[0]) && credential.Valid {
			t.Error("Generated username did not match for:", credential.Username)
		}
		if bcrypt_lib.CompareHashAndPassword([]byte(result[1]), []byte(credential.Password)) != nil && credential.Valid {
			t.Error("Generated hash is not the equivalent for password:", credential.Password)
		}
	}
}

func TestDerivePassword(t *testing.T) {
	expectations := map[string]string{
		`{{derivePassword 1 "long" "password" "user" "example.com"}}`:    "ZedaFaxcZaso9*",
		`{{derivePassword 2 "long" "password" "user" "example.com"}}`:    "Fovi2@JifpTupx",
		`{{derivePassword 1 "maximum" "password" "user" "example.com"}}`: "pf4zS1LjCg&LjhsZ7T2~",
		`{{derivePassword 1 "medium" "password" "user" "example.com"}}`:  "ZedJuz8$",
		`{{derivePassword 1 "basic" "password" "user" "example.com"}}`:   "pIS54PLs",
		`{{derivePassword 1 "short" "password" "user" "example.com"}}`:   "Zed5",
		`{{derivePassword 1 "pin" "password" "user" "example.com"}}`:     "6685",
	}

	for tpl, result := range expectations {
		out, err := runRaw(tpl, nil)
		if err != nil {
			t.Error(err)
		}
		if 0 != strings.Compare(out, result) {
			t.Error("Generated password does not match for", tpl)
		}
	}
}

// NOTE(bacongobbler): this test is really _slow_ because of how long it takes to compute
// and generate a new crypto key.
func TestGenPrivateKey(t *testing.T) {
	// test that calling by default generates an RSA private key
	tpl := `{{genPrivateKey ""}}`
	out, err := runRaw(tpl, nil)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(out, "RSA PRIVATE KEY") {
		t.Error("Expected RSA PRIVATE KEY")
	}
	// test all acceptable arguments
	tpl = `{{genPrivateKey "rsa"}}`
	out, err = runRaw(tpl, nil)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(out, "RSA PRIVATE KEY") {
		t.Error("Expected RSA PRIVATE KEY")
	}
	tpl = `{{genPrivateKey "dsa"}}`
	out, err = runRaw(tpl, nil)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(out, "DSA PRIVATE KEY") {
		t.Error("Expected DSA PRIVATE KEY")
	}
	tpl = `{{genPrivateKey "ecdsa"}}`
	out, err = runRaw(tpl, nil)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(out, "EC PRIVATE KEY") {
		t.Error("Expected EC PRIVATE KEY")
	}
	tpl = `{{genPrivateKey "ed25519"}}`
	out, err = runRaw(tpl, nil)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(out, "PRIVATE KEY") {
		t.Error("Expected PRIVATE KEY")
	}
	// test bad
	tpl = `{{genPrivateKey "bad"}}`
	out, err = runRaw(tpl, nil)
	if err != nil {
		t.Error(err)
	}
	if out != "Unknown type bad" {
		t.Error("Expected type 'bad' to be an unknown crypto algorithm")
	}
	// ensure that we can base64 encode the string
	tpl = `{{genPrivateKey "rsa" | b64enc}}`
	_, err = runRaw(tpl, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestRandBytes(t *testing.T) {
	tpl := `{{randBytes 12}}`
	out, err := runRaw(tpl, nil)
	if err != nil {
		t.Error(err)
	}

	bytes, err := base64.StdEncoding.DecodeString(out)
	if err != nil {
		t.Error(err)
	}
	if len(bytes) != 12 {
		t.Error("Expected 12 base64-encoded bytes")
	}

	out2, err := runRaw(tpl, nil)
	if err != nil {
		t.Error(err)
	}

	if out == out2 {
		t.Error("Expected subsequent randBytes to be different")
	}
}

func TestUUIDGeneration(t *testing.T) {
	tpl := `{{uuidv4}}`
	out, err := runRaw(tpl, nil)
	if err != nil {
		t.Error(err)
	}

	if len(out) != 36 {
		t.Error("Expected UUID of length 36")
	}

	out2, err := runRaw(tpl, nil)
	if err != nil {
		t.Error(err)
	}

	if out == out2 {
		t.Error("Expected subsequent UUID generations to be different")
	}
}

func TestBuildCustomCert(t *testing.T) {
	ca, _ := NewFunctionHandler().GenerateCertificateAuthority("example.com", 365)
	tpl := fmt.Sprintf(
		`{{- $ca := buildCustomCert "%s" "%s"}}
{{- $ca.Cert }}`,
		base64.StdEncoding.EncodeToString([]byte(ca.Cert)),
		base64.StdEncoding.EncodeToString([]byte(ca.Key)),
	)
	out, err := runRaw(tpl, nil)
	if err != nil {
		t.Error(err)
	}

	tpl2 := fmt.Sprintf(
		`{{- $ca := buildCustomCert "%s" "%s"}}
{{- $ca.Cert }}`,
		base64.StdEncoding.EncodeToString([]byte("fail")),
		base64.StdEncoding.EncodeToString([]byte(ca.Key)),
	)
	out2, _ := runRaw(tpl2, nil)

	assert.Equal(t, out, ca.Cert)
	assert.NotEqual(t, out2, ca.Cert)
}

func TestGenCA(t *testing.T) {
	testGenCA(t, nil)
}

func TestGenCAWithKey(t *testing.T) {
	for _, keyAlgo := range fastCertKeyAlgos {
		t.Run(keyAlgo, func(t *testing.T) {
			testGenCA(t, &keyAlgo)
		})
	}
}

func testGenCA(t *testing.T, keyAlgo *string) {
	const cn = "foo-ca"

	var genCAExpr string
	if keyAlgo == nil {
		genCAExpr = "genCA"
	} else {
		genCAExpr = fmt.Sprintf(`genPrivateKey "%s" | genCAWithKey`, *keyAlgo)
	}

	tpl := fmt.Sprintf(
		`{{- $ca := %s "%s" 365 }}
{{ $ca.Cert }}
`,
		genCAExpr,
		cn,
	)
	out, err := runRaw(tpl, nil)
	if err != nil {
		t.Error(err)
	}
	assert.Contains(t, out, beginCertificate)
	assert.Contains(t, out, endCertificate)

	decodedCert, _ := pem.Decode([]byte(out))
	assert.Nil(t, err)
	cert, err := x509.ParseCertificate(decodedCert.Bytes)
	assert.Nil(t, err)

	assert.Equal(t, cn, cert.Subject.CommonName)
	assert.True(t, cert.IsCA)
}

func TestGenSelfSignedCert(t *testing.T) {
	testGenSelfSignedCert(t, nil)
}

func TestGenSelfSignedCertWithKey(t *testing.T) {
	for _, keyAlgo := range fastCertKeyAlgos {
		t.Run(keyAlgo, func(t *testing.T) {
			testGenSelfSignedCert(t, &keyAlgo)
		})
	}
}

func testGenSelfSignedCert(t *testing.T, keyAlgo *string) {
	const (
		cn   = "foo.com"
		ip1  = "10.0.0.1"
		ip2  = "10.0.0.2"
		dns1 = "bar.com"
		dns2 = "bat.com"
	)

	var genSelfSignedCertExpr string
	if keyAlgo == nil {
		genSelfSignedCertExpr = "genSelfSignedCert"
	} else {
		genSelfSignedCertExpr = fmt.Sprintf(`genPrivateKey "%s" | genSelfSignedCertWithKey`, *keyAlgo)
	}

	tpl := fmt.Sprintf(
		`{{- $cert := %s "%s" (list "%s" "%s") (list "%s" "%s") 365 }}
{{ $cert.Cert }}`,
		genSelfSignedCertExpr,
		cn,
		ip1,
		ip2,
		dns1,
		dns2,
	)

	out, err := runRaw(tpl, nil)
	if err != nil {
		t.Error(err)
	}
	assert.Contains(t, out, beginCertificate)
	assert.Contains(t, out, endCertificate)

	decodedCert, _ := pem.Decode([]byte(out))
	assert.Nil(t, err)
	cert, err := x509.ParseCertificate(decodedCert.Bytes)
	assert.Nil(t, err)

	assert.Equal(t, cn, cert.Subject.CommonName)
	assert.Equal(t, 1, cert.SerialNumber.Sign())
	assert.Equal(t, 2, len(cert.IPAddresses))
	assert.Equal(t, ip1, cert.IPAddresses[0].String())
	assert.Equal(t, ip2, cert.IPAddresses[1].String())
	assert.Contains(t, cert.DNSNames, dns1)
	assert.Contains(t, cert.DNSNames, dns2)
	assert.False(t, cert.IsCA)
}

func TestGenSignedCert(t *testing.T) {
	testGenSignedCert(t, nil, nil)
}

func TestGenSignedCertWithKey(t *testing.T) {
	for _, caKeyAlgo := range fastCertKeyAlgos {
		for _, certKeyAlgo := range fastCertKeyAlgos {
			t.Run(fmt.Sprintf("%s-%s", caKeyAlgo, certKeyAlgo), func(t *testing.T) {
				testGenSignedCert(t, &caKeyAlgo, &certKeyAlgo)
			})
		}
	}
}

func testGenSignedCert(t *testing.T, caKeyAlgo, certKeyAlgo *string) {
	const (
		cn   = "foo.com"
		ip1  = "10.0.0.1"
		ip2  = "10.0.0.2"
		dns1 = "bar.com"
		dns2 = "bat.com"
	)

	var genCAExpr, genSignedCertExpr string
	if caKeyAlgo == nil {
		genCAExpr = "genCA"
	} else {
		genCAExpr = fmt.Sprintf(`genPrivateKey "%s" | genCAWithKey`, *caKeyAlgo)
	}
	if certKeyAlgo == nil {
		genSignedCertExpr = "genSignedCert"
	} else {
		genSignedCertExpr = fmt.Sprintf(`genPrivateKey "%s" | genSignedCertWithKey`, *certKeyAlgo)
	}

	tpl := fmt.Sprintf(
		`{{- $ca := %s "foo" 365 }}
{{- $cert := %s "%s" (list "%s" "%s") (list "%s" "%s") 365 $ca }}
{{ $cert.Cert }}
`,
		genCAExpr,
		genSignedCertExpr,
		cn,
		ip1,
		ip2,
		dns1,
		dns2,
	)
	out, err := runRaw(tpl, nil)
	if err != nil {
		t.Error(err)
	}

	assert.Contains(t, out, beginCertificate)
	assert.Contains(t, out, endCertificate)

	decodedCert, _ := pem.Decode([]byte(out))
	assert.Nil(t, err)
	cert, err := x509.ParseCertificate(decodedCert.Bytes)
	assert.Nil(t, err)

	assert.Equal(t, cn, cert.Subject.CommonName)
	assert.Equal(t, 1, cert.SerialNumber.Sign())
	assert.Equal(t, 2, len(cert.IPAddresses))
	assert.Equal(t, ip1, cert.IPAddresses[0].String())
	assert.Equal(t, ip2, cert.IPAddresses[1].String())
	assert.Contains(t, cert.DNSNames, dns1)
	assert.Contains(t, cert.DNSNames, dns2)
	assert.False(t, cert.IsCA)
}

func TestEncryptDecryptAES(t *testing.T) {
	tpl := `{{"plaintext" | encryptAES "secretkey" | decryptAES "secretkey"}}`
	if err := runt(tpl, "plaintext"); err != nil {
		t.Error(err)
	}
}
